package config

type Configuracao struct {
	Porta        string
	TipoAmbiente TipoAmbiente
}

type TipoAmbiente string

func CriarConfiguracao() *Configuracao {
	return &Configuracao{
		Porta:        "8080",
		TipoAmbiente: TipoAmbienteDesenvolvimento,
	}
}

const (
	TipoAmbienteDesenvolvimento TipoAmbiente = "DEV"
	TipoAmbienteProducao        TipoAmbiente = "PRODUCAO"
)

func StringParaTipoAmbiente(valor string) (TipoAmbiente, bool) {
	switch valor {
	case string(TipoAmbienteDesenvolvimento):
		return TipoAmbienteDesenvolvimento, true
	case string(TipoAmbienteProducao):
		return TipoAmbienteProducao, true
	default:
		return "", false
	}
}
