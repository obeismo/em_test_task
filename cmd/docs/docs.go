// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

// const docTemplate = `{
//     "schemes": {{ marshal .Schemes }},
//     "swagger": "2.0",
//     "info": {
//         "description": "{{escape .Description}}",
//         "title": "{{.Title}}",
//         "contact": {},
//         "version": "{{.Version}}"
//     },
//     "host": "{{.Host}}",
//     "basePath": "{{.BasePath}}",
//     "paths": {}
// }`

// SwaggerInfo holds exported Swagger Info so clients can modify it
const docсTemplate = `
	{
  "swagger": "2.0",
  "info": {
    "title": "Car API",
    "version": "1.0",
    "description": "API for managing cars"
  },
  "host": "localhost:8000",
  "basePath": "/",
  "paths": {
    "/cars": {
      "post": {
        "summary": "Add a new car",
        "description": "Add a new car to the database",
        "produces": ["application/json"],
        "parameters": [
          {
            "name": "car",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Car"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "car_id: ID of the newly added car",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "get": {
        "summary": "Get cars",
        "description": "Get a list of cars based on query parameters",
        "produces": ["application/json"],
        "parameters": [
          {
            "name": "reg_num",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mark",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "model",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "year",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "page",
            "in": "query",
            "required": true,
            "type": "string"
          },
          {
            "name": "page_size",
            "in": "query",
            "required": true,
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "List of cars",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Car"
              }
            }
          }
        }
      }
    },
    "/cars/{id}": {
      "put": {
        "summary": "Update a car",
        "description": "Update a car in the database",
        "produces": ["application/json"],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "updateFields",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "message: successfully updated",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete a car",
        "description": "Delete a car from the database",
        "produces": ["application/json"],
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "message: Vehicle successfully deleted",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "People": {
      "type": "object",
      "properties": {
        "id": {
          "type": "int"
        },
        "name": {
          "type": "string"
        },
        "surname": {
          "type": "string"
        },
        "patronymic": {
          "type": "string"
        }
      }
    },
    "Car": {
      "type": "object",
      "properties": {
        "id": {
          "type": "int"
        },
        "reg_num": {
          "type": "string"
        },
        "mark": {
          "type": "string"
        },
        "model": {
          "type": "string"
        },
        "year": {
          "type": "int"
        },
        "owner": {
          "$ref": "#/definitions/People"
        }
      }
    }
  }
}
`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Car API",
	Description:      "API for managing cars and owners",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docсTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}