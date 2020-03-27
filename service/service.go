package service

import (
	"github.com/ha-t2/di-sample/dao"
	"math"
)

type ProductService struct {
}

// dao(Data Access Object)を呼び出して、税込価格を取得する。
func (s *ProductService) GetPriceWithTax(id int) int {
	d := dao.NewProductDao()
	product := d.Get(id)
	return int(math.Floor( float64(product.Price) * 1.1))
}
