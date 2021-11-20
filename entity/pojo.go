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

//定义返回值
type JsonResult struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type MONConfig struct {
	Id               int    `json:"id"`
	Addr             string `json:"addr"`
	Status           int    `json:"status"`
	UserId           string `json:"userId"`
	StartBlockNumber string `json:"startBlockNumber"`
	CurBlockNumber   string `json:"curBlockNumber"`
	Count            string `json:"count"`
	NewTXCount       string `json:"newTXCount"`
}
type MonConfigList struct {
	Data   []MONConfig `json:"data"`
	Total  int         `json:"total"`
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
}

type TxMsg struct {
	Id              int `json:"id"`
	MonitorConfigId int `json:"monitorConfigId"`
	Addr            int `json:"addr"`
	FromTo          int `json:"fromTo"`
	Amount          int `json:"amount"`
	BlockHash       int `json:"blockHash"`
	BlockNumber     int `json:"blockNumber"`
	Hash            int `json:"hash"`
	TimeStamp       int `json:"timeStamp"`
	Error           int `json:"error"`
}
