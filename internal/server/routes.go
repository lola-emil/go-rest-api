package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"example.com/contact/internal/auth"
	"example.com/contact/internal/contact"
	"example.com/contact/internal/pkg/middleware"
	"example.com/contact/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

// SPAHandler serves a single page application.
func SPAHandler(staticPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Join internally call path.Clean to prevent directory traversal
		path := filepath.Join(staticPath, r.URL.Path)

		// check whether a file exists or is a directory at the given path
		fi, err := os.Stat(path)
		if os.IsNotExist(err) || fi.IsDir() {

			// set cache control header to prevent caching
			// this is to prevent the browser from caching the index.html
			// and serving old build of SPA App
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

			// file does not exist or path is a directory, serve index.html
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}

		if err != nil {
			// if we got an error (that wasn't that the file doesn't exist) stating the
			// file, return a 500 internal server error and stop
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set cache control header to serve file for a year
		// static files in this case need to be cache busted
		// (usualy by appending a hash to the filename)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

		// otherwise, use http.FileServer to serve the static file
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
	})
}

func (s *Server) RegisterRoutes() *chi.Mux {
	viteURL, _ := url.Parse("http://localhost:5173")
	viteProxy := httputil.NewSingleHostReverseProxy(viteURL)

	r := chi.NewRouter()

	// r.Use(csrf.Protect(
	// 	[]byte("32-byte-long-auth-key"),
	// 	csrf.Secure(true),
	// 	csrf.SameSite(csrf.SameSiteStrictMode),
	// ))

	r.Get("/csrf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.WriteHeader(http.StatusNoContent)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Mount("/", auth.RegisterModule(s.db.GetInstance()))
	})

	r.Route("/api", func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)

		r.Mount("/users", user.RegisterModule(s.db.GetInstance()))
		r.Mount("/contacts", contact.RegisterModule(s.db.GetInstance()))
	})

	// Serve SPA
	if os.Getenv("ENV") == "dev" {
		r.Handle("/*", viteProxy)
	} else {
		r.Handle("/*", SPAHandler("./frontend/dist"))
	}

	return r
}

func (s *Server) CorsMiddlewareWrapper(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-TOKEN")
		w.Header().Set("Access-Control-Allow-Credentials", "false")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})

}
