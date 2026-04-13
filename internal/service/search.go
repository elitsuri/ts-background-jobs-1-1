package service

import "database/sql"

type SearchService struct{ db *sql.DB }

func NewSearchService(db *sql.DB) *SearchService { return &SearchService{db: db} }

func (s *SearchService) Search(q, sType string, limit int) (map[string]interface{}, error) {
	result := map[string]interface{}{"items": []interface{}{}, "users": []interface{}{}}
	like := "%" + q + "%"
	if sType == "items" || sType == "all" {
		rows, err := s.db.Query(`SELECT id,name,description FROM items WHERE name ILIKE $1 OR description ILIKE $1 LIMIT $2`, like, limit)
		if err == nil {
			defer rows.Close()
			var items []map[string]interface{}
			for rows.Next() {
				var id int64; var name, desc string
				_ = rows.Scan(&id, &name, &desc)
				items = append(items, map[string]interface{}{"id":id,"name":name,"description":desc})
			}
			result["items"] = items
		}
	}
	return result, nil
}
