package main

import (
	"awesomeProject1/blocks"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	blockchain blocks.Blockchain
}

func NewApi(blockchain blocks.Blockchain) *Api {
	return &Api{
		blockchain: blockchain,
	}
}

func (a *Api) Start() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"resp": "pong",
		})
	})
	r.POST("/transactions/new", a.newTransaction)
	r.GET("/chain", a.getChain)

	r.Run(":8001")
}

func (a *Api) getChain(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"chain": a.blockchain.Chain(),
	})
}
func (a *Api) newTransaction(ctx *gin.Context) {
	var transaction blocks.Transaction
	ctx.BindJSON(&transaction)
	a.blockchain.NewTransaction(transaction)
	ctx.JSON(http.StatusOK, gin.H{
		"transaction": "created",
	})
}
