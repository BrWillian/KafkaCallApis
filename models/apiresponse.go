package model

// Rede Capacete
type CapaceteApiResponse struct {
	SemCapacete bool `json:"semCapacete"`
}

// Rede Classificador
type ClassificacaoReponse struct {
	Detections []Detections `json:"Detections"`
}
type Boxes struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
type Detections struct {
	Id         int    `json:"id"`
	Confidence string `json:"confidence"`
	Boxes      Boxes  `json:"boxes"`
}

// Rede OCR
type OcrApiResponse struct {
	Placa string `json:"placa"`
}
