package v1alpha2

import (
	apiextensionv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CRDPlural   string = "seldondeployments"
	CRDGroup    string = "machinelearning.seldon.io"
	CRDVersion  string = "v1alpha2"
	FullCRDName string = CRDPlural + "." + CRDGroup
)

func CreateCRD(clientset apiextension.Interface) error {
	crd := &apiextensionv1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: FullCRDName},
		Spec: apiextensionv1beta1.CustomResourceDefinitionSpec{
			Group:   CRDGroup,
			Version: CRDVersion,
			Scope:   apiextensionv1beta1.NamespaceScoped,
			Names: apiextensionv1beta1.CustomResourceDefinitionNames{
				Plural: CRDPlural,
				// Kind: reflect.Type
			},
		},
	}

	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}
