package store

import (
	"github.com/albakov/go-todo/internal/app/store/repo"
)

type Store struct {
	itemRepo *repo.ItemRepo
	Path     string `toml:"path_to_json"`
}

func (s *Store) Item() *repo.ItemRepo {
	if s.itemRepo == nil {
		s.itemRepo = &repo.ItemRepo{
			Path: s.Path,
		}
	}

	return s.itemRepo
}
