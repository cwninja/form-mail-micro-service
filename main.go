package main

import (
	m "github.com/keighl/mandrill"
	"log"
	"net/http"
	"os"
)

func main() {
	client := m.ClientWithKey(os.Getenv("MANDRILL_KEY"))
	subject := os.Getenv("SUBJECT")
	to := os.Getenv("TO")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		raw_message := r.PostFormValue("message")
		from := r.PostFormValue("from")
		log.Println("sending message")
		message := &m.Message{}
		message.AddRecipient(to, "", "to")
		message.FromEmail = from
		message.Subject = subject
		message.Text = raw_message

		if _, err := client.MessagesSend(message); err != nil {
			log.Println(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}
