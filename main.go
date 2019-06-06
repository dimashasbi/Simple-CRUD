package main

import (
	"M-Gate/controller"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := &controller.App{}
	app.Initialize()
	app.Run(":9482")

}
