package repository

import (
	"database/sql"
	"errors"
	"testing"

	gmock "godemo/repository/mock_repository"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	q string
}

func (md *mockDB) Select(dest interface{}, query string, args ...interface{}) error {
	md.q = query
	return nil
}

func (md *mockDB) Get(dest interface{}, query string, args ...interface{}) error {
	md.q = query
	return nil
}

func (md *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	md.q = query
	return nil, nil
}

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Select(gomock.Any(), gomock.Any()).Return(nil)

	query := NewCustomerRepository(mockSv)
	results, err := query.GetAll()
	assert.NotNil(t, results)
	assert.NoError(t, err)
}

func TestGetCustomer_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Select(gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	query := NewCustomerRepository(mockSv)
	results, err := query.GetAll()
	assert.Nil(t, results)
	assert.Error(t, err)
}

func TestGetCustomerById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	query := NewCustomerRepository(mockSv)
	results, err := query.GetById(1)
	assert.NotNil(t, results)
	assert.NoError(t, err)
}

func TestGetCustomerById_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("database error"))

	query := NewCustomerRepository(mockSv)
	results, err := query.GetById(1)
	assert.Nil(t, results)
	assert.Error(t, err)
}

func TestInsert_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	result := sqlmock.NewResult(1, 0)
	mockSv.EXPECT().Exec(gomock.Any(), gomock.Any()).Return(result, nil)

	query := NewCustomerRepository(mockSv)
	customer := Customer{}
	results, err := query.Insert(customer)
	assert.NotNil(t, results)
	assert.NoError(t, err)
}

func TestInsert_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)
	mockSv.EXPECT().Exec(gomock.Any(), gomock.Any()).Return(nil, errors.New("database error"))

	query := NewCustomerRepository(mockSv)
	customer := Customer{}
	results, err := query.Insert(customer)
	assert.Nil(t, results)
	assert.Error(t, err)
}

func TestInsert_Success_Resut_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)

	result := sqlmock.NewResult(1, 0)

	mockSv.EXPECT().Exec(gomock.Any(), gomock.Any()).Return(result, nil)

	query := NewCustomerRepository(mockSv)
	customer := Customer{}
	results, err := query.Insert(customer)
	assert.NotNil(t, results)
	assert.NoError(t, err)
}

func TestInsert__Success_Resut_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSv := gmock.NewMockDB(ctrl)

	// result := sqlmock.NewResult(1, 0)
	result := sqlmock.NewErrorResult(sql.ErrNoRows)

	mockSv.EXPECT().Exec(gomock.Any(), gomock.Any()).Return(result, nil)

	query := NewCustomerRepository(mockSv)
	customer := Customer{}
	results, err := query.Insert(customer)
	assert.NotNil(t, results)
	assert.NoError(t, err)
}
