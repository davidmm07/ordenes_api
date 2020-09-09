package types

import "github.com/graphql-go/graphql"

var ClientType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Cliente",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla cliente",
		},
		"usuario": &graphql.Field{
			Type:        graphql.String,
			Description: "usuario del cliente",
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: "password del cliente",
		},
		"token": &graphql.Field{
			Type:        graphql.String,
			Description: "codigo generado para ingreso",
		},
	},
})

var TechnicianType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Tecnico",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla tecnico",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "nombre del tecnico",
		},
	},
})

var KindOfRequestType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "TipoSolicitud",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla tipo de solicitud",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "nombre del tipo de solicitud",
		},
	},
})

var StatusType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Estado",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla estado",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "nombre del estado de la solicitud",
		},
	},
})
var ServiceType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Servicio",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla servicio",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "nombre del servicio",
		},
		"descripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "descripcion del servicio",
		},
	},
})
var ScoreType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Calificacion",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla calificacion",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "nombre de la clasificacion",
		},
	},
})

var RequestType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Solicitud",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla solicitud",
		},
		"descripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "información detallada de la solicitud",
		},
		"fecha": &graphql.Field{
			Type:        graphql.String,
			Description: "fecha de la solicitud",
		},
		"ticket": &graphql.Field{
			Type:        graphql.String,
			Description: "codigo generado",
		},
		"calificacion": &graphql.Field{
			Type:        ScoreType,
			Description: "calificación del servicio",
		},
		"tipoSolicitud": &graphql.Field{
			Type: KindOfRequestType,
		},
		"cliente": &graphql.Field{
			Type: ClientType,
		},
		"estado": &graphql.Field{
			Type: StatusType,
		},
		"servicio": &graphql.Field{
			Type: ServiceType,
		},
	},
})

var RequestbyTechnicianType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "solicitud_tecnico",
	Description: "",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "identificador de la tabla solicitud_tecnico",
		},
		"solicitud": &graphql.Field{
			Type: RequestType,
		},
		"tecnico": &graphql.Field{
			Type: TechnicianType,
		},
	},
})
