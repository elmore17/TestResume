package models

type User struct {
	Id    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}

type Book struct {
	Id     uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Title  string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId"`
}

type InputBook struct {
	Title  string
	UserId uint
}

type UserInput struct {
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}

type UserUpdate struct {
	Name  string
	Email string
}
