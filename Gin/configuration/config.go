package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

var path string

type Config struct {
	ConnectionString string
	Port             int
	LogToFile        bool
	LogPath          string
}

func SetPath(_path string) {
	fmt.Println("config path ", _path)
	path = _path
}

func GetConfiguration() Config {
	fmt.Printf("Loading configuration @ %v\n", path)
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("ERROR: Cannot deserialize the file due to %v", err)
		return Config{}
	}

	port := viper.Get("Port")
	connectionString := viper.Get("ConnectionString")
	fmt.Println("port and connection string ", port, connectionString)

	err = viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("ERROR: Cannot deserialize the file due to %v", err)
		return Config{}
	}
	fmt.Printf("config finally es  %v", config)
	return config
}
