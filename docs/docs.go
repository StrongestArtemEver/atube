// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.example.com/support",
            "email": "support@example.com"
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
        "/ping": {
            "get": {
                "description": "Проверяет работоспособность сервера",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Ping endpoint",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/videos": {
            "get": {
                "description": "Возвращает список всех доступных видео",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Video"
                ],
                "summary": "Получить список видео",
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Video"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Bookmark": {
            "type": "object",
            "properties": {
                "addedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                },
                "videoID": {
                    "type": "integer"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                },
                "videoID": {
                    "type": "integer"
                }
            }
        },
        "models.Video": {
            "type": "object",
            "properties": {
                "bookmarks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bookmark"
                    }
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "description": "Duration in seconds",
                    "type": "integer"
                },
                "fileURL": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "likesCount": {
                    "type": "integer"
                },
                "thumbnailURL": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "uploadDate": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "viewsCount": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Documentation",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
