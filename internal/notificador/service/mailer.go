package service

type Mailer interface {
	Build() (string, string)
	Receiver() string
}
