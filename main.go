package main

import "MGateAllCA/controller"

func main() {

	app := &controller.App{}
	app.Initialize()
	app.Run(":9482")

}
