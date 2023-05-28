package entity

type ResponseCreate struct {
	Status int    `json:"status"`
	Id     string `json:"id"`
}

type ResponseMakeFriend struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}
