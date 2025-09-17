package main

import (
	"fmt"
	"os"
)

/**เช็คว่ามี DB_HOST มีค่าอะไรหรือไม่*/

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host: %s \nport: %s\nuser: %s\npassword:%s\ndbname: %s", host, port, user, password, name)
	fmt.Println(conSt)
}
