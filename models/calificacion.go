package models

type Calificacion struct {
	ID     int
	Nombre string
}

// TableName overrides the table name used by Calificacion to `ordenes.calificacion`
func (Calificacion) TableName() string {
	return "ordenes.calificacion"
}
