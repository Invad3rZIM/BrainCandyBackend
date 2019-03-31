package main

type Message struct {
	Num  int
	Body string

	From     int
	To       int
	SecretID int

	TimeStamp string
}
