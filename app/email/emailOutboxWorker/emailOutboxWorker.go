package emailOutboxWorker

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/email/sendEmailFactory"
	"log"
	"time"
)

func Start(ac *appContext.AppCtx, sleepTime time.Duration) chan bool {
	stopChan := make(chan bool)

	go func() {
		log.Println("Starting email outbox worker")

		for {
			select {
			case <-stopChan:
				log.Println("Email outbox worker stopped")
				return
			default:
				processEmails(ac, sleepTime)
				time.Sleep(sleepTime)
			}
		}
	}()

	return stopChan
}
func processEmails(ac *appContext.AppCtx, sleepTime time.Duration) {
	log.Println("Getting unsent emails")
	emails, err := ac.EmailOutbox.GetUnsentEmails()
	if err != nil {
		log.Printf("Error getting unsent emails: %v", err)
		return
	}

	log.Printf("Found %d unsent emails", len(emails))

	for _, email := range emails {
		log.Printf("Sending email: %v", email)
		// Mark email as sent after successful sending
		uow, err := ac.UowFactory.Begin()
		if err != nil {
			log.Printf("Error beginning unit of work: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		sendEmail := sendEmailFactory.FromReqCtx(nil)

		err = sendEmail.SendEmail(uow, email)
		if err != nil {
			log.Printf("Error sending email: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		defer uow.Rollback()

		err = ac.EmailOutbox.MarkAsSent(uow, email)
		if err != nil {
			log.Printf("Error marking email as sent: %v", err)
			uow.Rollback()
			continue
		}
		uow.Commit()

		log.Printf("Email sent: %v", email)
	}

	log.Println("No more unsent emails")
}
