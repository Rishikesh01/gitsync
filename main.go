package main

import (
	"github.com/Rishikesh01/gitsync/dto"
	"github.com/Rishikesh01/gitsync/service"
	"log"
	"net/http"
)

func main() {
	sync := make(chan dto.SyncGit, 1000)
	client := new(http.Client)
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		log.Fatal(req)
	}
	service.NewSyncService(client, req, sync)
	service.NewGitService(sync)
}
