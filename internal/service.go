package inernal

import (
	m "github.com/vova1001/krios_proj/models"
)

type partService struct {
	repo *partRepo
}

func NewService(repo *partRepo) *partService {
	return &partService{repo: repo}
}

func (s *partService) CreateObj(Obj m.Object) {
	s.repo.AddObjFromDB(Obj)
}
