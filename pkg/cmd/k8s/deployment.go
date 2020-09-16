// Note: the example only works with the code within the same release/branch.
package k8s

import (
	"context"
	"fmt"

	"github.com/pableeee/processor/pkg/internal/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

//DeploymentManager K8s deployment wrapper interface
type DeploymentManager interface {
	CreateDeployment(cfg, namespace, image, name string) (string, error)
	DeleteDeployment(cfg, namespace, name string) error
}

//DeploymentManagerImpl DeploymentManager implementation
type DeploymentManagerImpl struct {
}

//CreateDeployment creates a kubernetes deployment with the given parameters
func (dp *DeploymentManagerImpl) CreateDeployment(cfg, namespace, image, name string) (string, error) {

	namespace, client, err := k8s.NewConfigSetup(cfg, namespace)

	if err != nil {
		return "", err
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	deployment := createDeploymentFromTemplate(namespace, image, name)

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := client.Resource(deploymentRes).Namespace(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetName())

	return "foobar", err
}

/*
Don't kill ne future me ;)

func (dp *DeploymentManagerImpl) listDeployments(client dynamic.Interface, deploymentRes schema.GroupVersionResource, namespace string) {
	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := client.Resource(deploymentRes).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		if err != nil || !found {
			fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
			continue
		}
		fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
	}
}
*/

//DeleteDeployment deletes the specified deployment
func (dp *DeploymentManagerImpl) DeleteDeployment(cfg, namespace, name string) error {
	namespace, client, err := k8s.NewConfigSetup(cfg, namespace)

	if err != nil {
		return err
	}

	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	err = client.Resource(deploymentRes).Namespace(namespace).Delete(context.TODO(), name, deleteOptions)

	return err
}

func createDeploymentFromTemplate(namespace, image, name string) *unstructured.Unstructured {
	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
				"labels": map[string]interface{}{
					"app": name,
				},
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": name,
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": name,
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  name,
								"image": image,
								/*								"ports": []map[string]interface{}{
																{
																	"name":          "http",
																	"protocol":      "TCP",
																	"containerPort": 80,
																},
															},*/
							},
						},
					},
				},
			},
		},
	}
	return deployment
}
