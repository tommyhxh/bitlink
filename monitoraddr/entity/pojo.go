package entity

type USER struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Pwd  string `json:"pwd"`
}

//定义列表返回值
type UserList struct {
	Data   []USER `json:"data"`
	Total  int    `json:"total"`
	Status bool   `json:"status"`
}

type MONConfig struct {
	Id               int    `json:"id"`
	Addr             string `json:"addr"`
	Status           string `json:"status"`
	UserId           string `json:"user_id"`
	StartBlockNumber string `json:"start_block_number"`
	CurBlockNumber   string `json:"cur_block_number"`
	Count            string `json:"count"`
	NewTXCount       string `json:"new_tx_count"`
}
type MonConfigList struct {
	Data   []MONConfig `json:"data"`
	Total  int         `json:"total"`
	Status bool        `json:"status"`
}

type TxMsg struct {
	Id              int `json:"id"`
	MonitorConfigId int `json:"monitor_config_id"`
	Addr            int `json:"addr"`
	FromTo          int `json:"from_to"`
	Amount          int `json:"amount"`
	BlockHash       int `json:"block_hash"`
	BlockNumber     int `json:"block_number"`
	Hash            int `json:"hash"`
	TimeStamp       int `json:"time_stamp"`
	Error           int `json:"error"`
}
