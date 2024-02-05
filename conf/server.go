package conf

import (
	"fmt"
	"truckconf"
)

type Server struct {
	truckconf.EnableConf

	Port         int      `mapstructure:"port"`
	AllowOrigins []string `mapstructure:"allow-origins"`
}

func (Server) Valid() bool {
	return true
}

func (s Server) String() string {
	return fmt.Sprintf(":%d", s.Port)
}
