package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "record_code"
)

var DB *gorm.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname) // 连接数据库
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("database connection failed:", err)
		return
	}

	fmt.Println("okokokoko")

	DB = db
}
