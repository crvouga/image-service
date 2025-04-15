package emailOutboxWorker

import (
	"imageresizerservice/app/ctx"
	"log"
	"time"
)

func Start(appCtx *ctx.AppCtx, sleepTime time.Duration) chan bool {
	stopChan := make(chan bool)

	go func() {
		log.Println("Starting email outbox worker")

		for {
			select {
			case <-stopChan:
				log.Println("Email outbox worker stopped")
				return
			default:
				processEmails(appCtx, sleepTime)
				time.Sleep(sleepTime)
			}
		}
	}()

	return stopChan
}
func processEmails(appCtx *ctx.AppCtx, sleepTime time.Duration) {
	log.Println("Getting unsent emails")
	emails, err := appCtx.EmailOutbox.GetUnsentEmails()
	if err != nil {
		log.Printf("Error getting unsent emails: %v", err)
		return
	}

	log.Printf("Found %d unsent emails", len(emails))

	for _, email := range emails {
		log.Printf("Sending email: %v", email)
		// Mark email as sent after successful sending
		uow, err := appCtx.UowFactory.Begin()
		if err != nil {
			log.Printf("Error beginning unit of work: %v", err)
			time.Sleep(sleepTime)
			continue
		}
		err = appCtx.SendEmail.SendEmail(uow, email)
		if err != nil {
			log.Printf("Error sending email: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		defer uow.Rollback()

		err = appCtx.EmailOutbox.MarkAsSent(uow, email)
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
