package models

type Solicitud struct {
	ID            int
	Descripcion   string
	Fecha         string
	Ticket        string
	Calificacion  Calificacion  `json:"calificacion" gorm:"foreignKey:ID"`
	TipoSolicitud TipoSolicitud `json:"tipoSolicitud" gorm:"foreignKey:ID"`
	Cliente       Cliente       `json:"cliente" gorm:"foreignKey:ID"`
	Estado        Estado        `json:"estado" gorm:"foreignKey:ID"`
	Servicio      Servicio      `json:"servicio" gorm:"foreignKey:ID"`
}

// TableName sobrescribe la tabla por defecto de Solicitud(s)
func (Solicitud) TableName() string {
	return "ordenes.solicitud"
}
