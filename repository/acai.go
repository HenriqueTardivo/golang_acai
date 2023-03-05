package repository

import (
	"goapp/database"
	"goapp/models"
	"strconv"
)

func CreateAcai() {

}

func ListAcais(id int) ([]models.Acai, error) {
	query := "select * from acai "

	if id > 0 {
		query += " where id_acai = " + strconv.Itoa(id)
	}

	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	acais := make([]models.Acai, 0)

	for rows.Next() {
		acai := models.Acai{}

		err = rows.Scan(&acai.id_acai)
	}

	return acais, nil
}
