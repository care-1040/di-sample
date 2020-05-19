package service

import (
	"github.com/ha-t2/di-sample/repo"
)

type ProductService struct {
}

// repo(Repository)を呼び出して、存在チェックをする。
func (s *ProductService) Exist(id int) bool {
	d := repo.NewProductRepo()
	product := d.Get(id)
	return product.Id == id
}
