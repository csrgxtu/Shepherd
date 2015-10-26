package inits

import (
  "gopkg.in/mgo.v2"
)

var Session *mgo.Session

func init() {
  var err error
	Session, err = mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	Session.SetMode(mgo.Monotonic, true)
}
