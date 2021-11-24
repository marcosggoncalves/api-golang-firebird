package models

type Contrato struct {
	Tipo       string  `json:"tipo, omitempty"`
	Idcliente  string  `json:"idcliente, omitempty"`
	Contrato   string  `json:"contrato, omitempty"`
	Nome       string  `json:"nome, omitempty"`
	Titular    string  `json:"titular, omitempty"`
	CpfCnpj    *string `json:"cpf_cnpj, omitempty"`
	UltPg      string  `json:"ult_pg, omitempty"`
	DiasAtraso string  `json:"dias_atraso, omitempty"`
	Status     string  `json:"status, omitempty"`
}

// response format
type Response struct {
	Status   bool     `json:"status,omitempty"`
	Message  string   `json:"message,omitempty"`
	Contrato Contrato `json:"contrato"`
}
