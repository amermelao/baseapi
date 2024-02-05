package conf

import (
	"fmt"
	"truckconf"
)

type ValidJWT struct {
	truckconf.EnableConf

	Port int    `mapstructure:"port"`
	Url  string `mapstructure:"url"`
}

func (v ValidJWT) Valid() bool {

	return true
}

func (v ValidJWT) ValidUrl() string {
	return fmt.Sprintf("%s:%d/valid", v.Url, v.Port)
}
