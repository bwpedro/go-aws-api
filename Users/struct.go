package main

type User struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Usuario   string `json:"usuario"`
	Senha     string `json:"senha"`
}
