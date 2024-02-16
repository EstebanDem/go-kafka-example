package service

// NewUserMessage represents the message from the message broker when a new user is registered
type NewUserMessage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (n NewUserMessage) EventName() string {
	return "new_user_registered"
}

// EmailMessage represents the message from the message broker when a email is validated from email-service or another
type EmailMessage struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}

func (e EmailMessage) EventName() string {
	return "email_validated"
}
