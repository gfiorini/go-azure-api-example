package config

import "github.com/spf13/viper"

type Config struct {
	MongoURI                string `mapstructure:"MONGOURI"`
	MongoDbAdminDatabase    string `mapstructure:"MONGO_DB_ADMIN"`
	MongoDbScoreDatabase    string `mapstructure:"MONGO_DB_SCORE_DATABASE"`
	MongoDbAlbumsCollection string `mapstructure:"MONGO_DB_ALBUMS_COLLECTION"`
	MongoDbScoresCollection string `mapstructure:"MONGO_DB_SCORES_COLLECTION"`

	ServerPort string `mapstructure:"FUNCTIONS_CUSTOMHANDLER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
