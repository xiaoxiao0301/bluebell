// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.jixiaoxiao.com",
        "contact": {
            "name": "Xiao",
            "url": "https://www.jixiaoxiao.com",
            "email": "simplexiaoxiao@gmail.com"
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
        "/categories": {
            "get": {
                "description": "分类列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "分类列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CategoryRow"
                        }
                    }
                }
            }
        },
        "/category": {
            "post": {
                "description": "新建分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "新建分类",
                "parameters": [
                    {
                        "description": "社区信息",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommon"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "get": {
                "description": "分类详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "分类详情",
                "parameters": [
                    {
                        "type": "string",
                        "default": "3730413906300928",
                        "description": "分类ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CategoryModel"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登陆",
                "parameters": [
                    {
                        "description": "登陆信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseLoginSuccess"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "刷新access_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token"
                ],
                "summary": "刷新token",
                "parameters": [
                    {
                        "description": "jwt验证信息",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamRefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseRefreshToken"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "注册信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommon"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller._ResponseCommon": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "string"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseLoginSuccess": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "jwt access_token 验证使用",
                    "type": "string"
                },
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "jwt refresh_token 刷新token",
                    "type": "string"
                },
                "user_id": {
                    "description": "登陆用户ID",
                    "type": "string"
                },
                "username": {
                    "description": "登陆用户昵称",
                    "type": "string"
                }
            }
        },
        "controller._ResponseRefreshToken": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "jwt access_token 验证使用",
                    "type": "string"
                },
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "jwt refresh_token 刷新token",
                    "type": "string"
                }
            }
        },
        "models.CategoryModel": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string",
                    "example": "0"
                },
                "category_name": {
                    "type": "string"
                },
                "created_time": {
                    "$ref": "#/definitions/models.LocalTime"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "updated_time": {
                    "$ref": "#/definitions/models.LocalTime"
                }
            }
        },
        "models.CategoryRow": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.LocalTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "models.ParamCategory": {
            "type": "object",
            "required": [
                "category_name",
                "introduction"
            ],
            "properties": {
                "category_name": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                }
            }
        },
        "models.ParamLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ParamRefreshToken": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.ParamSignUp": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "re_password": {
                    "type": "string"
                },
                "username": {
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
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "BlueBell Api",
	Description: "使用gin开发简单帖子展示系统",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
