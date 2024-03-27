// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/odontologos/{id}": {
            "get": {
                "description": "Obtiene un odontólogo por su id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Obtener odontólogo por id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del odontólogo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    },
                    "400": {
                        "description": "ID inválido",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Odontólogo no encontrado",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Odontologo": {
            "type": "object",
            "required": [
                "apellido",
                "matricula",
                "nombre"
            ],
            "properties": {
                "apellido": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "matricula": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "turnos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Turno"
                    }
                }
            }
        },
        "domain.Turno": {
            "type": "object",
            "required": [
                "descripcion",
                "fechaHora",
                "odontologoId",
                "pacienteId"
            ],
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "fechaHora": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "odontologoId": {
                    "type": "integer"
                },
                "pacienteId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Desafío Go Backend 3",
	Description:      "Esta API maneja Odontólogos, Pacientes y Turnos",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}