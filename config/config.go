package config

import (
	"awesomeProject/enums"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// RegularTokenLifetime refers to token lifetime of regular.
var RegularTokenLifetime string

// CTLTokenLifetime refers to token lifetime of ctl.
var CTLTokenLifetime string

// PrivateKey refers to rsa private key .
var PrivateKey string

// Publickey refers to rsa public key .
var Publickey string

// RunMode refers to run mode.
var RunMode string

// ServerPort refers to server port.
var ServerPort string

// DatabaseConnectionString refers to database connection string.
var DatabaseConnectionString string

// Database refers to database options.
var Database string

// DbUsername refers to database name.
var DbUsername string

// DbPassword refers to database password.
var DbPassword string

// DbServer refers to database server ip.
var DbServer string

// DbPort refers to database server port.
var DbPort string

// DatabaseName refers to database name.
var DatabaseName string

func InitEnv() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = string(enums.DEVELOP)
	}

	if RunMode != string(enums.PRODUCTION) {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)

	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database = os.Getenv("DATABASE")
	PrivateKey = os.Getenv("PRIVATE_KEY")
	Publickey = os.Getenv("PUBLIC_KEY")
	RegularTokenLifetime = os.Getenv("REGULAR_TOKEN_LIFETIME")
	CTLTokenLifetime = os.Getenv("CTL_TOKEN_LIFETIME")

	if Database == enums.MONGO {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort + "/?authSource=admin&readPreference=primary&directConnection=true&ssl=false"
	}
}
