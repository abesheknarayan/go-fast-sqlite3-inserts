package models

type User struct {
	Id     uint32 `gorm:"column:id;primary_key;not null" json:"id"`
	Area   string `gorm:"column:area" json:"area"`
	Age    uint32 `gorm:"column:age;not null" json:"age"`
	Active uint32 `gorm:"column:active;not null" json:"active"`
}
