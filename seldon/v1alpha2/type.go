package v1alpha2

import meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SeldonDeployment struct {
	meta_v1.TypeMeta   `json:"inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               SeldonDeploymentSpec   `json:"spec"`
	Status             SeldonDeploymentStatus `json:"status,omitempty"`
}

type SeldonDeploymentSpec struct {
	Name        string `json:"cert"`
	OautKey     string `json:"oauth_key"`
	OauthSecret string `json:"oauth_secret"`
}

type SeldonDeploymentStatus struct {
	State string `json:"state,omitempty"`
}

type SeldonDeploymentList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []SeldonDeployment `json:"items"`
}
