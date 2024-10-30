package main

import (
	"anograph/graph"
	"anograph/internal/coba"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	// Konfigurasi koneksi MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/cobagraph?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Auto migrate untuk memastikan tabel dibuat
	if err := db.AutoMigrate(&coba.Coba{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Inisialisasi repository
	cobaRepo := coba.NewCobaRepository(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Pastikan repository terpasang ke dalam Resolver
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			CobaRepo: cobaRepo,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
