package users

type UserModel struct {
	Id      uint    `gorm:"primary_key;unique;not null;comment:用户id"`
	Name    string  `gorm:"size:10"`
	Age     int     `gorm:"default:0"`
	Address *string `gorm:"type:varchar(255)"`
	Sex     byte    `gorm:"default:0"`
}
