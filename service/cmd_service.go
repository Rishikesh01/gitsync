package service

import (
	"errors"
	"fmt"
	"github.com/Rishikesh01/gitsync/dto"
	"log"
)

type CmdlineService interface {
	Args(string)
}

func NewCmdlineService(logger *log.Logger, com CommunicationService,
	gitService GitService, ioService YamlService) (CmdlineService, error) {
	if logger == nil || com == nil || gitService == nil || ioService == nil {
		return nil, errors.New("one of the params is nil")
	}
	return &cmdService{logger: logger, com: com, gitService: gitService, ioService: ioService}, nil
}

type cmdService struct {
	logger     *log.Logger
	com        CommunicationService
	gitService GitService
	ioService  YamlService
}

// gitsync add
func (c *cmdService) Args(args string) {
	validArgs := []string{"sync", "add", "sync && add"}

	switch args {
	case validArgs[0]:
		c.com.Sync()
		break
	case validArgs[1]:
		val, err := c.ioService.CheckYamlInCWD()
		if err != nil {
			log.Println(err)
		}
		c.com.AddNewGitRepos([]dto.Project{val})
		break
	case validArgs[2]:
		c.com.Sync()
		val, err := c.ioService.CheckYamlInCWD()
		if err != nil {
			log.Println(err)
		}
		c.com.AddNewGitRepos([]dto.Project{val})
		break
	default:
		fmt.Println("wrong args")
	}

}
