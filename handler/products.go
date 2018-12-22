package handler

import (
	"../db"
	"../model"
	"github.com/Jeffail/gabs"
	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var bot *tgbotapi.BotAPI

var chatID = int64(-237118654)

func GetBotInstance() *tgbotapi.BotAPI {

	if bot == nil {

		bot,_ = tgbotapi.NewBotAPI("796664752:AAFfcoC-QWTqUq_jm-3b6yH7A3MoZiPIOSg")
	}

	return bot
}

var NewProductHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var newProduct model.Product

	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)

	newProduct.Title = ClearString(json.Path("title").String())
	newProduct.PhotoURL = ClearString(json.Path("photoUrl").String())
	newProduct.Contacts = ClearString(json.Path("contacts").String())
	newProduct.Description = ""

	product, err := db.InsertNewProduct(&newProduct)
	if err != nil {
		log.Println("error adding product")
	}

	chatBot := GetBotInstance()

	sendMessage(*chatBot, chatID, randomNewProductMessage(newProduct.Title))

	responseJSON(product, response, request)
})

var GetNotApprovedProducts = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	products, err := db.GetNotApprovedProducts()
	if err != nil {
		log.Println("error getting not approved products", err)
	}

	responseJSON(products, response, request)
})

var GetApprovedProducts = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	products, err := db.GetApprovedProducts()
	if err != nil {
		log.Println("error getting approved products", err)
	}

	responseJSON(products, response, request)
})

var ApproveProduct = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)

	id, err := strconv.Atoi(json.Path("id").String())
	if err != nil {
		log.Println("can't parse id")
		return
	}


	product, err := db.ApproveProduct(id)
	if err != nil {
		log.Println("error update product")
	}

	responseJSON(product, response, request)
})

var DeleteProduct = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)

	id, err := strconv.Atoi(json.Path("id").String())
	if err != nil {
		log.Println("can't parse id")
		return
	}


	product, err := db.DeleteProduct(id)
	if err != nil {
		log.Println("error delete product")
	}

	responseJSON(product, response, request)
})

var GetProductByID = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(strings.ToLower(mux.Vars(request)["id"]))
	if err != nil {
		log.Println("can't parse id")
		return
	}

	product, err := db.GetProductByID(int64(id))
	if err != nil {
		log.Println("error get product")
	}

	responseJSON(product, response, request)
})

func sendMessage(bot tgbotapi.BotAPI, chatID int64, messageText string) {
	msg := tgbotapi.NewMessage(chatID, messageText)

	bot.Send(msg)
}

func randomNewProductMessage(title string) string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"новый итем в чулане - " + title, "очередную хуйню прислали - " + title, "еще штука за 2р - " + title,
		"очередное дерьмицо - " + title, "нужен модератор, в чулане - " + title, "кто посмотрит? в чулане - " + title}
	str2 := str[rand.Intn(len(str))]
	return str2
}
