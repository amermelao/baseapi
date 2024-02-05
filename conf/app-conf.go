package conf

import (
	"encoding/json"
	"log/slog"
	"truckconf"
)

type AppConf struct {
	Server   Server
	ValidJWT ValidJWT
	DB       DB
}

func NewAppConfiguratoin() AppConf {

	var conf AppConf

	if err := truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf.DB),
		truckconf.OneSubset(&conf.Server),
		truckconf.OneSubset(&conf.ValidJWT),
	); err != nil {
		slog.Error("read conf", slog.String("error", err.Error()))
	}

	return conf
}

func (c AppConf) String() string {
	a, _ := json.MarshalIndent(c, "", " ")
	return string(a)
}
