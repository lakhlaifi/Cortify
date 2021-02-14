package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// Configuration struct for settings
type Configuration struct {
	// Logging
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	// Database
	DBAddrs string `json:"dbAddrs"`
	DBName  string `json:"dbName"`
	DBUser  string `json:"dbUser"`
	DBPwd   string `json:"dbPwd"`

	// Tokens
	// AuthAddr  string `json:"authAddr"`
	// JWTSecret string `json:"jwtSecret"`
	// Isuser    string `json:"isuser"`
}

// Config share global config
var (
	Config *Configuration
)

// Status Text
const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

// Status Code
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000
	StatusMismatch    = 10
)

// LoadConfig : Load config from config file
func LoadConfig() error {
	// Path to the JSON config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// Log Formatter &TextFormatter{}
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}