package model

type KafkaMessage struct {
	DataHora          string `json:"dataHora"`
	CodigoEquipamento string `json:"codigoEquipamento"`
	Placa             string `json:"placa"`
	Classificacao     int    `json:"classificacao"`
	Rodovia           string `json:"rodovia"`
	Km                string `json:"km"`
	Metro             string `json:"metro"`
	Municipio         string `json:"municipio"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Faixa             int    `json:"faixa"`
	CodigoLocal       string `json:"codigoLocal"`
	Sentido           string `json:"sentido"`
	CaminhoImagem     string `json:"caminhoImagem"`
}

func NewKafkaMessage() *KafkaMessage {
	return &KafkaMessage{}
}
