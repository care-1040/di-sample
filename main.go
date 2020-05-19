package main

import (
	"fmt"
	"github.com/ha-t2/di-sample/service"
)

func main() {
	// 商品をidで取得して存在するかを表示する
	s := service.ProductService{}
	fmt.Println(s.Exist(1))
	fmt.Println(s.Exist(2))
}
