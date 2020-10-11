package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/**
 *
 */
var Db *sql.DB
func ConnerDB(){
	fmt.Println("连接数据库")
	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"

	db, err := sql.Open(dbDriver, connUrl)
	if err != nil { // err不为nil，表示数据连接时出现了错误，程序就在此中断就可以，
		//早解决，早解决
		panic("数据库连接错误")
	}
	//为全局赋值
	Db = db
}