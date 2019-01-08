package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexeyKremsa/CustomRedis/config"
	"github.com/AlexeyKremsa/CustomRedis/storage"
)

type server struct {
	storage *storage.Storage
}

func ServerStart(storage *storage.Storage, cfg *config.Config) {
	cr := &server{
		storage: storage,
	}
	router := newRouter(cr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}
