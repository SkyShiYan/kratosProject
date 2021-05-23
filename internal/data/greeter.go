package data

import (
	"context"
	"spaco_go/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper("data/greeter", logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) GetGreeter(context context.Context, name string) (int, error) {
	s := &GreeterEntity{
		Name: name,
	}
	result := r.data.db.Where("name = ?", name).First(&s).Error
	if result == gorm.ErrInvaildDB {
		panic("错误的数据库")
	} else if result == gorm.ErrRecordNotFound {
		// s := biz.NoDataErr{
		// 	Err: result,
		// }
		return 0, errors.NotFound("GreeterEntity", name, "名称未找到")
	}

	return *&s.Age, result
}

type GreeterEntity struct {
	gorm.Model
	Name string
	Age  int
}
