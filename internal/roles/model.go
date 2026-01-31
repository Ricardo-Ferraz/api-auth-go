package roles

type Role struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	UserId int
}
