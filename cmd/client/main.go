package main

import (
	"context"
	"flag"
	"github.com/Hagaz0/LinkShortener/pkg/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Введите флаг post или get, а так же ссылку или короткую ссылку соответственно")
	}
	if flag.Arg(0) != "get" && flag.Arg(0) != "post" {
		log.Fatal("Введен неккоректный флаг. Флаги: get, post")
	}
	f := flag.Arg(0)
	link := flag.Arg(1)
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewLinkShorterClient(conn)
	var res *api.AddResponse
	if f == "post" {
		res, err = c.Post(context.Background(), &api.AddRequest{Link: link})
	} else if f == "get" {
		res, err = c.Get(context.Background(), &api.AddRequest{Link: link})
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetResult())
}
