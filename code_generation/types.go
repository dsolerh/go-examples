package main

type SomePubStruct struct {
	PubProp     int `json:"pub-prop"`
	privateProp int
}

type somePrivate struct{}
