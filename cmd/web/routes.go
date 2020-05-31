package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.home)))))
	mux.Get("/about", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.about)))))
	mux.Get("/snippet/create", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.createSnippetForm))))))
	mux.Post("/snippet/create", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.createSnippet))))))
	mux.Get("/snippet/:id", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.showSnippet)))))

	mux.Get("/user/signup", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.signupUserForm)))))
	mux.Get("/user/profile", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.userProfile))))))
	mux.Post("/user/signup", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.signupUser)))))
	mux.Get("/user/login", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.loginUserForm)))))
	mux.Post("/user/login", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.loginUser)))))
	mux.Post("/user/logout", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.logoutUser))))))
	mux.Get("/user/change-password", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.changePasswordForm))))))
	mux.Post("/user/change-password", app.session.Enable(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.changePassword))))))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(secureHeaders(mux))
}
