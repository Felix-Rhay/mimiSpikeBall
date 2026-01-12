package main

import (
	"fmt"
	"log"
	"mimiSpikeBall/internal/db"
	"mimiSpikeBall/internal/handlers"
	"mimiSpikeBall/internal/repository"
	"mimiSpikeBall/internal/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, configError := db.LoadConfig()
	if configError != nil {
		log.Fatal(configError)
	}

	app, err := buildApp(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8090", app))
}

func buildApp(config *db.Config) (*http.ServeMux, error) {
	dbConn, err := db.NewMySQLFromURI(config.DatabaseURL)

	if err != nil {
		return nil, fmt.Errorf("erreur en essayant d'établir la connexion à la base de données: ", err)
	}

	log.Println("Connexion BD OK")

	//initialisation des services
	compteService := services.NewCompteService(repository.NewCompteRepository(dbConn))
	compteHandler := handlers.NewCompteHandler(compteService)

	//handle les appels http
	http.HandleFunc("POST /compte/ObtenirComptes", compteHandler.ObtenirComptes)
	http.HandleFunc("POST /compte/LoginUser", compteHandler.LoginUser)

	return http.DefaultServeMux, nil
}

func handlePOST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}
