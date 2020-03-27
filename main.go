package main

import (
	"fmt"
	"github.com/ha-t2/di-sample/service"
)

func main() {
	// 商品をidで取得して税込価格を表示する
	s := service.ProductService{}
	fmt.Println(s.GetPriceWithTax(1))
}
