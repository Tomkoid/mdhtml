package httpserver

import (
	"embed"
	"fmt"

	"codeberg.org/Tomkoid/mdhtml/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets/default.css
//go:embed assets/reload.js
//go:embed assets/prism.js
//go:embed assets/prism.css
var f embed.FS
var defaultCSSData, defaultCSSDataErr = f.ReadFile("assets/default.css")
var reloadData, reloadDataErr = f.ReadFile("assets/reload.js")
var prismJSData, prismJSErr = f.ReadFile("assets/prism.js")
var prismCSSData, prismCSSErr = f.ReadFile("assets/prism.css")

var History = []string{}

func SetReload() {
	History = append(History, "reload")
}

func HttpServer(args models.Args) {
	app := echo.New()
	app.HideBanner = true
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	if args.Debug {
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${method}: ${uri}, status=${status}\n",
		}))
	}

	// Create a new router
	router := app.Group("") // root group

	setupRoutes(router, args)

	app.Logger.Fatal(app.Start(fmt.Sprintf("%s:%d", args.ServerHostname, args.ServerPort)))
}
