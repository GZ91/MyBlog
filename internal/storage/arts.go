package storage

import "github.com/GZ91/MyBlog/internal/models"

func (s *Storage) GetArts() ([]models.Art, error) {
	rows, err := s.db.Query("SELECT name, text FROM articles ORDER BY time DESC")
	if err != nil {
		return nil, err
	}
	var data models.Art
	var retData []models.Art
	for rows.Next() {
		rows.Scan(&data.Name, &data.Text)
		retData = append(retData, data)
	}
	return retData, nil
}
