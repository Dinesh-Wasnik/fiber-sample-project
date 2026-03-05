package models

import "time"

type DemoModel struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	RegisterModel(&DemoModel{})
}

// func (GameTable) DropColumns() []string {
// 	return []string{
// 		//columns to drop
// 	}
// }
