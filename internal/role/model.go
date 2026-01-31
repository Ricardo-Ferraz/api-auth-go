package role

type Role struct {
	Id     int64  `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	UserId int64  `gorm:"column:user_id"`
}
