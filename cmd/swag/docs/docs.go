// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "github.com/anonystick/go-ecommerce-backend-go",
        "contact": {
            "name": "TEAM TIPSGO",
            "url": "github.com/anonystick/go-ecommerce-backend-go",
            "email": "tipsgo@gmail.com"
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
        "/auth/register": {
            "post": {
                "description": "When user is registered send otp to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth management"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authcommandrequest.UserRegistratorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/auth/save-account": {
            "post": {
                "description": "When user has registered send otp to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth management"
                ],
                "summary": "User Base Registration",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authcommandrequest.SaveAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/auth/verify-account": {
            "post": {
                "description": "When user is verified otp from email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth management"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authcommandrequest.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/transportation/search-trips": {
            "get": {
                "description": "Get list trips",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transportation"
                ],
                "summary": "Get list trips",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Departure date",
                        "name": "departure_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "From location",
                        "name": "from_location",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "To location",
                        "name": "to_location",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number (default 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size (default 10)",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transportationqueryresponse.GetListTripsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Returns user details data based on ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user details by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authcommandrequest.SaveAccountRequest": {
            "type": "object",
            "required": [
                "confirm_pass",
                "email",
                "name",
                "password",
                "token"
            ],
            "properties": {
                "birthday": {
                    "description": "birthday",
                    "type": "string"
                },
                "confirm_pass": {
                    "description": "confirm password",
                    "type": "string"
                },
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "gender": {
                    "description": "gender",
                    "type": "integer",
                    "enum": [
                        0,
                        1,
                        2
                    ]
                },
                "name": {
                    "description": "name",
                    "type": "string"
                },
                "password": {
                    "description": "password",
                    "type": "string",
                    "minLength": 8
                },
                "phone": {
                    "description": "phone",
                    "type": "string"
                },
                "token": {
                    "description": "token",
                    "type": "string"
                }
            }
        },
        "authcommandrequest.UserRegistratorRequest": {
            "type": "object",
            "required": [
                "email",
                "purpose"
            ],
            "properties": {
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "purpose": {
                    "description": "purpose (TEST_USER, CUSTOMER, ADMIN, etc.)",
                    "type": "string",
                    "enum": [
                        "TEST_USER",
                        "CUSTOMER",
                        "ADMIN"
                    ]
                }
            }
        },
        "authcommandrequest.VerifyOTPRequest": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "otp": {
                    "description": "otp",
                    "type": "string"
                }
            }
        },
        "response.ErrorResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "transportationqueryresponse.GetListTripsResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "trips": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/transportationqueryresponse.Trip"
                    }
                }
            }
        },
        "transportationqueryresponse.Trip": {
            "type": "object",
            "properties": {
                "arrival_date": {
                    "type": "string"
                },
                "departure_date": {
                    "type": "string"
                },
                "from_location": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "to_location": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8002",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "API Documentation Ecommerce Backend SHOPDEVGO",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
