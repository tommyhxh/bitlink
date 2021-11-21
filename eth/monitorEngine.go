package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"monitoraddr/dao"
	"monitoraddr/entity"
	"time"
)

type Monitor struct {
	Addr          map[common.Address]bool
	Status        bool
	StartBlockNum uint64
	CurBlockNum   uint64
}

func (m Monitor) InitConfig() {
	configs := dao.QueryFromMonConfigDBAll()
	for _, s := range configs {
		m.Addr[common.HexToAddress(s)] = true
	}
}

func (m Monitor) AddrChange(config entity.MONConfig) {
	addr := common.HexToAddress(config.Addr)
	_, ok := m.Addr[addr]
	if config.Status == 1 {
		m.Addr[addr] = true
	} else {
		if ok {
			m.Addr[addr] = false
		}
	}
}
func (m Monitor) StatusChange() {
	m.Status = !m.Status
}

func (m Monitor) Exec() {
	for true {
		time.Sleep(time.Second * 10)
		if m.Status {
			trans := GetTransaction(m.CurBlockNum, m.Addr)
			for _, tran := range trans {
				fmt.Println(tran.Addr)
				fmt.Println(tran.FromTo)
			}
		}
	}
}

func Test() {

	addrs := map[common.Address]bool{common.HexToAddress("0x234D060Be1E7e078eDf0D3c9bD0b77b8266a4245"): true}

	var i uint64
	for i = 626560; i <= 626566; i++ {
		trans := GetTransaction(i, addrs)
		for _, tran := range trans {
			fmt.Println(tran.Addr)
			fmt.Println(tran.FromTo)
		}
	}
}
