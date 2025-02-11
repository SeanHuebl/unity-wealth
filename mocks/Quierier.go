// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	database "github.com/seanhuebl/unity-wealth/internal/database"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Quierier is an autogenerated mock type for the Quierier type
type Querier struct {
	mock.Mock
}

// CreateDeviceInfo provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateDeviceInfo(ctx context.Context, arg database.CreateDeviceInfoParams) (string, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateDeviceInfo")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateDeviceInfoParams) (string, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateDeviceInfoParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateDeviceInfoParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateRefreshToken(ctx context.Context, arg database.CreateRefreshTokenParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateRefreshTokenParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTransaction provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateTransaction(ctx context.Context, arg database.CreateTransactionParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateTransactionParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, params
func (_m *Querier) CreateUser(ctx context.Context, params database.CreateUserParams) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateUserParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransactionById provides a mock function with given fields: ctx, arg
func (_m *Querier) DeleteTransactionById(ctx context.Context, arg database.DeleteTransactionByIdParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransactionById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DeleteTransactionByIdParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDetailedCategories provides a mock function with given fields: ctx
func (_m *Querier) GetDetailedCategories(ctx context.Context) ([]database.DetailedCategory, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDetailedCategories")
	}

	var r0 []database.DetailedCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.DetailedCategory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.DetailedCategory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.DetailedCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailedCategoryId provides a mock function with given fields: ctx, name
func (_m *Querier) GetDetailedCategoryId(ctx context.Context, name string) (int64, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetDetailedCategoryId")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceInfoByUser provides a mock function with given fields: ctx, arg
func (_m *Querier) GetDeviceInfoByUser(ctx context.Context, arg database.GetDeviceInfoByUserParams) (string, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceInfoByUser")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.GetDeviceInfoByUserParams) (string, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.GetDeviceInfoByUserParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.GetDeviceInfoByUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrimaryCategories provides a mock function with given fields: ctx
func (_m *Querier) GetPrimaryCategories(ctx context.Context) ([]database.PrimaryCategory, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPrimaryCategories")
	}

	var r0 []database.PrimaryCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.PrimaryCategory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.PrimaryCategory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.PrimaryCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Querier) GetUserByEmail(ctx context.Context, email string) (database.GetUserByEmailRow, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 database.GetUserByEmailRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.GetUserByEmailRow, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.GetUserByEmailRow); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(database.GetUserByEmailRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTransactionByID provides a mock function with given fields: ctx, arg
func (_m *Querier) GetUserTransactionByID(ctx context.Context, arg database.GetUserTransactionByIDParams) (database.GetUserTransactionByIDRow, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetUserTransactionByID")
	}

	var r0 database.GetUserTransactionByIDRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionByIDParams) (database.GetUserTransactionByIDRow, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionByIDParams) database.GetUserTransactionByIDRow); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.GetUserTransactionByIDRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.GetUserTransactionByIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTransactionsFirstPage provides a mock function with given fields: ctx, arg
func (_m *Querier) GetUserTransactionsFirstPage(ctx context.Context, arg database.GetUserTransactionsFirstPageParams) ([]database.GetUserTransactionsFirstPageRow, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetUserTransactionsFirstPage")
	}

	var r0 []database.GetUserTransactionsFirstPageRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionsFirstPageParams) ([]database.GetUserTransactionsFirstPageRow, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionsFirstPageParams) []database.GetUserTransactionsFirstPageRow); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.GetUserTransactionsFirstPageRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.GetUserTransactionsFirstPageParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTransactionsPaginated provides a mock function with given fields: ctx, arg
func (_m *Querier) GetUserTransactionsPaginated(ctx context.Context, arg database.GetUserTransactionsPaginatedParams) ([]database.GetUserTransactionsPaginatedRow, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetUserTransactionsPaginated")
	}

	var r0 []database.GetUserTransactionsPaginatedRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionsPaginatedParams) ([]database.GetUserTransactionsPaginatedRow, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.GetUserTransactionsPaginatedParams) []database.GetUserTransactionsPaginatedRow); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.GetUserTransactionsPaginatedRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.GetUserTransactionsPaginatedParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeToken provides a mock function with given fields: ctx, arg
func (_m *Querier) RevokeToken(ctx context.Context, arg database.RevokeTokenParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for RevokeToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.RevokeTokenParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTransactionByID provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateTransactionByID(ctx context.Context, arg database.UpdateTransactionByIDParams) (database.UpdateTransactionByIDRow, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransactionByID")
	}

	var r0 database.UpdateTransactionByIDRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateTransactionByIDParams) (database.UpdateTransactionByIDRow, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateTransactionByIDParams) database.UpdateTransactionByIDRow); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.UpdateTransactionByIDRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.UpdateTransactionByIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTx provides a mock function with given fields: tx
func (_m *Querier) WithTx(tx *sql.Tx) *database.Queries {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for WithTx")
	}

	var r0 *database.Queries
	if rf, ok := ret.Get(0).(func(*sql.Tx) *database.Queries); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Queries)
		}
	}

	return r0
}

// NewQuierier creates a new instance of Quierier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuierier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
