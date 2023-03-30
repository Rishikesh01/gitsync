package service

import (
	"encoding/json"
	"gitsync/dto"
	"log"
	"net/http"
)

type SyncService interface {
	Sync()
}

func NewSyncService(client *http.Client, req *http.Request, sync chan dto.SyncGit) SyncService {
	return &syncService{client: client, req: req, sync: sync}
}

type syncService struct {
	client *http.Client
	req    *http.Request
	sync   chan dto.SyncGit
}

func (s *syncService) Sync() {
	do, err := s.client.Do(s.req)
	if err != nil {
		log.Println(err)
	}
	var projects []dto.Project
	defer do.Body.Close()
	if err := json.NewDecoder(do.Body).Decode(&projects); err != nil {
		log.Println(err)
	}

	for _, val := range projects {
		s.sync <- dto.SyncGit{
			GitLink:   val.GithubLink,
			ParentDir: val.ParentDir,
		}
	}

}
