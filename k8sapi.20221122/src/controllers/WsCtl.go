package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/wscore"
	"log"
)

//@Controller
type WsCtl struct {}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func(this *WsCtl) Connect(c *gin.Context) string    {
	client,err:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)  //升级
	if err!=nil{
		log.Println(err)
		return err.Error()
	}else{
		wscore.ClientMap.Store(client)
		return "success"
	}

}
func(this *WsCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ws",this.Connect)
}
func(this *WsCtl) Name() string{
	return "WsCtl"
}
