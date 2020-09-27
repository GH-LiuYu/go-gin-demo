package initialize
//该文件作用-读取yaml配置为结构体赋值
import (
	"demo/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
const defaultConfigFile = "config.yaml"//指定配置文件路径
func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.DEMO_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.DEMO_CONFIG); err != nil {//Unmarshal 读取yaml 文件为结构体对应字段赋值
		fmt.Println(err)
	}
	global.DEMO_VP = v
}
