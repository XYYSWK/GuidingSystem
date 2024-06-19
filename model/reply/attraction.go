package reply

type Attraction struct {
	ID         int64  `json:"id,omitempty" gorm:"primary_key"`
	Name       string `json:"name,omitempty" gorm:"UNIQUE"`
	Describe   string `json:"describe,omitempty" gorm:"type:text"`
	CemeteryID int64  `json:"cemetery_id"`
}
