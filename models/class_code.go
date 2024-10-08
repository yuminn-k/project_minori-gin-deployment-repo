package models

type ClassCode struct {
	ID     uint    `gorm:"primaryKey;size:255;autoIncrement"`
	Code   string  `gorm:"size:10;not null"`
	Secret *string `gorm:"size:20"`
	CID    uint    `gorm:"column:cid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UID    uint    `gorm:"column:uid;not null"` // User ID
	Class  Class   `gorm:"foreignKey:CID;constraint:OnDelete:CASCADE"`
	User   User    `gorm:"foreignKey:UID"`
}
