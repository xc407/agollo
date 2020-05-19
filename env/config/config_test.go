package config

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/tevid/gohamcrest"
	"github.com/xc407/agollo/v3/utils"
)

var (
	appConfig = getTestAppConfig()
)

func getTestAppConfig() *AppConfig {
	jsonStr := `{
    "appId": "test",
    "cluster": "dev",
    "namespaceName": "application",
    "ip": "localhost:8888",
    "releaseKey": "1"
	}`
	c, _ := Unmarshal([]byte(jsonStr))

	return c.(*AppConfig)
}

func TestGetIsBackupConfig(t *testing.T) {
	config := appConfig.GetIsBackupConfig()
	Assert(t, config, Equal(true))
}

func TestGetBackupConfigPath(t *testing.T) {
	config := appConfig.GetBackupConfigPath()
	Assert(t, config, Equal("/app/"))
}

func TestSetNextTryConnTime(t *testing.T) {
	appConfig.SetNextTryConnTime(10)

	Assert(t, int(appConfig.NextTryConnTime), GreaterThan(int(time.Now().Unix())))
}

func Unmarshal(b []byte) (interface{}, error) {
	appConfig := &AppConfig{
		Cluster:          "default",
		NamespaceName:    "application",
		IsBackupConfig:   true,
		BackupConfigPath: "/app/",
	}
	err := json.Unmarshal(b, appConfig)
	if utils.IsNotNil(err) {
		return nil, err
	}

	return appConfig, nil
}

func TestGetHost(t *testing.T) {
	ip := appConfig.IP
	host := appConfig.GetHost()
	Assert(t, host, Equal("http://localhost:8888/"))

	appConfig.IP = "http://baidu.com"
	host = appConfig.GetHost()
	Assert(t, host, Equal("http://baidu.com/"))

	appConfig.IP = "http://163.com/"
	host = appConfig.GetHost()
	Assert(t, host, Equal("http://163.com/"))

	appConfig.IP = ip
}

func TestAppConfig_IsConnectDirectly(t *testing.T) {
	backup := appConfig.NextTryConnTime

	appConfig.NextTryConnTime = 0
	isConnectDirectly := appConfig.IsConnectDirectly()
	Assert(t, isConnectDirectly, Equal(false))

	appConfig.NextTryConnTime = time.Now().Unix() + 10
	isConnectDirectly = appConfig.IsConnectDirectly()
	Assert(t, isConnectDirectly, Equal(true))

	appConfig.NextTryConnTime = backup
}
