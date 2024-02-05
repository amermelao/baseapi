package conf

import (
	"fmt"
	"truckconf"
)

type DB struct {
	truckconf.EnableConf

	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	SSL      bool   `mapstructure:"ssl"`
}

func (db DB) Valid() bool {
	switch {
	case db.Host == "":
		return false
	case db.User == "":
		return false
	case db.Password == "":
		return false
	case db.Name == "":
		return false
	case db.Port == 0:
		return false
	}
	return true
}

func (db DB) ConnectionString() string {
	sslMode := ""
	if !db.SSL {
		sslMode = "?sslmode=disable"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		sslMode,
	)
}
