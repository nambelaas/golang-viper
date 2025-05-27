package belajargolangviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJson(t *testing.T){
	config:= viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Belajar Golang Viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "Salman", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.showsql"))
	assert.Equal(t, 5432, config.GetInt("database.port"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
}

func TestYaml(t *testing.T){
	config:= viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Belajar Golang Viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "Salman", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.showsql"))
	assert.Equal(t, 5432, config.GetInt("database.port"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
}

func TestEnv(t *testing.T){
	config:= viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "Salman", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOWSQL"))
	assert.Equal(t, 5432, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
}

func TestENV(t *testing.T){
	config:= viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "Salman", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOWSQL"))
	assert.Equal(t, 5432, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}