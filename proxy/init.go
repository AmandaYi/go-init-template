package proxy

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"reflect"
	"yy-price-api/models"
)

var (
	User *userProxy
	mysqlPrefix string
)

//链接数据库
func init()  {
	mysqlPrefix = os.Getenv("MYSQL_PREFIX")
	// 增加表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return mysqlPrefix + defaultTableName
	}
	mysqlConfig := os.Getenv("MYSQL_HOST")
	session, err := gorm.Open("mysql",mysqlConfig)
	if err != nil {
		panic(err)
		return
	}

	// 不允许在表名后面增加s
	session.SingularTable(true)
	User = &userProxy{baseProxy{session.Table(mysqlPrefix + "user"), reflect.TypeOf((*models.User)(nil)).Elem()}}
	session.AutoMigrate(&models.User{}) // 同步用户表
}

//迁移方案
func installSyncTable(session *gorm.DB){

}
