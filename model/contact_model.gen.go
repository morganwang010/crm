// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ContactModel represents a contact model.
type ContactModel struct {
	db *gorm.DB
}

// Contact represents a contact struct data.
type Contact struct {
	Company   string    `gorm:"column:company" json:"company"`
	Focal     string    `gorm:"column:focal" json:"focal"`
	Mobile    string    `gorm:"column:mobile" json:"mobile"`
	Startdate time.Time `gorm:"column:startdate" json:"startdate"`
	Enddate   time.Time `gorm:"column:enddate" json:"enddate"`
	Mount     float64   `gorm:"column:mount" json:"mount"`
	Mail      string    `gorm:"column:mail" json:"mail"`
	OrgName   string    `gorm:"column:org_name" json:"orgName"`
	Id        int32     `gorm:"column:id;primaryKey" json:"id"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (Contact) TableName() string {
	return "contact"
}

// NewContactModel returns a new contact model.
func NewContactModel(db *gorm.DB) *ContactModel {
	return &ContactModel{db: db}
}

// Create creates  contact data.
func (m *ContactModel) Create(ctx context.Context, data ...*Contact) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	db := m.db.WithContext(ctx)
	list := data[:]
	return db.Create(&list).Error
}

// FindAllContact is generated from sql:
// select * from public.contact;
func (m *ContactModel) FindAllContact(ctx context.Context) ([]*Contact, error) {
	var result []*Contact
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Find(&result)
	return result, db.Error
}
