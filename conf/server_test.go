package conf

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"truckconf"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	viper.SetConfigType("TOML") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
[server]
port = 8081
allow-origins = ["www.me.com", "localhost:3000"]
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var conf Server
	truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf),
	)

	assert.Equal(t, 8081, conf.Port)
	assert.Equal(t, "www.me.com:localhost:3000", strings.Join(conf.AllowOrigins, ":"))
	assert.Equal(t, ":8081", fmt.Sprint(conf))
}
