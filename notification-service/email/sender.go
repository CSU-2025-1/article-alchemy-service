package email

import (
	"gopkg.in/mail.v2"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func sendEmail(messageContent string, email string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	emailPassword := os.Getenv("EMAIL_PASSWORD")

	message := mail.NewMessage()

	message.SetHeader("From", "article.alchemy25@gmail.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Краткое содержание готово.")
	message.SetBody("text/plain", messageContent)

	// Пароля пока нет, но он обязательно будет...
	d := mail.NewDialer("smtp.mail.ru", 587,
		"article.alchemy25@gmail.com", emailPassword)

	if err := d.DialAndSend(message); err != nil {
		log.Println("Ошибка при отравке сообщения:", err)
	} else {
		log.Println("Сообщение отправлено")
	}
}
