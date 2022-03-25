package api

import "github.com/gin-gonic/gin"

// StartServerDetached start the http rest API.
func StartServerDetached() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1Acc := v1.Group("/accounts")
	v1Acc.GET("/:address", accountFullInfo)
	v1Acc.GET("/:address/transactions", accountTxs)

	v1Txs := v1.Group("/transactions")
	v1Txs.POST("", submitTxs)
	v1Txs.GET("/:txHash", txStatus)
	v1Txs.GET("/:txHash/data", txData)

	v1Blk := v1.Group("/blocks")
	v1Blk.GET("", getBlocks)
	v1Blk.GET("/:blockNumber", blockById)
	v1Blk.GET("/:blockNumber/transactions", blockTransactions)

}

func blockTransactions(context *gin.Context) {

}

func blockById(context *gin.Context) {

}

func getBlocks(context *gin.Context) {

}

func txData(c *gin.Context) {
	c.JSON(
		500, gin.H{
			"tx_data": 1,
		},
	)
}

func txStatus(c *gin.Context) {
	c.JSON(
		500, gin.H{
			"tx_status": 1,
		},
	)
}

func submitTxs(c *gin.Context) {
	c.JSON(
		500, gin.H{
			"submit_tx": 1,
		},
	)
}

func accountTxs(c *gin.Context) {
	c.JSON(
		500, gin.H{
			"acc_txs": 1,
		},
	)
}

func accountFullInfo(c *gin.Context) {
	c.JSON(
		500, gin.H{
			"accInfo": 1,
		},
	)
}
