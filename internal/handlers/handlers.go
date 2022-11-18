package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/mahdiZarepoor/url-shortener/internal/repository"
	"github.com/mahdiZarepoor/url-shortener/pkg"
)

type App struct {
	Router *mux.Router
	RedisClient *redis.Client
}

func NewApp() App{
	router := mux.NewRouter()
	redisClient := repository.InitializeStore()
	app := App{
		Router : router,
		RedisClient: redisClient,
	}

	app.RegisterRoutes()
	return app
}

func (a *App)RegisterRoutes () {
	a.Router.Methods(http.MethodPost).Path("/register").HandlerFunc(a.HandleRegisterUrl)
	a.Router.Methods(http.MethodGet).Path("/{shortUrl}").HandlerFunc(a.HandleRetrieveUrl)
}

func (a *App)HandleRegisterUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if len(url) == 0 {
		jsonResponse(w, "url is not set", http.StatusBadRequest)
		return 
	}

	shortUrl := pkg.GetMD5Hash(url)
	repository.RegisterUrl(a.RedisClient, shortUrl, url)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"shortUrl": shortUrl})
}

func (a *App)HandleRetrieveUrl(w http.ResponseWriter, r *http.Request ) {
	shortUrl := mux.Vars(r)["shortUrl"]
	originalUrl , err := repository.RetrieveUrl(a.RedisClient, shortUrl)
	if err != nil {
		jsonResponse(w, "error in data retrieve", http.StatusInternalServerError)	
		return 
	}
	http.Redirect(w, r, originalUrl, http.StatusSeeOther)
}

