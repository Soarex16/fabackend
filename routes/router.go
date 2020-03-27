package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/soarex16/fabackend/middlewares"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	//HandlerFunc http.HandlerFunc
	//Private bool
}

type Routes []Route

func New() *mux.Router {
	router := mux.NewRouter()

	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	router.Use(middlewares.Logger)

	router.
		Methods("GET").
		Path("/").
		Name("index").
		Handler(handler)

	return router
	// TODO: в хендлеры надо как-то изящно пробрасывать контекст приложения, который содержит
	// конфигурацию
	// контекст бд
	// токены юзеров?
	// может как-то завернуть хендлеры? видел такое в https://talks.golang.org/2012/
	// как-то надо сделать struct embedding

	// очень надо постараться изолироваться от фреймворка

	// контекст приложения
	// контекст контроллера
	// у нас должен быть контекст запроса
}

var routes = Routes{
	Route{
		"AddUserAchievement",
		strings.ToUpper("Post"),
		"/achievements/{userId}",
		//AddUserAchievement,
	},

	Route{
		"GetUserAchievements",
		strings.ToUpper("Get"),
		"/achievement/{userId}",
		//GetUserAchievements,
	},

	Route{
		"GetAllCources",
		strings.ToUpper("Get"),
		"/courses",
		//GetAllCources,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/users/login",
		//LoginUser,
	},

	Route{
		"RefreshTokens",
		strings.ToUpper("Post"),
		"/users/token/refresh",
		//RefreshTokens,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/users",
		//CreateUser,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/users/{username}",
		//DeleteUser,
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/users/{username}",
		//GetUserByName,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/users/login",
		//LoginUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/users/{username}",
		//UpdateUser,
	},
}
