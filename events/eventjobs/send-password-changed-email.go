package eventjobs

import (
	"fmt"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/models"
)

var SendPasswordChangedEmail core.EventJob = func(event *core.Event, c *core.Context) {
	go func() {
		mailer := c.GetMailer()
		logger := c.GetLogger()

		user, ok := event.Payload["user"].(models.User)
		if !ok {
			logger.Error("[SendPasswordChangedEmail job] invalid user")
			return
		}
		mailer.SetFrom(core.EmailAddress{Name: "GoCondor", Address: "mail@example.com"})
		mailer.SetTo([]core.EmailAddress{
			{
				Name: user.Name, Address: user.Email,
			},
		})
		mailer.SetSubject("Password Changed")
		body := fmt.Sprintf("Hi %v, \nYour password have been changed. \nThanks.", user.Name)
		mailer.SetPlainTextBody(body)
		mailer.Send()
	}()
}
