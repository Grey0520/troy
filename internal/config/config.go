package config

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// Config 结构体定义了应用的所有配置项
type Config struct {
	IP        string
	Domain    string
	Passwords []string
}

// LoadConfig 函数用于加载配置文件
func LoadConfig(path string) (*Config, error) {
	// 设置配置文件路径
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	// 设置环境变量前缀
	viper.SetEnvPrefix("TROY")

	// 设置键名替换规则
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 自动绑定环境变量
	viper.AutomaticEnv()

	// 获取配置值，优先从环境变量读取
	cfg := &Config{
		IP:        viper.GetString("ip"),
		Domain:    viper.GetString("domain"),
		Passwords: viper.GetStringSlice("passwords"),
	}

	return cfg, nil
}

// getPublicIPv4 函数用于获取公网的 IPv4 地址
func getPublicIPv4() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get public IP")
	}

	var result struct {
		IP string `json:"ip"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.IP, nil
}

func init() {
	// 获取公网的 IPv4 地址
	ip, err := getPublicIPv4()
	if err != nil {
		log.Fatalf("Failed to get local IPv4 address: %v", err)
	}

	// 设置默认值
	viper.SetDefault("ip", ip)
	viper.SetDefault("domain", ip+".nip.io")
	viper.SetDefault("passwords", []string{"123456"})
}
