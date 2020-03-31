package routes

import (
	"fmt"
	"net/http"

	ctx "github.com/soarex16/fabackend/context"
	"github.com/soarex16/fabackend/handlers"
)

// Route - type for defining all route information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Private     bool
	HandlerFunc http.HandlerFunc
}

// Routes - collection of routes
type Routes []Route

// GetAll - returns all routes for registration in router
func GetAll(appCtx *ctx.Context) *Routes {
	var routes = Routes{
		Route{"Index", "GET", "/", false, Index},
		Route{"AddUserAchievement", "POST", "/achievements/{userId}", true, handlers.AddUserAchievement(appCtx)},
		Route{"GetUserAchievements", "GET", "/achievements/{userId}", true, handlers.GetUserAchievements(appCtx)},

		Route{"GetAllCources", "GET", "/courses", false, handlers.GetAllCources(appCtx)},

		Route{"LoginUser", "POST", "/users/login", false, handlers.LoginUser(appCtx)},
		Route{"RefreshTokens", "POST", "/users/login/refresh", false, handlers.RefreshTokens(appCtx)},

		Route{"CreateUser", "POST", "/users", false, handlers.CreateUser(appCtx)},
		Route{"GetUserByName", "GET", "/users/{username}", true, handlers.GetUserByName(appCtx)},
		Route{"UpdateUser", "PUT", "/users/{username}", true, handlers.UpdateUser(appCtx)},
		Route{"DeleteUser", "DELETE", "/users/{username}", true, handlers.DeleteUser(appCtx)},
	}

	return &routes
}

//TODO: в авторизацию надо пробрасывать jwt secret

// Index - stub for index route
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
