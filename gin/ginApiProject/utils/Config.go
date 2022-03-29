package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	Database DatabaseConfig `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  string `json:"show_sql"`
}

// LoadConfig 从配置文件中载入json字符串
func LoadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}

	mainConfig := &Config{}
	err = json.Unmarshal(buf, mainConfig)

	if err != nil {
		return nil, err
	}

	return mainConfig, nil
}

var path = "/Users/liaojiayun/Desktop/gogo_learn/gin/ginApiProject/config/app.json"

func GetConfig() *Config {
	//pwd, _ := os.Getwd() // 获取到当前目录，相当于python里的os.Getcwd()
	//
	////文件路径拼接
	//path := filepath.Join(pwd, "/gin/ginApiProject/config/app.json")

	config, err := LoadConfig(path)

	fmt.Println("config: ", config)
	fmt.Println("path: ", path)

	if err != nil {
		panic("获取配置文件失败！")
		return nil
	}

	return config
}
