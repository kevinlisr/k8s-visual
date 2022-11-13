package configs

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8sapi/src/core"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
	PodHandler *core.PodHandler `inject:"-"`
}
func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}
//初始化客户端
func(*K8sConfig) InitClient() *kubernetes.Clientset{
	config, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	//config:=&rest.Config{
	//	Host:"http://192.168.31.61:8009",
 	//}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	//fmt.Println(config)
	return c

}
//初始化Informer
func(this *K8sConfig) InitInformer() informers.SharedInformerFactory{
	fact:=informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer:=fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	podInformer:=fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(this.PodHandler)

	fact.Start(wait.NeverStop)

	return fact
}