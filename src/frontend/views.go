package frontend

import (
	"aion/config"
	"aion/datasources/nasa"
	"io"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func renderTemplate(w io.Writer, name string, data interface{}, c echo.Context, only bool) error {
	baseTemps := []string{
		"frontend/templates/base.html",
		"frontend/templates/sidebar.html",
	}
	if only {
		baseTemps = []string{
			"frontend/templates/" + name + ".html",
		}
	} else {
		baseTemps = append(baseTemps, "frontend/templates/"+name+".html")
	}
	tmpl, err := template.ParseFiles(baseTemps...)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}
func StartEchoServer() {
	e := echo.New()
	subgroup := e.Group("/app")

	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	subgroup.GET("/nasa-apod/", listNasaApodFiles)
	// http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir(filepath.Join(config.BaseDir, config.BaseDataDir)))))
	e.Static("/data/", filepath.Join(config.BaseDir, config.BaseDataDir))

	e.Logger.Fatal(e.Start(":1323"))
}

func listNasaApodFiles(c echo.Context) error {
	nasaPhotos, err := nasa.GetNasaApod()
	if err != nil {
		return err
	}
	return renderTemplate(c.Response(), "nasa-apod", map[string]interface{}{
		"photos": nasaPhotos,
	}, c, false)
}
