# di sample
## Run Mysql
`docker-compose up -d --build`
## Set up environment
`dep ensure`
## Build and Run app
`go build . && ./di-sample`

## 問題
- mysql serverがあってかつ、適切にデータを置かないとserviceのテストできない。
- docker-compose せずにserviceのtestを通そう！
