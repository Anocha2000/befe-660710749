package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" /*Library*/
)

/**เช็คว่ามี DB_HOST มีค่าอะไรหรือไม่*/

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

/*Create ตัวแปร db ชี้ไปที่ sql.DB*/
var db *sql.DB

func initDB() {
	var err error
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host: %s \nport: %s\nuser: %s\npassword:%s\ndbname: %s sslmode=disable", host, port, user, password, name)
	//fmt.Println(conSt)

	db, err := sql.Open("postgres", conSt)

	if err != nil {
		log.Fatal("Failed to open database")
		/*Print Log แล้ว ออกโปรแกรมเลย ตรง terminal*/
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect database")
		/*เช็คว่าเชื่อม database ได้ไหม*/
	}
	log.Println("succesfully to connece to database")

}

func main() {
	initDB()

}
