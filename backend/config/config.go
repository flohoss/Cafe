package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/unjx-de/go-api-auth"
	"github.com/unjx-de/go-database"
	"github.com/unjx-de/go-folder"
	"reflect"
	"strings"
)

const StorageDir = "storage/"
const TemplatesDir = "templates/"

type Config struct {
	Port         int
	AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"`
	Swagger      bool
	Bookmarks    bool
	LogLevel     string         `mapstructure:"LOG_LEVEL"`
	Database     database.MySQL `mapstructure:"MYSQL"`
	Auth         auth.Auth
}

var C = Config{}

func readConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithField("error", err).Fatal("Failed opening config file")
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&C, viper.DecodeHook(passwordHookFunc()))
	if err != nil {
		logrus.WithField("error", err).Fatal("Failed reading environment variables")
	}
	logrus.WithField("file", viper.ConfigFileUsed()).Info("Initializing configuration")
}

func passwordHookFunc() mapstructure.DecodeHookFuncType {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(C.Auth.Password) {
			raw := data.(string)
			if raw == "" {
				return [32]byte{}, nil
			}
			return auth.HashPassword(raw), nil
		} else if t == reflect.TypeOf(C.Auth.Secret) {
			raw := data.(string)
			if raw == "" {
				return auth.SecretGenerator(), nil
			}
			return []byte(raw), nil
		} else if t == reflect.TypeOf(C.AllowedHosts) {
			raw := data.(string)
			if raw == "" {
				return []string{}, nil
			}
			return strings.Split(raw, ","), nil
		}
		return data, nil
	}
}

func configLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
}

func setLogLevel() {
	logLevel, err := logrus.ParseLevel(C.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logLevel)
	}
	logrus.WithField("logLevel", logLevel.String()).Debug("Log level set")
}

func createFolderStructure() {
	folders := []string{StorageDir, TemplatesDir}
	err := folder.CreateFolders(folders, 0755)
	if err != nil {
		logrus.WithField("error", err).Fatal("Failed creating folders")
	}
	logrus.WithField("folders", folders).Debug("Folders created")
}

func init() {
	configLogger()
	readConfig()
	setLogLevel()
	createFolderStructure()
}
