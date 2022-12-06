package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

type getUsuarioCallback func(*sql.Rows) (modelos.Usuario, error)

func (repositorio Usuarios) execute(query string, args ...any) (sql.Result, error) {
	statement, err := repositorio.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repositorio Usuarios) queryUsuarios(callback getUsuarioCallback, query string, args ...any) ([]modelos.Usuario, error) {
	rows, err := repositorio.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []modelos.Usuario
	for rows.Next() {
		usuario, err := callback(rows)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick,
		nomeOuNick,
	)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if err := linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorId(id uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id=?",
		id,
	)

	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if err := linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(usuarioId uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome=?, nick=?, email=? where id=?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioId); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) Deletar(usuarioId uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id=?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuarioId); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.Id, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioId, seguidorId uint64) error {
	if _, err := repositorio.execute(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
		usuarioId,
		seguidorId,
	); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioId, seguidorId uint64) error {
	if _, err := repositorio.execute(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
		usuarioId,
		seguidorId,
	); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	usuarios, err := repositorio.queryUsuarios(
		func(r *sql.Rows) (modelos.Usuario, error) {
			var usuario modelos.Usuario
			if err := r.Scan(
				&usuario.Id,
				&usuario.Nome,
				&usuario.Nick,
				&usuario.Email,
				&usuario.CriadoEm,
			); err != nil {
				return modelos.Usuario{}, err
			}

			return usuario, nil
		},

		`select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u join seguidores s on u.id=s.seguidor_id
		where s.usuario_id = ?`,
		usuarioId,
	)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	usuarios, err := repositorio.queryUsuarios(
		func(r *sql.Rows) (modelos.Usuario, error) {
			var usuario modelos.Usuario
			if err := r.Scan(
				&usuario.Id,
				&usuario.Nome,
				&usuario.Nick,
				&usuario.Email,
				&usuario.CriadoEm,
			); err != nil {
				return modelos.Usuario{}, err
			}

			return usuario, nil
		},

		`select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u join seguidores s on u.id=s.usuario_id
		where s.seguidor_id  = ?`,
		usuarioId,
	)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	usuarios, err := repositorio.queryUsuarios(
		func(r *sql.Rows) (modelos.Usuario, error) {
			var usuario modelos.Usuario
			if err := r.Scan(&usuario.Senha); err != nil {
				return modelos.Usuario{}, err
			}
			return usuario, nil
		},

		"select senha from usuarios where id = ?",
		usuarioId,
	)
	if err != nil {
		return "", err
	}

	return usuarios[0].Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioId uint64, senha string) error {
	if _, err := repositorio.execute(
		"update usuarios set senha = ? where id = ?",
		senha,
		usuarioId,
	); err != nil {
		return err
	}

	return nil
}
