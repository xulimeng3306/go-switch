package features

import (
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/xulimeng/go-switch/config"
)

const Exit = "exit"

func Switch() {
	versions := []string{}
	if config.Conf.LocalGos == nil {
		config.Conf.LocalGos = make([]config.GosVersion, 0)
	}
	for _, vInfo := range config.Conf.LocalGos {
		versions = append(versions, vInfo.Version)
	}

	versions = append(versions, Exit)
	prompt := promptui.Select{
		Label: "Choose You Want Switch Version",
		Items: versions,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	if result == Exit {
		return
	}

	goRootPath := filepath.Join(config.GosPath, result)

	// TODO: support windows
	if config.SystemEnv == config.Windows {
		UpdateGoEnvWindows(goRootPath)
		return
	} else if config.SystemEnv == config.Mac || config.SystemEnv == config.Linux {
		UpdateGoEnvUnix(goRootPath)
	}
	config.Conf.GoRoot = goRootPath
	config.Conf.SaveConfig()
}
