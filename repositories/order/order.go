package repositories

import (
	"context"
	"errors"
	"fmt"
	errWrap "order-service/common/error"
	errConstants "order-service/constants/error"
	errOrder "order-service/constants/error/order"
	"order-service/domain/dto"
	"order-service/domain/models"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type IOrderRepository interface {
	FindAllWithPagination(context.Context, *dto.OrderRequestParam) ([]models.Order, int64, error)
	FindByUserID(context.Context, string) ([]models.Order, error)
	FindByUUID(context.Context, string) (*models.Order, error)
	Create(context.Context, *gorm.DB, *models.Order) (*models.Order, error)
	Update(context.Context, *gorm.DB, *models.Order, uuid.UUID) error
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) FindAllWithPagination(ctx context.Context, param *dto.OrderRequestParam) ([]models.Order, int64, error) {
	var (
		orders []models.Order
		total  int64
		sort   string
	)

	if param.SortColumn != nil {
		sort = fmt.Sprintf("%s %s", *param.SortColumn, *param.SortOrder)
	} else {
		sort = "created_at desc"
	}

	limit := param.Limit
	offfset := (param.Page - 1) * limit
	err := o.db.WithContext(ctx).Limit(limit).Offset(offfset).Order(sort).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	err = o.db.WithContext(ctx).Model(&models.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstants.ErrSQLError)
	}

	return orders, total, nil
}

func (o *OrderRepository) FindByUUID(ctx context.Context, uuid string) (*models.Order, error) {
	var order models.Order
	err := o.db.WithContext(ctx).Where("uuid = ?", uuid).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errOrder.ErrOrderNotFound)
		}
		return nil, errWrap.WrapError(errConstants.ErrSQLError)
	}

	return &order, nil
}

func (o *OrderRepository) FindByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	var orders []models.Order
	err := o.db.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errOrder.ErrOrderNotFound)
		}
		return nil, errWrap.WrapError(errConstants.ErrSQLError)
	}

	return orders, nil
}

func (o *OrderRepository) incrementCode(ctx context.Context) (*string, error){
	var (
		order *models.Order
		result string
		today = time.Now().Format("20260102")
	)

	err := o.db.WithContext(ctx).Order("id desc").First(&order).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errOrder.ErrOrderNotFound)
		}
	}

	if order.ID != 0 {
		orderCode := order.Code
		spiltOrderName, _ := strconv.Atoi(orderCode[4:9])
		code := spiltOrderName + 1
		result = fmt.Sprintf("ORD-%05d-%s", code, today)
	}else {
		result = fmt.Sprintf("ORD-%5d-%s", 1,today)
	}

	return &result, nil
}

func (o *OrderRepository) Create(ctx context.Context, tx *gorm.DB, order *models.Order) (*models.Order, error) {
	code , err := o.incrementCode(ctx)
	if err != nil {
		return nil, err
	}
	
	orderModels := &models.Order{
		UUID: uuid.New(),
		Code: *code,
		UserID: order.UserID,
		Amount: order.Amount,
		Date: order.Date,
		Status: order.Status,
		IsPaid: order.IsPaid,
		PaidAt: order.PaidAt,
	}

	err = tx.WithContext(ctx).Create(orderModels).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstants.ErrSQLError)
	}

	return orderModels, nil
}

func (o *OrderRepository) Update(ctx context.Context, tx *gorm.DB ,request *models.Order, uuid uuid.UUID) error {
	err := tx.WithContext(ctx).Where("uuid = ?", uuid).Updates(request).Error
	if err != nil {
		return errWrap.WrapError(errConstants.ErrSQLError)
	}

	return nil
}