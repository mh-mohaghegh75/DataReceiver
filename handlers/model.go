package handler

import "time"

type UserId string

type User struct {
	ID        UserId    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp WITH TIME ZONE; default: NOW(); NOT NULL"`
	Quota     uint      `gorm:"type:VARCHAR; NOT NULL" json:"quota"`
	Datas     []Data
}

type Data struct {
	ID        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp WITH TIME ZONE; default: NOW(); NOT NULL"`
	UserID    UserId    `gorm:"type:VARCHAR; NOT NULL" json:"userId"`
}

type UsedQouta struct {
	UserID string
	Count  int
}
