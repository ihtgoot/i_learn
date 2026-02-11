package repository

import (
	"time"

	"github.com/ihtgoot/i_learn/Section_3/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertBanglowRestriction(r models.BanglowRestriction) error
	SearchAvailibilityByDate(start time.Time, end time.Time, banglowID int) (bool, error)
	SearchAvailibilityByDateForAllBanglows(start, end time.Time) ([]models.Banglow, error)
	GetBanglowByID(id int) (models.Banglow, error)
}
