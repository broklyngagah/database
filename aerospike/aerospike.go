package aerospike

import (
	_ "fmt"
	as "github.com/aerospike/aerospike-client-go"
)

type Info struct {
	Hostname string
	Port     int
}

func (i *Info) Connect() (*as.Client, error) {
	//fmt.Println(i)
	client, err := as.NewClient(i.Hostname, i.Port)
	if err != nil {
		panic(err)
	}

	return client, err
}
