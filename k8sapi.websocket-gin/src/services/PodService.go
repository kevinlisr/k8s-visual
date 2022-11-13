package services



type PodService struct {
	PodMap *PodMapStruct `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (this *PodService)ListByNs(ns string) interface{}{
	//return this.PodMap.ListByNS(ns)
	return this.PodMap.ListByNs(ns)
}

//func (this *PodService)ListByNs(ns string) []*corev1.Pod {
//	return this.ListByNs(ns)
//}