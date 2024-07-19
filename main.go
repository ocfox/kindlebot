package main

import (
	"context"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func IsKindleFormat(update *models.Update) bool {
	if update.Message.Document != nil {
		return strings.HasSuffix(update.Message.Document.FileName, ".pdf") ||
			strings.HasSuffix(update.Message.Document.FileName, ".epub")
	}
	return false
}

func GetFile(ctx context.Context, b *bot.Bot, fileID string) *models.File {
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: fileID,
	})
	if err != nil {
		panic(err)
	}
	return file
}

func GetFileLink(ctx context.Context, b *bot.Bot, fileID string) string {
	file := GetFile(ctx, b, fileID)
	return b.FileDownloadLink(file)
}

func HandleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !IsKindleFormat(update) {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Not kindle format",
		})
		if err != nil {
			panic(err)
		}
		return
	}

	fileLink := GetFileLink(ctx, b, update.Message.Document.FileID)

	envs := GetEnvs()
	mail := EmailAccount{
		MailAddress: envs.SendMail,
		Password:    envs.Password,
		Server:      envs.Server,
	}

	println("downloading file")
	attachment := DownloadFile(File{
		Name: update.Message.Document.FileName,
		Link: fileLink,
	})

	println("sending mail")

	SendMail(mail, "test@ocfox.me", attachment)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fileLink,
	})
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(HandleMessage),
	}

	if IsEnvsEmpty(GetEnvs()) {
		panic("missing envs")
	}

	b, err := bot.New("", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
