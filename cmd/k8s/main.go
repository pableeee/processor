package main

import (
	"bufio"
	"fmt"
	"os"

	guuid "github.com/google/uuid"
	"github.com/pableeee/processor/pkg/cmd/k8s"
)

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func main() {
	deployment := k8s.DeploymentManagerImpl{}
	id := guuid.New()

	res, err := deployment.CreateDeployment("", "default", "nginx", id.String())

	if err != nil {
		fmt.Println("Hubo un error")
		os.Exit(1)
	}

	fmt.Println("Deployment Created")
	fmt.Println(res)

	service := k8s.ServiceManagerImpl{}
	var port k8s.ServiceResponse

	port, err = service.CreateService("", "default", "nginx", 80)
	fmt.Println("Service Created")
	fmt.Println(port)

	prompt()

	err = service.DeleteService("", "default", "nginx")

	if err != nil {
		fmt.Println("Hubo un error")
		os.Exit(1)
	}
	fmt.Println("Service Deleted")

	err = deployment.DeleteDeployment("", "default", "nginx")

	if err != nil {
		fmt.Println("Hubo un error")
		os.Exit(1)
	}
	fmt.Println("Deployment Deleted")
}
