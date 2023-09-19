package main

import (
	"fmt"
	"time"
)

type Member struct {
	ID                          int       `json:"id"`
	Email                       string    `json:"name"`
	Password                    string    `json:"password"`
	Title                       string    `json:"title"`
	FirstName                   string    `json:"first_name"`
	LastName                    string    `json:"last_name"`
	AreaCode                    string    `json:"area_code"`
	Phone                       string    `json:"phone"`
	BirthdayMonth               int       `json:"birthday_month"`
	Lang                        int       `json:"lang"`
	IsEmailSubscription         bool      `json:"is_email_subscription"`
	VipNo                       string    `json:"vip_no"`
	MiraplusMembershipID        int       `json:"miraplus_membership_id"`
	TsuiHangVillageMembershipID int       `json:"tsuihangvillage_membership_id"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	DeletedAt                   time.Time `json:"deleted_at"`
}

func (Member) TableName() string {
	return `members`
}

func (a Member) Example() Member {
	fmt.Println("working")

	return a
}

func (a Member) FindOne() Member {

	db := NewDatabase()
	connDB := db.ConnDB()
	sqlDB, _ := connDB.DB()
	defer sqlDB.Close()

	var member Member

	connDB.Find(&member, Member{ID: a.ID})

	return member
}
