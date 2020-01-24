package models

// Article -- article struct
type Article struct {
	ID    uint   `gorm:"primary_key"`
	Title string `gorm:"column:title;type:varchar(50);not null" json:"title"`
	Body  string `gorm:"column:body;type:varchar(255)" json:"body"`
}
