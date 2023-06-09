package service

import "webtodo/pkg/repository"

type Authorization interface {
}

type Todo interface {
}

type Service struct {
	Authorization
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
