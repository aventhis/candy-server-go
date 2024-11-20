// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Candy Server",
    "version": "1.0.0"
  },
  "paths": {
    "/buy_candy": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "buyCandy",
        "parameters": [
          {
            "description": "summary of the candy order",
            "name": "order",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "money",
                "candyType",
                "candyCount"
              ],
              "properties": {
                "candyCount": {
                  "description": "number of candy",
                  "type": "integer"
                },
                "candyType": {
                  "description": "kind of candy",
                  "type": "string"
                },
                "money": {
                  "description": "amount of money put into vending machine",
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "purchase succesful",
            "schema": {
              "type": "object",
              "properties": {
                "change": {
                  "type": "integer"
                },
                "thanks": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "some error in input data",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          },
          "402": {
            "description": "not enough money",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Candy Server",
    "version": "1.0.0"
  },
  "paths": {
    "/buy_candy": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "buyCandy",
        "parameters": [
          {
            "description": "summary of the candy order",
            "name": "order",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "money",
                "candyType",
                "candyCount"
              ],
              "properties": {
                "candyCount": {
                  "description": "number of candy",
                  "type": "integer"
                },
                "candyType": {
                  "description": "kind of candy",
                  "type": "string"
                },
                "money": {
                  "description": "amount of money put into vending machine",
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "purchase succesful",
            "schema": {
              "type": "object",
              "properties": {
                "change": {
                  "type": "integer"
                },
                "thanks": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "some error in input data",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          },
          "402": {
            "description": "not enough money",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}
