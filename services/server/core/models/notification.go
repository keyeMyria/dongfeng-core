package models

import (
	"time"
)

// Notification notification entity
type Notification struct {
	ID         int64     `dapper:"id,primarykey,autoincrement,table=notifications"`
	UserID     string    `dapper:"user_id"`
	CustomCode string    `dapper:"custom_code"`
	Details    string    `dapper:"details"`
	Link       string    `dapper:"link"`
	CategoryID int64     `dapper:"category_id"`
	Read       int64     `dapper:"read"`
	Time       time.Time `dapper:"created_at"`
}

// Category parse category by id
func (n *Notification) Category() string {
	switch n.CategoryID {
	case 1:
		return `成长档案`
	case 2:
		return `体格检查表`
	case 3:
		return `仓库管理`
	case 4:
		return `学校资产管理`
	case 5:
		return `营养膳食`
	case 6:
		return `出勤管理`
	case 7:
		return `系统通知`
	default:
		return ``
	}
}
