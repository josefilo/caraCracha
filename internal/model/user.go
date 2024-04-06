package model

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email	 string `gorm:"unique"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	BirthDate string `gorm:"not null"`
	CreatedAt string `gorm:"autoCreateTime"`
	UpdatedAt string `gorm:"autoUpdateTime"`
}

// Deixe no singular
func (User) TableName() string {
	return "user"
}
