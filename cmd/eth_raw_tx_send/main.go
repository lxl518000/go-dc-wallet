// 发送交易
package main

import (
	"go-dc-wallet/app"
	"go-dc-wallet/eth"
)

func main() {
	app.EnvCreate()
	defer app.EnvDestroy()

	eth.CheckRawTxSend()
}