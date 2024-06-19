package reply

type Cemetery struct {
	ID          int64        `json:"id" gorm:"primary_key"`
	Name        string       `json:"name" gorm:"UNIQUE"`
	Describe    string       `json:"describe" gorm:"type:text"`
	Information string       `json:"information" gorm:"type:text"`
	Image       string       `json:"image" gorm:"varchar(255)"`
	Website     string       `json:"website" gorm:"varchar(255)"`
	Coordinates string       `json:"coordinates" gorm:"varchar(255)"`
	Attractions []Attraction `json:"attractions" gorm:"foreignKey:cemetery_id"`
}

type Cemeteries struct {
	CemeteriesInfo []*CemeteryInfo
}

type CemeteryInfo struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"UNIQUE"`
	Image string `json:"image" gorm:"varchar(255)"`
}

type GetCemeteryByID struct {
	CemeteryInfo Cemetery
	Attraction   []*Attraction
}
