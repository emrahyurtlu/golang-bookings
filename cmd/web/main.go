package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/emrahyurtlu/go-course/cmd/pkg/config"
	"github.com/emrahyurtlu/go-course/cmd/pkg/handlers"
	"github.com/emrahyurtlu/go-course/cmd/pkg/render"
)

const portNumber string = ":8090"
var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime =  24 * time.Hour
	// All the session is stored in the cookies.
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Require https
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	message := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(message)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}