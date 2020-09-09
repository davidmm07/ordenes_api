package models

type Cliente struct {
	ID       int
	Usuario  string
	Password string
}

// TableName overrides the table name used by Cliente to `ordenes.cliente`
func (Cliente) TableName() string {
	return "ordenes.cliente"
}
