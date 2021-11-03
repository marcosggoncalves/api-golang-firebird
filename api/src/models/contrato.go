package models

type Contrato struct {
	Tipo  string `json:"tipo"`
	Idcliente  string  `json:"idcliente"`
	Contrato  string  `json:"contrato"`
	Nome  string  `json:"nome"`
	Titular string    `json:"titular"`
	CpfCnpj  string    `json:"cpf_cnpj"`
	UltMesRefPg string `json:"ult_mes_ref_pg"`
	DiasAtraso  string    `json:"dias_atraso"`
	Status string `json:"status"`
}
