package controllers

import (
	"github.com/revel/revel"
	"golang-bbs/app/controllers"
	"golang-bbs/app/models"
)

type ApiV1Terms struct {
	ApiV1Controller
}

func (c ApiV1Terms) Index() revel.Result {
	terms := []models.Term{}

	if err := controllers.DB.Order("id desc").Find(&terms).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	r := Response{terms}
	return c.RenderJSON(r)
}

func (c ApiV1Terms) Show(id int) revel.Result {
	term := &models.Term{}

	if err := controllers.DB.First(&term, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{term}
	return c.RenderJSON(r)
}

func (c ApiV1Terms) Create() revel.Result {
	term := &models.Term{}

	if err := c.Params.BindJSON(term); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := controllers.DB.Create(term).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	r := Response{term}
	return c.RenderJSON(r)
}

func (c ApiV1Terms) Delete(id int) revel.Result {
	term := models.Term{}

	if err := controllers.DB.First(&term, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := controllers.DB.Delete(&term).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	r := Response{"success"}
	return c.RenderJSON(r)
}
