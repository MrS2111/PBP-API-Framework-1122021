package controllers

import (
    "my-app/app"
    "my-app/app/models"
    "github.com/revel/revel"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) TestDBConnection() revel.Result {
    var users []models.User
    if err := app.DB.Find(&users).Error; err != nil {
        return c.RenderText("Failed to connect to database: " + err.Error())
    }

    return c.RenderJSON(users)
}