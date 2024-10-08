package config

import (
	"os"
)

type Env string

const (
	// Linux, Windows, Mac 系统环境类型
	Linux   Env = "linux"
	Windows Env = "windows"
	Mac     Env = "drawin" // mac os
)

var (

	// go-switch 的文件夹名
	GoSwitchDir = ".go-switch"
	// 真正保存go 版本的文件夹名
	SaveGoDir = "gos"
	// golang 压缩包解压之后的文件夹名称
	UnzipGoDir = "go"

	// 不同系统默认的 go 安装路径
	LinuxGoPath   = os.Getenv("HOME")
	WindowsGoPath = os.Getenv("USERPROFILE")
	MacGoPath     = os.Getenv("HOME")

	SystemEnv     Env
	SystemArch    string
	RootPath      string
	GosPath       string
	TempUnzipPath string
	GoEnvFilePath string
)
