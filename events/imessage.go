package event

//IMessage interface for generic messaage type
type IMessage interface {
	Ack()
	Nack()
}

//Ack the incoming message
func Ack(m IMessage) {
	m.Ack()
}

//Nack the incoming message
func Nack(m IMessage) {
	m.Ack()
}
