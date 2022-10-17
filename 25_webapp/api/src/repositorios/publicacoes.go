package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

type getPublicacaoCallback func(*sql.Rows) (modelos.Publicacao, error)

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) execute(query string, args ...any) (sql.Result, error) {
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

func (repositorio Publicacoes) queryPublicacoes(callback getPublicacaoCallback, query string, args ...any) ([]modelos.Publicacao, error) {
	rows, err := repositorio.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publucacoes []modelos.Publicacao
	for rows.Next() {
		publicacao, err := callback(rows)
		if err != nil {
			return nil, err
		}

		publucacoes = append(publucacoes, publicacao)
	}

	return publucacoes, nil
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	resultado, err := repositorio.execute(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?);",
		publicacao.Titulo,
		publicacao.Conteudo,
		publicacao.AutorId,
	)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIdInserido), err
}

func (repositorio Publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {
	publicacoes, err := repositorio.queryPublicacoes(
		func(r *sql.Rows) (modelos.Publicacao, error) {
			var publicacao modelos.Publicacao
			if err := r.Scan(
				&publicacao.Titulo,
				&publicacao.Conteudo,
				&publicacao.AutorId,
				&publicacao.Curtidas,
				&publicacao.CriadaEm,
				&publicacao.AutorNick,
			); err != nil {
				return modelos.Publicacao{}, err
			}

			return publicacao, nil
		},

		`
		select distinct p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criadaEm, u.nick
		from publicacoes p
		join usuarios u on u.id=p.autor_id
		join seguidores s on s.usuario_id=u.id
		where u.id = ? or s.seguidor_id = ?
		order by p.id desc`,
		usuarioId,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	publicacoes, err := repositorio.queryPublicacoes(
		func(r *sql.Rows) (modelos.Publicacao, error) {
			var publicacao modelos.Publicacao
			if err := r.Scan(
				&publicacao.Titulo,
				&publicacao.Conteudo,
				&publicacao.AutorId,
				&publicacao.Curtidas,
				&publicacao.CriadaEm,
				&publicacao.AutorNick,
			); err != nil {
				return modelos.Publicacao{}, err
			}
			return publicacao, nil
		},

		`
		select p.titulo, p.conteudo, p.autor_id, p.curtidas, p.criadaEm, u.nick
		from publicacoes p
		join usuarios u on u.id=p.autor_id
		where p.id = ?`,
		publicacaoId,
	)

	if err != nil {
		return modelos.Publicacao{}, err
	}

	return publicacoes[0], nil
}
