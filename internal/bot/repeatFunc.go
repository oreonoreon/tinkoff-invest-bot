package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

// Вызов переданной функции раз в сутки в указанное время.
func callAt(hour, min, sec int, f func(), D time.Duration) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	// Вычисляем время первого запуска.
	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	durationUntilStart := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(durationUntilStart)
		for {
			f()
			// Следующий запуск через сутки или через D.
			time.Sleep(D) // периуд повтора time.Hour * 24 или любой другой периуд повтора

			//m = <-time.After(D)

		}
	}()

	return nil
}

type Xrr struct {
	BOT     *tgbotapi.BotAPI
	C       *Cache
	Updates tgbotapi.UpdatesChannel
}

func newXrr(b *tgbotapi.BotAPI, c *Cache, updates tgbotapi.UpdatesChannel) *Xrr {
	return &Xrr{BOT: b, C: c, Updates: updates}
}
func callAtStartOfProgram(f func(), D time.Duration) {
	go func() {
		for {
			f()
			// Следующий запуск через через D.
			time.Sleep(D) // периуд повтора time.Hour * 24 или любой другой периуд повтора
		}
	}()
}
