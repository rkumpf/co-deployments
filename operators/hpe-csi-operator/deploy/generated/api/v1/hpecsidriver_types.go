/*

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HPECSIDriverSpec defines the desired state of HPECSIDriver
type HPECSIDriverSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Image Pull Policy for HPE CSI driver images
	ImagePullPolicy string `json:"imagePullPolicy"`
	// Flavor of the CO orchestrator
	Flavor string `json:"flavor"`
	// Default logLevel for HPE CSI driver deployments
	LogLevel string `json:"logLevel"`
	// Disable automatic installation of iSCSI/Multipath packages, default: false
	DisableNodeConformance bool `json:"disableNodeConformance"`
	// Iscsi parameters to be configured
	Iscsi IscsiInfo `json:"iscsi"`
	// Registry prefix for HPE CSI driver images, default: quay.io
	Registry string `json:"registry"`
	// Enables and configures the CSI info metrics
	InfoMetrics InfoMetricsInfo `json:"infoMetrics,omitempty"`
}

// IscsiInfo defines different Iscsi parameters which can be configured
type IscsiInfo struct {
	// iSCSI CHAP username
	ChapUser string `json:"chapUser,omitempty"`
	// iSCSI CHAP password
	ChapPassword string `json:"chapPassword,omitempty"`
}

// InfoMetricsServiceInfo specifies properties of the service definition for the CSI info metrics
type InfoMetricsServiceInfo struct {
	// The type of service created to provide access to the info metrics:
	// ClusterIP (the default) for access solely from within the cluster
	// or NodePort to enable access from outside the cluster
	Type string `json:"type,omitempty"`
	// The port exposed by the info metrics service within the cluster
	// (default 9090)
	Port int `json:"port,omitempty"`
	// The external node port at which info metrics are served, when the
	// service type is NodePort
	NodePort int `json:"nodePort,omitempty"`
	// Labels to be added to the info metrics service, for example to add
	// target labels in a ServiceMonitor scrape configuration
	CustomLabels map[string]string `json:"customLabels,omitempty"`
}

// InfoMetricsInfo specifies whether the CSI info metrics container is
// enabled and how its service is defined
type InfoMetricsInfo struct {
	// Specifies whether CSI info metrics are provided
	Enabled bool `json:"enabled,omitempty"`
	// Specifies properties of the CSI info metrics service
	Service InfoMetricsServiceInfo `json:"service,omitempty"`
}

type HelmAppConditionType string
type ConditionStatus string
type HelmAppConditionReason string

type HelmAppCondition struct {
	Type    HelmAppConditionType   `json:"type"`
	Status  ConditionStatus        `json:"status"`
	Reason  HelmAppConditionReason `json:"reason,omitempty"`
	Message string                 `json:"message,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type HelmAppRelease struct {
	Name     string `json:"name,omitempty"`
	Manifest string `json:"manifest,omitempty"`
}

// HpecsidriverStatus defines the observed state of Hpecsidriver
type HPECSIDriverStatus struct {
	// HPE CSI Driver helm release status
	Conditions []HelmAppCondition `json:"conditions"`
	// HPE CSI Driver helm release
	DeployedRelease *HelmAppRelease `json:"deployedRelease,omitempty"`
}

// +kubebuilder:object:root=true

// HPECSIDriver is the Schema for the hpecsidrivers API
// +kubebuilder:subresource:status
type HPECSIDriver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HPECSIDriverSpec   `json:"spec,omitempty"`
	Status HPECSIDriverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPECSIDriverList contains a list of Hpecsidriver
type HPECSIDriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPECSIDriver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HPECSIDriver{}, &HPECSIDriverList{})
}
