package main

import (
	"fmt"
	pb "github.com/itswcg/micro-demo/user-srv/proto/user"
	"github.com/micro/go-micro"
	"log"
)

const (
	user     = ""
	password = ""
	host     = ""
	port     = ""
	dbName   = ""
)

func main() {
	db, err := CreateConnection(user, password, host, port, dbName)
	defer db.Close()

	if err != nil {
		log.Fatalf("Cound not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenHandler{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()
	_ = pb.RegisterUserServiceHandler(srv.Server(), &handler{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
