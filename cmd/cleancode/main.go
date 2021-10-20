package main

import (
	"fmt"
	api "gorestapi/pkg/api/author"
	bookApi "gorestapi/pkg/api/book"
	memberApi "gorestapi/pkg/api/member"
	repository "gorestapi/pkg/repository/author"
	bookRepository "gorestapi/pkg/repository/book"
	memberRepository "gorestapi/pkg/repository/member"
	service "gorestapi/pkg/service/author"
	bookService "gorestapi/pkg/service/book"
	memberService "gorestapi/pkg/service/member"
	"os"

	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	a := App{}

	// Initialize storage
	a.initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	// Initialize routes
	a.routes()

	// Start server
	a.run(":3000")
}

func (a *App) initialize(host, port, username, password, dbname string) {
	var err error

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// repository := post.NewRepository(a.DB)

	a.Router = mux.NewRouter()
}
func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
func (a *App) routes() {
	authorAPI, bookAPI, memberAPI := InitPostAPI(a.DB)
	a.Router.HandleFunc("/authors", authorAPI.FindAllAuthors()).Methods("GET")
	a.Router.HandleFunc("/authors", authorAPI.CreateAuthor()).Methods("POST")
	a.Router.HandleFunc("/authors/{id:[0-9]+}", authorAPI.FindById()).Methods("GET")
	a.Router.HandleFunc("/authors/{id:[0-9]+}", authorAPI.UpdateAuthor()).Methods("PUT")
	a.Router.HandleFunc("/authors/{id:[0-9]+}", authorAPI.DeleteAuthor()).Methods("DELETE")

	a.Router.HandleFunc("/books", bookAPI.FindAllBooks()).Methods("GET")
	a.Router.HandleFunc("/books", bookAPI.CreateBook()).Methods("POST")
	a.Router.HandleFunc("/books/{id:[0-9]+}", bookAPI.FindById()).Methods("GET")
	a.Router.HandleFunc("/books/{id:[0-9]+}", bookAPI.UpdateBook()).Methods("PUT")
	a.Router.HandleFunc("/books/{id:[0-9]+}", bookAPI.DeleteBook()).Methods("DELETE")

	a.Router.HandleFunc("/members", memberAPI.FindAllMembers()).Methods("GET")
	a.Router.HandleFunc("/members", memberAPI.CreateMember()).Methods("POST")
	a.Router.HandleFunc("/members/{id:[0-9]+}", memberAPI.FindById()).Methods("GET")
	a.Router.HandleFunc("/members/{id:[0-9]+}", memberAPI.UpdateMember()).Methods("PUT")
	a.Router.HandleFunc("/members/{id:[0-9]+}", memberAPI.DeleteMember()).Methods("DELETE")
}
func InitPostAPI(db *gorm.DB) (api.AuthorAPI, bookApi.BookAPI, memberApi.MemberAPI) {
	authorRepository := repository.NewAuthor(db)
	bookRepository := bookRepository.NewBook(db)
	memberRepository := memberRepository.NewMember(db)
	authorService := service.NewAuthorService(authorRepository)
	bookService := bookService.NewBookService(bookRepository)
	memberService := memberService.NewMemberService(memberRepository)
	authorAPI := api.NewAuthorAPI(authorService)
	bookAPI := bookApi.NewBookAPI(bookService)
	memberAPI := memberApi.NewMemberAPI(memberService)
	authorAPI.Migrate()
	bookAPI.Migrate()
	memberAPI.Migrate()
	return authorAPI, bookAPI, memberAPI
}
