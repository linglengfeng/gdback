package mysqldb

import (
	"fmt"
	"gdback/config"
	"gdback/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

const Successed = "r1"
const Failed = "r2"

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
}

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

// func RegisterInfo() ([]byte, error) {
// 	var results []Register
// 	query := "select * from register"
// 	DB.Raw(query).Scan(&results)
// 	return json.Marshal(results)
// }

func RegisterInfo() [](map[string]any) {
	// var results []Register
	var results [](map[string]any)
	query := "select * from register"
	DB.Raw(query).Scan(&results)
	return results
}

func UserInsert(account, password string) bool {
	results := Successed
	query := "call sp_user_insert(?, ?)"
	tx := DB.Raw(query, account, password).Scan(&results)
	logger.Debug("UserInsert, account:%v, password:%v, result:%v, tx:%v", account, password, results, tx)
	return results == Successed
}

func UserLogin(account, password string) bool {
	results := Successed
	query := "call sp_user_login(?, ?)"
	tx := DB.Raw(query, account, password).Scan(&results)
	logger.Debug("UserLogin, account:%v, password:%v, result:%v, tx:%v", account, password, results, tx)
	return results == Successed
}

func UserIsExist(account string) bool {
	results := 0
	query := "select count(*) from user where account = ?"
	tx := DB.Raw(query, account).Scan(&results)
	logger.Debug("UserIsExist, account:%v, result:%v, tx:%v", account, results, tx)
	return results == 1
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
