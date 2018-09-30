package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewSqlHandler() SqlHandler {
	db, err := gorm.Open("mysql", "root:password@tcp(0.0.0.0:3306)/tsk?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return *sqlHandler
}
