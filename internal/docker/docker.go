package docker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// CheckDockerInstallation 检查宿主机是否安装了Docker，并且是否有权限操作Docker
func CheckDockerInstallation() (bool, error) {
	cmd := exec.Command("docker", "version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("Docker is not installed or you don't have permission to run Docker: %v", err)
	}
	return true, nil
}

// InstallDocker 安装Docker
func InstallDocker() error {
	cmd := exec.Command("sh", "-c", "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to install Docker: %v", err)
	}
	return nil
}

// RunDockerContainer 运行Docker容器
func RunDockerContainer(image string, ports map[string]string) (bool, error) {
	portMappings := ""
	for hostPort, containerPort := range ports {
		portMappings += fmt.Sprintf("-p %s:%s ", hostPort, containerPort)
	}
	cmdStr := fmt.Sprintf("docker run -d %s %s", portMappings, image)
	cmd := exec.Command("sh", "-c", cmdStr)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("Failed to run Docker container: %v", err)
	}
	containerID := strings.TrimSpace(out.String())
	fmt.Println("Container ID:", containerID)
	if containerID == "" {
		return false, fmt.Errorf("Failed to get container ID")
	}
	return true, nil
}
