package main

import (
	"fmt"
	"github.com/Rishikesh01/gitsync/dto"
	"github.com/Rishikesh01/gitsync/service"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"strings"
)

type env struct {
	SyncProjectEndpoint string `required:"true" split_words:"true"`
	AddProjectEndpoint  string `required:"true" split_words:"true"`
}

func main() {
	var env env
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}
	logger := log.Default()
	sync := make(chan dto.SyncGit, 1000)
	client := new(http.Client)
	com, err := service.NewCommunicationService(&service.CommunicationConfig{
		Client:              client,
		SyncProjectEndpoint: env.SyncProjectEndpoint,
		AddProjectEndpoint:  env.AddProjectEndpoint,
		OutSync:             sync,
		Logger:              logger,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	gitService, err := service.NewGitService(sync, logger)
	if err != nil {
		log.Fatal(err)
		return
	}
	cmd, err := service.NewCmdlineService(logger, com, gitService)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(os.Args[1:])
	cmd.Args(strings.Join(os.Args[1:], " "))

}
