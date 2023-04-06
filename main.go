package main

import (
	"github.com/Rishikesh01/gitsync/dto"
	"github.com/Rishikesh01/gitsync/service"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type env struct {
	BaseUrl             string `required:"true" split_words:"true"`
	SyncProjectEndpoint string `required:"true" split_words:"true"`
	AddProjectEndpoint  string `required:"true" split_words:"true"`
}

func main() {
	var env env
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}
	sync := make(chan dto.SyncGit, 1000)
	client := new(http.Client)
	_, err := service.NewCommunicationService(&service.CommunicationConfig{
		Client:              client,
		SyncProjectEndpoint: env.SyncProjectEndpoint,
		AddProjectEndpoint:  env.AddProjectEndpoint,
		OutSync:             sync,
	})
	if err != nil {
		return
	}

	service.NewGitService(sync, nil)
}
