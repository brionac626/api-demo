package models

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	_defaultPage      = 0
	_defaultPageLimit = 10
)

type GetArticlesReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (ga *GetArticlesReq) CheckPaginationValue() {
	if ga.Page < 0 {
		ga.Page = _defaultPage
	}

	if ga.Limit < 0 {
		ga.Limit = _defaultPageLimit
	}
}

type GetArticlesResp struct {
	Articles []Article `json:"articles,omitempty"`
	Total    int       `json:"total"`
}

type CreateArticlesReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ModifyArticlesReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
