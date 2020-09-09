package models

type Servicio struct {
	ID          int
	Nombre      string
	Descripcion string
}

// TableName sobrescribe la tabla por defecto de Servicio(s)
func (Servicio) TableName() string {
	return "ordenes.servicio"
}
