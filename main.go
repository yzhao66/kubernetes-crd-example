package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/martin-helmich/kubernetes-crd-example/controllers"
)

var namespace = "default"
var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	beego.BConfig.CopyRequestBody = true
	beego.Router("/updata", new(controllers.UpdataController), "get,post:Update")
	beego.Run(":80")

	/*v1alpha1.AddToScheme(scheme.Scheme)
	kubeConfig, err := utils.KubeConfig()
	clientSet, err := clientV1alpha1.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}

	projects, err := clientSet.Projects("default").List(metav1.ListOptions{})

	devices, err := clientSet.Devices("default").List(metav1.ListOptions{})


	if err != nil {
		panic(err)
	}
	fmt.Printf("devices found: %+v\n", devices)
	fmt.Printf("projects found: %+v\n", projects)

	store := WatchResources(clientSet)

	for {
		projectsFromStore := store.List()
		fmt.Printf("project in store: %d\n", len(projectsFromStore))

		time.Sleep(2 * time.Second)
	}*/


}
