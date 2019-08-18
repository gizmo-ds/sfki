package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"sfki/model"

	"github.com/labstack/echo"
)

var (
	config struct {
		Addr      string
		AccessKey string
	}
)

func init() {
	bytes, err := ioutil.ReadFile(filepath.Join(model.ROOT, "config/config.yaml"))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.POST("/graphql", func(c echo.Context) error {
		var form struct {
			Query string `json:'query'`
		}
		if err := c.Bind(&form); err != nil {
			return err
		}
		c.JSON(200, model.ExecuteQuery(form.Query))
		return nil
	})
	e.POST("/update", func(c echo.Context) error {
		if c.FormValue("access_key") != config.AccessKey {
			c.String(401, "Not Authorized")
			return nil
		}
		model.PostLoading()
		model.LinkLoading()
		model.AboutLoading()
		return nil
	})
	e.GET("/about", func(c echo.Context) error {
		c.JSON(200, model.About)
		return nil
	})
	e.Logger.Fatal(e.Start(config.Addr))
}
