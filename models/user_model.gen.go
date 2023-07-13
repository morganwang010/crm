// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserModel represents a user model.
type UserModel struct {
	db *gorm.DB
}

// User represents a user struct data.
type User struct {
	Id       uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`         // The username
	Password string    `gorm:"column:password" json:"password"` // The  user password
	Mobile   string    `gorm:"column:mobile" json:"mobile"`     // The mobile phone number
	Gender   string    `gorm:"column:gender" json:"gender"`     // gender,male|female|unknown
	Nickname string    `gorm:"column:nickname" json:"nickname"` // The nickname
	Type     int8      `gorm:"column:type" json:"type"`         // The user type, 0:normal,1:vip, for test golang keyword
	CreateAt time.Time `gorm:"column:create_at" json:"createAt"`
	UpdateAt time.Time `gorm:"column:update_at" json:"updateAt"`
}

// FindOneWhereParameter is a where parameter structure.
type FindOneWhereParameter struct {
	IdEqual uint64
}

// FindLimitWhereParameter is a where parameter structure.
type FindLimitWhereParameter struct {
	IdGT uint64
}

// FindLimitLimitParameter is a limit parameter structure.
type FindLimitLimitParameter struct {
	Count int
}

// CountWhereParameter is a where parameter structure.
type CountWhereParameter struct {
	IdGT uint64
}

// CountResult is a count result.
type CountResult struct {
	Count sql.NullInt64 `gorm:"column:count" json:"count"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (CountResult) TableName() string {
	return "user"
}

// UpdateOneWhereParameter is a where parameter structure.
type UpdateOneWhereParameter struct {
	IdEqual uint64
}

// DeleteOneWhereParameter is a where parameter structure.
type DeleteOneWhereParameter struct {
	IdEqual uint64
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	db := m.db.WithContext(ctx)
	list := data[:]
	return db.Create(&list).Error
}

// FindOne is generated from sql:
// select * from user where id = ? limit 1;
func (m *UserModel) FindOne(ctx context.Context, where FindOneWhereParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindLimit is generated from sql:
// select * from user where id > ? limit ?;
func (m *UserModel) FindLimit(ctx context.Context, where FindLimitWhereParameter, limit FindLimitLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// Count is generated from sql:
// select count(id) AS count from user where id > ?;
func (m *UserModel) Count(ctx context.Context, where CountWhereParameter) (*CountResult, error) {
	var result = new(CountResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`count(id) AS count`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// UpdateOne is generated from sql:
// update user set name = ? where id = ?;
func (m *UserModel) UpdateOne(ctx context.Context, data *User, where UpdateOneWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Model(&User{})
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Updates(map[string]interface{}{
		"name": data.Name,
	})
	return db.Error
}

// DeleteOne is generated from sql:
// delete from user where id = ? limit 1;
func (m *UserModel) DeleteOne(ctx context.Context, where DeleteOneWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Limit(1)
	db = db.Delete(&User{})
	return db.Error
}