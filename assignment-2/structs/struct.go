package structs

import (
	"time"
)

type Order struct {
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `gorm:"type:VARCHAR(50);not null" json:"customerName"`
	OrderedAt    time.Time `gorm:"type:timestamp;autoCreateTime" json:"orderedAt"`
	Items        []Items   `json:"items" gorm:"foreignKey:order_id;constraint:OnDelete:CASCADE"`
}

type Items struct {
	ItemId      uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `gorm:"null" json:"OrderId"`
}
