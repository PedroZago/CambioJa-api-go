package models

type DadosAutenticacao struct {
	UsuarioID string `json:"usuarioID"`
	Token     string `json:"token"`
}
