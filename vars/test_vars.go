package vars

type ExchangeTest struct {
	Name    string
	Kind    string
	AutoDel bool
}

type RabbitTest struct {
	URL string
}

type QueueTest struct {
	Name     string
	Key      string
	Exchange string
	AutoDel  bool
}
