package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/user", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		return context.JSON(http.StatusCreated, u)
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)

	// static files
	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start("0.0.0.0:5555"))
}

func getUser(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func show(ctx echo.Context) error {
	team := ctx.QueryParam("team")
	member := ctx.QueryParam("member")
	return ctx.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func save(ctx echo.Context) error {
	name := ctx.FormValue("name")
	avatar, err := ctx.FormFile("avatar")
	if err != nil {
		return err
	}
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != err {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return ctx.HTML(http.StatusOK, "<b>Thank you!" + name + "</br>")
}