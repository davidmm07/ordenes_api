package schema

import (
	"errors"
	"strconv"

	"github.com/davidmm07/ordenes_api/auth"
	"github.com/davidmm07/ordenes_api/models"
	"github.com/davidmm07/ordenes_api/resolvers"
	"github.com/davidmm07/ordenes_api/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

type Ordenes struct {
	Ordenes graphql.Schema
}

func (o *Ordenes) InitSchema() graphql.Schema {

	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"connection": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "successful", nil
				},
			},
			"clientes": &graphql.Field{
				Type: graphql.NewList(types.ClientType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolvers.GetClients(), nil
				},
			},
			"solicitudes_tecnico": &graphql.Field{
				Type: graphql.NewList(types.RequestbyTechnicianType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "identificador tecnico",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"fecha": &graphql.ArgumentConfig{
						Description: "fecha con las solicitudes para el tecnico",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return resolvers.GetRequestByTechnician(id, p.Args["fecha"]), nil
				},
			},
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"loginCliente": &graphql.Field{
				Type:        graphql.String,
				Description: "crea un nuevo JWT token mediante la validación del usuario ",
				Args: graphql.FieldConfigArgument{
					"usuario": &graphql.ArgumentConfig{
						Description: "username",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Description: "password",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					username, _ := params.Args["usuario"].(string)
					password, _ := params.Args["password"].(string)

					err := resolvers.ValidarDatosCliente(username, password)

					if err != nil {
						return nil, gqlerrors.FormatError(errors.New("Datos invalidos"))
					}
					token, err := auth.CrearToken(username, password)
					if err != nil {
						return nil, err
					}

					return token, nil

				},
			},
			"crearSolicitud": &graphql.Field{
				Type:        types.RequestType,
				Description: "crear nueva solicitud",
				Args: graphql.FieldConfigArgument{
					"descripcion": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"fecha": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ticket": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"calificacion": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"tipoSolicitud": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"cliente": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"estado": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"servicio": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// validando token

					isValid, err := auth.ValidarToken(params.Context.Value("token").(string))
					if err != nil {
						return nil, err
					}

					if !isValid {
						return nil, gqlerrors.FormatError(errors.New("token invalido"))
					}
					descripcion, _ := params.Args["descripcion"].(string)
					fecha, _ := params.Args["fecha"].(string)
					calificacion, _ := params.Args["calificacion"].(int)
					tipoSolicitud, _ := params.Args["tipoSolicitud"].(int)
					cliente, _ := params.Args["cliente"].(int)
					estado, _ := params.Args["estado"].(int)
					servicio, _ := params.Args["servicio"].(int)
					solicitud := models.Solicitud{
						Descripcion: descripcion,
						Fecha:       fecha,
						Calificacion: models.Calificacion{
							ID: calificacion,
						},
						TipoSolicitud: models.TipoSolicitud{
							ID: tipoSolicitud,
						},
						Cliente: models.Cliente{
							ID: cliente,
						},
						Estado: models.Estado{
							ID: estado,
						},
						Servicio: models.Servicio{
							ID: servicio,
						},
					}
					return resolvers.CrearSolicitud(solicitud), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		panic(err)
	}
	o.Ordenes = schema
	return o.Ordenes
}
