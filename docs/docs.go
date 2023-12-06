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
        "/api/user/ChangePassWord": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "description": "账号, 密码, 确认密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ChangePassWord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户修改密码,返回包括修改成功，失败",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/GetUserInfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "获取用户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/GetUserList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取系统用户列表",
                "parameters": [
                    {
                        "description": "页码, 每页大小",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "分页获取用户列表,返回包括列表,总数,页码,每页数量",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.PageResult"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/captcha": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.SysCaptchaResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回包括用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册账号",
                "parameters": [
                    {
                        "description": "用户名, 昵称, 密码, 角色ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户注册账号,返回包括用户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.SysUser": {
            "type": "object",
            "properties": {
                "activeColor": {
                    "description": "活跃颜色",
                    "type": "string"
                },
                "baseColor": {
                    "description": "基础颜色",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "enable": {
                    "description": "用户是否被冻结 1正常 2冻结",
                    "type": "integer"
                },
                "headerImg": {
                    "description": "用户头像",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "passWord": {
                    "description": "用户登录密码",
                    "type": "string"
                },
                "phone": {
                    "description": "用户手机号",
                    "type": "string"
                },
                "sideMode": {
                    "description": "用户侧边主题",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userName": {
                    "description": "用户登录名",
                    "type": "string"
                },
                "uuid": {
                    "description": "用户UUID",
                    "type": "string"
                }
            }
        },
        "request.ChangePassWord": {
            "type": "object",
            "properties": {
                "confirmPassword": {
                    "type": "string",
                    "example": "确认密码"
                },
                "passWord": {
                    "type": "string",
                    "example": "密码"
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "properties": {
                "keyword": {
                    "description": "关键字",
                    "type": "string"
                },
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页大小",
                    "type": "integer"
                }
            }
        },
        "request.UserLogin": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "验证码",
                    "type": "string"
                },
                "captchaId": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "电子邮箱"
                },
                "enable": {
                    "type": "string",
                    "example": "int 是否启用"
                },
                "headerImg": {
                    "type": "string",
                    "example": "头像链接"
                },
                "nickName": {
                    "type": "string",
                    "example": "昵称"
                },
                "passWord": {
                    "type": "string",
                    "example": "密码"
                },
                "phone": {
                    "type": "string",
                    "example": "电话号码"
                },
                "userName": {
                    "type": "string",
                    "example": "用户名"
                }
            }
        },
        "response.PageResult": {
            "type": "object",
            "properties": {
                "list": {},
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.SysCaptchaResponse": {
            "type": "object",
            "properties": {
                "captchaId": {
                    "type": "string"
                },
                "captchaLength": {
                    "type": "integer"
                },
                "openCaptcha": {
                    "type": "boolean"
                },
                "picPath": {
                    "type": "string"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/model.SysUser"
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
