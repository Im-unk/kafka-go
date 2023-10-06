package models

type Data struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
