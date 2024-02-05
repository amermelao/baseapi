package conf

import (
	"encoding/json"
	"log/slog"

	"github.com/amermelao/simpleconf"
)

type AppConf struct {
	Server Server
}

func NewAppConfiguratoin() AppConf {

	var conf AppConf

	if err := simpleconf.ReadConfiguration(
		simpleconf.OneSubset(&conf.Server),
	); err != nil {
		slog.Error("read conf", slog.String("error", err.Error()))
	}

	return conf
}

func (c AppConf) String() string {
	a, _ := json.MarshalIndent(c, "", " ")
	return string(a)
}
