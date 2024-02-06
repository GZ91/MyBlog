package service

import "github.com/GZ91/MyBlog/internal/models"

func (s *Service) GetArts() ([]models.Art, error) {
	return s.NodeStorage.GetArts()
}
