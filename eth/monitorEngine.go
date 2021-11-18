package eth

import (
	"monitoraddr/entity"
	"time"
)

type monitor struct {
	addr        map[string]bool
	status      bool
	curBlockNum int
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
