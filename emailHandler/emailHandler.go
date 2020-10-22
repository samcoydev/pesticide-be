package emailHandler

import (
	gomail "gopkg.in/gomail.v2"
	log "pesticide/logHandler"
	"pesticide/models/user"
)

var fromName = "[emailHandler.go]"
var smtpServer = "samcodesthings.com"

func SendEmail(_user user.User) {
	m := gomail.NewMessage()
	m.SetHeader("From", "testautomatedemail@samcodesthings.com")
	m.SetHeader("To", _user.Email)
	m.SetAddressHeader("Cc", _user.Email, _user.FirstName)
	m.SetHeader("Subject", "Successfully registered!")
	m.SetBody("text/html", "Hello "+_user.FirstName+"! Thank you for registering for Pesticide!")

	d := gomail.NewDialer(smtpServer, 465, "samcoy", "0dXDR4mMSivx")

	if err := d.DialAndSend(m); err != nil {
		log.Err(fromName, "Something went wrong sending an email!")
	}

	log.Info(fromName, "Email sent to: "+_user.Email)
}
