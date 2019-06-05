package main

import "M-Gate/controller"

func main() {

	app := &controller.App{}
	app.Initialize()
	app.Run(":9482")

}
