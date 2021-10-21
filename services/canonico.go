package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"

	models "github.com/brwillian/kafka-consumer-api/models"
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func GetResult(msg models.KafkaMessage) models.Canonico {
	var canonico models.Canonico

	ocrResult := ConsumeOcrApi(ReadImage(msg.CaminhoImagem))

	classificacaoResult := ConsumeClassificadorApi(ReadImage(msg.CaminhoImagem))
	capaceteResult := ConsumeCapaceteApi(ReadImage(msg.CaminhoImagem))

	var ocr models.OcrApiResponse
	var classificacao models.ClassificacaoReponse
	var capacete models.CapaceteApiResponse

	json.Unmarshal(ocrResult, &ocr)

	json.Unmarshal(classificacaoResult, &classificacao)

	json.Unmarshal(capaceteResult, &capacete)

	if capaceteResult == nil {
		canonico.SemCapacete = nil
	} else {
		canonico.SemCapacete = new(int)
		*canonico.SemCapacete = bool2int(capacete.SemCapacete)
	}

	canonico.Placa = &msg.Placa
	if msg.Placa == "0" {
		canonico.Placa = nil
	}
	canonico.PlacaOcr = ocr.Placa
	canonico.CaminhoImagem = msg.CaminhoImagem
	canonico.CodigoEquipamento = msg.CodigoEquipamento
	canonico.Faixa = msg.Faixa
	canonico.Classificacao = msg.Classificacao
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", msg.DataHora)
	formatted := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	canonico.DataHora = formatted
	return canonico
}

func SaveDb(data models.Canonico) {

	db, err := sqlx.Open("godror", `user="ANTT_OCORRENCIA" password="anttocorrencia" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=DTF-LBDEXP-DEV.datatraffic.com.br)(PORT=1521))(CONNECT_DATA=(Service_name=xe)))"`)
	defer db.Close()

	rows, err := db.Query("select SQ_CLASSIFICACAO_PASSAGEM.nextval from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data.Id)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	sqlStatement := `INSERT INTO TBCLASSIFICACAO_PASSAGEM (classificacao_passagem_id, data_hora, codigo_equipamento,
		faixa, placa, placa_ocr, classificacao_id, classificacaoia_id, sem_capacete, caminho_imagem)
		VALUES (:classificacao_passagem_id, :data_hora, :codigo_equipamento, :faixa, :placa, :placa_ocr, :classificacao_id, :classificacaoia_id, :sem_capacete, :caminho_imagem)`

	_, err = db.Exec(sqlStatement, data.Id, data.DataHora, data.CodigoEquipamento, data.Faixa,
		data.Placa, data.PlacaOcr, data.Classificacao, data.ClassificacaoIa, data.SemCapacete,
		data.CaminhoImagem)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("deu Bom!")

}
