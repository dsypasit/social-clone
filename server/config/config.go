package config

import (
	"os"
	"strconv"
)

type env func(key string) string

type cfg struct {
	getEnv env
}

func New() *cfg {
	return &cfg{getEnv: os.Getenv}
}

type Config struct {
	Server       Server
	DBConnection string
}

type Server struct {
	Hostname string
	Port     int
}

const (
	cHostname = "HOSTNAME"
	cPort     = "PORT"

	cDBConnection = "DATABASE_URL"
)

const (
	dPort         = 1323
	dDBConnection = "postgresql://postgres:password@localhost:5432/banking?sslmode=disable"
)

func (c *cfg) All() Config {
	return Config{
		Server: Server{
			Hostname: c.envString(cHostname, ""),
			Port:     c.envInt(cPort, dPort),
		},
		DBConnection: c.envString(cDBConnection, dDBConnection),
	}
}

func (c *cfg) SetEnvGetter(overrideEnvGetter env) {
	c.getEnv = overrideEnvGetter
}

func (c *cfg) envString(key, defaultValue string) string {
	val := c.getEnv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (c *cfg) envInt(key string, defaultValue int) int {
	v := c.getEnv(key)

	val, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return val
}

func (c *cfg) envBool(key string, defaultValue bool) bool {
	v := c.getEnv(key)

	val, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return val
}
