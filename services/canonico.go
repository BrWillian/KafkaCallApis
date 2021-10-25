package service

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
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
func sortClassificacao(result models.ClassificacaoReponse) int {
	m := make(map[float64]int)

	if len(result.Detections) == 0 {
		return -1
	}

	for _, v := range result.Detections {
		s, _ := strconv.ParseFloat(v.Confidence, 2)
		m[s] = v.Id
	}
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	return m[keys[0]]
}

func GetResult(msg models.KafkaMessage) models.Canonico {
	var canonico models.Canonico

	ocrResult := ConsumeOcrApi(ReadImage(msg.CaminhoImagem))

	classificacaoResult := ConsumeClassificadorApi(ReadImage(msg.CaminhoImagem))

	if classificacaoResult != nil {
		var classificacao models.ClassificacaoReponse

		json.Unmarshal(classificacaoResult, &classificacao)

		sorted_class := sortClassificacao(classificacao)
		canonico.ClassificacaoIa = &sorted_class

		if sorted_class == 1 {
			capaceteResult := ConsumeCapaceteApi(ReadImage(msg.CaminhoImagem))

			var capacete models.CapaceteApiResponse
			canonico.SemCapacete = nil
			if capaceteResult != nil {
				json.Unmarshal(capaceteResult, &capacete)
			}
			canonico.SemCapacete = new(int)
			*canonico.SemCapacete = bool2int(capacete.SemCapacete)

		} else {
			canonico.SemCapacete = new(int)
			*canonico.SemCapacete = -1
		}
	}

	var ocr models.OcrApiResponse

	json.Unmarshal(ocrResult, &ocr)

	canonico.Placa = &msg.Placa
	if msg.Placa == "0" {
		canonico.Placa = nil
	}
	canonico.PassouOcr = 0
	if ocrResult != nil {
		canonico.PassouOcr = 1
	}
	canonico.PlacaOcr = ocr.Placa
	canonico.CaminhoImagem = msg.CaminhoImagem
	canonico.CodigoEquipamento = msg.CodigoEquipamento
	canonico.Faixa = msg.Faixa
	canonico.Classificacao = msg.Classificacao
	t, _ := time.Parse(time.RFC3339, msg.DataHora)
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
		faixa, placa, placa_ocr, classificacao_id, classificacaoia_id, sem_capacete, caminho_imagem, ocr)
		VALUES (:classificacao_passagem_id, :data_hora, :codigo_equipamento, :faixa, :placa, :placa_ocr, :classificacao_id, :classificacaoia_id, :sem_capacete, :caminho_imagem, :ocr)`

	_, err = db.Exec(sqlStatement, data.Id, data.DataHora, data.CodigoEquipamento, data.Faixa,
		data.Placa, data.PlacaOcr, data.Classificacao, data.ClassificacaoIa, data.SemCapacete,
		data.CaminhoImagem, data.PassouOcr)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("deu Bom!")

}
