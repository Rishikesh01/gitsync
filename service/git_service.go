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

func NewGitService(sync chan dto.SyncGit) GitService {
	return &gitService{sync: sync}
}

type gitService struct {
	sync chan dto.SyncGit
}

func (g *gitService) CloneRepo() {
	for val := range g.sync {
		clone := exec.Command("git", fmt.Sprintf("clone %s %s", val.GitLink, val.ParentDir))
		_, err := clone.Output()
		if err != nil {
			log.Println(err)
		}
	}
}
