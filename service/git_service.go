package service

import (
	"errors"
	"fmt"
	"github.com/Rishikesh01/gitsync/dto"
	"log"
	"os/exec"
)

type GitService interface {
	CloneRepo()
}

type GitConfig struct {
}

func NewGitService(InSync chan dto.SyncGit, Logger *log.Logger) (GitService, error) {
	if InSync == nil || Logger == nil {
		return nil, errors.New("one of the parameters in empty")
	}
	return &gitService{inSync: InSync, logger: Logger}, nil
}

type gitService struct {
	inSync chan dto.SyncGit
	logger *log.Logger
}

func (g *gitService) CloneRepo() {
	for val := range g.inSync {
		clone := exec.Command("git", fmt.Sprintf("clone %s %s", val.GitLink, val.ParentDir))
		_, err := clone.Output()
		if err != nil {
			g.logger.Println(err)
		}
	}
}
