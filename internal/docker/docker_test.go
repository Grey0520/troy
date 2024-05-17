package docker

import (
	"errors"
	"os/exec"
	"testing"
)

var execCommand = exec.Command

func TestCheckDockerInstallation(t *testing.T) {
	// 测试 Docker 已安装且有权限运行的情况
	t.Run("Docker is installed and has permission", func(t *testing.T) {
		isInstalled, err := CheckDockerInstallation()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !isInstalled {
			t.Error("Expected Docker to be installed, but it's not")
		}
	})

	// 模拟 Docker 未安装或没有权限运行的情况
	t.Run("Docker is not installed or no permission", func(t *testing.T) {
		// 保存原始的 execCommand 函数
		originalExecCommand := execCommand
		defer func() { execCommand = originalExecCommand }()

		// 定义一个模拟的 execCommand 函数
		mockExecCommand := func(name string, arg ...string) *exec.Cmd {
			return &exec.Cmd{
				Path: "dummy",
				Err:  errors.New("simulated error"),
			}
		}

		// 将模拟函数赋值给 execCommand
		execCommand = mockExecCommand

		isInstalled, err := CheckDockerInstallation()
		if err == nil {
			t.Error("Expected an error, but got none")
		}
		if isInstalled {
			t.Error("Expected Docker to not be installed, but it is")
		}
	})
}
