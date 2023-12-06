// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/clans": {
            "get": {
                "description": "Get List of Clan",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Get All Clan",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Clan"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Clan",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Create Clan",
                "parameters": [
                    {
                        "description": "the body to create new clan",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ClanInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Clan"
                        }
                    }
                }
            }
        },
        "/clans/{id}": {
            "get": {
                "description": "Get one Clan by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Get a Clan by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Clan Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Clan"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Clan by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Delete a Clan by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Clan Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Clan by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Update Clan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Clan Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to Update new clan",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ClanInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Clan"
                        }
                    }
                }
            }
        },
        "/clans/{id}/shinobies": {
            "get": {
                "description": "Get all Shinobies by Clan by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clan"
                ],
                "summary": "Get Shinobies by Clan by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Clan Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Shinobi"
                            }
                        }
                    }
                }
            }
        },
        "/jutsus": {
            "get": {
                "description": "Get List of Jutsu",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jutsu"
                ],
                "summary": "Get All Jutsu",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Jutsu"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Jutsu",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jutsu"
                ],
                "summary": "Create Jutsu",
                "parameters": [
                    {
                        "description": "the body to create new Jutsu",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.JutsuInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Jutsu"
                        }
                    }
                }
            }
        },
        "/jutsus/{id}": {
            "get": {
                "description": "Get one Jutsu by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jutsu"
                ],
                "summary": "Get a Jutsu by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Jutsu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Jutsu"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Jutsu by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jutsu"
                ],
                "summary": "Delete a Jutsu by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Jutsu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Jutsu by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jutsu"
                ],
                "summary": "Update Jutsu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Jutsu Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to Update new Jutsu",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.JutsuInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Jutsu"
                        }
                    }
                }
            }
        },
        "/nature-types": {
            "get": {
                "description": "Get List of NatureType",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NatureType"
                ],
                "summary": "Get All NatureType",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.NatureType"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new NatureType",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NatureType"
                ],
                "summary": "Create NatureType",
                "parameters": [
                    {
                        "description": "the body to create new NatureType",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.NatureTypeInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.NatureType"
                        }
                    }
                }
            }
        },
        "/nature-types/{id}": {
            "get": {
                "description": "Get one NatureType by Id",
                "tags": [
                    "NatureType"
                ],
                "summary": "Get a NatureType by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NatureType Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "2": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.NatureType"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one NatureType by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NatureType"
                ],
                "summary": "Delete a NatureType by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NatureType Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update NatureType by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NatureType"
                ],
                "summary": "Update NatureType",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NatureType Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to Update new NatureType",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.NatureTypeInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.NatureType"
                        }
                    }
                }
            }
        },
        "/nature-types/{id}/jutsus": {
            "get": {
                "description": "Get all Jutsus by NatureType by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NatureType"
                ],
                "summary": "Get Jutsus by NatureType by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NatureType Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Jutsu"
                            }
                        }
                    }
                }
            }
        },
        "/shinobies": {
            "get": {
                "description": "Get List of Shinobi",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Get All Shinobi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Shinobi"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new Shinobi",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Create Shinobi",
                "parameters": [
                    {
                        "description": "the body to create new Shinobi",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ShinobiInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Shinobi"
                        }
                    }
                }
            }
        },
        "/shinobies/{id}": {
            "get": {
                "description": "Get one Shinobi by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Get a Shinobi by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shinobi Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Shinobi"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Shinobi by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Delete a Shinobi by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shinobi Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Shinobi by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Update Shinobi",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shinobi Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to Update new Shinobi",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ShinobiInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Shinobi"
                        }
                    }
                }
            }
        },
        "/shinobies/{id}/nature-types": {
            "get": {
                "description": "Get all NatureTypes by Shinobi by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shinobi"
                ],
                "summary": "Get NatureTypes by Shinobi by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shinobi Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.NatureType"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ClanInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.JutsuInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "jutsu_name": {
                    "type": "string"
                },
                "nature_type_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.NatureTypeInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "shinobi_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.ShinobiInput": {
            "type": "object",
            "properties": {
                "clan_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Clan": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Jutsu": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nature_type_id": {
                    "description": "put the NatureTypeID so it can GET Jutsu by NatureType ID",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.NatureType": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "shinobi_id": {
                    "description": "put the ShinobiID so it can GET NatureType by ShinobiID\nShinobiID will called by shinobiController",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Shinobi": {
            "type": "object",
            "properties": {
                "clan_id": {
                    "description": "put the ClanID so it can GET Shinobi by Clan ID",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
