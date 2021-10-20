package api

import (
	"database/sql/driver"
	"encoding/json"
	model "gorestapi/pkg/model/author"
	repository "gorestapi/pkg/repository/author"
	author "gorestapi/pkg/service/author"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func dbSetup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	DB, _ := gorm.Open("postgres", db)
	DB.LogMode(true)
	return DB, mock
}

func routersSetup(api AuthorAPI) *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/authors", api.FindAllAuthors()).Methods("GET")
	return apiRouter
}
func apiSetup(db *gorm.DB) AuthorAPI {
	postRepository := repository.NewAuthor(db)
	postService := author.NewAuthorService(postRepository)
	postAPI := NewAuthorAPI(postService)
	// postAPI.Migrate()
	return postAPI
}
func TestFindAllPosts(t *testing.T) {
	w := httptest.NewRecorder()

	// Initialize mock db
	mockDB, mock := dbSetup()

	// Initialize api
	api := apiSetup(mockDB)

	// Initialize router
	r := routersSetup(api)

	// Generate dummy post
	var posts []model.Author
	post := model.Author{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		FullName:  "Dummy Title",
		Email:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	posts = append(posts, post)

	// Add dummy post to dummy db table
	rows := sqlmock.
		NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "full_name", "email"}).
		AddRow(post.ID, post.CreatedAt, post.UpdatedAt, post.DeletedAt, post.FullName, post.Email)

	// Define expected query
	const sqlSelectOne = `SELECT * FROM "posts"`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
		WillReturnRows(rows)

	// Make request
	r.ServeHTTP(w, httptest.NewRequest("GET", "/authors", nil))

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code, "Did not get expected HTTP status code, got")

	// Unmarshall response
	var resultPosts []model.Author
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&resultPosts); err != nil {
		t.Error(err)
	}
	resultPosts[0].FullName = "Dummy Title Ankara"

	// Compare response and table data
	assert.Nil(t, deep.Equal(posts, resultPosts))

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
