package vars

var RabbitURL = "amqp://guest:guest@localhost:5672/"

var Messages = []string{
	"Hi Ivan",
	"Hi Kirill",
	"Hi Sasha",
}

type MessageInChan struct {
	Message []byte
	Err     error
}
