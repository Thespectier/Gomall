package module

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex;type:varchar(127) not null"`
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}
