// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package config

import (
	"log"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	CER      *CERConfig
	Database *DatabaseConfig
	Logging  *LoggingConfig
	Website  *WebsiteConfig
}

type CERConfig struct {
	Host     string
	Password string
	Username string
	Refresh  int
}

type DatabaseConfig struct {
	AutoMigrate bool
	Database    string
	Driver      string
	Host        string
	Limit       uint32
	Password    string
	Path        string
	Port        int
	Username    string
	SSL         string
}

type LoggingConfig struct {
	Compress bool
	Level    string
	MaxAge   uint32
	MaxSize  uint32
	Name     string
	Path     string
}

type WebsiteConfig struct {
	Port       int
	Title      string
	SucessURL  string
	FailureURL string
}

func SetDefaults() {
	// Set defaults for the LoggingConfig
	viper.SetDefault("logging.compress", true)
	viper.SetDefault("logging.level", "debug")
	viper.SetDefault("logging.maxAge", 7)
	viper.SetDefault("logging.maxSize", 10)
	viper.SetDefault("logging.name", "go-cer.log")
	viper.SetDefault("logging.path", "./go-cer/log")

	// Set defaults for the DatabaseConfig
	viper.SetDefault("database.autoMigrate", true)
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.path", "./go-cer.db")
	viper.SetDefault("database.limit", 100)

	// Set defaults for the CERConfig
	viper.SetDefault("cer.refresh", 60)

	// Set defaults for the WebsiteConfig
	viper.SetDefault("website.port", 8080)
	viper.SetDefault("website.title", "Go-CER")
	viper.SetDefault("website.successURL", "https://www.google.com")
	viper.SetDefault("website.failureURL", "https://www.google.com")

}

func GetLoggerFromGlobalConfig() *LoggingConfig {
	loggerConfig := viper.Sub("logging")
	if loggerConfig == nil {
		log.Fatalf("No log settings found in config file")
		return nil
	}
	return &LoggingConfig{
		Compress: loggerConfig.GetBool("compress"),
		Level:    loggerConfig.GetString("level"),
		MaxAge:   loggerConfig.GetUint32("maxAge"),
		MaxSize:  loggerConfig.GetUint32("maxSize"),
		Name:     loggerConfig.GetString("name"),
		Path:     loggerConfig.GetString("path"),
	}
}

func GetCERFromGlobalConfig() *CERConfig {
	cerConfig := viper.Sub("cer")
	if cerConfig == nil {
		log.Fatalf("No CER settings found in config file")
		return nil
	}
	return &CERConfig{
		Host:     cerConfig.GetString("host"),
		Password: cerConfig.GetString("password"),
		Refresh:  cerConfig.GetInt("refresh"),
		Username: cerConfig.GetString("username"),
	}
}

func GetDatabaseFromGlobalConfig() *DatabaseConfig {
	databaseConfig := viper.Sub("database")
	if databaseConfig == nil {
		log.Fatalf("No database settings found in config file")
		return nil
	}
	return &DatabaseConfig{
		AutoMigrate: databaseConfig.GetBool("autoMigrate"),
		Database:    databaseConfig.GetString("database"),
		Driver:      databaseConfig.GetString("driver"),
		Host:        databaseConfig.GetString("host"),
		Limit:       databaseConfig.GetUint32("limit"),
		Password:    databaseConfig.GetString("password"),
		Path:        databaseConfig.GetString("path"),
		Port:        databaseConfig.GetInt("port"),
		Username:    databaseConfig.GetString("username"),
	}
}

func GetWebsiteFromGlobalConfig() *WebsiteConfig {
	websiteConfig := viper.Sub("website")
	if websiteConfig == nil {
		log.Fatalf("No website settings found in config file")
		return nil
	}
	return &WebsiteConfig{
		Port:       websiteConfig.GetInt("port"),
		SucessURL:  websiteConfig.GetString("sucessUrl"),
		FailureURL: websiteConfig.GetString("failureUrl"),
	}
}

func GetGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		CER:      GetCERFromGlobalConfig(),
		Database: GetDatabaseFromGlobalConfig(),
		Logging:  GetLoggerFromGlobalConfig(),
		Website:  GetWebsiteFromGlobalConfig(),
	}
}
