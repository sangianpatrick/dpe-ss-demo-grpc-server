package config

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Application struct {
		Location *time.Location
		Name     string
		Port     int
	}

	LogFormatter logrus.Formatter

	Mariadb struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
		DSN      string
	}
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = newConfig()
	}

	return config
}

func newConfig() *Config {
	c := new(Config)
	c.application()
	c.logFormatter()
	c.mariadb()

	return c
}

func (c *Config) application() {
	location, _ := time.LoadLocation("Asia/Jakarta")

	loc, err := time.LoadLocation(os.Getenv("APP_LOCATION"))
	if err == nil {
		location = loc
	}

	name := os.Getenv("APP_NAME")
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	c.Application.Location = location
	c.Application.Name = name
	c.Application.Port = port
}

func (c *Config) mariadb() {
	host := os.Getenv("MARIADB_HOST")
	port, _ := strconv.Atoi(os.Getenv("MARIADB_PORT"))
	username := os.Getenv("MARIADB_USERNAME")
	password := os.Getenv("MARIADB_PASSWORD")
	database := os.Getenv("MARIADB_DATABASE")

	urlValues := url.Values{}
	urlValues.Add("parseTime", "true")
	urlValues.Add("loc", c.Application.Location.String())
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, database, urlValues.Encode())

	c.Mariadb.Host = host
	c.Mariadb.Port = port
	c.Mariadb.Username = username
	c.Mariadb.Password = password
	c.Mariadb.Database = database
	c.Mariadb.DSN = dsn
}

func (c *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	c.LogFormatter = formatter
}
