package dbrepo

import (
	"context"
	"time"

	"github.com/ihtgoot/i_learn/Section_3/internal/models"
)

const dbTimeout = 3 * time.Second

// implement all the function that DatabaseRepo requires

func (m *postgresDBrepo) AllUsers() bool {
	return true
}

func (m *postgresDBrepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int

	stmt := `
		insert into reservation 
			(full_name,email,phone,start_date,end_date,banglow_id,created_at,updated_at)		
			values
				($1, $2, $3, $4, $5, $6, $7, $8) returning id
		`
	err := m.DB.QueryRowContext(ctx,
		stmt,
		res.Name,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.BanglowId,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// insertbanglow restriction places restriction into the databse
func (m *postgresDBrepo) InsertBanglowRestriction(r models.BanglowRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		insert into "banglowRestriction" 
			(start_date,end_date,banglow_id,reservation_id,created_at,updated_at,restriction_id)
		values
			($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := m.DB.ExecContext(ctx,
		stmt,
		r.StartDate,
		r.EndDate,
		r.BanglowId,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil

}

// searchavailibility finds reservation in a timeframe , true if there is availibiity false if not and is done for banglow basis
func (m *postgresDBrepo) SearchAvailibilityByDate(start time.Time, end time.Time, banglowID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			count(id)
		from
			"banglowRestriction"
		where 
			banglow_id = $1
			and 
			$2 <= end_date  
			and
			$3 >= start_date
	`

	var numRow int
	row := m.DB.QueryRowContext(ctx, query, banglowID, start, end)
	err := row.Scan(&numRow)
	if err != nil {
		return false, err
	}

	if numRow == 0 {
		return true, nil
	}

	return false, nil

}

func (m *postgresDBrepo) SearchAvailibilityByDateForAllBanglows(start, end time.Time) ([]models.Banglow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
	 		b.id , b.banglow_name
		from 
			"banglow" b
		where 
			b.id not in  
					(select 
						banglow_id 
					from
						"banglowRestriction" br
					where
						br.end_date >= $1
						and 
						br.start_date <=  $2
					);
	`
	var banglows []models.Banglow

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return banglows, err
	}

	for rows.Next() {
		var banglow models.Banglow

		err := rows.Scan(
			&banglow.ID,
			&banglow.BanglowName,
		)
		if err != nil {
			return banglows, nil
		}

		banglows = append(banglows, banglow)
	}

	if err = rows.Err(); err != nil {
		return banglows, err
	}

	return banglows, nil
}

// GEtBanglowById gets a bunglow by id
func (m *postgresDBrepo) GetBanglowByID(id int) (models.Banglow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			id, banglow_name, created_at, updated_at
		from
			banglow
		where
			id = $1
	`
	var banglow models.Banglow

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&banglow.ID,
		&banglow.BanglowName,
		&banglow.CreatedAt,
		&banglow.UpdatedAt,
	)
	if err != nil {
		return banglow, err
	}
	return banglow, nil
}
