package v1alpha2

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (c *SeldonDeploymentV1Alpha2Client) SeldonDeployments(namespace string) SeldonDeploymentInt {
	return &seldonDeploymentclient{
		client: c.restClient,
		ns:     namespace,
	}
}

type SeldonDeploymentV1Alpha2Client struct {
	restClient rest.Interface
}

type SeldonDeploymentInt interface {
	Create(obj *SeldonDeployment) (*SeldonDeployment, error)
	Update(obj *SeldonDeployment) (*SeldonDeployment, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	Get(name string) (*SeldonDeployment, error)
}

type seldonDeploymentclient struct {
	client rest.Interface
	ns     string
}

func (c *seldonDeploymentclient) Create(obj *SeldonDeployment) (*SeldonDeployment, error) {
	result := &SeldonDeployment{}
	err := c.client.Post().
		Namespace(c.ns).Resource("seldondeployments").Body(obj).Do().Into(result)
	return result, err
}

func (c *seldonDeploymentclient) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().Namespace(c.ns).Resource("seldondeployments").Name(name).Body(options).Do().Error()
}

func (c *seldonDeploymentclient) Get(name string) (*SeldonDeployment, error) {
	result := &SeldonDeployment{}
	err := c.client.Get().Namespace(c.ns).Resource("seldondeployments").Name(name).Do().Into(result)
	return result, err
}

func (c *seldonDeploymentclient) Update(obj *seldonDeployment) (*SeldonDeployment, error) {
	result := &SeldonDeployment{}
	err := c.client.Put().Namespace(c.ns).Resource("seldondeployments").Body(obj).Do().Into(result)
	return result, err
}
