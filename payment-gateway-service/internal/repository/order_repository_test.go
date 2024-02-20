package repository_test

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"github.com/timopattikawa/payment-gateway-service/internal/repository"
)

var orderDummy = &domain.Order{
	UUID:        uuid.New().String(),
	ProductId:   1,
	CostumerId:  1,
	TotalAmount: 45000,
}

func DBMockPrep() (*sql.DB, sqlmock.Sqlmock) {
	sqlDB, mockDb, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Someting error %s expected when mocking database", err)
	}
	return sqlDB, mockDb
}

func TestSaveOrder_SuccessToSave(t *testing.T) {
	sqlMock, mock := DBMockPrep()

	fakeOrderRepo := &repository.OrderRepositoryImpl{Db: sqlMock}
	query := `INSERT INTO 
	public.order(id, product_id, costumer_id, total_amount)
	VALUES ($1, $2, $3, $4)`

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().
		WithArgs(orderDummy.UUID, orderDummy.ProductId, orderDummy.CostumerId, orderDummy.TotalAmount).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := fakeOrderRepo.SaveRepository(*orderDummy)
	assert.NoError(t, err)
}

func TestFindOrderById_SuccessToFind(t *testing.T) {
	sqlMock, mock := DBMockPrep()

	fakeOrderRepo := &repository.OrderRepositoryImpl{Db: sqlMock}
	query := `SELECT id, product_id, costumer_id, total_amount
	FROM public.order WHERE id = $1;`

	rows := mock.NewRows([]string{"id", "product_id", "costumer_id", "total_amount"}).
		AddRow(orderDummy.UUID, orderDummy.ProductId, orderDummy.CostumerId, orderDummy.TotalAmount)

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectQuery().
		WithArgs(orderDummy.UUID).
		WillReturnRows(rows)

	order, err := fakeOrderRepo.FindOrderById(orderDummy.UUID)
	assert.NoError(t, err)
	assert.Equal(t, order.UUID, order.UUID)
}

func TestFindOrderById_FailToFind(t *testing.T) {
	sqlMock, mock := DBMockPrep()

	fakeOrderRepo := &repository.OrderRepositoryImpl{Db: sqlMock}
	query := `SELECT id, product_id, costumer_id, total_amount
	FROM public.order WHERE id = $1;`

	rows := mock.NewRows([]string{"id", "product_id", "costumer_id", "total_amount"}).
		AddRow(orderDummy.UUID, orderDummy.ProductId, orderDummy.CostumerId, orderDummy.TotalAmount)

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectQuery().
		WithArgs(orderDummy.UUID).
		WillReturnRows(rows)

	order, err := fakeOrderRepo.FindOrderById(uuid.New().String())
	assert.Error(t, err)
	assert.Equal(t, order, domain.Order{})
}
