package models

type SolicitudTecnico struct {
	ID        int
	Solicitud Solicitud `json:"solicitud" gorm:"foreignKey:ID"`
	Tecnico   Tecnico   `json:"tecnico" gorm:"foreignKey:ID"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by SolicitudTecnico to `ordenes.solicitud_tecnico`
func (SolicitudTecnico) TableName() string {
	return "ordenes.solicitud_tecnico"
}
