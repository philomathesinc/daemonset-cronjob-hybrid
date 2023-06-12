package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DaemonjobSpec defines the desired state of Daemonjob
type DaemonjobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Daemonjob. Edit daemonjob_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DaemonjobStatus defines the observed state of Daemonjob
type DaemonjobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Daemonjob is the Schema for the daemonjobs API
type Daemonjob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaemonjobSpec   `json:"spec,omitempty"`
	Status DaemonjobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DaemonjobList contains a list of Daemonjob
type DaemonjobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Daemonjob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Daemonjob{}, &DaemonjobList{})
}
