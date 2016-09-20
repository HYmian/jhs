package proxy

import (
	"fmt"
	"github.com/HYmian/jhs/hack"
	. "github.com/HYmian/jhs/mysql"
	"github.com/HYmian/jhs/sqlparser"
	"sort"
	"strings"
	"time"

	"github.com/siddontang/go-log/log"
)

func (c *Conn) handleShow(sql string, stmt *sqlparser.Show) error {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Section) {
	case "databases":
		r, err = c.handleShowDatabases()
	case "tables":
		r, err = c.handleShowTables(sql, stmt)
	case "proxy":
		r, err = c.handleShowProxy(sql, stmt)
	case "variables":
		r, err = c.handleShowVariables(stmt)
	default:
		err = fmt.Errorf("unsupport show %s now", sql)
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func (c *Conn) handleShowDatabases() (*Resultset, error) {
	dbs := make([]interface{}, 1)
	dbs[0] = c.server.database

	return c.buildSimpleShowResultset(dbs, "Database")
}

func (c *Conn) handleShowTables(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	if c.db == "" {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	s := c.server

	var tables []string
	for _, v := range s.rules {
		tables = append(tables, v.Table)
	}

	sort.Strings(tables)

	values := make([]interface{}, len(tables))
	for i := range tables {
		values[i] = tables[i]
	}

	return c.buildSimpleShowResultset(values, fmt.Sprintf("Tables_in_%s", s.database))
}

func (c *Conn) handleShowVariables(stmt *sqlparser.Show) (*Resultset, error) {
	s := c.server
	names := []string{"Variable_name", "Value"}
	const column = 2

	list := sqlparser.GetVariablesOpt(stmt.LikeOrWhere)
	rows := make([][]interface{}, 0, len(list))
	for _, name := range list {
		value, ok := s.GetMySqlVariable(name)
		if !ok {
			continue
		}
		var row []interface{} = make([]interface{}, 0, column)
		row = append(row, name)
		row = append(row, value)
		rows = append(rows, row)
	}
	return c.buildResultset(names, rows)
}

func (c *Conn) handleShowProxy(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	var err error
	var r *Resultset
	switch strings.ToLower(stmt.Key) {
	case "config":
		r, err = c.handleShowProxyConfig()
	case "status":
		r, err = c.handleShowProxyStatus(sql, stmt)
	default:
		err = fmt.Errorf("Unsupport show proxy [%v] yet, just support [config|status] now.", stmt.Key)
		log.Warn(err.Error())
		return nil, err
	}
	return r, err
}

func (c *Conn) handleShowProxyConfig() (*Resultset, error) {
	var names []string = []string{"Section", "Key", "Value"}
	var rows [][]string
	const (
		Column = 3
	)

	rows = append(rows, []string{"Global_Config", "Addr", c.server.cfg.Addr})
	rows = append(rows, []string{"Global_Config", "User", c.server.cfg.User})
	rows = append(rows, []string{"Global_Config", "Password", c.server.cfg.Password})
	rows = append(rows, []string{"Global_Config", "LogLevel", c.server.cfg.LogLevel})
	//rows = append(rows, []string{"Global_Config", "Schemas_Count", fmt.Sprintf("%d", len(c.server.schemas))})
	rows = append(rows, []string{"Global_Config", "Nodes_Count", fmt.Sprintf("%d", len(c.server.nodes))})

	rows = append(rows, []string{"Schemas", "DB", c.server.database})

	var nodeNames []string
	var nodeRows [][]string
	for name, node := range c.server.nodes {
		nodeNames = append(nodeNames, name)
		var nodeSection = fmt.Sprintf("Schemas[%s]-Node[ %v ]", c.server.database, name)

		if node.master != nil {
			nodeRows = append(nodeRows, []string{nodeSection, "Master", node.master.String()})
		}

		if node.slave != nil {
			nodeRows = append(nodeRows, []string{nodeSection, "Slave", node.slave.String()})
		}
		nodeRows = append(nodeRows, []string{nodeSection, "Last_Master_Ping", fmt.Sprintf("%v", time.Unix(node.lastMasterPing, 0))})

		nodeRows = append(nodeRows, []string{nodeSection, "Last_Slave_Ping", fmt.Sprintf("%v", time.Unix(node.lastSlavePing, 0))})

		nodeRows = append(nodeRows, []string{nodeSection, "down_after_noalive", fmt.Sprintf("%v", node.downAfterNoAlive)})

	}
	rows = append(rows, []string{fmt.Sprintf("Schemas[%s]", c.server.database), "Nodes_List", strings.Join(nodeNames, ",")})

	//var defaultRule = schema.rule.DefaultRule
	//if defaultRule.DB == db {
	//	if defaultRule.DB == db {
	//		rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Default", db),
	//			"Default_Table", defaultRule.String()})
	//	}
	//}
	//for tb, r := range schema.rule.Rules {
	//	if r.DB == db {
	//		rows = append(rows, []string{fmt.Sprintf("Schemas[%s]_Rule_Table", db),
	//			fmt.Sprintf("Table[ %s ]", tb), r.String()})
	//	}
	//}

	rows = append(rows, nodeRows...)

	var values [][]interface{} = make([][]interface{}, len(rows))
	for i := range rows {
		values[i] = make([]interface{}, Column)
		for j := range rows[i] {
			values[i][j] = rows[i][j]
		}
	}

	return c.buildResultset(names, values)
}

func (c *Conn) handleShowProxyStatus(sql string, stmt *sqlparser.Show) (*Resultset, error) {
	// TODO: handle like_or_where expr
	return nil, nil
}

func (c *Conn) buildSimpleShowResultset(values []interface{}, name string) (*Resultset, error) {

	r := new(Resultset)

	field := &Field{}

	field.Name = hack.Slice(name)
	field.Charset = 33
	field.Type = MYSQL_TYPE_VAR_STRING

	r.Fields = []*Field{field}

	var row []byte
	var err error

	for _, value := range values {
		row, err = formatValue(value)
		if err != nil {
			return nil, err
		}
		r.RowDatas = append(r.RowDatas,
			PutLengthEncodedString(row))
	}

	return r, nil
}
