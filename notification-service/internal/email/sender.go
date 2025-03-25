package email

import (
	"gopkg.in/mail.v2"
	"os"
	"strings"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/domain/notification"
)

func Send(email string, url string, data []notification.Chapter) error {

	content, err := formatContent(url, data)
    if err != nil {
        return err
    }

	message := mail.NewMessage()

	message.SetHeader("From", os.Getenv("SMTP_USER"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Краткое содержание готово.")
	message.SetBody("text/plain", content)

	d := mail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(message)
}


func formatContent(url string, data []notification.Chapter) (string, error) {
    var builder strings.Builder
    builder.WriteString("<h1>Ваше краткое содержание готово</h1>")
    builder.WriteString(fmt.Sprintf("<p><a href='%s'>Ссылка на статью</a></p>", url))
    builder.WriteString("<ul>")
    
    for i, chapter := range data {
        builder.WriteString(fmt.Sprintf("<li><strong>%d. %s</strong><ul>", i+1, chapter.Title))
        for j, point := range chapter.Points {
            builder.WriteString(fmt.Sprintf("<li>%d.%d. %s</li>", i+1, j+1, point))
        }
        builder.WriteString("</ul></li>")
    }
    
    builder.WriteString("</ul>")
    return builder.String(), nil
}