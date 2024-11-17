// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "url": "http://www.swagger.io/support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/device": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Create a new device",
                "parameters": [
                    {
                        "description": "Device to create",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateDeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/device-usage": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find all device usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device Usage"
                ],
                "summary": "Find all device usage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceUsageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new device usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device Usage"
                ],
                "summary": "Create a new device usage",
                "parameters": [
                    {
                        "description": "Device usage to create",
                        "name": "device_usage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateDeviceUsageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceUsageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/device/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get a device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the device to get",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Update a device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the device to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Device to update",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateDeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Delete a device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the device to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/devices": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get all devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeviceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Log in a user using email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Request Body",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "Register Request Body",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/report": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Send device usage report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Email"
                ],
                "summary": "Send device usage report",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.EmailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/suggestion": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get suggestions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Suggestion"
                ],
                "summary": "Get suggestions",
                "parameters": [
                    {
                        "description": "City to get suggestions",
                        "name": "city",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SuggestionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SuggestionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user-usage": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find user usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Usage"
                ],
                "summary": "Find user usage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserUsageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new user usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Usage"
                ],
                "summary": "Create a new user usage",
                "parameters": [
                    {
                        "description": "User usage to create",
                        "name": "user_usage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserUsageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserUsageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        },
        "/weather": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get weather by city and date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Weather"
                ],
                "summary": "Get weather by city and date",
                "parameters": [
                    {
                        "description": "City of the weather",
                        "name": "city",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateWeatherRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WeatherResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.BaseResponse": {
            "description": "BaseResponse is the base response for all the endpoints",
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "request.CreateDeviceRequest": {
            "description": "CreateDeviceRequest is the request for the create device endpoint",
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "power": {
                    "type": "number"
                }
            }
        },
        "request.CreateDeviceUsageRequest": {
            "description": "CreateDeviceUsageRequest is the request for the create device-usage endpoint",
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "request.CreateUserUsageRequest": {
            "description": "CreateUserUsageRequest is the request for the create user-usage endpoint",
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                }
            }
        },
        "request.CreateWeatherRequest": {
            "description": "CreateWeatherRequest is the request for the create weather endpoint",
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "description": "LoginRequest is the request for the login endpoint",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.RegisterRequest": {
            "description": "RegisterRequest is the request for the register endpoint",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.SuggestionRequest": {
            "description": "SuggestionRequest is the request for the suggestion endpoint",
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                }
            }
        },
        "request.UpdateDeviceRequest": {
            "description": "UpdateDeviceRequest is the request for the update device endpoint",
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "power": {
                    "type": "number"
                }
            }
        },
        "response.AuthResponse": {
            "description": "AuthResponse is the response for the auth controller",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "response.DeviceResponse": {
            "description": "DeviceResponse is the response for the device endpoint",
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "power": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "response.DeviceUsageResponse": {
            "description": "DeviceUsageResponse is the response for the device-usage endpoint",
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "integer"
                },
                "duration": {
                    "type": "number"
                },
                "end_time": {
                    "type": "string"
                },
                "energy_consumed": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "response.EmailResponse": {
            "description": "EmailResponse is the response for the email endpoint",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.SuggestionResponse": {
            "description": "SuggestionResponse is the response for the suggestion endpoint",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.UserUsageResponse": {
            "description": "UserUsageResponse is the response for the user usage endpoint",
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "total_cost": {
                    "type": "number"
                },
                "total_energy": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "response.WeatherResponse": {
            "description": "WeatherResponse is the response for the weather endpoint",
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "condition": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "humidity": {
                    "type": "number"
                },
                "temperature": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://52.65.161.24:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Energia API Mini Project",
	Description:      "This is a simple API for managing energy usage in a household.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}