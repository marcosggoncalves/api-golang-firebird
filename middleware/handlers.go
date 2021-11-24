package middleware

import (
	"database/sql"
	"encoding/json"
	"firebird-golang/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/nakagami/firebirdsql"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("firebirdsql", os.Getenv("FIREBIRD_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

func GetContrato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	contrato, err := GetContratoQuery(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// format the response message
	res := models.Response{
		Status:   true,
		Message:  "Contrato localizado!",
		Contrato: contrato,
	}

	json.NewEncoder(w).Encode(res)
}

func GetContratoQuery(id int64) (models.Contrato, error) {
	// create the postgres db connection
	db := createConnection()

	defer db.Close()

	var contrato models.Contrato

	sqlStatement := fmt.Sprintf(`SELECT 'TITULAR' AS TIPO, C.IDCLIENTE,C.CONTRATO, left(C.NOME, 15), left(C.NOME, 15) AS TITULAR,C.CPF_CNPJ,
		(SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
		WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
		ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC) AS ULT_PG,
		CASE 
			WHEN C.SITUACAO = 4 THEN 0
			WHEN C.DATA_CONTRATO >= DATEADD(1 - EXTRACT(DAY FROM CURRENT_DATE) DAY TO CURRENT_DATE ) THEN 0
			ELSE
				((CURRENT_DATE) - (SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
				WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
				ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC))
		END AS DIAS_ATRASO,
		CASE
			WHEN C.SITUACAO = 4 THEN 'ATIVO'
			WHEN C.DATA_CONTRATO >= DATEADD(1 - EXTRACT(DAY FROM CURRENT_DATE) DAY TO CURRENT_DATE ) THEN 'ATIVO'
			WHEN ((CURRENT_DATE) - (SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
		WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
		ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC)) < 90 THEN 'ATIVO'
		ELSE 'INATIVO'
		END AS STATUS
		FROM CLIENTE C
		WHERE C.CONTRATO = '%d'
		AND ((C.SITUACAO =1) OR (C.SITUACAO =4))
		UNION
		SELECT  'DEPENDENTE' AS TIPO , C.IDCLIENTE,C.CONTRATO, left(D.NOME,15) , left(C.NOME, 15) AS TITULAR, D.CPF AS CPF_CNPJ,
		(SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
		WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
		ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC) AS ULT_PG,
		CASE 
			WHEN C.SITUACAO = 4 THEN 0
			WHEN C.DATA_CONTRATO >= DATEADD(1 - EXTRACT(DAY FROM CURRENT_DATE) DAY TO CURRENT_DATE ) THEN 0
			ELSE
				((CURRENT_DATE) - (SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
				WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
				ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC))
		END AS DIAS_ATRASO,
		CASE
			WHEN C.SITUACAO = 4 THEN 'ATIVO'
			WHEN C.DATA_CONTRATO >= DATEADD(1 - EXTRACT(DAY FROM CURRENT_DATE) DAY TO CURRENT_DATE ) THEN 'ATIVO'
			WHEN ((CURRENT_DATE) - (SELECT FIRST(1) PC.DATA_VENCIMENTO FROM PARCELAS_CLIENTE PC
		WHERE PC.CLIENTE_ID=C.IDCLIENTE AND PC.DATA_PAGAMENTO IS NOT NULL
		ORDER BY PC.DATA_VENCIMENTO DESC , PC.DATA_PAGAMENTO DESC)) < 90 THEN 'ATIVO'
		ELSE 'INATIVO'
		END AS STATUS
		FROM DEPENDENTE D
		LEFT JOIN CLIENTE C ON C.IDCLIENTE=D.CLIENTE_ID
		WHERE C.CONTRATO = '%d'
		AND D.SITUACAO=1
		AND D.DATA_FALECIMENTO IS NULL AND D.SITUACAO = 1 
		AND D.TIPO_DEPENDENTE = 1 
		AND ((C.SITUACAO =1) OR (C.SITUACAO =4))`, int64(id), int64(id))

	// execute the sql statement
	row := db.QueryRow(sqlStatement)

	err := row.Scan(
		&contrato.Tipo,
		&contrato.Idcliente,
		&contrato.Contrato,
		&contrato.Nome,
		&contrato.Titular,
		&contrato.CpfCnpj,
		&contrato.UltPg,
		&contrato.DiasAtraso,
		&contrato.Status,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return contrato, nil
	case nil:
		return contrato, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return contrato, err
}
