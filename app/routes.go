package app

import (
	"github.com/soarex16/fabackend/handlers"
	"github.com/soarex16/fabackend/routes"
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
	}

	users := handlers.UsersHandler{
		Users: app.Store.Users,
	}

	return &routes.Routes{
		routes.Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: index.Index},
		routes.Route{Name: "AddUserAchievement", Method: "POST", Pattern: "/users/:username/achievements", Private: true, HandlerFunc: achievements.AddUserAchievement},
		routes.Route{Name: "GetUserAchievements", Method: "GET", Pattern: "/users/:username/achievements", Private: true, HandlerFunc: achievements.GetUserAchievements},

		routes.Route{Name: "GetAllCources", Method: "GET", Pattern: "/courses", HandlerFunc: courses.GetAllCources},

		routes.Route{Name: "LoginUser", Method: "POST", Pattern: "/auth/login", HandlerFunc: auth.Login},
		routes.Route{Name: "RefreshTokens", Method: "POST", Pattern: "/auth/refresh", HandlerFunc: auth.RefreshTokens},

		routes.Route{Name: "CreateUser", Method: "POST", Pattern: "/users", HandlerFunc: users.CreateUser},
		routes.Route{Name: "GetUserByName", Method: "GET", Pattern: "/users/:username", Private: true, HandlerFunc: users.GetUserByName},
		routes.Route{Name: "UpdateUser", Method: "PUT", Pattern: "/users/:username", Private: true, HandlerFunc: users.UpdateUser},
		routes.Route{Name: "DeleteUser", Method: "DELETE", Pattern: "/users/:username", Private: true, HandlerFunc: users.DeleteUser},
	}
}
