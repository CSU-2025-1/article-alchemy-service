package email

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/domain/notification"
	"gopkg.in/mail.v2"
)

func Send(email string, url string, data json.RawMessage, errorMsg string) error {
	var content string
	var topic string

	if errorMsg != "" {
		content, topic = formatError(url, errorMsg)
	} else {
		var chapters []notification.Chapter
		if len(data) > 0 {
            if err := json.Unmarshal(data, &chapters); err != nil {
				var bodyStr string
                if err := json.Unmarshal(data, &bodyStr); err == nil {
					if err := json.Unmarshal([]byte(bodyStr), &chapters); err != nil {
						return fmt.Errorf("failed to unmarshal body string: %w", err)
					}
				} else {
					return fmt.Errorf("failed to unmarshal body: %w", err)
				}
            }
        }
		content, topic = formatContent(url, chapters)
	}

	message := mail.NewMessage()

	message.SetHeader("From", os.Getenv("SMTP_USER"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", topic)
	message.SetBody("text/html", content)

	d := mail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(message)
}

func formatError(url string, errorMsg string) (string, string) {
	var builder strings.Builder
    builder.WriteString("<h2>Возникла ошибка при генерации краткого содержания :(</h2>")
    builder.WriteString(fmt.Sprintf("<p><a href='%s'>Ссылка на статью</a></p>", url))
	builder.WriteString(fmt.Sprintf("<p>Ошибка: %s</p>", errorMsg))
	return builder.String(), "Ошибка при генерации содержания"
}


func formatContent(url string, data []notification.Chapter) (string, string) {
    var builder strings.Builder
    builder.WriteString("<h2>Ваше краткое содержание готово</h2>")
    builder.WriteString(fmt.Sprintf("<p><a href='%s'>Ссылка на статью</a></p>", url))
    builder.WriteString("<ul>")
    
    for _, chapter := range data {
        builder.WriteString(fmt.Sprintf("<li><strong> %s</strong><ul>", chapter.Title))
        for _, point := range chapter.Points {
            builder.WriteString(fmt.Sprintf("<li> %s</li>", point))
        }
        builder.WriteString("</ul></li>")
    }
    
    builder.WriteString("</ul>")
    return builder.String(), "Краткое содержание готово"
}