package src

import (
	"context"
	"fmt"
	"github.com/Hagaz0/LinkShortener/pkg/api"
	"github.com/Hagaz0/LinkShortener/pkg/memory"
	"log"
)

type GRPCServer struct{}

func (s *GRPCServer) Post(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	link := req.GetLink()
	if !IsValidUrl(link) {
		return nil, fmt.Errorf("Введен неккоректный URL")
	}
	var shortLink string
	if memory.Flag == "postgres" {
		db := DB{}
		err := db.NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}
		short, err := db.GetShortLink(link)
		if err == nil && short != "" {
			return &api.AddResponse{Result: short}, nil
		}
		shortLink = Shorting()
		err = db.InsertNewLink(link, shortLink)
		if err != nil {
			log.Fatal(err)
		}
	} else if memory.Flag == "in_memory" {
		if short, ok := memory.InMemoryLinkShort[link]; ok {
			return &api.AddResponse{Result: short}, nil
		}
		shortLink = Shorting()
		memory.InMemoryLinkShort[link] = shortLink
		memory.InMemoryShortLink[shortLink] = link
	}
	return &api.AddResponse{Result: shortLink}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	link := req.GetLink()
	if memory.Flag == "postgres" {
		db := DB{}
		err := db.NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}
		original, err := db.GetOriginalLink(link)
		if err != nil {
			log.Fatal(err)
		} else if original == "" {
			return nil, fmt.Errorf("Ссылка отсутствует")
		}
		return &api.AddResponse{Result: original}, nil
	} else if memory.Flag == "in_memory" {
		if original, ok := memory.InMemoryShortLink[link]; ok {
			return &api.AddResponse{Result: original}, nil
		} else {
			return nil, fmt.Errorf("Ссылка отсутствует")
		}
	}
	return nil, nil
}
