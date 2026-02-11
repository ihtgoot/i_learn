package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/driver"
	"github.com/ihtgoot/i_learn/Section_3/internal/form"
	"github.com/ihtgoot/i_learn/Section_3/internal/helper"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/ihtgoot/i_learn/Section_3/internal/rendrer"
	"github.com/ihtgoot/i_learn/Section_3/internal/repository"
	"github.com/ihtgoot/i_learn/Section_3/internal/repository/dbrepo"
)

// get ip of user
func getClintIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	fmt.Println(r.RemoteAddr)
	ip := net.ParseIP(host)
	if ip == nil {
		fmt.Println(ip)
		return host
	}

	if ip.To4() != nil {
		fmt.Println(ip.To4())
		return host
	}
	fmt.Println(ip)
	return host
}

// Reepository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// repo the repository used by the handler
var Repo *Repository

// newwwRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQl, a),
	}
}

// New Handlers set the repsitory for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// handeler of home page
func (h *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// h.DB.AllUsers() just test

	// store remote IP of user
	remoteIP := getClintIP(r)
	h.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	rendrer.Template(w, r, "home-page.tpml", &models.TemplateData{})
}

// handels about page
func (a *Repository) About(w http.ResponseWriter, r *http.Request) {
	sidekickmap := make(map[string]string)
	sidekickmap["morty"] = "ooh , wee"
	remoteIP := a.App.Session.GetString(r.Context(), "remote_ip")
	sidekickmap["remote_ip"] = remoteIP
	rendrer.Template(w, r, "about-page.tpml", &models.TemplateData{})
}

// handels Contact page
func (a *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	rendrer.Template(w, r, "contact-page.tpml", &models.TemplateData{})
}

// handels Erermite page
func (a *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	rendrer.Template(w, r, "eremite-page.tpml", &models.TemplateData{})
}

// handels Couple page
func (a *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	rendrer.Template(w, r, "couple-page.tpml", &models.TemplateData{})
}

// handels Family page
func (a *Repository) Family(w http.ResponseWriter, r *http.Request) {
	rendrer.Template(w, r, "family-page.tpml", &models.TemplateData{})
}

// handels Reservation page GET request
func (a *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	rendrer.Template(w, r, "check-availability-page.tpml", &models.TemplateData{})
}

// handels Reservation page POST request
func (a *Repository) POST_Reservation(w http.ResponseWriter, r *http.Request) {

	sd := r.Form.Get("startingDate")
	ed := r.Form.Get("endingDate")
	fmt.Println(sd, ed)
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	if endDate.Before(startDate) {
		http.Error(w, "end before start", http.StatusBadRequest)
		return
	}

	banglows, err := a.DB.SearchAvailibilityByDateForAllBanglows(startDate, endDate)
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println(err)
		return
	}

	fmt.Println(startDate, endDate, banglows)

	if len(banglows) == 0 {
		a.App.Session.Put(r.Context(), "error", ":( nothing available")
		fmt.Println(":( nothing available")
		http.Redirect(w, r, "/reservation", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["banglows"] = banglows

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	a.App.Session.Put(r.Context(), "reservation", res)

	rendrer.Template(w, r, "choose-bunglow-page.tpml", &models.TemplateData{
		Data: data,
	})
	//w.Write([]byte(fmt.Sprintf("Success you fucking send something, start : %s, end : %s , %v", startDate, endDate, banglows)))
}

// handels resrervation json it retur s json (json : javascript obect resource)

type jsonResponse struct {
	OK         bool   `json:"ok"`
	Message    string `json:"messgae`
	Bangloe_id string `json:"banglow_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

func (a *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {

	banglow_id, err := strconv.Atoi(r.Form.Get("banglow_id"))
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println("git nothing ", err)
		return
	}

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println("start : nothing ", err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println("end :nothing ", err)
		return
	}

	avail, err := a.DB.SearchAvailibilityByDate(startDate, endDate, banglow_id)
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println("sql chuda ", err)
		resp := jsonResponse{
			OK:      false,
			Message: "error querying db",
		}
		output, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			helper.ServerError(w, err)
			fmt.Println("chuda  response", err)
			return
		}
		w.Header().Set("content-Type", "application/json") // tell browser what the dtatype is o that it has a header
		w.Write(output)
		return
	}

	resp := jsonResponse{
		OK:         avail,
		Message:    "",
		StartDate:  sd,
		EndDate:    ed,
		Bangloe_id: strconv.Itoa(banglow_id),
	}

	output, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helper.ServerError(w, err)
		fmt.Println("chuda  response", err)
		return
	}

	// log.Println(string(out))
	w.Header().Set("content-Type", "application/json") // tell browser what the dtatype is o that it has a header
	w.Write(output)
	return
}

// handels make-reservation page get request
func (a *Repository) Make_Reservation(w http.ResponseWriter, r *http.Request) {

	res, ok := a.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		a.App.Session.Put(r.Context(), "error", "Missing Paramaert in URL")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	banglow, err := a.DB.GetBanglowByID(res.BanglowId)
	if err != nil {
		fmt.Println("Missing Paramaert in URL")
		a.App.Session.Put(r.Context(), "error", "Missing Paramaert in URL")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Banglow.BanglowName = banglow.BanglowName

	a.App.Session.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")

	stringsMap := make(map[string]string)
	stringsMap["start_date"] = sd
	stringsMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	rendrer.Template(w, r, "make-reservation-page.tpml", &models.TemplateData{
		Form:      form.New(nil),
		Data:      data,
		StringMap: stringsMap,
	})
}

// handels make-reservationPOSt page psot request : it hadels pist request fom reservation form
func (a *Repository) Make_ReservationPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	//err = errors.New("Epic fail") //test error logging
	if err != nil {
		log.Println(err)
		helper.ServerError(w, err)
		return
	}

	sd := r.Form.Get("startingDate")
	ed := r.Form.Get("endingDate")

	// Date and time in go:
	// 2023-12-31 -- 01/02 03:04:05PM -- `06 -0700 -- 12/31 03:04:05PM `23 -0700

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	banglowID, err := strconv.Atoi(r.Form.Get("vacationHome"))

	reservation := models.Reservation{
		Name:      r.Form.Get("full_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		BanglowId: banglowID,
	}

	form := form.New(r.PostForm)
	// That list is how the form knows it only needs those three fields.
	// The posted form (r.PostForm) may contain more fields, but Required only checks the names you pass it.
	// Add/remove names in that call to change required fields.
	form.Required("full_name", "email", "phone")
	//form.Has("full_name")
	// form.Has("email")
	// form.Has("phone")
	form.MinLength("full_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		rendrer.Template(w, r, "make-reservation-page.tpml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	log.Println(reservation)
	newReservationID, err := a.DB.InsertReservation(reservation)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	log.Println(newReservationID)
	restriction := models.BanglowRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		BanglowId:     banglowID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}
	log.Println(restriction)

	err = a.DB.InsertBanglowRestriction(restriction)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	a.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-overview", http.StatusSeeOther)
}

// ReservationOverview displays the reservation summery page
func (a *Repository) ReservationOverview(w http.ResponseWriter, r *http.Request) {

	reservation, ok := a.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		//a.App.ErrorLog.Println("could not get session")
		a.App.Session.Put(r.Context(), "error", "no reservation in session")
		log.Println("could not get item from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	a.App.Session.Remove(r.Context(), "reservation")

	banglow, err := a.DB.GetBanglowByID(reservation.BanglowId)
	if err != nil {
		fmt.Println("Missing Paramaert in URL")
		a.App.Session.Put(r.Context(), "error", "Missing Paramaert in URL")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	reservation.Banglow.BanglowName = banglow.BanglowName

	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")

	stringsMap := make(map[string]string)
	stringsMap["start_date"] = sd
	stringsMap["end_date"] = ed

	rendrer.Template(w, r, "reservation-overview-page.tpml", &models.TemplateData{
		Data:      data,
		StringMap: stringsMap,
	})
}

// displays the list of available banglow and lets the user choose a banglow
func (a *Repository) ChooseBanglow(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	banglowID, err := strconv.Atoi(exploded[2])
	if err != nil {
		fmt.Println("Missing Paramaert in URL 1")
		a.App.Session.Put(r.Context(), "error", "Missing Paramaert in URL")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res, ok := a.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		fmt.Println("Missing Paramaert in URL 2")
		a.App.Session.Put(r.Context(), "error", "Missing Paramaert in URL")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	res.BanglowId = banglowID

	a.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

}

// takes url parameter from request -> build reservation -> redire t to make-reservation
func (a *Repository) BookBanglow(w http.ResponseWriter, r *http.Request) {
	log.Println("FULL URL:", r.URL.String())
	banglow_ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		log.Print(err)
		helper.ServerError(w, err)
		return
	}

	var res models.Reservation

	banglow, err := a.DB.GetBanglowByID(banglow_ID)
	if err != nil {
		a.App.Session.Put(r.Context(), "error", "cannot find banglow")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Banglow.BanglowName = banglow.BanglowName
	res.BanglowId = banglow_ID
	res.StartDate = startDate
	res.EndDate = endDate

	a.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

// repo pattern : make use of packages across applications , this hellps us in developement : we can centrally turn our cache on/off

// ❯ go run ./
// package github.com/ihtgoot/i_learn/Section_3/cmd/web
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/rendrer
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers: import cycle not allowed
