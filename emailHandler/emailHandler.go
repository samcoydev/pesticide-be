package emailHandler

import (
	gomail "gopkg.in/gomail.v2"
	log "pesticide/logHandler"
	"pesticide/models/user"
)

var fromName = "[emailHandler.go]"

func SendEmail(_user user.User) {
	m := gomail.NewMessage()
	m.SetHeader("From", "testautomatedemail@samcodesthings.com")
	m.SetHeader("To", _user.Email)
	m.SetAddressHeader("Cc", _user.Email, _user.FirstName)
	m.SetHeader("Subject", "Testing my emails!")
	m.SetBody("text/html", "Hello "+_user.FirstName)

	d := gomail.NewDialer("samcodesthings.com", 465, "samcoy", "0dXDR4mMSivx")

	if err := d.DialAndSend(m); err != nil {
		log.Err(fromName, "Something went wrong sending an email!")
	}

	log.Info(fromName, "Email sent to: "+_user.Email)
}
