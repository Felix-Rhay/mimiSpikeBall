package repository

import (
	"database/sql"
	"mimiSpikeBall/internal/models"
)

type CompteRepository struct {
	db *sql.DB
}

func NewCompteRepository(db *sql.DB) *CompteRepository {
	return &CompteRepository{db: db}
}

func (r *CompteRepository) ObtenirComptes() (*[]models.Compte, error) {
	rows, err := r.db.Query("SELECT login FROM COMPTE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Compte
	for rows.Next() {
		var u models.Compte
		if err := rows.Scan(&u.Nom); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &users, nil
}

func (r *CompteRepository) LoginUser(entrant *models.LoginUserRequest) (bool, error) {
	rows := r.db.QueryRow("SELECT count(*) FROM COMPTE WHERE login = ? AND mot_de_passe = ?", entrant.Login, entrant.Password)

	var count int
	scanError := rows.Scan(&count)

	if scanError != nil {
		return false, scanError
	}

	return count > 0, nil
}
