package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Rishikesh01/gitsync/dto"
	"log"
	"net/http"
)

type CommunicationService interface {
	AddNewGitRepos(projects []dto.Project)
	Sync()
}

type CommunicationConfig struct {
	Client              *http.Client
	SyncProjectEndpoint string
	AddProjectEndpoint  string
	OutSync             chan dto.SyncGit
	Logger              *log.Logger
}

func NewCommunicationService(client *http.Client,
	syncProjectEndpoint string,
	addProjectEndpoint string,
	outSync chan dto.SyncGit,
	logger *log.Logger) (CommunicationService, error) {

	if client == nil || syncProjectEndpoint == "" || addProjectEndpoint == "" || outSync == nil || logger == nil {
		return nil, errors.New("one of the value is nil")
	}
	return &comService{
		client:              client,
		syncProjectEndpoint: syncProjectEndpoint,
		addProjectEndpoint:  addProjectEndpoint,
		outSync:             outSync,
		logger:              logger,
	}, nil
}

type comService struct {
	client              *http.Client
	syncProjectEndpoint string
	addProjectEndpoint  string
	outSync             chan dto.SyncGit
	logger              *log.Logger
}

func (s *comService) AddNewGitRepos(project []dto.Project) {
	arr, err := json.Marshal(project)
	if err != nil {
		s.logger.Println("ERROR:", err)
	}
	req, err := http.NewRequest(http.MethodPost, s.addProjectEndpoint, bytes.NewBuffer(arr))
	if err != nil {
		s.logger.Println("ERROR:", err)
	}

	do, err := s.client.Do(req)
	if err != nil {
		s.logger.Println("ERROR:", err)
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		s.logger.Println("ERROR:", do.StatusCode)
	}
}

func (s *comService) Sync() {
	req, err := http.NewRequest(http.MethodPost, s.addProjectEndpoint, nil)
	if err != nil {
		s.logger.Println("ERROR:", err)
	}

	do, err := s.client.Do(req)
	if err != nil {
		s.logger.Println("ERROR:", err)
	}
	var projects []dto.Project
	defer do.Body.Close()
	if err := json.NewDecoder(do.Body).Decode(&projects); err != nil {
		s.logger.Println(err)
	}
	if do.StatusCode != http.StatusOK {
		s.logger.Println("ERROR:", do.StatusCode)
	}
	for _, val := range projects {
		s.outSync <- dto.SyncGit{
			GitLink:   val.GithubLink,
			ParentDir: val.ParentDir,
		}
	}

}
