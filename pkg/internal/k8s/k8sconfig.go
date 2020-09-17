package k8s

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var ErrorNoConfig = fmt.Errorf("no kubeconfig provided")

func NewConfigSetup(cfg string, namespace string) (string, dynamic.Interface, error) {
	var kubeconfig string

	if cfg != "" {
		kubeconfig = cfg
	} else {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			return "", nil, ErrorNoConfig
		}
	}

	if namespace == "" {
		namespace = "default"
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return namespace, client, err
}
