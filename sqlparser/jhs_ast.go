package sqlparser

import "fmt"

type JHSStatement interface {
	GetTableName() string
	SetTableName(string, int)
	IStatement()
	SQLNode
}

func (s *Select) GetTableName() string {
	return String(s.From)
}

func (s *Select) SetTableName(name string, index int) {
	s.From = TableExprs{NewJHSTableName(name, index)}
}

func (i *Insert) GetTableName() string {
	return string(i.Table.Name)
}

func (i *Insert) SetTableName(name string, index int) {
	i.Table.Name = []byte(mantage(name, index))
}

func (u *Update) GetTableName() string {
	return string(u.Table.Name)
}

func (u *Update) SetTableName(name string, index int) {
	u.Table.Name = []byte(mantage(name, index))
}

func (d *Delete) GetTableName() string {
	return string(d.Table.Name)
}

func (d *Delete) SetTableName(name string, index int) {
	d.Table.Name = []byte(mantage(name, index))
}

func (r *Replace) GetTableName() string {
	return string(r.Table.Name)
}

func (r *Replace) SetTableName(name string, index int) {
	r.Table.Name = []byte(mantage(name, index))
}

func mantage(name string, index int) string {
	if index == -1 {
		return name
	} else {
		return fmt.Sprintf("%s_%d", name, index)
	}
}

type JHSTableName struct {
	Name string
}

func (*JHSTableName) ITableExpr() {}

func (j *JHSTableName) Format(buf *TrackedBuffer) {
	buf.Fprintf("%s", j.Name)
}

func NewJHSTableName(name string, index int) *JHSTableName {
	j := new(JHSTableName)
	j.Name = mantage(name, index)
	return j
}

type Call struct {
	Procedure *TableName
	Values    ValExprs
}

func (*Call) IStatement() {}

func (c *Call) Format(buf *TrackedBuffer) {
	buf.Fprintf("call %v(%v)", c.Procedure, c.Values)
}

func (c *Call) GetTableName() string {
	return string(c.Procedure.Name)
}

func (c *Call) SetTableName(name string, index int) {}
