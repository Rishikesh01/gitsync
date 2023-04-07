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
	gitService GitService) (CmdlineService, error) {
	if logger == nil || com == nil || gitService == nil {
		return nil, errors.New("one of the params is nil")
	}
	return &cmdService{logger: logger, com: com, gitService: gitService}, nil
}

type cmdService struct {
	logger     *log.Logger
	com        CommunicationService
	gitService GitService
}

// gitsync add
func (c *cmdService) Args(args string) {
	validArgs := []string{"sync", "add", "sync && add"}

	switch args {
	case validArgs[0]:
		c.com.Sync()
		break
	case validArgs[1]:
		c.interActiveCommand()
	case validArgs[2]:
		c.com.Sync()
		c.interActiveCommand()
	default:
		fmt.Println("wrong args")
	}

}

func (c *cmdService) interActiveCommand() {
	var projects []dto.Project
	fmt.Print("Enter the following details")
	var exit string
	for exit != "y" {
		project := dto.Project{}
		fmt.Print("Project name:")
		fmt.Scan(&project.ProjectName)
		fmt.Println("Github link:")
		fmt.Scan(&project.GithubLink)
		fmt.Println("Is the project private on github:")
		fmt.Scan(&project.IsActive)
		fmt.Println("Directory path of project src:")
		fmt.Scan(&project.ParentDir)
		fmt.Print("Is the project currently active:")
		fmt.Scanln(&project.IsActive)
		fmt.Println("Are you doing adding projects[y/n]:")
		fmt.Scan(&exit)
		projects = append(projects, project)
	}

	c.com.AddNewGitRepos(projects)
}
