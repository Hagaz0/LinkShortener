package main

import (
	"flag"
	"github.com/Hagaz0/LinkShortener/pkg/api"
	"github.com/Hagaz0/LinkShortener/pkg/memory"
	"github.com/Hagaz0/LinkShortener/pkg/src"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Выберите используемое хранилище (postgres, in_memory)")
	}
	if flag.Arg(0) != "postgres" && flag.Arg(0) != "in_memory" {
		log.Fatal("Введено неккоректное значение хранилища. Доступные хранилища: postgres, in_memory")
	}
	memory.Flag = flag.Arg(0)
	s := grpc.NewServer()
	srv := &src.GRPCServer{}
	api.RegisterLinkShorterServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
