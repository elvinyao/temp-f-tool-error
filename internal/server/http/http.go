package http

import (
	"fmt"
	_ "focalboard-tool/docs"
	"focalboard-tool/internal/conf"
	"focalboard-tool/internal/middleware"
	"focalboard-tool/internal/service"
	"focalboard-tool/library/log"
	"focalboard-tool/pkg/errors"
	"net/http"
	"os"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var (
	srv *service.Service
)

type apiroutes struct {
	routerGroup *gin.Engine
}

func (r apiroutes) Run(addr ...string) error {
	return r.routerGroup.Run()
}

// New init
func New(c *conf.Config, s *service.Service) (httpSrv *http.Server) {
	gin.SetMode(gin.ReleaseMode)
	routerRes := route()
	readTimeout := conf.Conf.App.ReadTimeout
	writeTimeout := conf.Conf.App.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.Conf.App.HttpPort)
	maxHeaderBytes := 1 << 20
	httpSrv = &http.Server{
		Addr:           endPoint,
		Handler:        routerRes.routerGroup,
		ReadTimeout:    time.Duration(readTimeout),
		WriteTimeout:   time.Duration(writeTimeout),
		MaxHeaderBytes: maxHeaderBytes,
	}
	srv = s
	ginpprof.Wrapper(routerRes.routerGroup)
	go func() {
		// service connections
		log.Info("HTTP服务器启动中...",
			zap.String("address", endPoint),
			zap.Duration("read_timeout", time.Duration(readTimeout)),
			zap.Duration("write_timeout", time.Duration(writeTimeout)))

		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("HTTP服务器启动失败",
				zap.Error(err),
				zap.String("address", endPoint),
				zap.String("component", "http_server"))
			// 不使用panic，而是记录错误并让主程序处理
			os.Exit(1)
		}
	}()
	return
}

func route() apiroutes {
	r := apiroutes{
		routerGroup: gin.New(),
	}

	// 添加全局中间件
	r.routerGroup.Use(gin.Recovery())
	r.routerGroup.Use(middleware.RequestID())
	r.routerGroup.Use(middleware.ErrorHandler())

	r.routerGroup.NoRoute(func(c *gin.Context) {
		c.Error(errors.ConfigResourceNotFound(c.Request.URL.Path, nil))
	})
	r.routerGroup.NoMethod(func(c *gin.Context) {
		c.Error(errors.NewConfigError("api_errors", "method_not_allowed", nil, nil, nil))
	})

	// Swagger 文档
	r.routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由分组
	v1 := r.routerGroup.Group("/api/v1", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
		"user":  "user",
	}))
	r.addFocalboardApi(v1)

	return r
}
