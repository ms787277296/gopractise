package server

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"github.com/golang/glog"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func JsonTransform() (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		person := &Person{
			Name:   fmt.Sprintf("Name%d", i),
			Age:    i,
			Gender: "Man",
		}
		persons = append(persons, person)
	}

	jsonPerson, err := json.Marshal(persons)
	if err != nil {
		glog.Error(err)
		return
	}

	err = ioutil.WriteFile("~/Desktop/tmp/person.txt", jsonPerson, 0755)
	if err != nil {
		glog.Error(err)
		return
	}
	//glog.Info(string(jsonPerson))

	return
}
