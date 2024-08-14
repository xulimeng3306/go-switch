package features

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xulimeng/go-switch/config"
	"github.com/xulimeng/go-switch/utils"
)

// UpdateGoEnvUnix 更新 Unix 系统的环境变量
func UpdateGoEnvUnix(goRoot string) {
	// set GOROOT
	sh := JudgeZshOrBash()
	goRootCmd := fmt.Sprintf("export GOROOT=%s", goRoot)
	pathCmd := "export PATH=$PATH:$GOROOT/bin"
	if config.GoEnvFilePath != "" {
		if err := utils.TruncateFile(config.ConnectPathWithEnv(config.SystemEnv, config.GoEnvFilePath, []string{"system"})); err != nil {
			panic(err)
		}
		addEnvironmentVariable(config.GoEnvFilePath, goRootCmd)
		addEnvironmentVariable(config.GoEnvFilePath, pathCmd)
	}
	if !config.Conf.Init {
		var configFile string
		switch sh {
		case "zsh":
			configFile := os.Getenv("HOME") + "/.zshrc"
			if err := reloadZshCOnfig("zsh", configFile); err != nil {
				fmt.Printf("Failed to reload zsh config: %v\n", err)
				panic(err)
			}
		case "bash":
			configFile := os.Getenv("HOME") + "/.bashrc"
			if config.SystemEnv == config.Mac {
				configFile = os.Getenv("HOME") + "/.bash_profile"
			}
			if err := reloadZshCOnfig("bash", configFile); err != nil {
				fmt.Printf("Failed to reload bash config: %v\n", err)
				panic(err)
			}
		default:
			fmt.Println("Not support shell")
		}
		if configFile != "" && config.GoEnvFilePath != "" {
			addEnvironmentVariable(configFile, fmt.Sprintf("source %s", config.GoEnvFilePath))
		}
	}

}

// addEnvironmentVariable 添加环境变量
func addEnvironmentVariable(configFile, line string) {
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", configFile, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == line {
			found = true
			break
		}
	}

	if !found {
		if _, err := file.WriteString(line + "\n"); err != nil {
			fmt.Printf("Failed to write to %s: %v\n", configFile, err)
		} else {
			fmt.Printf("Added '%s' to %s\n", line, configFile)
		}
	} else {
		fmt.Printf("Line '%s' already exists in %s\n", line, configFile)
	}
}

func reloadZshCOnfig(shCmd string, shPath string) error {
	cmd := exec.Command(shCmd, "-c", fmt.Sprintf("source %s", shPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// JudgeZshOrBash 判断当前 shell 类型
func JudgeZshOrBash() string {
	// 获取 SHELL 环境变量
	shell := os.Getenv("SHELL")
	if shell == "" {
		fmt.Println("SHELL environment variable is not set")
		return ""
	}

	currentShell := ""
	shellSplit := strings.Split(shell, "/")
	if len(shellSplit) > 0 {
		currentShell = shellSplit[len(shellSplit)-1]
	}
	// 根据 shell 类型执行不同操作
	if strings.Contains(currentShell, "zsh") {
		return "zsh"
	} else if strings.Contains(currentShell, "bash") {
		return "bash"
	}
	return ""
}
