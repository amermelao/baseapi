package conf

import (
	"bytes"
	"errors"
	"testing"
	"truckconf"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	viper.SetConfigType("TOML") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
[db]
user = "postgress"
password = "docker"
port = 5432
host = "localhost"
name = "me"
ssl = true
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var conf DB
	truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf),
	)

	assert.Equal(t, 5432, conf.Port)
	assert.Equal(t, "localhost", conf.Host)
	assert.Equal(t, "docker", conf.Password)
	assert.Equal(t, "postgress", conf.User)
	assert.True(t, conf.SSL)
}

func TestDBInvalid(t *testing.T) {
	viper.Reset()
	viper.SetConfigType("TOML") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
[db]
password = "docker"
port = 5432
host = "localhost"
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var conf DB
	err := truckconf.ReadConfiguration(
		truckconf.OneSubset(&conf),
	)

	assert.True(t, errors.Is(err, truckconf.ErrInvalidConfiguration))
	assert.ErrorIs(t, err, truckconf.ErrInvalidConfiguration)

}
