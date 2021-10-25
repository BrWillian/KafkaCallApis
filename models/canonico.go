package model

import (
	"time"
)

type Canonico struct {
	Id                int       `db:"classificacao_passagem_id"`
	DataHora          time.Time `db:"data_hora"`
	CodigoEquipamento string    `db:"codigo_equipamento"`
	Faixa             int       `db:"faixa"`
	Placa             *string   `db:"placa"`
	PlacaOcr          string    `db:"placa_ocr"`
	Classificacao     int       `db:"classificacao_id"`
	ClassificacaoIa   *int      `db:"classificacaoia_id"`
	SemCapacete       *int      `db:"sem_capacete"`
	CaminhoImagem     string    `db:"caminho_imagem"`
	PassouOcr         int       `db:"ocr"`
}
