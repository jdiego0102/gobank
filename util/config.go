package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config almacena toda la configuración de la aplicación.
// Los valores son leídos por viper desde el archivo de configuración o variables de entorno.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymetricKey    string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig lee la configuración desde el archivo o variables de entorno.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
