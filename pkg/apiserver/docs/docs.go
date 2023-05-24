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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/clusters": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "List all clusters. Use query parameters to filter results. Auth is required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "List all clusters",
                "operationId": "v1-get-clusters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by region",
                        "name": "region",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by environment",
                        "name": "environment",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter since last updated (RFC3339)",
                        "name": "lastUpdated",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset to start pagination search results (default is 0)",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "The number of results per page (default is 200)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_apiserver_web_handler_v1.clusterList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/v1/clusters/{name}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get an cluster. Auth is required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Get an cluster",
                "operationId": "v1-get-cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the cluster to get",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ClusterSpec"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/v2/clusters": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "List all or a subset of clusters. Use conditions to filter clusters based on their fields.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "List clusters",
                "operationId": "v2-get-clusters",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Filter conditions",
                        "name": "conditions",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset to start pagination search results (default is 0)",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "The number of results per page (default is 200)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_apiserver_web_handler_v2.clusterList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        },
        "/v2/clusters/{name}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get an cluster. Auth is required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Get an cluster",
                "operationId": "v2-get-cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the cluster to get",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ClusterSpec"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Update a cluster. Auth is required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Patch a cluster",
                "operationId": "v2-patch-cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the cluster to patch",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ClusterSpec"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.Error": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "github.com_adobe_cluster-registry_pkg_apiserver_web_handler_v1.clusterList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ClusterSpec"
                    }
                },
                "itemsCount": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "more": {
                    "type": "boolean"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "github.com_adobe_cluster-registry_pkg_apiserver_web_handler_v2.clusterList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ClusterSpec"
                    }
                },
                "itemsCount": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "more": {
                    "type": "boolean"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "pkg_apiserver_web_handler_v1.clusterList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ClusterSpec"
                    }
                },
                "itemsCount": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "more": {
                    "type": "boolean"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "pkg_apiserver_web_handler_v2.clusterList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.ClusterSpec"
                    }
                },
                "itemsCount": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "more": {
                    "type": "boolean"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "v1.APIServer": {
            "type": "object",
            "properties": {
                "certificateAuthorityData": {
                    "description": "Information about K8s Api CA Cert",
                    "type": "string"
                },
                "endpoint": {
                    "description": "Information about K8s Api Endpoint\n+kubebuilder:validation:Required",
                    "type": "string"
                }
            }
        },
        "v1.AllowedOnboardingTeam": {
            "type": "object",
            "properties": {
                "gitTeams": {
                    "description": "List of git teams",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ldapGroups": {
                    "description": "List of ldap groups",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Name of the team\n+kubebuilder:validation:Required",
                    "type": "string"
                }
            }
        },
        "v1.Capacity": {
            "type": "object",
            "properties": {
                "clusterCapacity": {
                    "type": "integer"
                },
                "clusterProvisioning": {
                    "type": "integer"
                },
                "lastUpdated": {
                    "type": "string"
                }
            }
        },
        "v1.ClusterSpec": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "The cloud account associated with the cluster\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "allowedOnboardingTeams": {
                    "description": "Git teams and/or LDAP groups that are allowed to onboard and deploy on the cluster",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.AllowedOnboardingTeam"
                    }
                },
                "apiServer": {
                    "description": "Information about K8s API endpoint and CA cert\n+kubebuilder:validation:Required",
                    "$ref": "#/definitions/v1.APIServer"
                },
                "businessUnit": {
                    "description": "The BU that owns the cluster\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "capabilities": {
                    "description": "List of cluster capabilities",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "Capacity cluster information",
                    "$ref": "#/definitions/v1.Capacity"
                },
                "cloudProviderRegion": {
                    "description": "The cloud provider standard region\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "cloudType": {
                    "description": "The cloud provider\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "environment": {
                    "description": "Cluster environment\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "extra": {
                    "description": "Extra information, not necessary related to the cluster",
                    "$ref": "#/definitions/v1.Extra"
                },
                "k8sInfraRelease": {
                    "description": "K8s Infrastructure release information\n+kubebuilder:validation:Required",
                    "$ref": "#/definitions/v1.K8sInfraRelease"
                },
                "lastUpdated": {
                    "description": "Timestamp when cluster information was updated",
                    "type": "string"
                },
                "managingOrg": {
                    "description": "The Org that is responsible for the cluster operations\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "name": {
                    "description": "Cluster name\n+kubebuilder:validation:Required\n+kubebuilder:validation:MaxLength=64\n+kubebuilder:validation:MinLength=3",
                    "type": "string"
                },
                "offering": {
                    "description": "The Offering that the cluster is meant for\n+kubebuilder:validation:Required",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "peerVirtualNetworks": {
                    "description": "Information about Virtual Networks manual peered with the cluster",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.PeerVirtualNetwork"
                    }
                },
                "phase": {
                    "description": "Cluster phase\n+kubebuilder:validation:Required\n+kubebuilder:validation:Enum=Building;Testing;Running;Upgrading",
                    "type": "string"
                },
                "region": {
                    "description": "Cluster internal region name\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "registeredAt": {
                    "description": "Timestamp when cluster was registered in Cluster Registry\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "shortName": {
                    "description": "Cluster name, without dash\n+kubebuilder:validation:Required\n+kubebuilder:validation:MaxLength=64\n+kubebuilder:validation:MinLength=3",
                    "type": "string"
                },
                "status": {
                    "description": "Cluster status\n+kubebuilder:validation:Required\n+kubebuilder:validation:Enum=Inactive;Active;Deprecated;Deleted",
                    "type": "string"
                },
                "tags": {
                    "description": "Cluster tags that were applied",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "tiers": {
                    "description": "List of tiers with their associated information\n+kubebuilder:validation:Required",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Tier"
                    }
                },
                "type": {
                    "description": "The type of the cluster",
                    "type": "string"
                },
                "virtualNetworks": {
                    "description": "Virtual Private Networks information\n+kubebuilder:validation:Required",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.VirtualNetwork"
                    }
                }
            }
        },
        "v1.Extra": {
            "type": "object",
            "properties": {
                "domainName": {
                    "description": "Name of the domain\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "ecrIamArns": {
                    "description": "List of IAM Arns",
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "egressPorts": {
                    "description": "Egress ports allowed outside of the namespace",
                    "type": "string"
                },
                "extendedRegion": {
                    "description": "ExtendedRegion information",
                    "type": "string"
                },
                "lbEndpoints": {
                    "description": "Load balancer endpoints\n+kubebuilder:validation:Required",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "loggingEndpoints": {
                    "description": "Logging endpoints\n+kubebuilder:validation:Required",
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                },
                "nfsInfo": {
                    "description": "NFS information",
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "v1.K8sInfraRelease": {
            "type": "object",
            "properties": {
                "gitSha": {
                    "description": "GitSha of the release",
                    "type": "string"
                },
                "lastUpdated": {
                    "description": "When the release was applied on the cluster",
                    "type": "string"
                },
                "release": {
                    "description": "Release name",
                    "type": "string"
                }
            }
        },
        "v1.PeerVirtualNetwork": {
            "type": "object",
            "properties": {
                "cidrs": {
                    "description": "Remote Virtual Netowrk CIDRs",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "Remote Virtual Netowrk ID",
                    "type": "string"
                },
                "ownerID": {
                    "description": "Cloud account of the owner",
                    "type": "string"
                }
            }
        },
        "v1.Tier": {
            "type": "object",
            "properties": {
                "containerRuntime": {
                    "description": "Container runtime\n+kubebuilder:validation:Required\n+kubebuilder:validation:Enum=docker;cri-o",
                    "type": "string"
                },
                "enableKataSupport": {
                    "description": "EnableKataSupport",
                    "type": "boolean"
                },
                "instanceType": {
                    "description": "Type of the instances\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "kernelParameters": {
                    "description": "KernelParameters",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "Instance K8s labels",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "maxCapacity": {
                    "description": "Max number of instances\n+kubebuilder:validation:Required",
                    "type": "integer"
                },
                "minCapacity": {
                    "description": "Min number of instances\n+kubebuilder:validation:Required",
                    "type": "integer"
                },
                "name": {
                    "description": "Name of the tier\n+kubebuilder:validation:Required",
                    "type": "string"
                },
                "taints": {
                    "description": "Instance K8s taints",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.VirtualNetwork": {
            "type": "object",
            "properties": {
                "cidrs": {
                    "description": "CIDRs used in this VirtualNetwork\n+kubebuilder:validation:Required",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "Virtual private network Id\n+kubebuilder:validation:Required",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://127.0.0.1:8080",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Cluster Registry API",
	Description:      "Cluster Registry API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
