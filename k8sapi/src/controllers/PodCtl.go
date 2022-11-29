package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/services"
)

type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}
func(this *PodCtl) GetAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.PodService.ListByNs(ns),
	}

}
func(this *PodCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/pods",this.GetAll)
}
func(*PodCtl) Name() string{
	return "PodCtl"
}