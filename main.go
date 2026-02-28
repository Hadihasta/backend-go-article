package main

import (
	"backend-article-portal/database"
	"backend-article-portal/handlers"
	"backend-article-portal/middlewares"
	"backend-article-portal/repositories"
	"backend-article-portal/services"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
	APiKey string `mapstructure:"API_KEY"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
		APiKey: viper.GetString("API_KEY"),
	}


	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	
	apiKeyMiddleware := middlewares.APIKEY(config.APiKey)
	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	http.HandleFunc("/api/v1/article", middlewares.CORS(middlewares.Logger(postHandler.HandlePosts)))
	http.HandleFunc("/api/v1/article/", middlewares.CORS(middlewares.Logger(postHandler.HandlePostByID)))
	

	http.HandleFunc("/api/v1/health",apiKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Content-Type", "application/json")
		
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	}),
	)

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
