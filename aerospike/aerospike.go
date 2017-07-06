package aerospike

import (
	as "github.com/aerospike/aerospike-client-go"
)

type Info struct {
	Hostname string
	Port     int
}

func (i *Info) Connect() (*as.Client, error) {
	client, err := as.NewClient(i.Hostname, i.Port)
	if err != nil {
		panic(err)
	}

	return client, err
}
