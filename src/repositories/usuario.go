package repositories

import (
	"api/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) CriarUsuario(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO usuario (nome, email, senha, sexo) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Email, usuario.Senha, usuario.Sexo)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) AtualizarUsuario(usuarioID uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare("UPDATE usuario SET nome = ?, email = ?, sexo = ? WHERE usuarioID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, usuario.Sexo, usuarioID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) DeletarUsuario(usuarioID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuario WHERE usuarioID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuarioID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarUsuarioPorID(usuarioID uint64) (models.Usuario, error) {
	linha, err := repositorio.db.Query("SELECT usuarioID, nome, email, sexo, dataCriacao FROM usuario WHERE usuarioID = ?", usuarioID)

	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(
			&usuario.UsuarioID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Sexo,
			&usuario.DataCriacao,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) BuscarTodosUsuarios() ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query("SELECT usuarioID, nome, email, sexo, dataCriacao FROM usuario")

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.UsuarioID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Sexo,
			&usuario.DataCriacao,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, err := repositorio.db.Query("SELECT usuarioID, senha FROM usuario WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.UsuarioID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, err := repositorio.db.Query("SELECT senha FROM usuario WHERE usuarioID = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, err := repositorio.db.Prepare("UPDATE usuario SET senha = ? WHERE usuarioID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senha, usuarioID); err != nil {
		return err
	}

	return nil
}
