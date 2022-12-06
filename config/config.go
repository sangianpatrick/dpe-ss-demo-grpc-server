package config

import (
	"fmt"
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
		Name string
		Port int
	}

	LogFormatter logrus.Formatter
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

	return c
}

func (c *Config) application() {

	name := os.Getenv("APP_NAME")
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	c.Application.Name = name
	c.Application.Port = port
}

func (c *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	c.LogFormatter = formatter
}
