package main

import (
	"M-GateAllCA/process"
	"github.com/gin-gonic/gin"
)

func main() {

	process := process.trxProcess{data: data}
	router := gin.Default()

	router.GET("/api/billers", process.getBiller)

}
