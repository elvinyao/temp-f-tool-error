package conf

import (
	"fmt"
	"focalboard-tool/internal/appconst"
	"focalboard-tool/library/log"
	xhttp "focalboard-tool/library/net/apiclient"
	xtime "focalboard-tool/library/time"
	"focalboard-tool/pkg/errors"
	systemlog "log"

	"github.com/BurntSushi/toml"
)

var (
	ConfPath *string
	Conf     = initConfigSetting()
)

type Config struct {
	// App
	App  *AppSetting
	Auth *AuthSetting
	//Boardschema
	BoardSchema      *BoardSchemaSetting
	Userkintaistatus *UserkintaistatusSetting

	// Log
	Log *log.LogConfig
	// HttpClient
	HttpClient *HTTPClientSetting
	// Error
	Error *ErrorSetting
}

func initConfigSetting() *Config {

	return &Config{
		App:         initAppSetting(),
		BoardSchema: initBoardSchemaSetting(),
		Error:       initErrorSetting(),
	}
}

func initAppSetting() *AppSetting {
	return &AppSetting{
		HttpPort: appconst.HttpPort,
		Version:  appconst.AppVersion,
		AppName:  appconst.AppName,
	}
}

type AppSetting struct {
	HttpPort             int
	Version              string
	AppName              string
	RunMode              string
	CardStatusPropName   string
	CardAsleadIdPropName string
	ReadTimeout          xtime.Duration
	WriteTimeout         xtime.Duration
}
type AuthSetting struct {
	Username string
	Password string
}

type UserkintaistatusSetting struct {
	Name  string
	Emoji string
}

func initBoardSchemaSetting() *BoardSchemaSetting {
	return &BoardSchemaSetting{
		CardStatusPropName:   appconst.CardStatusPropName,
		CardAsleadIDPropName: appconst.CardAsleadIDPropName,
	}
}

type HTTPClientSetting struct {
	FocalboardClient *xhttp.ClientConfig
	MattermostClient *xhttp.ClientConfig
}

type BoardSchemaSetting struct {
	CardStatusPropName        string
	CardAsleadIDPropName      string
	CardGroupCategoryPropName string
	Props                     []string
}

func (bss *BoardSchemaSetting) AddProp() {
	if len(bss.CardAsleadIDPropName) > 0 {
		bss.Props = append(bss.Props, bss.CardAsleadIDPropName)
	}
	if len(bss.CardStatusPropName) > 0 {
		bss.Props = append(bss.Props, bss.CardStatusPropName)
	}
	if len(bss.CardGroupCategoryPropName) > 0 {
		bss.Props = append(bss.Props, bss.CardGroupCategoryPropName)
	}

}

func Init(fileName *string, extrapath *string) (err error) {
	// 加载主配置文件
	ConfPath, err = appconst.ConfigPathGenerator(fileName, extrapath)

	if err != nil {
		pathelems := appconst.CurrentOsConfigPath()
		orderedPathElems := []string{}
		orderedPathElems = append(orderedPathElems, *extrapath)
		orderedPathElems = append(orderedPathElems, pathelems...)
		return fmt.Errorf("generate config path failed: %w, search path %v", err, orderedPathElems)
	}
	_, err = toml.DecodeFile(*ConfPath, &Conf)
	if _, err = toml.DecodeFile(*ConfPath, &Conf); err != nil {
		return fmt.Errorf("decode config file: %w", err)
	}
	if RunWithoutAuth() {
		systemlog.Println("Run without auth")
	}
	Conf.BoardSchema.AddProp()

	// 加载错误配置文件
	if Conf.Error != nil && Conf.Error.ConfigFile != "" {
		errorConfigPath := Conf.Error.ConfigFile
		err = errors.LoadErrorConfig(errorConfigPath)
		if err != nil {
			panic(err)
		}
	}

	return
}

func RunWithoutAuth() bool {
	return len(Conf.Auth.Username) == 0 && len(Conf.Auth.Password) == 0
}

func initErrorSetting() *ErrorSetting {
	return &ErrorSetting{
		ConfigFile: "configs/errors.yaml",
	}
}

type ErrorSetting struct {
	ConfigFile string
}

// GetSystemID returns the system ID (appName) from configuration
// func GetSystemID() string {
// 	if Conf.App != nil {
// 		return Conf.App.AppName
// 	}
// 	return appconst.AppName // fallback to default if config not loaded
// }
