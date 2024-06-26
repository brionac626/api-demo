// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/brionac626/api-demo/models"
	mock "github.com/stretchr/testify/mock"
)

// ArticlesRepoMock is an autogenerated mock type for the ArticlesRepo type
type ArticlesRepoMock struct {
	mock.Mock
}

type ArticlesRepoMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ArticlesRepoMock) EXPECT() *ArticlesRepoMock_Expecter {
	return &ArticlesRepoMock_Expecter{mock: &_m.Mock}
}

// DeleteArticle provides a mock function with given fields: ctx, id
func (_m *ArticlesRepoMock) DeleteArticle(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArticlesRepoMock_DeleteArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteArticle'
type ArticlesRepoMock_DeleteArticle_Call struct {
	*mock.Call
}

// DeleteArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ArticlesRepoMock_Expecter) DeleteArticle(ctx interface{}, id interface{}) *ArticlesRepoMock_DeleteArticle_Call {
	return &ArticlesRepoMock_DeleteArticle_Call{Call: _e.mock.On("DeleteArticle", ctx, id)}
}

func (_c *ArticlesRepoMock_DeleteArticle_Call) Run(run func(ctx context.Context, id string)) *ArticlesRepoMock_DeleteArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ArticlesRepoMock_DeleteArticle_Call) Return(_a0 error) *ArticlesRepoMock_DeleteArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ArticlesRepoMock_DeleteArticle_Call) RunAndReturn(run func(context.Context, string) error) *ArticlesRepoMock_DeleteArticle_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllArticles provides a mock function with given fields: ctx, author, page, limit
func (_m *ArticlesRepoMock) FindAllArticles(ctx context.Context, author string, page int64, limit int64) ([]models.Article, int64, error) {
	ret := _m.Called(ctx, author, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindAllArticles")
	}

	var r0 []models.Article
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) ([]models.Article, int64, error)); ok {
		return rf(ctx, author, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []models.Article); ok {
		r0 = rf(ctx, author, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) int64); ok {
		r1 = rf(ctx, author, page, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int64, int64) error); ok {
		r2 = rf(ctx, author, page, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ArticlesRepoMock_FindAllArticles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllArticles'
type ArticlesRepoMock_FindAllArticles_Call struct {
	*mock.Call
}

// FindAllArticles is a helper method to define mock.On call
//   - ctx context.Context
//   - author string
//   - page int64
//   - limit int64
func (_e *ArticlesRepoMock_Expecter) FindAllArticles(ctx interface{}, author interface{}, page interface{}, limit interface{}) *ArticlesRepoMock_FindAllArticles_Call {
	return &ArticlesRepoMock_FindAllArticles_Call{Call: _e.mock.On("FindAllArticles", ctx, author, page, limit)}
}

func (_c *ArticlesRepoMock_FindAllArticles_Call) Run(run func(ctx context.Context, author string, page int64, limit int64)) *ArticlesRepoMock_FindAllArticles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64), args[3].(int64))
	})
	return _c
}

func (_c *ArticlesRepoMock_FindAllArticles_Call) Return(_a0 []models.Article, _a1 int64, _a2 error) *ArticlesRepoMock_FindAllArticles_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ArticlesRepoMock_FindAllArticles_Call) RunAndReturn(run func(context.Context, string, int64, int64) ([]models.Article, int64, error)) *ArticlesRepoMock_FindAllArticles_Call {
	_c.Call.Return(run)
	return _c
}

// FindArticle provides a mock function with given fields: ctx, id
func (_m *ArticlesRepoMock) FindArticle(ctx context.Context, id string) (*models.Article, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindArticle")
	}

	var r0 *models.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Article, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Article); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticlesRepoMock_FindArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindArticle'
type ArticlesRepoMock_FindArticle_Call struct {
	*mock.Call
}

// FindArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ArticlesRepoMock_Expecter) FindArticle(ctx interface{}, id interface{}) *ArticlesRepoMock_FindArticle_Call {
	return &ArticlesRepoMock_FindArticle_Call{Call: _e.mock.On("FindArticle", ctx, id)}
}

func (_c *ArticlesRepoMock_FindArticle_Call) Run(run func(ctx context.Context, id string)) *ArticlesRepoMock_FindArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ArticlesRepoMock_FindArticle_Call) Return(_a0 *models.Article, _a1 error) *ArticlesRepoMock_FindArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticlesRepoMock_FindArticle_Call) RunAndReturn(run func(context.Context, string) (*models.Article, error)) *ArticlesRepoMock_FindArticle_Call {
	_c.Call.Return(run)
	return _c
}

// InsertNewArticle provides a mock function with given fields: ctx, article
func (_m *ArticlesRepoMock) InsertNewArticle(ctx context.Context, article models.Article) error {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for InsertNewArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArticlesRepoMock_InsertNewArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertNewArticle'
type ArticlesRepoMock_InsertNewArticle_Call struct {
	*mock.Call
}

// InsertNewArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - article models.Article
func (_e *ArticlesRepoMock_Expecter) InsertNewArticle(ctx interface{}, article interface{}) *ArticlesRepoMock_InsertNewArticle_Call {
	return &ArticlesRepoMock_InsertNewArticle_Call{Call: _e.mock.On("InsertNewArticle", ctx, article)}
}

func (_c *ArticlesRepoMock_InsertNewArticle_Call) Run(run func(ctx context.Context, article models.Article)) *ArticlesRepoMock_InsertNewArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Article))
	})
	return _c
}

func (_c *ArticlesRepoMock_InsertNewArticle_Call) Return(_a0 error) *ArticlesRepoMock_InsertNewArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ArticlesRepoMock_InsertNewArticle_Call) RunAndReturn(run func(context.Context, models.Article) error) *ArticlesRepoMock_InsertNewArticle_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateArticle provides a mock function with given fields: ctx, article
func (_m *ArticlesRepoMock) UpdateArticle(ctx context.Context, article models.Article) error {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for UpdateArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArticlesRepoMock_UpdateArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateArticle'
type ArticlesRepoMock_UpdateArticle_Call struct {
	*mock.Call
}

// UpdateArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - article models.Article
func (_e *ArticlesRepoMock_Expecter) UpdateArticle(ctx interface{}, article interface{}) *ArticlesRepoMock_UpdateArticle_Call {
	return &ArticlesRepoMock_UpdateArticle_Call{Call: _e.mock.On("UpdateArticle", ctx, article)}
}

func (_c *ArticlesRepoMock_UpdateArticle_Call) Run(run func(ctx context.Context, article models.Article)) *ArticlesRepoMock_UpdateArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Article))
	})
	return _c
}

func (_c *ArticlesRepoMock_UpdateArticle_Call) Return(_a0 error) *ArticlesRepoMock_UpdateArticle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ArticlesRepoMock_UpdateArticle_Call) RunAndReturn(run func(context.Context, models.Article) error) *ArticlesRepoMock_UpdateArticle_Call {
	_c.Call.Return(run)
	return _c
}

// NewArticlesRepoMock creates a new instance of ArticlesRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticlesRepoMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticlesRepoMock {
	mock := &ArticlesRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
