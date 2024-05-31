// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KeyAllowedApplications struct {
	/* The package name of the application. */
	PackageName string `json:"packageName"`

	/* The SHA1 fingerprint of the application. For example, both sha1 formats are acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format is the latter. */
	Sha1Fingerprint string `json:"sha1Fingerprint"`
}

type KeyAndroidKeyRestrictions struct {
	/* A list of Android applications that are allowed to make API calls with this key. */
	AllowedApplications []KeyAllowedApplications `json:"allowedApplications"`
}

type KeyApiTargets struct {
	/* Optional. List of one or more methods that can be called. If empty, all methods for the service are allowed. A wildcard (*) can be used as the last symbol. Valid examples: `google.cloud.translate.v2.TranslateService.GetSupportedLanguage` `TranslateText` `Get*` `translate.googleapis.com.Get*` */
	// +optional
	Methods []string `json:"methods,omitempty"`

	/* The service for this restriction. It should be the canonical service name, for example: `translate.googleapis.com`. You can use [`gcloud services list`](/sdk/gcloud/reference/services/list) to get a list of services that are enabled in the project. */
	Service string `json:"service"`
}

type KeyBrowserKeyRestrictions struct {
	/* A list of regular expressions for the referrer URLs that are allowed to make API calls with this key. */
	AllowedReferrers []string `json:"allowedReferrers"`
}

type KeyIosKeyRestrictions struct {
	/* A list of bundle IDs that are allowed when making API calls with this key. */
	AllowedBundleIds []string `json:"allowedBundleIds"`
}

type KeyRestrictions struct {
	/* The Android apps that are allowed to use the key. */
	// +optional
	AndroidKeyRestrictions *KeyAndroidKeyRestrictions `json:"androidKeyRestrictions,omitempty"`

	/* A restriction for a specific service and optionally one or more specific methods. Requests are allowed if they match any of these restrictions. If no restrictions are specified, all targets are allowed. */
	// +optional
	ApiTargets []KeyApiTargets `json:"apiTargets,omitempty"`

	/* The HTTP referrers (websites) that are allowed to use the key. */
	// +optional
	BrowserKeyRestrictions *KeyBrowserKeyRestrictions `json:"browserKeyRestrictions,omitempty"`

	/* The iOS apps that are allowed to use the key. */
	// +optional
	IosKeyRestrictions *KeyIosKeyRestrictions `json:"iosKeyRestrictions,omitempty"`

	/* The IP addresses of callers that are allowed to use the key. */
	// +optional
	ServerKeyRestrictions *KeyServerKeyRestrictions `json:"serverKeyRestrictions,omitempty"`
}

type KeyServerKeyRestrictions struct {
	/* A list of the caller IP addresses that are allowed to make API calls with this key. */
	AllowedIps []string `json:"allowedIps"`
}

type APIKeysKeySpec struct {
	/* Human-readable display name of this key that you can modify. The maximum length is 63 characters. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Key restrictions. */
	// +optional
	Restrictions *KeyRestrictions `json:"restrictions,omitempty"`
}

type KeyObservedStateStatus struct {
	/* Output only. Unique id in UUID4 format. */
	// +optional
	Uid *string `json:"uid,omitempty"`
}

type APIKeysKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   APIKeysKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *KeyObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapikeyskey;gcpapikeyskeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/directcrd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIKeysKey is the Schema for the apikeys API
// +k8s:openapi-gen=true
type APIKeysKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIKeysKeySpec   `json:"spec,omitempty"`
	Status APIKeysKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIKeysKeyList contains a list of APIKeysKey
type APIKeysKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKeysKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIKeysKey{}, &APIKeysKeyList{})
}
