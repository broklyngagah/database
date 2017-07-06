package mongo

import (
	"gopkg.in/mgo.v2"
)

type Info struct {
	Hostname string
	Database string
}

func (i *Info) Connect() (*mgo.Session, error) {
	session, err := mgo.Dial(i.Hostname)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
