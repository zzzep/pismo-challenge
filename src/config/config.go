package config

import "fmt"

const dbUser = "giuseppe"
const dbPassword = "1q!Q"
const dbHost = "10.30.1.22"
const dbPort = "3306"
const dbName = "pismo"
const dbCharset = "utf8mb4"
const dbTemplate = "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"

func GetDatabaseConnection() string {
	return fmt.Sprintf(dbTemplate, dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)
}
