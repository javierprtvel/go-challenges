package repository

import "github.mpi-internal.com/javier-porto/learning-go/domain"

type AdRepository interface {
	Persist(ad domain.Ad) domain.Ad
	FindById(id string) domain.Ad
	Slice() []domain.Ad
}
