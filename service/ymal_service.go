package service

import (
	"github.com/Rishikesh01/gitsync/dto"
	"gopkg.in/yaml.v2"
	"os"
)

type YamlService interface {
	CheckYamlInCWD() (dto.Project, error)
}

func NewYamlService() YamlService {
	return &yamlService{fileName: "gitsync.ymal"}
}

type yamlService struct {
	fileName string
}

func (i *yamlService) CheckYamlInCWD() (dto.Project, error) {
	var project dto.Project
	yamlFile, err := os.ReadFile(i.fileName)
	if err != nil {
		return dto.Project{}, err
	}
	err = yaml.Unmarshal(yamlFile, &project)
	if err != nil {
		return dto.Project{}, err
	}

	return project, nil

}
