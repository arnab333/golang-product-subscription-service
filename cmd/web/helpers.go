package main

func (app *Config) sendEmail(msg Message) {
	app.WG.Add(1)
	app.Mailer.MailerChan <- msg
}
