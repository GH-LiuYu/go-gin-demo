package initialize
//该文件作用-为建立数据库连接
import (
	"demo/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Mysql()  {
	admin :=global.DEMO_CONFIG.Mysql
	db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@tcp("+admin.Path+")/"+admin.Dbname+"?"+admin.Config)
	if  err != nil {
		fmt.Println("打印错误:", err,admin)
		global.DEMO_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.DEMO_DB = db
		global.DEMO_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.DEMO_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.DEMO_DB.LogMode(admin.LogMode)
	}
}
