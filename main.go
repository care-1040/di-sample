package main

import (
	"fmt"
	"github.com/ha-t2/di-sample/repo"
	"github.com/ha-t2/di-sample/service"
)

func main() {
	r := repo.NewProductRepo()
	s := service.ProductService{Repo: r}
	fmt.Println(s.Exist(1))
	fmt.Println(s.Exist(2))
}
