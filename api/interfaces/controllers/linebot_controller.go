package controllers

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	// "regexp"

	// "github.com/watarun54/spbill-api/server/domain"
	"github.com/watarun54/spbill-api/server/interfaces/database"
	"github.com/watarun54/spbill-api/server/usecase"
)

type LinebotController struct {
	UserInteractor  usecase.UserInteractor
}

func NewLinebotController(sqlHandler database.SqlHandler) *LinebotController {
	return &LinebotController{
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *LinebotController) GetTest(c Context) (err error) {
	c.JSON(200, NewResponse("success"))
	return
}

func (controller *LinebotController) Post(c Context) (err error) {
	bot, err := linebot.New(
		os.Getenv("LINEBOT_CHANNEL_SECRET"),
		os.Getenv("LINEBOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(500, NewError(err))
		} else {
			c.JSON(500, NewError(err))
		}
		return
	}

	for _, event := range events {
		lineID := event.Source.UserID
		GroupID := event.Source.GroupID
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println("lineID", lineID)
				fmt.Println("GroupID", GroupID)
				fmt.Println(message.Text)
				resMessage := linebot.NewTextMessage(message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
					log.Printf("send error: %v", err)
				}
			// 	user, _ := controller.UserInteractor.UserByLineID(lineID)
			// 	if user.ID == 0 {
			// 		message.Text = fmt.Sprintf("以下の Line ID を登録してください\n\n%v", lineID)
			// 	} else {
			// 		// URLを抽出
			// 		re, _ := regexp.Compile(`http(s)://[\w\d/%#$&?()~_.=+-]+`)
			// 		url := string(re.Find([]byte(message.Text)))
			// 		com := domain.Paper{
			// 			UserID: user.ID,
			// 			URL:    url,
			// 		}
			// 		c.Bind(&com)
			// 		title, err := controller.ScrapeHandler.GetTitleFromURL(com.URL)
			// 		if err != nil {
			// 			com.Text = "Faild to scrape title"
			// 		} else {
			// 			com.Text = title
			// 		}
			// 		_, err = controller.PaperInteractor.Add(com)
			// 		if err != nil {
			// 			message.Text = "記事の登録に失敗しました"
			// 		}
			// 		message.Text = "記事の登録に成功しました"
			// 	}
			// 	resMessage := linebot.NewTextMessage(message.Text)
			// 	if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
			// 		log.Printf("send error: %v", err)
			// 	}
			}
		}
	}

	c.JSON(200, NewResponse("success"))
	return
}
