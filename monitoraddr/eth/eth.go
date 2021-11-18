package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func Eth() {

	// 引入第三方包，方法连接接口
	cli, err := ethclient.Dial("http://10.10.10.203:8545")

	// 找到最新的区块
	// a, b := cli.BlockNumber(context.Background())
	// fmt.Printf("a=%d,b=%#v\n", a, b)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// var address string
	// address = "0x45d92b6707754D7E787D1554fAe5F922A1C35EB1"

	var blockIdx uint64
	blockIdx = 626555
	addrFrom := common.HexToAddress("0x234D060Be1E7e078eDf0D3c9bD0b77b8266a4245")
	// addrTo := common.HexToAddress("0x234D060Be1E7e078eDf0D3c9bD0b77b8266a4245")
	for blockIdx <= 626566 {

		// 查询链的当前高度
		currentHeight, err := cli.BlockNumber(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 10)
			continue
		}

		// 如果初始值大于当前高度，就等待十秒
		if blockIdx >= currentHeight {
			fmt.Printf("sleep.....10s \n")
			time.Sleep(time.Second * 10)
			continue
		}

		blockIdx = blockIdx + 1
		// 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等
		// context.Background() 返回一个空的Context 我们可以用这个 空的 Context 作为 goroutine 的root 节点
		block, err := cli.BlockByNumber(context.Background(), big.NewInt(int64(blockIdx)))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 打印块的高度
		fmt.Printf("block height: %v \n", block.Number())

		// 循环遍历,返回字段
		// 调用Transactions方法来读取块中的事务，该方法返回一个Transaction类型的列表。 然后，重复遍历集合并获取有关事务的任何信息
		for k, tx := range block.Body().Transactions {
			// fmt.Printf("To():           %T \n", msg.To())
			// fmt.Printf("GasPrice():     %v \n", msg.GasPrice())
			// fmt.Printf("GasFeeCap():    %v \n", msg.GasFeeCap())
			// fmt.Printf("GasTipCap():    %v \n", msg.GasTipCap())
			// fmt.Printf("Value():        %v \n", msg.Value())
			// fmt.Printf("Gas():          %v \n", msg.Gas())
			// fmt.Printf("Nonce():        %v \n", msg.Nonce())
			// fmt.Printf("Data():         %x \n", msg.Data())
			// fmt.Printf("AccessList():   %v \n", msg.AccessList())

			// 为了读取发送方的地址，我们需要在事务上调用AsMessage，它返回一个Message类型，其中包含一个返回sender（from）地址的函数。 AsMessage方法需要EIP155签名者，这个我们从客户端拿到链ID。
			msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), nil)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Printf("To():           %T \n", msg.To())

			if msg.From() == addrFrom {
				// fmt.Printf("idx: %v  tx: %v to: %v  value: %v	ChainId: %v	  Hash:%v \n", k, tx.Hash(), tx.To(), tx.Value(), tx.ChainId(), tx.Hash())
				fmt.Printf("ChainId:        %v \n", tx.ChainId())
				fmt.Printf("所在块的高度: 	%v \n", block.Number())
				fmt.Printf("TxIdx:          %v \n", k)
				fmt.Printf("From():         %v     %T \n", msg.From(), msg.From())
				fmt.Printf("To:		%v \n", tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
				fmt.Printf("Hash:		%v \n", tx.Hash)              // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				fmt.Printf("Value:		%v \n", tx.Value().String()) // 10000000000000000
				fmt.Printf("Gas:		%v \n", tx.Gas())
				fmt.Printf("GasPrice:	%v \n", tx.GasPrice().Uint64()) // 102000000000
				fmt.Printf("GasFeeCap():    %v \n", tx.GasFeeCap())
				fmt.Printf("GasTipCap():    %v \n", tx.GasTipCap())
				fmt.Printf("Nonce:		%v \n", tx.Nonce()) // 110644
				fmt.Printf("Data:		%v \n", tx.Data())   // []
				fmt.Printf("IsFake():       %v \n", msg.IsFake())
				fmt.Printf("AccessList():   %v \n", tx.AccessList())

				//每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志，以及为“1”（成功）或“0”（失败）的事件结果状态。
				receipt, err := cli.TransactionReceipt(context.Background(), tx.Hash())
				fmt.Printf("tx: %v status: %v \n", tx.Hash(), receipt.Status)

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				// Json Marshal：将数据编码成json字符串receiptAsJson, err := json.Marshal(receipt)
				// Json Unmarshal：将json字符串解码到相应的数据结构
				// json.MarshalIndent将receipt收据编码成json字符串
				receiptAsJson, err := json.MarshalIndent(receipt, "", "\t")

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Printf("result: %s \n", receiptAsJson)

				// value, err := cli.BalanceAt(context.Background(), *tx.To(), nil)

				// if err != nil {
				//  fmt.Println(err.Error())
				//  return
				// }

				// fmt.Printf("addr: %v value:%v \n", tx.To(), value)

				// if err != nil {
				// 	fmt.Println(err.Error())
				// 	return
				// }
			}

		}
	}
}
