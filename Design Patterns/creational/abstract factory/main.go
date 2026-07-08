package main

import "fmt"

func main() {
	fmt.Println("email noification")
	emailService := NewNotificationService(&EmailFactory{})
	emailService.notify("alice@example.com", "your order has been shipped")
	fmt.Println("")

	smsService := NewNotificationService(&SmsFactory{})
	smsService.notify("9876543210", "you have a missed call")
}

type Message interface {
	setContent(to, body string)
	format() string
}

type Sender interface {
	send(message Message)
}

type EmailMessage struct {
	to   string
	body string
}

func (m *EmailMessage) setContent(to, body string) {
	m.to = to
	m.body = body
}

func (m *EmailMessage) format() string {
	return "email to <" + m.to + ">:" + m.body
}

type EmailSender struct{}

func (s *EmailSender) send(message Message) {
	fmt.Println("sending vis smtp: " + message.format())
}

type SmsMessage struct {
	to   string
	body string
}

func (m *SmsMessage) setContent(to, body string) {
	m.to = to
	if len(body) > 150 {
		m.body = body[:150]
	} else {
		m.body = body
	}
}

func (m *SmsMessage) format() string {
	return "SMS to" + m.to + ": " + m.body
}

type SmsSender struct{}

func (s *SmsSender) send(message Message) {
	fmt.Println("sending via carrier api: " + message.format())
}

type NotificationFactory interface {
	createMessage() Message
	createSender() Sender
}

type EmailFactory struct{}

func (f *EmailFactory) createMessage() Message { return &EmailMessage{} }

func (f *EmailFactory) createSender() Sender { return &EmailSender{} }

type SmsFactory struct{}

func (f *SmsFactory) createMessage() Message { return &SmsMessage{} }
func (f *SmsFactory) createSender() Sender   { return &SmsSender{} }

type NotificationService struct {
	factory NotificationFactory
}

func NewNotificationService(factory NotificationFactory) *NotificationService {
	return &NotificationService{factory: factory}
}

func (s *NotificationService) notify(to, body string) {
	message := s.factory.createMessage()
	message.setContent(to, body)
	sender := s.factory.createSender()
	sender.send(message)
}
