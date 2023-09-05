package mysqldb

import (
	"encoding/json"
	"fmt"
	"gdback/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Register struct {
	Id            uint64
	Uid           uint64
	Remote        string
	Ip            string
	Imei          string
	Os            string
	Model         string
	App_id        string
	Channel_id    string
	Register_at   uint64
	Register_type int
} // insert into register values(2,2,"remote","ip","imei","os","model","app_id","channel_id",1,1);

func Start() {
	// 设置数据库连接参数
	// dsn := "username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	username := config.Config.GetString("mysql.user")
	password := config.Config.GetString("mysql.password")
	ip := config.Config.GetString("mysql.ip")
	port := config.Config.GetString("mysql.port")
	database := config.Config.GetString("mysql.db")
	dsn := username + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	mysqlconfig := &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}
	db, err := gorm.Open(mysql.Open(dsn), mysqlconfig)

	if err != nil {
		panic("无法连接到数据库：" + err.Error() + dsn)
	}
	DB = db
}

func RegisterInfo() ([]byte, error) {
	var results []Register
	query := "select * from register"
	DB.Raw(query).Scan(&results)
	return json.Marshal(results)
}

func Test() {
	aaa := DB.Exec("SHOW DATABASES")
	var results1 []Register
	DB.Find(&results1)
	for _, user := range results1 {
		fmt.Println("uuuuuuuuuuu results1 find", user)
	}

	var results []Register
	query := "select * from register"
	DB.Raw(query).Scan(&results)
	for _, user := range results {
		fmt.Println("uuuuuuuuuuu results Raw", user)
	}

	var results3 []Register
	query1 := "select * from register where id = ?"
	conditionValue := 1
	DB.Raw(query1, conditionValue).Scan(&results3)
	for _, user := range results3 {
		fmt.Println("uuuuuuuuuuu results Raw3", user)
	}

	fmt.Println("mysql Test, aaa:", aaa)
}
