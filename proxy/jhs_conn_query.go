package proxy

import (
	"jhs/client"
	. "jhs/mysql"
	"jhs/router"
	"jhs/sqlparser"
	"sync"

	"github.com/siddontang/go-log/log"
)

func (c *Conn) jhs_handleSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	bindVars := makeBindVars(args)

	conns, err := c.jhs_getShardConns(true, stmt, bindVars)
	if err != nil {
		return err
	} else if conns == nil {
		log.Warn("%s 未能找到合适的节点", sql)
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*Result

	rs, err = c.jhs_executeInShard(conns, true, stmt, args)

	if err == nil {
		err = c.mergeSelectResult(rs, stmt)
	}

	return err
}

func (c *Conn) jhs_handleExec(stmt sqlparser.JHSStatement, sql string, args []interface{}) error {
	bindVars := makeBindVars(args)

	conns, err := c.jhs_getShardConns(false, stmt, bindVars)
	if err != nil {
		return err
	} else if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*Result

	if len(conns) == 1 {
		rs, err = c.jhs_executeInShard(conns, false, stmt, args)
	} else {
		//for multi nodes, 2PC simple, begin, exec, commit
		//if commit error, data maybe corrupt
		for {
			if err = c.jhs_beginShardConns(conns); err != nil {
				break
			}

			if rs, err = c.jhs_executeInShard(conns, false, stmt, args); err != nil {
				break
			}

			err = c.jhs_commitShardConns(conns)
			break
		}
	}

	c.jhs_closeShardConns(conns, err != nil)

	if err == nil {
		err = c.jhs_mergeExecResult(rs)
	}

	return err
}

func (c *Conn) jhs_executeInShard(
	conns map[int]*client.SqlConn,
	isSelect bool,
	stmt sqlparser.JHSStatement,
	args []interface{},
) (
	[]*Result,
	error,
) {
	var wg sync.WaitGroup
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))

	f := func(rs []interface{}, i int, sql string, co *client.SqlConn) {
		//log.Info(sql)
		r, err := co.Execute(sql, args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		if isSelect && !c.isInTransaction() {
			co.Close()
		}

		wg.Done()
	}

	i := 0
	table := stmt.GetTableName()
	for k, co := range conns {
		stmt.SetTableName(table, k)
		go f(rs, i, sqlparser.String(stmt), co)
		i++
	}

	wg.Wait()

	var err error
	r := make([]*Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, err
}

func (c *Conn) jhs_getShardConns(isSelect bool, stmt sqlparser.Statement,
	bindVars map[string]interface{}) (map[int]*client.SqlConn, error) {
	nodes, err := c.jhs_getShardMap(stmt, bindVars)
	if err != nil {
		return nil, err
	} else if nodes == nil {
		return nil, nil
	}

	conns := make(map[int]*client.SqlConn)

	var co *client.SqlConn
	for k, v := range nodes {
		co, err = c.getConn(v, isSelect)
		if err != nil {
			break
		}

		conns[k] = co
	}

	return conns, err
}

func (c *Conn) jhs_getShardMap(stmt sqlparser.Statement,
	bindVars map[string]interface{}) (map[int]*Node, error) {
	if c.db == "" {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	ns, err := sqlparser.JHS_GetStmtShardMap(stmt, router.NewRouter(c.server.rules, c.server.default_), bindVars)
	if err != nil {
		return nil, err
	}

	if len(ns) == 0 {
		return nil, nil
	}

	n := make(map[int]*Node)
	for k, v := range ns {
		n[k] = c.server.getNode(v)
	}
	return n, nil
}

func (c *Conn) jhs_beginShardConns(conns map[int]*client.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Begin(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) jhs_commitShardConns(conns map[int]*client.SqlConn) error {
	if c.isInTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) jhs_closeShardConns(conns map[int]*client.SqlConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, co := range conns {
		if rollback {
			co.Rollback()
		}

		co.Close()
	}
}

func (c *Conn) jhs_mergeExecResult(rs []*Result) error {
	if len(rs) == 1 {
		r := rs[0]
		if r.InsertId > 0 {
			c.lastInsertId = int64(r.InsertId)
		}

		c.affectedRows = int64(r.AffectedRows)

		if r.Resultset != nil {
			rt := r.Resultset
			status := c.status | r.Status

			if err := c.writeResultset(status, rt); err != nil {
				return err
			}
		}

		r.Status = c.status
		return c.writeOK(r)
	} else {
		return c.mergeExecResult(rs)
	}
}
