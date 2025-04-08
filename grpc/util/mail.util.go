package grpcutil

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmailWhenCreateNewEmployee(recipientEmail string, firstname string, username string) error {
	// SMTP server details
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	authEmail := "noreply.freightflex@gmail.com"
	authPassword := "utck wxru gmxm tzyi"

	// Email message
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("HMS System <%s>", authEmail)) // Sets the display name
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "HMS System Notification - Your account has been created")
	//Content is like that but with username and password = 12345678@X
	m.SetBody("text/html", fmt.Sprintf("Hello %s, <br><br> Your account has been created successfully. <br><br> Your username is: %s <br> Your password is: 12345678@X <br><br> Please login to the system and change your password. <br><br> Best regards, <br> HMS System", firstname, username))

	// Attach a file (optional)
	// m.Attach("path/to/file.pdf")

	// Send the email
	d := gomail.NewDialer(smtpHost, smtpPort, authEmail, authPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	fmt.Println("Email sent successfully to:", recipientEmail)
	return nil
}
