package router

import (
	"fmt"
	"strings"
)

type Rule struct {
	Table string
	Key   string

	Type string

	Nodes []string
	//by mian
	Shard JHSShard
}

func (r *Rule) FindNode(key interface{}) string {
	i := r.Shard.FindForKey(key)
	return r.Nodes[i]
}

func (r *Rule) FindNodeIndex(key interface{}) int {
	return r.Shard.FindForKey(key)
}

//by mian
func (r *Rule) JHSFindNodeIndex(key interface{}) (int, int) {
	return r.Shard.JHSFindForKey(key)
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s.%s?key=%v&shard=%s&nodes=%s",
		"r.DB", r.Table, r.Key, r.Type, strings.Join(r.Nodes, ", "))
}

func NewDefaultRule(node string) *Rule {
	var r *Rule = &Rule{
		Type:  NoneRuleType,
		Nodes: []string{node},
	}
	return r
}

func (r *Router) GetRule(table string) *Rule {
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

type Router struct {
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
}

func NewRouter(rules map[string]*Rule, rule *Rule) *Router {
	router := new(Router)
	router.Rules = rules
	router.DefaultRule = rule

	return router
}
