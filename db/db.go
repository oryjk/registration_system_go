package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB = nil

func init() {
	var err error
	dsn := "root:beifa888@tcp(49.234.55.170:3306)/registration_system?charset=utf8mb4&parseTime=True&loc=Local" // 更新为实际的数据库信息
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}

type MyBool bool

func (mb *MyBool) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan MyBool: %v", value)
	}
	*mb = b[0] == 1
	return nil
}
