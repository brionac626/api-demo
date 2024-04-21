package models

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	_defaultPage      = 1
	_defaultPageLimit = 10
)

type GetArticlesReq struct {
	ID     string `query:"id"`
	Author string `param:"author"`
	Page   int64  `query:"page"`
	Limit  int64  `query:"limit"`
}

func (ga *GetArticlesReq) CheckPaginationValue() {
	if ga.Page < 0 {
		ga.Page = _defaultPage
	}

	if ga.Limit <= 0 {
		ga.Limit = _defaultPageLimit
	}
}

type GetArticlesResp struct {
	Articles []Article `json:"articles,omitempty"`
	Total    int64     `json:"total"`
}

type CreateArticlesReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ModifyArticlesReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
