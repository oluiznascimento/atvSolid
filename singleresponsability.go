package main

import "fmt"

// Email representa a estrutura de um email.
type Email struct {
    Subject string
    Body    strins
}

// EmailSender é responsável por enviar emails.
type EmailSender struct{}

// SendEmail envia um email.
func (es *EmailSender) SendEmail(email Email) {
    // Implementação para enviar o email
    fmt.Printf("Email sent with subject: %s\n", email.Subject)
}

// NotificationService é responsável por enviar notificações.
type NotificationService struct{}

// SendNotification envia uma notificação.
func (ns *NotificationService) SendNotification(notification string) {
    // Implementação para enviar a notificação
    fmt.Printf("Notification sent: %s\n", notification)
}

// NotificationManager é responsável por coordenar o envio de emails e notificações.
type NotificationManager struct {
    emailSender       *EmailSender
    notificationSrvc *NotificationService
}

// NewNotificationManager cria uma nova instância de NotificationManager.
func NewNotificationManager() *NotificationManager {
    return &NotificationManager{
        emailSender:       &EmailSender{},
        notificationSrvc: &NotificationService{},
    }
}

// Notify envia um email e uma notificação.
func (nm *NotificationManager) Notify(email Email, notification string) {
    nm.emailSender.SendEmail(email)
    nm.notificationSrvc.SendNotification(notification)
}

func main() {
    nm := NewNotificationManager()
    email := Email{Subject: "Hello", Body: "This is a test email."}
    nm.Notify(email, "New notification received.")
}
