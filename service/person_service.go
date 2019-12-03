package service

import (
	"github.com/ipan97/go-dig-sample/config"
	"github.com/ipan97/go-dig-sample/model"
	"github.com/ipan97/go-dig-sample/repository"
)

type PersonService struct {
	config     *config.Config
	repository *repository.PersonRepository
}

func (service *PersonService) FindAll() []*model.Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*model.Person{}
}

func NewPersonService(config *config.Config, repository *repository.PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}
