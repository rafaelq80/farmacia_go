package model

type UsuarioLogin struct {
	ID      uint   	`json:"id"`
	Nome    string 	`json:"nome"`
	Usuario string 	`json:"usuario"`
	Senha   string 	`json:"senha"`
	Foto    string 	`json:"foto"`
	Role    string  `json:"role"` // Role
	Token   string 	`json:"token"`
}
