package initialize

import (
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/global"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/errors"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/spf13/viper"
)

func LoadConfig() *setting.Config {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		errors.Must(err, "Error loading configuration")
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// configure structur
	var config setting.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	// save global
	global.Config = config
	return &config
}
