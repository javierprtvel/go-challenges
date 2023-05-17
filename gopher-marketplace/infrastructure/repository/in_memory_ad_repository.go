package repository

import (
	"github.mpi-internal.com/javier-porto/learning-go/domain"
	"time"
)

type InMemoryAdRepository map[string]domain.Ad

func (ar InMemoryAdRepository) Persist(ad domain.Ad) domain.Ad {
	if ad.Date == (time.Time{}) {
		ad.Date = time.Now()
	}
	ar[ad.Id] = ad
	return domain.Ad{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Date:        ad.Date,
	}
}

func (ar InMemoryAdRepository) FindById(id string) domain.Ad {
	ad := ar[id]
	return domain.Ad{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Date:        ad.Date,
	}
}

const sliceMaxSize = 5

func (ar InMemoryAdRepository) Slice() []domain.Ad {
	s := make([]domain.Ad, 0)
	i := 0
	for _, ad := range ar {
		s = append(s, ad)
		i++
		if i == sliceMaxSize {
			break
		}
	}
	return s
}
