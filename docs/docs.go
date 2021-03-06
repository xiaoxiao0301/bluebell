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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "社区列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "社区列表",
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
                "description": "新建社区",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "新建社区",
                "parameters": [
                    {
                        "description": "社区",
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
                "description": "社区详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "社区详情",
                "parameters": [
                    {
                        "type": "string",
                        "default": "3730413906300928",
                        "description": "社区ID",
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
        "/category/{id}/posts": {
            "get": {
                "description": "获取某个社区下的所有帖子",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区"
                ],
                "summary": "获取某个社区下的所有帖子",
                "parameters": [
                    {
                        "type": "string",
                        "default": "3730413906300928",
                        "description": "社区ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostModel"
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
        "/post": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "存储帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "存储帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "帖子",
                        "name": "posts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamPost"
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
        "/post/vote": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户可以给帖子投赞成或者反对票",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子投票",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "投票",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamVote"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostListDetail"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "get": {
                "description": "帖子详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子详情",
                "parameters": [
                    {
                        "type": "string",
                        "default": "3765906580705280",
                        "description": "帖子ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostModel"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "description": "帖子列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页大小",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostListDetail"
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
        },
        "/v2/posts": {
            "get": {
                "description": "可以根据发帖时间和帖子分数来获取帖子列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子"
                ],
                "summary": "帖子列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页大小",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "time",
                        "description": "排序依据, time 时间 score 得分",
                        "name": "order",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "asc",
                        "description": "升序还是降序 asc 升序 desc 降序",
                        "name": "sorts",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostListDetail"
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
                    "description": "社区ID",
                    "type": "string",
                    "example": "0"
                },
                "category_name": {
                    "description": "社区名称",
                    "type": "string"
                },
                "created_time": {
                    "description": "社区创建时间",
                    "$ref": "#/definitions/models.LocalTime"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "introduction": {
                    "description": "社区简介",
                    "type": "string"
                },
                "updated_time": {
                    "description": "社区更新时间",
                    "$ref": "#/definitions/models.LocalTime"
                }
            }
        },
        "models.CategoryRow": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "社区ID",
                    "type": "integer"
                },
                "name": {
                    "description": "社区名称",
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
                    "description": "社区名称",
                    "type": "string"
                },
                "introduction": {
                    "description": "社区简介",
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
                    "description": "密码",
                    "type": "string",
                    "example": "12456"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "jack"
                }
            }
        },
        "models.ParamPost": {
            "type": "object",
            "required": [
                "category_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "发帖作者",
                    "type": "integer"
                },
                "category_id": {
                    "description": "社区ID",
                    "type": "string",
                    "example": "0"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "post_id": {
                    "description": "帖子ID",
                    "type": "integer"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
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
                    "description": "令牌",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "刷新令牌",
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
                    "description": "密码",
                    "type": "string"
                },
                "re_password": {
                    "description": "确认密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "models.ParamVote": {
            "type": "object",
            "required": [
                "post_id",
                "value"
            ],
            "properties": {
                "post_id": {
                    "description": "投票的帖子id",
                    "type": "string",
                    "example": "0"
                },
                "value": {
                    "description": "投票结果， 1 赞成 0 取消 -1 反对",
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "models.PostListDetail": {
            "type": "object",
            "properties": {
                "author_id": {
                    "description": "帖子作者ID",
                    "type": "string",
                    "example": "0"
                },
                "category_id": {
                    "description": "社区ID",
                    "type": "string",
                    "example": "0"
                },
                "category_name": {
                    "description": "社区名称",
                    "type": "string"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "created_time": {
                    "description": "社区创建时间",
                    "$ref": "#/definitions/models.LocalTime"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "introduction": {
                    "description": "社区简介",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "post_id": {
                    "description": "帖子ID",
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                },
                "updated_time": {
                    "description": "社区更新时间",
                    "$ref": "#/definitions/models.LocalTime"
                },
                "user_id": {
                    "type": "string",
                    "example": "0"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.PostModel": {
            "type": "object",
            "properties": {
                "author_id": {
                    "description": "帖子作者ID",
                    "type": "string",
                    "example": "0"
                },
                "category_id": {
                    "description": "帖子社区ID",
                    "type": "string",
                    "example": "0"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "created_time": {
                    "description": "帖子创建时间",
                    "$ref": "#/definitions/models.LocalTime"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "post_id": {
                    "description": "帖子ID",
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                },
                "updated_time": {
                    "description": "帖子更新时间",
                    "$ref": "#/definitions/models.LocalTime"
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
