package config

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/xulimeng/go-switch/utils"
)

func LoadConfig() {
	if Conf == nil {
		Conf = &Config{}
	}
	mt, err := toml.DecodeFile(RootPath+"/config/config.toml", Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(mt)
	if Conf.GoSwitchPath == "" && RootPath != "" {
		Conf.GoSwitchPath = RootPath
	}
}

func InitConfigFile() {
	_, err := os.Stat(RootPath + "/config/config.toml")
	if os.IsExist(err) {
		return
	}
	if exists, create := utils.ExistsPath(RootPath + "/config/"); !exists && !create {
		panic("RootPath not exists")
	}

	_, err = os.Create(RootPath + "/config/config.toml")
	if err != nil {
		panic(err)
	}
}

func InitSystemVars() {

	os := runtime.GOOS
	switch os {
	case "linux":
		SystemEnv = Linux
		RootPath = LinuxGoPath + GoSwitchDir
		GosPath = RootPath + "/" + SaveGoDir
		TempUnzipPath = GosPath + "/" + UnzipGoDir
	case "windows":
		SystemEnv = Windows
		userNameCurr, err := user.Current()
		if err != nil {
			panic(err)
		}
		RootPath = WindowsGoPath + userNameCurr.Username + "\\" + GoSwitchDir
		GosPath = RootPath + "\\" + SaveGoDir
		TempUnzipPath = GosPath + "\\" + UnzipGoDir
	case "darwin":
		SystemEnv = Mac
		RootPath = MacGoPath + GoSwitchDir
		GosPath = RootPath + "/" + SaveGoDir
		TempUnzipPath = GosPath + "/" + UnzipGoDir
	}
	SystemArch = runtime.GOARCH
}
