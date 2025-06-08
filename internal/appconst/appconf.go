package appconst

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	HttpPort              = 8080
	AppName               = "focalboard-tool"
	AppVersion            = "0.0.1"
	DefaultConfigFileName = "focalboard-tool-config.toml"
	BasePathV1            = "/api/v1"
	ParamLengthLimit      = 5
)

const (
	CardStatusPropName   = "status"
	CardAsleadIDPropName = "aslead"
)

const (
	IsWindows = (runtime.GOOS == "windows")
)

const (
	ConfigPath  = "."
	ConfigPath2 = "./config"
	ConfigPath3 = "/etc/" + AppName + "/conf.d/"
	ConfigPath4 = "/etc/" + AppName
	ConfigPath5 = "c:/" + AppName
	ConfigPath6 = "d:/" + AppName
)

func CurrentOsConfigPath() []string {
	configPaths := []string{}
	configPaths = append(configPaths, ConfigPath, ConfigPath2, ConfigPath3, ConfigPath4)
	if IsWindows {
		configPaths = append(configPaths, ConfigPath5, ConfigPath6)
	}
	return configPaths
}
func ConfigPathGenerator(fileName *string, expath *string) (combinedPath *string, err error) {
	configPaths := []string{}
	if len(*expath) > 0 {
		configPaths = append(configPaths, *expath)
	}
	configPaths = append(configPaths, CurrentOsConfigPath()...)
	for _, path := range configPaths {
		fullPath := filepath.Join(path, *fileName)
		if fileExists(fullPath) && (combinedPath == nil || len(*combinedPath) == 0) {
			combinedPath = &fullPath

		}
	}
	if combinedPath == nil || len(*combinedPath) == 0 {
		err = fmt.Errorf("config file not found in %v", configPaths)
		return nil, err
	} else {
		return
	}

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}
