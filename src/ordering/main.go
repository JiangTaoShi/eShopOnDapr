package ordering

import (
	"fmt"
	"github.com/JiangTaoShi/eShopOnDapr/ordering/router"
	"github.com/JiangTaoShi/eShopOnDapr/pkg/dapr"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	dapr.Setup()
}

func main() {
	gin.SetMode(ServerSetting.RunMode)

	routersInit := router.InitRouter()
	readTimeout := ServerSetting.ReadTimeout
	writeTimeout := ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("[error] start http server error")
	}
}
