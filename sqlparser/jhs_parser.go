package sqlparser

import (
	"github.com/HYmian/jhs/router"
	"strings"

	"github.com/siddontang/go-log/log"
)

func JHS_GetStmtShardMap(stmt Statement, r *router.Router, bindVars map[string]interface{}) (nodes map[int]string, err error) {
	defer handleError(&err)

	plan := getRoutingPlan(stmt, r)

	plan.bindVars = bindVars

	ns := plan.jhs_shardMapFromPlan()

	nodes = make(map[int]string)
	for k, v := range ns {
		nodes[k] = plan.rule.Nodes[v]
	}

	return nodes, nil
}

func (plan *RoutingPlan) jhs_shardMapFromPlan() map[int]int {
	if plan.criteria == nil {
		//正式运行时此处返回应该是nil
		return map[int]int{-1: 0}
	}

	if plan.rule.Type == router.NoneRuleType || plan.rule.Type == "" {
		if len(plan.fullList) != 1 {
			log.Error("默认节点数大于一")
			panic(NewParserError("invalid rule nodes num %d, must 1", plan.fullList))
		}
		return map[int]int{-1: 0}
	}

	switch criteria := plan.criteria.(type) {
	case Values:
		return plan.jhs_findInsertShard(criteria)
	case BoolExpr:
		return plan.jhs_routingAnalyzeBoolean(criteria)
	default:
		return nil
	}
}

func (plan *RoutingPlan) jhs_findInsertShard(vals Values) map[int]int {
	n_index, t_index := -1, -1
	for i := 0; i < len(vals); i++ {
		first_value_expression := vals[i].(ValTuple)[0]
		tindex, nindex := plan.jhs_findShard(first_value_expression)
		if n_index == -1 {
			n_index = nindex
			t_index = tindex
		} else if n_index != nindex {
			panic(NewParserError("insert has multiple shard targets"))
		}
	}
	log.Info("t_index: %d, n_index: %d", t_index, n_index)
	return map[int]int{t_index: n_index}
}

func (plan *RoutingPlan) jhs_routingAnalyzeBoolean(node BoolExpr) map[int]int {
	switch node := node.(type) {
	case *AndExpr:
		left := plan.jhs_routingAnalyzeBoolean(node.Left)
		right := plan.jhs_routingAnalyzeBoolean(node.Right)
		return unionMap(left, right)
	case *OrExpr:
		left := plan.jhs_routingAnalyzeBoolean(node.Left)
		right := plan.jhs_routingAnalyzeBoolean(node.Right)
		return unionMap(left, right)
	case *ParenBoolExpr:
		return plan.jhs_routingAnalyzeBoolean(node.Expr)
	case *ComparisonExpr:
		switch {
		case StringIn(node.Operator, "="):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.jhs_findConditionShard(node)
			}
		case StringIn(node.Operator, "in"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				return plan.jhs_findConditionShard(node)
			}
		}
	case *RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.jhs_findConditionShard(node)
		}
	}
	return nil
}

func (plan *RoutingPlan) jhs_findConditionShard(expr BoolExpr) map[int]int {
	var t_index, n_index int
	switch criteria := expr.(type) {
	case *ComparisonExpr:
		switch criteria.Operator {
		case "=":
			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				t_index, n_index = plan.jhs_findShard(criteria.Right)
				log.Info("t: %d, n: %d", t_index, n_index)
			} else {
				t_index, n_index = plan.jhs_findShard(criteria.Left)
			}
			return map[int]int{t_index: n_index}
		case "in":
			return plan.jhs_findShardMap(criteria.Right)
		}
	default:
		return nil
	}
	return nil
}

func (plan *RoutingPlan) jhs_findShardMap(valExpr ValExpr) map[int]int {
	shard_map := make(map[int]int)
	switch node := valExpr.(type) {
	case ValTuple:
		for _, n := range node {
			t_index, n_index := plan.jhs_findShard(n)
			shard_map[t_index] = n_index
		}
	}

	return shard_map
}

func (plan *RoutingPlan) jhs_findShard(valExpr ValExpr) (int, int) {
	value := plan.getBoundValue(valExpr)
	return plan.rule.JHSFindNodeIndex(value)
}

func GetVariablesOpt(expr Expr) []string {
	var list []string
	switch expr := expr.(type) {
	case *OrExpr:
		right := GetVariablesOpt(expr.Right)
		left := GetVariablesOpt(expr.Left)
		return unionStringList(right, left)
	case *ComparisonExpr:
		list = append(list, strings.Trim(String(expr.Right), "'"))
		return list
	default:
		panic(NewParserError("where opt hava and, are you kiding me?"))
	}
	return nil
}

func unionStringList(l1 []string, l2 []string) []string {
	if len(l1) == 0 {
		return l2
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]string, 0, len(l1)+len(l2))

	var i = 0
	for ; i < len(l1) && i < len(l2); i++ {
		l3 = append(l3, l1[i])
		l3 = append(l3, l2[i])
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	} else if i != len(l2) {
		l3 = append(l3, l2[i:]...)
	}

	return l3
}

func unionMap(m1 map[int]int, m2 map[int]int) map[int]int {
	if m1 == nil {
		return m2
	} else if m2 == nil {
		return m1
	}

	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
