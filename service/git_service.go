package service

import (
	"fmt"
	"github.com/Rishikesh01/gitsync/dto"
	"log"
	"os/exec"
)

type GitService interface {
	CloneRepo()
}

type GitConfig struct {
	InSync chan dto.SyncGit
}

func NewGitService(config *GitConfig) GitService {
	return &gitService{config}
}

type gitService struct {
	*GitConfig
}

func (g *gitService) CloneRepo() {
	for val := range g.InSync {
		clone := exec.Command("git", fmt.Sprintf("clone %s %s", val.GitLink, val.ParentDir))
		_, err := clone.Output()
		if err != nil {
			log.Println(err)
		}
	}
}
