package main

type Patient struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Sobrenome  string `json:"sobrenome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
	Nascimento string `json:"nascimento"`
	Sexo       string `json:"sexo"`
	Obs        string `json:"obs"`
}
