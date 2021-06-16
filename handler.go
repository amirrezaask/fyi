package fyi

type Handler func(Notification) (Response, error)

type Response interface {
	Raw() (string, error)
}
