package main

import (
	"fmt"
	"log"

	"github.com/juiiyang/troy/internal/config"
	"github.com/juiiyang/troy/internal/docker"
)

func main() {
	// 读配置，优先级：环境变量 > 配置文件 > 默认值
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Printf("Failed to load config: %v", err)
	}

	log.Printf("Loaded Config: %+v", cfg)

	installed, err := docker.CheckDockerInstallation()
	if err != nil {
		fmt.Println("Error:", err)
	} else if installed {
		fmt.Println("Docker is installed and running.")
		fmt.Println("running capgrey/caddy-trojan")
		image := "capgrey/caddy-trojan"
		ports := map[string]string{
			"80":  "80",
			"443": "443",
		}
		success, err := docker.RunDockerContainer(image, ports)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if success {
			fmt.Println("Docker container started successfully")
		} else {
			fmt.Println("Failed to start Docker container")
		}
	} else {
		fmt.Println("Docker is not installed or not running.")

	}
}
