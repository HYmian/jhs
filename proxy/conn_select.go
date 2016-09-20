package proxy

import (
	"bytes"
	"fmt"
	. "github.com/HYmian/jhs/mysql"
	"github.com/HYmian/jhs/sqlparser"
	"strings"
)

func (c *Conn) handleSimpleSelect(sql string, stmt *sqlparser.SimpleSelect) error {
	if len(stmt.SelectExprs) != 1 {
		return fmt.Errorf("support select one informaction function, %s", sql)
	}

	expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	if !ok {
		return fmt.Errorf("support select informaction function, %s", sql)
	}

	var r *Resultset
	var err error
	switch v := expr.Expr.(type) {
	case *sqlparser.FuncExpr:
		switch strings.ToLower(string(v.Name)) {
		case "last_insert_id":
			r, err = c.buildSimpleSelectResult(c.lastInsertId, v.Name, expr.As)
		case "row_count":
			r, err = c.buildSimpleSelectResult(c.affectedRows, v.Name, expr.As)
		case "version":
			r, err = c.buildSimpleSelectResult(ServerVersion, v.Name, expr.As)
		case "connection_id":
			r, err = c.buildSimpleSelectResult(c.connectionId, v.Name, expr.As)
		case "database":
			r, err = c.buildSimpleSelectResult(c.server.database, v.Name, expr.As)
		default:
			return fmt.Errorf("function %s not support", v.Name)
		}
	case *sqlparser.ColName:
		name := strings.TrimLeft(strings.ToLower(string(v.Name)), "@")
		if value, ok := c.server.GetMySqlVariable(name); ok {
			r, err = c.buildSimpleSelectResult(value, v.Name, expr.As)
		} else {
			fmt.Println(c.server.mysql_config)
			return fmt.Errorf("variable %s not support", v.Name)
		}
	}

	if err != nil {
		return err
	}

	return c.writeResultset(c.status, r)
}

func (c *Conn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {
	field := &Field{}

	field.Name = name

	if asName != nil {
		field.Name = asName
	}

	field.OrgName = name

	formatField(field, value)

	r := &Resultset{Fields: []*Field{field}}
	row, err := formatValue(value)
	if err != nil {
		return nil, err
	}
	r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))

	return r, nil
}

func (c *Conn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if c.db == "" {
		return NewDefaultError(ER_NO_DB_ERROR)
	}

	//nodeName := c.schema.rule.GetRule(table).Nodes[0]
	nodeName := c.server.rules[table].Nodes[0]

	n := c.server.getNode(nodeName)

	co, err := n.getMasterConn()
	if err != nil {
		return err
	}
	defer co.Close()

	//if err = co.UseDB(c.schema.db); err != nil {
	//	return err
	//}

	if fs, err := co.FieldList(table, wildcard); err != nil {
		return err
	} else {
		return c.writeFieldList(c.status, fs)
	}
}

func (c *Conn) writeFieldList(status uint16, fs []*Field) error {
	c.affectedRows = int64(-1)

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(status); err != nil {
		return err
	}
	return nil
}
