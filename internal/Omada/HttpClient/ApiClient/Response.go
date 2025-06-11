package ApiClient

import "strconv"

const DEFAULT_PAGE_SIZE = 100

type Page[T any] struct {
	TotalRows   int `json:"totalRows"`
	CurrentPage int `json:"currentPage"`
	CurrentSize int `json:"currentSize"`
	Data        []T `json:"data"`
}

type Response[T any] struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"msg"`
	Result    T      `json:"result"`
}

func AddPaginationParams(params map[string]string, page int) map[string]string {
	if params == nil {
		params = make(map[string]string)
	}
	params["page"] = strconv.Itoa(page)
	params["pageSize"] = strconv.Itoa(DEFAULT_PAGE_SIZE)
	return params
}

func (p *Page[T]) HasMorePages() bool {
	if p.CurrentPage <= 0 {
		return true
	}
	return p.TotalRows > (p.CurrentPage * p.CurrentSize)
}
