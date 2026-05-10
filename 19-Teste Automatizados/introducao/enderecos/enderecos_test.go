// teste de unidade
package enderecos

import (
	"testing"
)

func TestTipoEndereco(t *testing.T) {
	enderecosParaTeste := "Rua ABC"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecosParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Error("O tipo rcebido e diferente do esperado")
	}
}
