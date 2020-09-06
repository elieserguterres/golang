package produtos

type Produto struct {
	Id   int
	Nome string
}

func findAll() []Produto {

	conexao := db.conn()

	selectProdutos, err := conexao.Query("select * from servicos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id int
		var nome string

		err = selectProdutos.Scan(&id, &nome)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome

		produtos = append(produtos, p)
	}

	conexao.Close()

	return produtos
}
