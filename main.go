package main

import (
	"backend-article-portal/database"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)


type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

	if _, err := os.Stat(".env"); err == nil{
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err:= database.InitDB(config.DBConn)
	if err != nil{
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	port := config.Port
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Println("server running on", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}