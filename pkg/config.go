package pkg

import "github.com/ethereum/go-ethereum/common"

type config struct {
	addresses []common.Address
	topics    [][]common.Hash
}

func NewConfig() *config {
	return &config{
		addresses: make([]common.Address,0),
		topics:    make([][]common.Hash,0),
	}
}
func (c *config) AddAddressWithAddress(address common.Address)  {
	c.addresses = append(c.addresses,address)
}

func (c *config) AddAddressWithString(address string)  {
	c.addresses = append(c.addresses,common.HexToAddress(address))
}

func (c *config) AddTopicWithHash(topic common.Hash)  {
	c.topics = append(c.topics,[]common.Hash{topic})
}

func (c *config) AddTopicWithString(topic string)  {
	c.topics = append(c.topics,[]common.Hash{common.HexToHash(topic)})
}