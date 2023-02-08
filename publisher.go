package broker

type Publisher interface {
	Pub(msg *Message)
}
