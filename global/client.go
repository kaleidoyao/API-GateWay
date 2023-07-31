package global

import (
	"github.com/cloudwego/kitex/client/genericclient"
	"sync"
)

type ClientPool struct {
	pool sync.Map
}

func NewClientPool() *ClientPool {
	return &ClientPool{}
}

func (clientPool *ClientPool) GetClient(serviceName string) (genericclient.Client, error) {
	client, ok := clientPool.pool.Load(serviceName)
	if ok {
		return client.(genericclient.Client), nil
	}

	newClient, err := GenerateClient(serviceName)
	if err != nil {
		return nil, err
	}
	clientPool.pool.Store(serviceName, newClient)
	return newClient, nil
}
