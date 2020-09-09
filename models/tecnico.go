package models

type Tecnico struct {
	ID     int
	Nombre string
}

// TableName overrides the table name used by Tecnico to `ordenes.tecnico`
func (Tecnico) TableName() string {
	return "ordenes.tecnico"
}
