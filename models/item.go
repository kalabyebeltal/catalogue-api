package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Item struct with correct JSON tags
type Item struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Taxcode     string    `json:"taxcode"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) { //tx is used to represent the current db transaction GORM is using when calling this function
	if item.ID == uuid.Nil {
		item.ID = uuid.New()
	}
	return
}

func MigrateItem(db *gorm.DB) error {
	return db.AutoMigrate(&Item{})
}

// make sure the tabe being accessed it 'item' not 'items'
func (Item) TableName() string {
	return "item"
}
