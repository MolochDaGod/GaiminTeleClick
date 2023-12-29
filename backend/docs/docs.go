// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ButtonMania Team",
            "email": "team@buttonmania.win"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/stats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Handles API requests for getting room stats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "clientId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Button Type",
                        "name": "buttonType",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.GameRoomStats"
                        }
                    }
                }
            }
        },
        "/ws": {
            "get": {
                "summary": "Handles WebSocket connections",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client ID",
                        "name": "clientId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Button Type",
                        "name": "buttonType",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Telegram init data",
                        "name": "initData",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "protocol.GameRoomStats": {
            "type": "object",
            "properties": {
                "bestDuration": {
                    "type": "integer"
                },
                "countActive": {
                    "type": "integer"
                },
                "countLeaderboard": {
                    "type": "integer"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "buttonmania.win",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ButtonMania API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}