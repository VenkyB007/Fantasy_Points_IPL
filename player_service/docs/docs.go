// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/player": {
            "post": {
                "description": "Create a new  Player with the input payload",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Create a new  Player",
                "parameters": [
                    {
                        "description": "Player Details",
                        "name": "Player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Posted!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/player/{id}": {
            "delete": {
                "description": "Delete details of Player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Delete details of Player based on player id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Deleted!!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/player/{id}/score": {
            "post": {
                "description": "Post player scores with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scores"
                ],
                "summary": "Post scores",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Player Scores",
                        "name": "Score",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Stats"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Posted!!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Scores of Player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scores"
                ],
                "summary": "Delete Scores of Player based on player id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully Deleted!!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/players": {
            "get": {
                "description": "Get details of all Players",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get details of all Players from the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    }
                }
            }
        },
        "/players/scores": {
            "get": {
                "description": "Get Scores of all Players",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scores"
                ],
                "summary": "Get Scores of all Players from the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Score"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Player": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                }
            }
        },
        "models.Score": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "match": {
                    "type": "integer"
                },
                "runs": {
                    "type": "integer"
                },
                "wickets": {
                    "type": "integer"
                }
            }
        },
        "models.Stats": {
            "type": "object",
            "properties": {
                "match": {
                    "type": "integer"
                },
                "runs": {
                    "type": "integer"
                },
                "wickets": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8989",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Player Service API",
	Description:      "In this Service, You can Post, Get, Delete Player details",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
