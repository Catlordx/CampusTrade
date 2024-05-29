package mysql

import (
	"gorm.io/gorm"
	"time"
)

// Commodity 映射商品表
type Commodity struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Owner       string
	Img         string
}

// Order 映射订单表
type Order struct {
	OrderID       uint           `gorm:"primaryKey;autoIncrement"`
	CustomerID    uint           `gorm:"not null"`
	OrderDate     time.Time      `gorm:"not null"`
	TotalAmount   float64        `gorm:"not null"`
	Status        string         `gorm:"type:enum('pending','processing','shipped','completed','cancelled');default:'pending';not null"`
	PaymentMethod string         `gorm:"size:50"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// OrderItem 订单商品关联表
type OrderItem struct {
	ItemID    uint           `gorm:"primaryKey;autoIncrement"`
	OrderId   uint           `gorm:"not null"`
	ProductId uint           `gorm:"not null"`
	Quantity  uint           `gorm:"not null"`
	Price     float64        `gorm:"not null"`
	Order     Order          `gorm:"foreignKey:OrderID"`
	Product   Commodity      `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User 用户表模型
type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	RealName string `gorm:"not null"`
	//Password    string `gorm:"not null"`
	Password    []byte `gorm:"not null"`
	PhoneNumber string
	Role        string `gorm:"not null"`
}

// RolePermission 角色权限关联表
type RolePermission struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Role       string `gorm:"not null"`
	Permission string `gorm:"not null"`
}
