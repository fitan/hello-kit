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
        "/foo/{id}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "HelloService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hello.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
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
        "/say/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "HelloService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hello.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.SayRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/say1/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "HelloService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hello.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.SayRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sayhello/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "HelloService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hello.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.SayRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sayhello1/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "HelloService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hello.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/ent.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "UserService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "HostIPEQ",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "NamespaceEQ",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "ServiceNameEQ",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "name": "ServiceNameIn",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "_limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "_page",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "acs",
                            "desc"
                        ],
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sortBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/user.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user.GetListRes"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "UserService"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/user.SwagResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/ent.User"
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
        "ent.Pod": {
            "type": "object",
            "properties": {
                "cluster_name": {
                    "description": "ClusterName holds the value of the \"cluster_name\" field.\n集群标识",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "detail": {
                    "description": "Detail holds the value of the \"detail\" field.\n详细内容",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PodQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.PodEdges"
                },
                "host_ip": {
                    "description": "HostIP holds the value of the \"host_ip\" field.\n主机IP",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "message": {
                    "description": "Message holds the value of the \"message\" field.\n消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "Namespace holds the value of the \"namespace\" field.\n所属项目空间",
                    "type": "string"
                },
                "phase": {
                    "description": "Phase holds the value of the \"phase\" field.\n阶段",
                    "type": "string"
                },
                "pod_ip": {
                    "description": "PodIP holds the value of the \"pod_ip\" field.\nPodIP",
                    "type": "string"
                },
                "pod_name": {
                    "description": "PodName holds the value of the \"pod_name\" field.\npod名称",
                    "type": "string"
                },
                "reason": {
                    "description": "Reason holds the value of the \"reason\" field.\n原因",
                    "type": "string"
                },
                "resource_version": {
                    "description": "ResourceVersion holds the value of the \"resource_version\" field.\n资源版本号",
                    "type": "string"
                },
                "service_name": {
                    "description": "ServiceName holds the value of the \"service_name\" field.\n应用名称",
                    "type": "string"
                },
                "start_time": {
                    "description": "StartTime holds the value of the \"start_time\" field.\n启动时间",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.PodEdges": {
            "type": "object",
            "properties": {
                "servicetree": {
                    "description": "Servicetree holds the value of the servicetree edge.",
                    "$ref": "#/definitions/ent.SpiderDevTblServicetree"
                }
            }
        },
        "ent.SpiderDevTblServicetree": {
            "type": "object",
            "properties": {
                "aname": {
                    "description": "Aname holds the value of the \"aname\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "key": {
                    "description": "Key holds the value of the \"key\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "origin": {
                    "description": "Origin holds the value of the \"origin\" field.",
                    "type": "string"
                },
                "pnode": {
                    "description": "Pnode holds the value of the \"pnode\" field.",
                    "type": "integer"
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "Age holds the value of the \"age\" field.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "hello.SwagResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "traceId": {
                    "type": "string"
                }
            }
        },
        "types.SayRes": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.GetListRes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Pod"
                    }
                }
            }
        },
        "user.SwagResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "traceId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
