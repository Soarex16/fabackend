package app

import (
	"github.com/soarex16/fabackend/handlers"
	"github.com/soarex16/fabackend/middlewares"
	"github.com/soarex16/fabackend/routes"
	"net/http"
)

// InitHandlers - returns all routes for registration in router
func InitHandlers(app *App) *routes.Routes {
	// instantiate all handlers
	index := handlers.IndexHandler{}

	achievements := handlers.AchievementsHandler{
		Achievements: app.Store.Achievements,
		Users:        app.Store.Users,
	}

	courses := handlers.CoursesHandler{
		Courses: app.Store.Courses,
	}

	auth := handlers.AuthHandler{
		JwtSecret: app.Config.JwtSecret,
		Users:     app.Store.Users,
		Sessions:  app.Store.Sessions,
	}

	users := handlers.UsersHandler{
		Users: app.Store.Users,
	}

	sessions := app.Store.Sessions

	return &routes.Routes{
		routes.Route{Name: "Index", Method: "GET", Pattern: "/", Handler: http.HandlerFunc(index.Index)},

		routes.Route{
			Name:    "AddUserAchievement",
			Method:  "POST",
			Pattern: "/users/:username/achievements",
			Private: true,
			Handler: middlewares.Authorization(http.HandlerFunc(achievements.AddUserAchievement), sessions),
		},
		routes.Route{
			Name:    "GetUserAchievements",
			Method:  "GET",
			Pattern: "/users/:username/achievements",
			Private: true,
			Handler: middlewares.Authorization(http.HandlerFunc(achievements.GetUserAchievements), sessions),
		},

		routes.Route{Name: "GetAllCources", Method: "GET", Pattern: "/courses", Handler: http.HandlerFunc(courses.GetAllCources)},

		routes.Route{Name: "LoginUser", Method: "POST", Pattern: "/auth/login", Handler: http.HandlerFunc(auth.Login)},
		routes.Route{Name: "RefreshTokens", Method: "POST", Pattern: "/auth/refresh", Handler: http.HandlerFunc(auth.RefreshTokens)},

		routes.Route{
			Name:    "CreateUser",
			Method:  "POST",
			Pattern: "/users",
			Handler: http.HandlerFunc(users.CreateUser),
		},
		routes.Route{
			Name:    "GetUserByName",
			Method:  "GET",
			Pattern: "/users/:username",
			Private: true,
			Handler: middlewares.Authorization(http.HandlerFunc(users.GetUserByName), sessions),
		},
		routes.Route{
			Name:    "UpdateUser",
			Method:  "PUT",
			Pattern: "/users/:username",
			Private: true,
			Handler: middlewares.Authorization(http.HandlerFunc(users.UpdateUser), sessions),
		},
		routes.Route{
			Name:    "DeleteUser",
			Method:  "DELETE",
			Pattern: "/users/:username",
			Private: true,
			Handler: middlewares.Authorization(http.HandlerFunc(users.DeleteUser), sessions),
		},
	}
}
