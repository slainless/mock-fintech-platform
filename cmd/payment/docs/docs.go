// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json",
        "multipart/form-data",
        "application/x-www-form-urlencoded"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Aiman Fauzy"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/send": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Send amount to account UUID",
                "parameters": [
                    {
                        "description": "Send payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.SendPayload"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payment.SendResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/subscribe": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Subscribe to recurring payment",
                "parameters": [
                    {
                        "description": "Subscribe payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.SubscribePayload"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payment.SubscribeResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/unsubscribe": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Unsubscribe to recurring payment",
                "parameters": [
                    {
                        "description": "Unsubscribe payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.UnsubscribePayload"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payment.UnsubscribeResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/withdraw": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Withdraw amount from account",
                "parameters": [
                    {
                        "description": "Withdraw payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.WithdrawPayload"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authentication token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payment.WithdrawResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "payment.SendPayload": {
            "type": "object",
            "required": [
                "account",
                "amount",
                "dest"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer",
                    "maximum": 999999999999999,
                    "minimum": 1
                },
                "dest": {
                    "type": "string"
                }
            }
        },
        "payment.SendResponse": {
            "type": "object",
            "properties": {
                "transaction": {
                    "$ref": "#/definitions/platform.TransactionHistory"
                }
            }
        },
        "payment.SubscribePayload": {
            "type": "object",
            "required": [
                "account",
                "billing",
                "callback_data",
                "service"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "billing": {
                    "type": "string"
                },
                "callback_data": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                }
            }
        },
        "payment.SubscribeResponse": {
            "type": "object",
            "properties": {
                "payment": {
                    "$ref": "#/definitions/platform.RecurringPayment"
                },
                "transaction": {
                    "$ref": "#/definitions/platform.TransactionHistory"
                }
            }
        },
        "payment.UnsubscribePayload": {
            "type": "object",
            "required": [
                "payment_id"
            ],
            "properties": {
                "payment_id": {
                    "type": "string"
                }
            }
        },
        "payment.UnsubscribeResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "transaction": {
                    "$ref": "#/definitions/platform.TransactionHistory"
                }
            }
        },
        "payment.WithdrawPayload": {
            "type": "object",
            "required": [
                "account_id",
                "amount",
                "callback"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer",
                    "maximum": 999999999999999,
                    "minimum": 1
                },
                "callback": {
                    "type": "string"
                }
            }
        },
        "payment.WithdrawResponse": {
            "type": "object",
            "properties": {
                "transaction": {
                    "$ref": "#/definitions/platform.TransactionHistory"
                }
            }
        },
        "platform.RecurringPayment": {
            "type": "object",
            "properties": {
                "accountUUID": {
                    "type": "string"
                },
                "chargingMethod": {
                    "type": "integer"
                },
                "foreignID": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastCharge": {
                    "type": "string"
                },
                "schedulerType": {
                    "type": "integer"
                },
                "serviceID": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "platform.TransactionHistory": {
            "type": "object",
            "properties": {
                "accountUUID": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "destUUID": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mutation": {
                    "type": "integer"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "transactionDate": {
                    "type": "string"
                },
                "transactionNote": {
                    "type": "string"
                },
                "transactionType": {
                    "type": "integer"
                },
                "userUUID": {
                    "type": "string"
                },
                "uuid": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Payment Manager Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
