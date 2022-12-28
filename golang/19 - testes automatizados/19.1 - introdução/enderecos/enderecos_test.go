package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida ABC"
	TipoDeEnderecoEsperado := "avenida"
	TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if TipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			TipoDeEnderecoEsperado,
			TipoDeEnderecoRecebido)
	}

}
