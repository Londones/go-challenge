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
        "/association": {
            "post": {
                "description": "Create a new association with the provided details",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "associations"
                ],
                "summary": "Create association",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Postal Code",
                        "name": "cp",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Location\" \"Redirect location",
                        "schema": {
                            "type": "header"
                        }
                    },
                    "400": {
                        "description": "all fields are required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error creating association",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/annonce": {
            "post": {
                "description": "crée une annonce",
                "consumes": ["application/x-www-form-urlencoded"],
                "produces": ["application/json"],
                "tags": ["annonces"],
                "summary": "create annonce",
                "parameters": [
                    {
                        "type": "string",
                        "description": "description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Location\" \"Redirect location",
                        "schema": {
                            "type": "header"
                        }
                    },
                    "400": {
                        "description": "all fields are required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error creating association",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/{provider}": {
            "get": {
                "description": "Start the authentication process with the specified provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Begin authentication provider callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication Provider",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Authentication process started",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/{provider}/callback": {
            "get": {
                "description": "Completes user authentication with the specified provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authentication callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication Provider",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Authenticated user",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login with the given email and password",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Location\" \"Redirect location",
                        "schema": {
                            "type": "header"
                        },
                        "headers": {
                            "Set-Cookie": {
                                "type": "string",
                                "description": "jwt={token}; HttpOnly; SameSite=Lax; Expires={expiry}"
                            }
                        }
                    },
                    "400": {
                        "description": "email and password are required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "invalid password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "user not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout/": {
            "get": {
                "description": "Remove the JWT cookie and redirect to the success page",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Basic logout",
                "responses": {
                    "302": {
                        "description": "Location\" \"Redirect location",
                        "schema": {
                            "type": "header"
                        }
                    }
                }
            }
        },
        "/logout/{provider}": {
            "get": {
                "description": "Logout from the authentication provider and remove the JWT cookie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout from provider",
                "responses": {
                    "307": {
                        "description": "Location\" \"Redirect location",
                        "schema": {
                            "type": "header"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with the given email, password, name, address, cp, and city",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "CP",
                        "name": "cp",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Set-Cookie": {
                                "type": "string",
                                "description": "jwt={token}; HttpOnly; SameSite=Lax; Expires={expiry}"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Annonce": {
            "type": "object",
            "properties": {
                "cats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cats"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "favorite": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Favorite"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Rating"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.Cats": {
            "type": "object",
            "properties": {
                "annonceID": {
                    "type": "integer"
                },
                "behavior": {
                    "type": "string"
                },
                "birthDate": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastVaccine": {
                    "type": "string"
                },
                "lastVaccineName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "race": {
                    "type": "string"
                },
                "reserved": {
                    "type": "boolean"
                },
                "sex": {
                    "type": "string"
                },
                "sterilized": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Favorite": {
            "type": "object",
            "properties": {
                "annonceID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.Rating": {
            "type": "object",
            "properties": {
                "annonceID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mark": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.RoleName": {
            "type": "string",
            "enum": [
                "admin",
                "user"
            ],
            "x-enum-varnames": [
                "Admin",
                "UserRole"
            ]
        },
        "models.Roles": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "$ref": "#/definitions/models.RoleName"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "addressRue": {
                    "type": "string"
                },
                "annonce": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Annonce"
                    }
                },
                "associationID": {
                    "type": "integer"
                },
                "cp": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "favorite": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Annonce"
                    }
                },
                "googleID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "rating": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Rating"
                    }
                },
                "role": {
                    "$ref": "#/definitions/models.Roles"
                },
                "updatedAt": {
                    "type": "string"
                },
                "ville": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GO-challenge-PurrfectMatch",
	Description:      "Swagger de PurrfectMatch",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
