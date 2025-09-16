package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/meghraj/online-test-backend/internal/auth"
	"github.com/meghraj/online-test-backend/internal/db"
	"github.com/meghraj/online-test-backend/internal/graph/generated"
	"github.com/meghraj/online-test-backend/internal/graph/resolvers"
	"github.com/meghraj/online-test-backend/internal/models"
	"gorm.io/gorm"
)

func autoMigrate(d *gorm.DB) {
	if err := d.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Automigrate : ", err)
	}
}

func main() {
	d := db.Connect()
	defer db.Close(d)
	autoMigrate(d)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:5175", "http://localhost:5176"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	r.Use(auth.Middleware)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &resolvers.Resolver{DB: d}},
	))
	r.Handle("/query", srv)
	r.Handle("/", playground.Handler("GraphQL", "/query"))

	log.Println("GraphQL running on : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
