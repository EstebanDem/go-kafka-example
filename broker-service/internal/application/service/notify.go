package service

type NotifyService interface {
	Send(message any)
	GetInterestedApps() []InterestedApp
}

type InterestedApp struct {
	Name    string
	Address string
}
