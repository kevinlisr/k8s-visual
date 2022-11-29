package services

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/models"
)

type PodService struct {
	PodMap *PodMapStruct `inject:"-"`
	Common *CommonService `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (this *PodService)ListBy()string{
	return "hello"
}

func (this *PodService)ListByNs(ns string) interface{}{
	//return this.PodMap.ListByNS(ns)
	//return this.PodMap.ListByNs(ns)
	list := this.PodMap.ListByNs(ns)
	ret := make([]*models.Pod, 0)
	for _, pod := range list{
		ret = append(ret, &models.Pod{
			Name: pod.Name,
			NameSpace: pod.Namespace,
			Images: this.Common.GetImagesByPod(pod.Spec.Containers),
			NodeName: pod.Spec.NodeName,
			Phase: string(pod.Status.Phase),
			IsReady: this.Common.PodsIsReady(pod),
			IP: []string{pod.Status.PodIP,pod.Status.HostIP},
			// Message: core.EventMessage
			CreateTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return ret

}

func (this *PodService)PodListByNs(ns string) []*corev1.Pod {
	list := this.PodMap.ListByNs(ns)
	return list
}