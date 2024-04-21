package main

import (
	"database/sql"
	"github.com/brunoliveiradev/GoExpertPostGrad/gRPC/internal/database"
	"github.com/brunoliveiradev/GoExpertPostGrad/gRPC/internal/pb"
	"github.com/brunoliveiradev/GoExpertPostGrad/gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const defaultPort = "50051"

func main() {
	db := InitSQLiteDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	log.Printf("Server listening at http://localhost/%v", port)
	log.Fatal(grpcServer.Serve(lis))
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
