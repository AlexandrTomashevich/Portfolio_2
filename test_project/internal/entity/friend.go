package entity

type Friend struct {
	Id      string   `json:"-"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []Friend `json:"friends"`
}

type MakeFriend struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

type NewAge struct {
	Age int `json: "age"`
}
