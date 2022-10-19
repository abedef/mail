package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// Gets and returns environment variable or returns fallback if it does not exist
func defaultGetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// Gets environment variable or fatally fails
func fatalGetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("missing value for environment variable ", key)
	} else if value == "" {
		log.Print("empty string value provided for environment variable ", key)
	}
	return value
}

// SendMailFrom uses SMTP to send an email from the provided address to the
// specified recipient with the provided subject and body
//
// See https://zetcode.com/golang/email-smtp/ for source code
func SendMailFrom(from string, to string, subject string, body string) error {
	user := defaultGetEnv("SMTP_USER", "apikey")
	password := fatalGetEnv("SMTP_PASS")

	addr := fmt.Sprintf("%v:%v", defaultGetEnv("SMTP_HOST", "smtp.sendgrid.net"), defaultGetEnv("SMTP_PORT", "587"))
	host := defaultGetEnv("SMTP_HOST", "smtp.sendgrid.net")

	msg := []byte(fmt.Sprintf("From: %v\r\n"+
		"To: %v\r\n"+
		"Subject: %v\r\n\r\n"+
		"%v\r\n", from, to, subject, body))

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, []string{to}, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")

	return nil
}
