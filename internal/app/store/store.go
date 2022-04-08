package store

import (
	"github.com/albakov/go-todo/internal/app/store/repo"
)

type Store struct {
	itemRepo *repo.ItemRepo
}

func (s *Store) Item() *repo.ItemRepo {
	if s.itemRepo == nil {
		s.itemRepo = &repo.ItemRepo{}
	}

	return s.itemRepo
}
