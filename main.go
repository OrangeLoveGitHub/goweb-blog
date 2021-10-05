package main

import (
	"goweb-blog/model"
	"goweb-blog/routes"
)

func main() {
	//init db
	model.InitDb()
	routes.InitRouter()
}
