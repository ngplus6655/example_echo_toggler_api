package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal(".envファイルの読み込みに失敗しました")
	}

	connString := os.Getenv("DB_USERNAME")+":"+
		os.Getenv("DB_PASSWORD")+
		"@tcp("+ os.Getenv("DB_HOST") +":"+
		os.Getenv("DB_PORT")+")/"+
		os.Getenv("DB_DATABASE")+
		"?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err = gorm.Open("mysql", connString)
	if err != nil {
		logrus.Fatal(err)
	}
	return db, err
}