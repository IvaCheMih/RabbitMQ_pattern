package exchange

import (
	"fmt"
	"github.com/IvaCheMih/RabbitMQ_pattern/vars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclareExchange(t *testing.T) {
	var testExchange = vars.ExchangeTest{
		Name:    "exchange_test",
		Kind:    "direct",
		AutoDel: true,
	}

	_, err := DeclareExchange(testExchange.Name, testExchange.Kind, testExchange.AutoDel, vars.RabbitURL)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Should not produce an error")
	}

	assert.Nil(t, err)
	assert.Equal(t, nil, err)

}
