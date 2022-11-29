package core

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

//对deployments的集合进行定义
type DeploymentMap struct {
	data sync.Map  // [key string] []*v1.Deployment    key=>namespace
}
//添加
func(this *DeploymentMap) Add(dep *v1.Deployment){

	if list,ok:=this.data.Load(dep.Namespace);ok{
		list=append(list.([]*v1.Deployment),dep)
		this.data.Store(dep.Namespace,list)
	}else{
		this.data.Store(dep.Namespace,[]*v1.Deployment{dep})
	}
}
//更新
func(this *DeploymentMap) Update(dep *v1.Deployment) error {
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				list.([]*v1.Deployment)[i]=dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found",dep.Name)
}
// 删除
func(this *DeploymentMap) Delete(dep *v1.Deployment){
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				newList:= append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(dep.Namespace,newList)
				break
			}
		}
	}
}
func(this *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment,error){
	if list,ok:=this.data.Load(ns);ok {
		 return  list.([]*v1.Deployment),nil
	}
	return nil,fmt.Errorf("record not found")
}
func(this *DeploymentMap) GetDeployment(ns string,depname string) (*v1.Deployment,error){
	if list,ok:=this.data.Load(ns);ok {
		for _,item:=range list.([]*v1.Deployment){
			if item.Name==depname{
				return item,nil
			}
		}
	}
	return nil,fmt.Errorf("record not found")
}



// pod defind
type PodMapStruct struct {
	data sync.Map   // [key string] []*v1.Pod   key=>namespace
}

func (this *PodMapStruct) Get(ns string, podName string) *corev1.Pod {
	if list, ok := this.data.Load(ns);ok{
		for _,pod:= range list.([]*corev1.Pod){
			if pod.Name==podName{
				return pod
			}
		}
	}
	return nil
}
func (this *PodMapStruct) Add(pod *corev1.Pod)  {
	if list,ok:=this.data.Load(pod.Namespace);ok{
		list=append(list.([]*corev1.Pod),pod)
		this.data.Store(pod.Namespace,list)
	}else{
		this.data.Store(pod.Namespace,[]*corev1.Pod{pod})
	}
}

func (this *PodMapStruct) Update(pod *corev1.Pod) error {
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){
			if range_pod.Name==pod.Name{
				list.([]*corev1.Pod)[i]=pod
			}
		}
		return nil
	}
	return fmt.Errorf("pod-%s not found",pod.Name)
}

func (this *PodMapStruct) Delete(pod *corev1.Pod)  {
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){
			if range_pod.Name==pod.Name{
				newList:= append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				this.data.Store(pod.Namespace,newList)
				break
			}
		}
	}
}

func(this *PodMapStruct) ListByNS(ns string) []*corev1.Pod{
	if list,ok:=this.data.Load(ns);ok {
		return  list.([]*corev1.Pod)
	}
	return nil
}
