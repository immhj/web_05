package main

import (
	"web_05/apis"
	"web_05/dao"
)

func main() {
	dao.Initdatabase()
	apis.InitRouter()
}
