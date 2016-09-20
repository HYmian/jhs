package proxy

import (
	"fmt"
	"github.com/HYmian/jhs/router"
)

type Schema struct {
	db string

	nodes map[string]*Node

	rule *router.Router
}

func (s *Server) parseRules() error {
	s.rules = make(map[string]*router.Rule)

	for _, c := range s.cfg.ShardRule {
		r := new(router.Rule)
		r.Table = c.Table
		r.Key = c.Key
		r.Type = c.Type
		r.Nodes = c.Nodes

		if r.Type == router.HashRuleType {
			r.Shard = &router.JHSHashShard{NodeShardNum: len(r.Nodes), TableShardNum: c.ShardNum}
		} else if r.Type == router.RangeRuleType {
			rs, err := router.ParseNumShardingSpec(c.Range)
			if err != nil {
				return err
			}

			if len(rs) != len(r.Nodes) {
				return fmt.Errorf("range space %d not equal nodes %d", len(rs), len(r.Nodes))
			}
		} else {
			r.Type = router.NoneRuleType
		}

		s.rules[r.Table] = r
	}

	return nil
}
