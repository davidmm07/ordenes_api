package models

type TipoSolicitud struct {
	ID     int
	Nombre string
}

// TableName overrides the table name used by TipoSolicitud to `ordenes.tipo_solicitud`
func (TipoSolicitud) TableName() string {
	return "ordenes.tipo_solicitud"
}
