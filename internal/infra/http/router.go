package http

import (
	"fmt"
	"github.com/test_server/pkg/dbsettings"
	"github.com/upper/db/v4/adapter/postgresql"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			AddEventRoutes(&apiRouter, eventController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	//router.Get("/create", func(w http.ResponseWriter, r *http.Request) {
	//	t, err := template.ParseFiles("templates/create.html")
	//	if err != nil {
	//		fmt.Printf(err.Error())
	//	}
	//	t.Execute(w, nil)
	//})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouter.Get("/create", func(w http.ResponseWriter, r *http.Request) {
			t, err := template.ParseFiles("templates/create.html")
			if err != nil {
				fmt.Printf(err.Error())
			}
			t.Execute(w, nil)
		})
		apiRouter.Post("/save_event", func(w http.ResponseWriter, r *http.Request) {
			//Отримуємо данні з форми і записуємо їх у змінну
			name := r.FormValue("name")
			//Підключаємося до БД
			sess, err := postgresql.Open(dbsettings.Settings)
			if err != nil {
				fmt.Println("Open: ", err)
			}
			//Відкладенне відключення до БД
			defer sess.Close()
			//Добавление записи
			insert, err := sess.SQL().Query(fmt.Sprintf("INSERT INTO events (name) VALUES ('%s')", name))
			if err != nil {
				fmt.Println(err)
			}
			defer insert.Close()
			//Сторінка успіху
			t, err := template.ParseFiles("templates/saved.html")
			if err != nil {
				fmt.Printf(err.Error())
			}
			t.Execute(w, nil)

		})
	})
}
