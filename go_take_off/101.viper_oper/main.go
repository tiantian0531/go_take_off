package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Cfg struct {
	Addr string
	No   string
}

func main() {
	// WriteConfigFile()
	//ReadConfigFile()
	//MonitorConfigFile()
	// AutoEnv()
	ReadCmdLine()
}

// 基本写入读取操作
func WriteBasicUse() {

	viper.SetDefault("name", "yuxl")
	viper.Set("age", 26)

	name := viper.GetString("name")
	age := viper.GetInt("age")
	fmt.Println("配置名字是：", name)
	fmt.Println("配置年龄是：", age)

	viper.SetDefault("age", 18)
	fmt.Println("配置年龄是：", viper.GetInt("age"))
}

// 将配置写入文件
func WriteConfigFile() {

	viper.Set("info", Cfg{Addr: "湖北武汉", No: "001"})
	viper.Set("name", "yuxl")
	viper.SetConfigType("json")
	// 文件存在就不写入
	err := viper.SafeWriteConfigAs("config.json")

	if err != nil {
		fmt.Println("写入json配置错误", err)
	} else {
		fmt.Println("写入json配置成功")
	}

	viper.SetConfigType("yaml")
	// 文件存在就不写入
	yamlErr := viper.SafeWriteConfigAs("config.yaml")
	if yamlErr != nil {
		fmt.Println("写入yaml配置错误", yamlErr)
	} else {
		fmt.Println("写入yaml配置成功")
	}

	viper.SetConfigType("toml")
	// 文件存在就不写入
	tomlErr := viper.SafeWriteConfigAs("config.toml")
	if tomlErr != nil {
		fmt.Println("写入tomlErr配置错误", tomlErr)
	} else {
		fmt.Println("写入toml配置成功")
	}
}

// 读取配置
func ReadConfigFile() {
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")

	readErr := viper.ReadInConfig()
	if readErr != nil {
		fmt.Println("读取配置文件失败", readErr)
		return
	}

	info := viper.Get("info")
	name := viper.Get("name")

	fmt.Println("读取内容为", info)
	fmt.Println("读取内容为", name)
}

// 监听配置
func MonitorConfigFile() {

	// 先读取旧的值
	ov := viper.New()
	ov.SetConfigType("json")
	ov.SetConfigFile("config.json")
	ov.ReadInConfig()

	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变更", e.Name, e.Op)
		fmt.Println("配置文件发生变更前的值是:", ov.Get("info"))
		fmt.Println("配置文件发生变更后的值是:", viper.Get("info"))
	})

	fmt.Println("正在监听....")
	select {}
}

// 读取

func AutoEnv() {
	viper.AutomaticEnv()

	// 读取全部
	viper.SetEnvPrefix("")

	ACSvcPort := viper.Get("ACSvcPort")
	USERNAME := viper.Get("USERNAME")
	fmt.Println("读取到系统环境变量为：", ACSvcPort, USERNAME)

	viper.SetEnvPrefix("ENV")
	fmt.Println(viper.Get("version"))
}

func ReadCmdLine() {
	var ip string
	var port int
	pflag.StringVar(&ip, "ip", "", "程序运行的ip")
	pflag.IntVar(&port, "port", 0, "程序运行的端口")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	fmt.Println("程序运行读取ip为", viper.Get("ip"))
	fmt.Println("程序运行读取端口为", viper.Get("port"))

	// pflag 等于前面加--
	//go run main.go --ip 192.168.0.1 --port 5020
	// go run main.go -h        查看所有参数

}
