package controllers

type IScrapeHandler interface {
	GetTitleFromURL(url string) (string, error)
}
