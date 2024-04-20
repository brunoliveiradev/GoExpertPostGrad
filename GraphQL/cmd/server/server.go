package main

import (
	"database/sql"
	"github.com/brunoliveiradev/GoExpertPostGrad/GraphQL/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brunoliveiradev/GoExpertPostGrad/GraphQL/graph"
)

const defaultPort = "8080"

func main() {
	db := InitSQLiteDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: database.NewCategory(db),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func InitSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT, description TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS courses (id TEXT PRIMARY KEY, name TEXT, description TEXT, category_id TEXT, FOREIGN KEY(category_id) REFERENCES categories(id))")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
