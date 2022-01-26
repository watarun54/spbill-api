package controllers

import (
	"github.com/watarun54/serverless-skill-manager/server/domain"
	"regexp"
)

type ScrapeController struct {
	ScrapeHandler IScrapeHandler
}

func NewScrapeController(scrapeHandler IScrapeHandler) *ScrapeController {
	return &ScrapeController{
		ScrapeHandler: scrapeHandler,
	}
}

func (controller *ScrapeController) GetPaperTitle(c Context) (err error) {
	form := domain.ScrapeForm{}
	c.Bind(&form)
	// URLを抽出
	re, _ := regexp.Compile(`http(s)://[\w\d/%#$&?()~_.=+-]+`)
	form.URL = string(re.Find([]byte(form.URL)))
	title, err := controller.ScrapeHandler.GetTitleFromURL(form.URL)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(title))
	return
}
