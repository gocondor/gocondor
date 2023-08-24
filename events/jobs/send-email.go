package eventjobs

import (
	"fmt"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/models"
)

var SendEmail core.EventJob = func(event *core.Event, c *core.Context) {
	go func() {
		mailer := c.GetMailer()
		logger := c.GetLogger()

		user, ok := event.Payload["userStruct"].(models.User)
		if !ok {
			logger.Error("[SenEmail job] invalid userStruct")
			return
		}
		mailer.SetFrom(core.EmailAddress{Name: "GoCondor", Address: "mail@example.com"})
		mailer.SetTo([]core.EmailAddress{
			{
				Name: user.Name, Address: user.Email,
			},
		})
		mailer.SetSubject("Welcome To GoCondor")
		body := fmt.Sprintf("Hi %v, \nWelcome to GoCondor \nYour account have been created successfully. \nThanks.", user.Name)
		mailer.SetPlainTextBody(body)
		mailer.Send()
	}()
}
