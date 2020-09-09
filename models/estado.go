package models

type Estado struct {
	ID     int
	Nombre string
}

// TableName overrides the table name used by Estado to `ordenes.estado`
func (Estado) TableName() string {
	return "ordenes.estado"
}
