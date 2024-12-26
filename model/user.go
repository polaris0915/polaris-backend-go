package model

type User struct {
	Model
	Identity string `json:"identity" gorm:"uniqueIndex"`
	Account  string `json:"account" gorm:"not null;"`
	Password string `json:"password" gorm:"not null;"`
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	Name     string `json:"name"`
	Avatar   []byte `json:"avatar"`
	Profile  string `json:"profile"`
	Role     string `json:"role" gorm:"default:user;"`
}

func (u *User) TableName() string {
	return "user"
}
