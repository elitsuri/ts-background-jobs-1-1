package search

import ("database/sql"; "fmt"; "strings")

type Engine struct{ db *sql.DB }

func New(db *sql.DB) *Engine { return &Engine{db: db} }

type Result struct {
	Items []map[string]interface{} `json:"items"`
	Users []map[string]interface{} `json:"users"`
	Total int                      `json:"total"`
}

func (e *Engine) Search(q string, types []string, limit int) (*Result, error) {
	result := &Result{}
	like := "%" + strings.ToLower(q) + "%"
	for _, t := range types {
		switch t {
		case "items":
			rows, err := e.db.Query(fmt.Sprintf(`SELECT id, name, description FROM items WHERE LOWER(name) LIKE $1 LIMIT %d`, limit), like)
			if err == nil { defer rows.Close(); for rows.Next() { var id int64; var n, d string; _ = rows.Scan(&id, &n, &d); result.Items = append(result.Items, map[string]interface{}{"id":id,"name":n,"description":d}) } }
		case "users":
			rows, err := e.db.Query(fmt.Sprintf(`SELECT id, name, email FROM users WHERE LOWER(name) LIKE $1 OR LOWER(email) LIKE $1 LIMIT %d`, limit), like)
			if err == nil { defer rows.Close(); for rows.Next() { var id int64; var n, em string; _ = rows.Scan(&id, &n, &em); result.Users = append(result.Users, map[string]interface{}{"id":id,"name":n,"email":em}) } }
		}
	}
	result.Total = len(result.Items) + len(result.Users)
	return result, nil
}
