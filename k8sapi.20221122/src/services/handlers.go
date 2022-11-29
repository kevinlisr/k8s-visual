package services

import (
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)

//处理deployment 回调的handler
type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}
func(this *DepHandler) OnAdd(obj interface{}){
	this.DepMap.Add(obj.(*v1.Deployment))
	wscore.ClientMap.SendAllDepList(this.DepService.ListAll(obj.(*v1.Deployment).Namespace))
}
func(this *DepHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.DepMap.Update(newObj.(*v1.Deployment))
	if err!=nil{
		log.Println(err)
	}else{
		wscore.ClientMap.SendAllDepList(this.DepService.ListAll(newObj.(*v1.Deployment).Namespace))
	}
}
func(this *DepHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*v1.Deployment);ok{
		this.DepMap.Delete(d)
		wscore.ClientMap.SendAllDepList(this.DepService.ListAll(obj.(*v1.Deployment).Namespace))

	}
}

// pod相关的回调handler
type PodHandler struct {
	PodMap *PodMapStruct `inject:"-"`
}
func(this *PodHandler) OnAdd(obj interface{}){
	this.PodMap.Add(obj.(*corev1.Pod))
}
func(this *PodHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.PodMap.Update(newObj.(*corev1.Pod))
	if err!=nil{
		log.Println(err)
	}
}
func(this *PodHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Pod);ok{
		this.PodMap.Delete(d)
	}
}

// pod相关的回调handler
type NsHandler struct {
	NsMap *NsMapStruct `inject:"-"`
}
func(this *NsHandler) OnAdd(obj interface{}){
	this.NsMap.Add(obj.(*corev1.Namespace))
}
func(this *NsHandler) OnUpdate(oldObj, newObj interface{}){
	this.NsMap.Update(newObj.(*corev1.Namespace))
}
func(this *NsHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Namespace);ok{
		this.NsMap.Delete(d)
	}
}
