package main

import (
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
	com, err := service.NewCommunicationService(client, env.SyncProjectEndpoint, env.AddProjectEndpoint, sync, logger)
	if err != nil {
		log.Fatal(err)
		return
	}

	gitService, err := service.NewGitService(sync, logger)
	if err != nil {
		log.Fatal(err)
		return
	}
	ioService := service.NewYamlService()
	cmd, err := service.NewCmdlineService(logger, com, gitService, ioService)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Args(strings.Join(os.Args[1:], " "))

}
