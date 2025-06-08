package main

import (
	"context"
	"fmt"
	"focalboard-tool/internal/appconst"
	"focalboard-tool/internal/conf"
	"focalboard-tool/internal/server/http"
	"focalboard-tool/internal/service"
	"focalboard-tool/library/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"go.uber.org/zap"
)

var (
	ConfigFile          = kingpin.Flag("config", "Configuration name").Default("application.toml").Short('c').String()
	ExtraConfigFilePath = kingpin.Flag("expath", "Extra path to configuration file").Short('e').String()
)

func main() {
	kingpin.Parse()

	// 配置初始化 - 改进错误信息
	if err := conf.Init(ConfigFile, ExtraConfigFilePath); err != nil {
		fmt.Fprintf(os.Stderr, "配置初始化失败: %v\n", err)
		fmt.Fprintf(os.Stderr, "请检查配置文件路径和格式是否正确\n")
		os.Exit(1)
	}

	// 日志初始化 - 改进错误处理
	if err := log.Init(conf.Conf.Log); err != nil {
		fmt.Fprintf(os.Stderr, "日志系统初始化失败: %v\n", err)
		fmt.Fprintf(os.Stderr, "将使用标准输出进行日志记录\n")
		// 不退出，继续运行但使用标准输出
	}

	// 验证关键配置
	if err := validateCriticalConfig(); err != nil {
		log.Error("关键配置验证失败", zap.Error(err))
		os.Exit(1)
	}

	srv := service.New(conf.Conf)
	httpSrv := http.New(conf.Conf, srv)

	// 改进启动成功信息
	infomsg := fmt.Sprintf(appconst.AppName+" 启动成功 - 端口: %d, 模式: %s, 版本: %s, 配置: %s",
		conf.Conf.App.HttpPort, conf.Conf.App.RunMode, conf.Conf.App.Version, *conf.ConfPath)
	log.Info(infomsg)

	// 优雅关闭处理
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info(fmt.Sprintf("接收到信号: %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("开始优雅关闭应用...")

			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("HTTP服务器关闭失败",
					zap.Error(err),
					zap.String("timeout", "35s"))
			} else {
				log.Info("HTTP服务器已优雅关闭")
			}

			httpSrv.Close()
			cancel()

			log.Info(appconst.AppName + " 已安全退出")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
			log.Info("接收到SIGHUP信号，忽略处理")
		default:
			log.Warn("接收到未知信号", zap.String("signal", s.String()))
			return
		}
	}
}

// validateCriticalConfig 验证关键配置项
func validateCriticalConfig() error {
	if conf.Conf.HttpClient == nil {
		return fmt.Errorf("HTTP客户端配置缺失")
	}

	if conf.Conf.HttpClient.FocalboardClient == nil {
		return fmt.Errorf("Focalboard客户端配置缺失")
	}

	if conf.Conf.App.HttpPort <= 0 || conf.Conf.App.HttpPort > 65535 {
		return fmt.Errorf("HTTP端口配置无效: %d", conf.Conf.App.HttpPort)
	}

	return nil
}
