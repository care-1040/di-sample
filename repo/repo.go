package repo

import (
	. "github.com/ha-t2/di-sample/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepo struct {
	db *gorm.DB
}

// gormを通してmysqlにアクセスし、Productを取ってくる
func (t *ProductRepo) Get(id int) Product {
	p := Product{}
	t.db.First(&p, id)
	return p
}

// ProductRepoのコンストラクタ
func NewProductRepo() *ProductRepo {
	DBMS := "mysql"
	USER := "di"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	d := &ProductRepo{db}
	return d
}
