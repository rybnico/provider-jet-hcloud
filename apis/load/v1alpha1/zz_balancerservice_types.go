/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BalancerServiceHTTPObservation struct {
}

type BalancerServiceHTTPParameters struct {

	// +kubebuilder:validation:Optional
	Certificates []*int64 `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// +kubebuilder:validation:Optional
	CookieLifetime *int64 `json:"cookieLifetime,omitempty" tf:"cookie_lifetime,omitempty"`

	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// +kubebuilder:validation:Optional
	RedirectHTTP *bool `json:"redirectHttp,omitempty" tf:"redirect_http,omitempty"`

	// +kubebuilder:validation:Optional
	StickySessions *bool `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`
}

type BalancerServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BalancerServiceParameters struct {

	// +kubebuilder:validation:Optional
	DestinationPort *int64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// +kubebuilder:validation:Optional
	HTTP []BalancerServiceHTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// +kubebuilder:validation:Optional
	ListenPort *int64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerID *string `json:"loadBalancerId" tf:"load_balancer_id,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Proxyprotocol *bool `json:"proxyprotocol,omitempty" tf:"proxyprotocol,omitempty"`
}

type HTTPObservation struct {
}

type HTTPParameters struct {

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`

	// +kubebuilder:validation:Optional
	StatusCodes []*string `json:"statusCodes,omitempty" tf:"status_codes,omitempty"`

	// +kubebuilder:validation:Optional
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// +kubebuilder:validation:Required
	Interval *int64 `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Retries *int64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *int64 `json:"timeout" tf:"timeout,omitempty"`
}

// BalancerServiceSpec defines the desired state of BalancerService
type BalancerServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerServiceParameters `json:"forProvider"`
}

// BalancerServiceStatus defines the observed state of BalancerService.
type BalancerServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerService is the Schema for the BalancerServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloudjet}
type BalancerService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerServiceSpec   `json:"spec"`
	Status            BalancerServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerServiceList contains a list of BalancerServices
type BalancerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerService `json:"items"`
}

// Repository type metadata.
var (
	BalancerService_Kind             = "BalancerService"
	BalancerService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerService_Kind}.String()
	BalancerService_KindAPIVersion   = BalancerService_Kind + "." + CRDGroupVersion.String()
	BalancerService_GroupVersionKind = CRDGroupVersion.WithKind(BalancerService_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerService{}, &BalancerServiceList{})
}