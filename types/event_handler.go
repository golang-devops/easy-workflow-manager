package types

type EventHandler interface {
	Info(msg string)
	Error(msg string)
}
