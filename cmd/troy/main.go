package main

import (
	"log"

	"github.com/juiiyang/troy/internal/config"
)

func main() {
	// 读配置，优先级：环境变量 > 配置文件 > 默认值
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Printf("Failed to load config: %v", err)
	}

	log.Printf("Loaded Config: %+v", cfg)
}
