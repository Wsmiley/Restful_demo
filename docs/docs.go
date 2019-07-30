// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-29 14:09:12.0179309 +0800 CST m=+0.068005501

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "wsmiley29@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/players/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PlayerdataSerializer"
                        }
                    }
                }
            }
        },
        "/players/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "player id",
                        "name": "playerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PlayerdataSerializer"
                        }
                    }
                }
            }
        },
        "/teams/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TeamRankingSerializer"
                        }
                    }
                }
            }
        },
        "/teams/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "team id",
                        "name": "teamID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TeamRankingSerializer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.PlayerdataSerializer": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "delete_at": {
                    "type": "string"
                },
                "hitrate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "penaltyrate": {
                    "type": "string"
                },
                "penaltyshot": {
                    "type": "string"
                },
                "session": {
                    "type": "string"
                },
                "shot": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "threerate": {
                    "type": "string"
                },
                "threeshot": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "model.TeamRankingSerializer": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "defeat": {
                    "type": "string"
                },
                "delete_at": {
                    "type": "string"
                },
                "diffencewins": {
                    "type": "string"
                },
                "division": {
                    "type": "string"
                },
                "east": {
                    "type": "string"
                },
                "home_field": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lose_point": {
                    "type": "string"
                },
                "net_victory": {
                    "type": "string"
                },
                "point": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                },
                "visiting_field": {
                    "type": "string"
                },
                "win": {
                    "type": "string"
                },
                "winrate": {
                    "type": "string"
                },
                "wl_streak": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:5000",
	BasePath:    "/v1/api",
	Schemes:     []string{},
	Title:       "NBA-data API Demo",
	Description: "This is a server for NBA-data.",
}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
