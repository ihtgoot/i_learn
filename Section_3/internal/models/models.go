package models

import "time"

// // Reservation cotains reservation data
// type Reservation struct {
// 	Name  string
// 	Email string
// 	Phone string
// }

// user is the model of user data
type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// banglow is the model of banglow data
type Banglow struct {
	ID          int
	BanglowName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// restriction is the model of restriction data
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// reservations is the model of reservations data
type Reservation struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Password  string
	StartDate time.Time
	EndDate   time.Time
	BanglowId int
	CreatedAt time.Time
	UpdatedAt time.Time
	Banglow   Banglow
	Processed int
}

// banglowRestriction is the model of banglowRestriction data
type BanglowRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	BanglowId     int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Restriction   Restriction
	Banglow       Banglow
}
