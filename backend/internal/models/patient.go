package models

type Patient struct {
	FullName string `json:"fullname"`
	Birthday Date   `json:"birthday"`
	Gender   Gender `json:"gender"`
	GUID     string `json:"guid"`
}
