package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	FirstName      string
	LastName       string
	UserName       string
	TokenSanBox    string
	TokenRealTrade string
	ChUser         chan any
	WorkerOn       bool
	MessageId      int
}

type Cache struct {
	Users map[int64]User
}

func New() *Cache {
	users := make(map[int64]User)

	cache := Cache{
		Users: users,
	}
	return &cache
}
func (U *Cache) CheckNewUser(update tgbotapi.Update) {

	if _, ok := U.Users[update.SentFrom().ID]; !ok {
		user := User{
			FirstName: update.SentFrom().FirstName,
			LastName:  update.SentFrom().LastName,
			UserName:  update.SentFrom().UserName,
			ChUser:    make(chan any),
		}
		U.Users[update.SentFrom().ID] = user
		fmt.Println(U, ok)
	} else {
		fmt.Println("IT'S ALIVE! ", update.SentFrom().ID, U.Users[update.SentFrom().ID], ok)
	}

}

func (U *Cache) FullFillCache(ID int64) {

	if _, ok := U.Users[ID]; !ok {
		user := User{
			FirstName:      "",
			LastName:       "",
			UserName:       "",
			TokenSanBox:    "",
			TokenRealTrade: "",
			ChUser:         make(chan any),
			//WorkerOn: false,
			//MessageId: 0,//todo а если чёт не работает можно взглянуть сюда
		}
		U.Users[ID] = user
		fmt.Println(U, ok)

	} else {
		fmt.Println("IT'S ALIVE! ", U, ok)
	}
}
func (U *Cache) SetToken(ID int64, token string) {
	if v, ok := U.Users[ID]; ok {
		v.TokenSanBox = token
		U.Users[ID] = v
	}
}
func (U Cache) SetWorkerOn(ID int64, bool bool) {
	if v, ok := U.Users[ID]; ok {
		v.WorkerOn = bool
		U.Users[ID] = v
	}
}
func (U Cache) SetMessageId(ID int64, messageID int) {
	if v, ok := U.Users[ID]; ok {
		v.MessageId = messageID
		U.Users[ID] = v
	}
}
