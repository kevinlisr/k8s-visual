package services

import (
	"k8sapi/src/core"
)

type PodService struct {
	PodMap *core.PodMapStruct `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (this *PodService)ListByNs(ns string) interface{}{
	return this.PodMap.ListByNS(ns)
}

//func (this *PodService)ListByNs(ns string) []*corev1.Pod {
//	return this.ListByNs(ns)
//}