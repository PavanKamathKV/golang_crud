// database connection
package utility

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// const (
// 	host     = "localhost"
// 	port     = 3306
// 	user     = "root"
// 	password = "Kamath@123"
// 	dbname   = "golang_crud"
// )

func GetConnection() *gorm.DB {
	msqlInfo := "root:Kamath@123@tcp(127.0.0.1:3306)/golang_crud?charset=utf8mb4&parseTime=True&loc=Local"
	// msqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	db, err := gorm.Open("mysql", msqlInfo)

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")

	return db

}
