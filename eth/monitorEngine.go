package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"monitoraddr/entity"
	"time"
)

type monitor struct {
	addr          map[string]bool
	status        bool
	startBlockNum uint64
	curBlockNum   uint64
}

func (m monitor) initConfig(configs []entity.MONConfig) {
	for _, s := range configs {
		m.addr[s.Addr] = true
	}
}

func (m monitor) addrChange(config entity.MONConfig) {
	_, ok := m.addr[config.Addr]
	if config.Status == "1" {
		m.addr[config.Addr] = true
	} else {
		if ok {
			m.addr[config.Addr] = false
		}
	}
}
func (m monitor) statusChange() {
	m.status = !m.status
}

func (m monitor) exec() {
	for true {
		time.Sleep(time.Second * 10)
		if m.status {
			//获取区块
		}
	}
}

func Test() {
	addrs := []common.Address{common.HexToAddress("0x234D060Be1E7e078eDf0D3c9bD0b77b8266a4245")}
	var i uint64
	for i = 626560; i <= 626566; i++ {
		trans := GetTransaction(i, addrs)
		for _, tran := range trans {
			fmt.Println(tran.Addr)
			fmt.Println(tran.FromTo)
		}
	}
}
