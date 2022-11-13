package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/configs"
	"k8sapi/src/controllers"
	"net/http"
)

func cross() gin.HandlerFunc {
	return func (c *gin.Context)  {
		method := c.Request.Method
		if method != "" {
		//c.Header("Access-Control-Allow-Origin", "*") // ke jiang * ti huan wei zhi ding de yu ming
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Token")
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
	}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	    }
	}

}

func main() {
	service := goft.Ignite(cross()).Config(
		configs.NewK8sHandler(),  //1
		configs.NewK8sConfig(), //2
		configs.NewK8sMaps(), //3
		configs.NewServiceConfig(), //4
		).
		Mount("",controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			controllers.NewUserCtl(),
			).
		Attach(
			//middlewares.NewCrosMiddleware(),
			)
	    service.Static("/dashboard","./dist")
	    service.Static("/static","./dist/static")
		service.Launch()


}
