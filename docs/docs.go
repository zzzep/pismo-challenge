// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Giuseppe Fechio",
            "url": "http://github.com/zzzep/pismo-challenge",
            "email": "giuseppe.fechio@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new Account",
                "parameters": [
                    {
                        "description": "Document Number",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domains.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domains.Account"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "document_number": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/accounts/{accountId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountId",
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
                                    "$ref": "#/definitions/domains.Account"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "account_id": {
                                            "type": "integer"
                                        },
                                        "document_number": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "object",
                                "properties": {
                                    "message": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/accounts/{accountId}/transactions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List Transaction by Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountId",
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
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/domains.Transaction"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "account_id": {
                                                "type": "integer"
                                            },
                                            "document_number": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "object",
                                "properties": {
                                    "message": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new Transaction",
                "parameters": [
                    {
                        "description": "Transaction",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domains.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.Transaction"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.Account": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "document_number": {
                    "type": "string"
                }
            }
        },
        "domains.OperationType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "Purchase",
                "Installment",
                "Withdrawal",
                "Payment"
            ]
        },
        "domains.Transaction": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "operation_type_id": {
                    "$ref": "#/definitions/domains.OperationType"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pismo Challenge Giuseppe",
	Description:      "This is a Challenge made by Giuseppe to Pismo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
