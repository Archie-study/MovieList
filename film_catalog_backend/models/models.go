package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Movie struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(100);not null" json:"title" binding:"required"`
	Description string `json:"description"`
	ReleaseDate string `gorm:"type:date" json:"release_date"`
	Genre       string `gorm:"type:varchar(50)" json:"genre"`
	Director    string `gorm:"type:varchar(100)" json:"director"`
}

type Rating struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"` // ID pengguna yang memberikan rating
	MovieID uint `gorm:"not null"` // ID film yang dinilai
	Rating  int  `gorm:"check:rating >= 1 AND rating <= 5"`
	Review  string
}
