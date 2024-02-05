package conf

import (
	"bytes"
	"testing"
	"truckconf"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestValidJWT(t *testing.T) {
	viper.SetConfigType("TOML") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
[validjwt]
port = 8081
url = "localhost:3000"
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var conf ValidJWT
	truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf),
	)

	assert.Equal(t, 8081, conf.Port)
	assert.Equal(t, "localhost:3000", conf.Url)
	assert.True(t, conf.Enabled(), "should be enabled")
}

func TestJWTDisable(t *testing.T) {
	viper.SetConfigType("TOML") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
[validjwt]
port = 8081
url = "localhost:3000"
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var conf ValidJWT
	truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf),
	)

	assert.Equal(t, 8081, conf.Port)
	assert.Equal(t, "localhost:3000", conf.Url)
}
