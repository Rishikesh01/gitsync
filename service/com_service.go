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

func (c *CommunicationConfig) validate() error {
	if c.Client == nil || c.SyncProjectEndpoint == "" || c.AddProjectEndpoint == "" || c.OutSync == nil || c.Logger == nil {
		return errors.New("one of the value is nil")
	}
	return nil
}

func NewCommunicationService(com *CommunicationConfig) (CommunicationService, error) {
	if err := com.validate(); err != nil {
		return nil, err
	}
	return &comService{com}, nil
}

type comService struct {
	*CommunicationConfig
}

func (s *comService) AddNewGitRepos(project []dto.Project) {
	arr, err := json.Marshal(project)
	if err != nil {
		s.Logger.Println("ERROR:", err)
	}
	req, err := http.NewRequest(http.MethodPost, s.AddProjectEndpoint, bytes.NewBuffer(arr))
	if err != nil {
		s.Logger.Println("ERROR:", err)
	}

	do, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Println("ERROR:", err)
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		s.Logger.Println("ERROR:", do.StatusCode)
	}
}

func (s *comService) Sync() {
	req, err := http.NewRequest(http.MethodPost, s.AddProjectEndpoint, nil)
	if err != nil {
		s.Logger.Println("ERROR:", err)
	}

	do, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Println("ERROR:", err)
	}
	var projects []dto.Project
	defer do.Body.Close()
	if err := json.NewDecoder(do.Body).Decode(&projects); err != nil {
		s.Logger.Println(err)
	}
	if do.StatusCode != http.StatusOK {
		s.Logger.Println("ERROR:", do.StatusCode)
	}
	for _, val := range projects {
		s.OutSync <- dto.SyncGit{
			GitLink:   val.GithubLink,
			ParentDir: val.ParentDir,
		}
	}

}
