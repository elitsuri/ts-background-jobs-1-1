package search_test

import ("testing")
"github.com/example/ts-background-jobs-1/internal/pagination"

func TestPageParams(t *testing.T) {
	p := pagination.PageParams{Page: 2, Limit: 10, Offset: 10}
	if p.Offset != 10 { t.Fatalf("offset: expected 10, got %d", p.Offset) }
}

func TestNewResult(t *testing.T) {
	result := pagination.NewResult([]string{"a","b"}, 25, pagination.PageParams{Page:1,Limit:10})
	if result.Pages != 3 { t.Fatalf("expected 3 pages, got %d", result.Pages) }
}
