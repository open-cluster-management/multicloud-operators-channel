// +build !ignore_autogenerated

// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.Channel":       schema_pkg_apis_app_v1alpha1_Channel(ref),
		"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelGate":   schema_pkg_apis_app_v1alpha1_ChannelGate(ref),
		"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelList":   schema_pkg_apis_app_v1alpha1_ChannelList(ref),
		"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelSpec":   schema_pkg_apis_app_v1alpha1_ChannelSpec(ref),
		"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelStatus": schema_pkg_apis_app_v1alpha1_ChannelStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Channel(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Channel is the Schema for the channels API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelSpec", "github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_ChannelGate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ChannelGate defines criteria for promote to channel",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"labelSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_app_v1alpha1_ChannelList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ChannelList contains a list of Channel",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.Channel"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.Channel", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_ChannelSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ChannelSpec defines the desired state of Channel",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"make\" to regenerate code after modifying this file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pathName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"configMapRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"gates": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelGate"),
						},
					},
					"sourceNamespaces": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "pathName"},
			},
		},
		Dependencies: []string{
			"github.com/IBM/multicloud-operators-channel/pkg/apis/app/v1alpha1.ChannelGate", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_app_v1alpha1_ChannelStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ChannelStatus defines the observed state of Channel",
				Type:        []string{"object"},
			},
		},
	}
}