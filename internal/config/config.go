package configs

import "github.com/spf13/viper"

// Инициализация конфига
func InitConfig() error {
	viper.AddConfigPath("internal/config/") // Директория конфига
	viper.SetConfigName("config")           // Название файла
	viper.SetConfigType("yaml")             // Рассширение файла
	return viper.ReadInConfig()
}
