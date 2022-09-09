package models

type Form struct {
	ID           int    `json:"id"`
	UserURL      string `json:"userUrl" valid:"url"`
	NewURL       string `json:"newUrl " valid:"alphanumeric"`
	StatisticURL string `json:"statisticUrl"`
}

type Stat struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Time string `json:"time " `
}
