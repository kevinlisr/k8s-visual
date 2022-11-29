package services

import "k8sapi/src/models"

//@Service
type PodService struct {
	PodMap *PodMapStruct `inject:"-"`
	Common *CommonService `inject:"-"`
	EventMap *EventMapStruct `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}
func(this *PodService) ListByNs(ns string ) interface{}{
	podList:=this.PodMap.ListByNs(ns)
	ret:=make([]*models.Pod,0)
	for _,pod:=range podList{
		ret=append(ret,&models.Pod{
			Name:pod.Name,
			NameSpace:pod.Namespace,
			Images:this.Common.GetImagesByPod(pod.Spec.Containers),
			NodeName:pod.Spec.NodeName,
			Phase:string(pod.Status.Phase),// 阶段
			IsReady:this.Common.PosIsReady(pod), //是否就绪
			IP:[]string{pod.Status.PodIP,pod.Status.HostIP},
			Message:this.EventMap.GetMessage(pod.Namespace,"Pod",pod.Name),
			CreateTime:pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}