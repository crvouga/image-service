package emailOutboxWorker

import (
	"imageresizerservice/deps"
	"log"
	"time"
)

func Start(d *deps.Deps, sleepTime time.Duration) chan bool {
	stopChan := make(chan bool)

	go func() {
		log.Println("Starting email outbox worker")

		for {
			select {
			case <-stopChan:
				log.Println("Email outbox worker stopped")
				return
			default:
				processEmails(d, sleepTime)
				time.Sleep(sleepTime)
			}
		}
	}()

	return stopChan
}

func processEmails(d *deps.Deps, sleepTime time.Duration) {
	log.Println("Getting unsent emails")
	emails, err := d.EmailOutbox.GetUnsentEmails()
	if err != nil {
		log.Printf("Error getting unsent emails: %v", err)
		return
	}

	log.Printf("Found %d unsent emails", len(emails))

	for _, email := range emails {
		log.Printf("Sending email: %v", email)
		err := d.SendEmail.SendEmail(email)
		if err != nil {
			log.Printf("Error sending email: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		log.Printf("Email sent: %v", email)
	}

	log.Println("No more unsent emails")
}
