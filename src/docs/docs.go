// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-07-13 12:46:54.035043001 +0800 CST m=+0.082559848

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
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
        "version": "1.0"
    },
    "host": "petstore.swagger.io",
    "basePath": "/v2",
    "paths": {
        "/api/v1/login": {
            "post": {
                "description": "login by name and pw",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user login"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "your unique name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your pw",
                        "name": "bcrypt_pw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.DataRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    }
                }
            }
        },
        "/api/v1/signup": {
            "post": {
                "description": "register your account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user register"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "your unique name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your pw",
                        "name": "bcrypt_pw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your nickname",
                        "name": "nickname",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your phone_number",
                        "name": "phone_number",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "your group",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.DataRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/util.ErrorRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "util.DataRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "util.ErrorRes": {
            "type": "object",
            "properties": {
                "errMsg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}