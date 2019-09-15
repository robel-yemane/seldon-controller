package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"machinelearning.seldon.io/seldon/v1alpha2"

	"github.com/golang/glog"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/glog"
)

func main() {

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		fmt.Printf("Using default kubeconfig @ \"%s\"\n", *kubeconfig)
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	kubeClient, err := apiextension.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create client: %v", err)
	}

	//create crd
	err = v1alpha2.CreateCRD(kubeClient)
	if err != nil {
		glog.Fatalf("Failed to create crd: %v", err)
	}

	//Wait for the CRD to be created before we use it
	time.Sleep(5 * time.Second)

	//Create a new clientset which include our CRD schema
	crdclient, err := v1alpha2.NewClient(config)
	if err != nil {
		panic(err)
	}

	SeldonDeployment := &v1alpha2.SeldonDeployment{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "seldondeploymentobj",
			Labels: map[string]string{"app": "seldon", "name": "seldon-model"},
		},
		Spec: v1alpha2.SeldonDeploymentSpec{
			Name:        "test-deployment",
			OautKey:     "oauth-key",
			OauthSecret: "oauth-secret",
		},
		Status: v1alpha2.SeldonDeploymentStatus{
			State: "Available",
		},
	}

	//Create the SeldonDeployment object in the cluster
	resp, err := crdclient.SeldonDeployments("default").Create(SeldonDeployment)
	if err != nil {
		fmt.Printf("error occured while creating obj: %v\n", err)
	} else {
		fmt.Printf("obj creaed: %v\n", resp)
	}

	obj, err := crdclient.SeldonDeployments("default").Get(SeldonDeployment.ObjectMeta.Name)
	if err != nil {
		glog.Infof("error while getting the object %v\n", err)
	}
	fmt.Printf("SeldonDeployment obj found: \n%+v\n", obj)

	select {}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
