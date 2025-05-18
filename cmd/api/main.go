package main

import (
	"log"
	"net"
	"os"

	categorypbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/category/v1"
	"github.com/Ostap00034/course-work-backend-category-service/db"
	category "github.com/Ostap00034/course-work-backend-category-service/internal"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbString, exists := os.LookupEnv("CATEGORY_DB_CONN_STRING")
	if !exists {
		log.Fatal("not CATEGORY_DB_CONN_STRING in .env file")
	}
	client := db.NewClient(dbString)
	defer client.Close()

	repo := category.NewRepo(client)
	svc := category.NewService(repo)
	srv := category.NewServer(svc)

	lis, _ := net.Listen("tcp", ":50053")
	s := grpc.NewServer()
	categorypbv1.RegisterCategoryServiceServer(s, srv)

	log.Println("CategoryService on :50053")
	log.Fatal(s.Serve(lis))
}
