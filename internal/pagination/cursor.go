package pagination

import ("net/http"; "strconv")

type PageParams struct{ Page, Limit int; Offset int }

func FromRequest(r *http.Request) PageParams {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 { page = 1 }
	if limit < 1 || limit > 200 { limit = 20 }
	return PageParams{Page: page, Limit: limit, Offset: (page - 1) * limit}
}

type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Pages int64       `json:"pages"`
}

func NewResult(data interface{}, total int64, p PageParams) PageResult {
	pages := total / int64(p.Limit)
	if total%int64(p.Limit) > 0 { pages++ }
	return PageResult{Data: data, Total: total, Page: p.Page, Limit: p.Limit, Pages: pages}
}
