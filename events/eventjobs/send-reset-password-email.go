package eventjobs

import (
	"fmt"
	"os"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/models"
)

var SendResetPasswordEmail core.EventJob = func(event *core.Event, c *core.Context) {
	go func() {
		mailer := c.GetMailer()
		logger := c.GetLogger()

		user, ok := event.Payload["user"].(models.User)
		if !ok {
			logger.Error("[SendResetPasswordEmail job] invalid user")
			return
		}
		mailer.SetFrom(core.EmailAddress{Name: "GoCondor", Address: "mail@example.com"})
		mailer.SetTo([]core.EmailAddress{
			{
				Name: user.Name, Address: user.Email,
			},
		})

		mailer.SetSubject("Reset Password Link")
		hostname, err := os.Hostname()
		if err != nil {
			c.GetLogger().Error(err.Error())
		}
		resetPasswordLink := fmt.Sprintf("%v/reset-password/code/%v", hostname, c.CastToString(event.Payload["code"]))
		body := fmt.Sprintf("Hi %v, <br>Click the link below to reset your password <br><a href=\"%v\">Reset Password</a>. <br>Thanks.", user.Name, resetPasswordLink)
		mailer.SetHTMLBody(body)
		mailer.Send()
	}()
}
