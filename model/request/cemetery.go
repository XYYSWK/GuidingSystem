package request

type GetCemeteriesByName struct {
	Name string `json:"name"`
}

type GetCemeteryByID struct {
	ID int64 `json:"id"`
}

type Cemetery struct {
	Name        string `json:"name" form:"Name"`
	Describe    string `json:"describe"form:"Describe"`
	Information string `json:"information"form:"Information"`
	Image       string `json:"image" form:"Image"`
	Website     string `json:"website" form:"Website"`
	Coordinates string `json:"coordinates" form:"Coordinates"`
}
