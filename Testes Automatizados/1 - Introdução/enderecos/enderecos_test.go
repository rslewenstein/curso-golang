package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada da Posse", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
		//{"Avenida das Bromélias", "Estrada"}, // teste de erro
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

// implementando Testes unitários
// para realizar os testes, é obrigatório:
// O nome do arquivo deve ter o _test
// O nome da função deve começar com "Test"

// para executar, digite: go test (+ parametros)

// Parametros:
// (-v) = Mais detalhes
// (--cover) = exibe qual a porcentagem de cobertura dos testes
// (--coverprofile nome_arquivo.txt) = gera um arquivo com as informações
// (go tool cover --func=nome_arquivo.txt) = lê as informações do arquivo gerado
// (go tool cover --html=nome_arquivo.txt) = exibe as informações das linhas que não estão cobertas
