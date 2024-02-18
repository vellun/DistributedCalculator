package models

type Expression struct {
	Id         int    `json:"id"`
	Expression string `json:"expression"`
	Status     string `json:"status"` // process/complete
	Started_at int64  `json:"started_at"`
	Ended_at   int64  `json:"ended_at"`
	Result     string `json:"result"`
}
