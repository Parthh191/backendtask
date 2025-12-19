package repository

import (
	"database/sql"
	"github.com/Parthh191/backendtask/internal/models"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (name, dob, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, dob, created_at, updated_at
	`

	err:=r.db.QueryRow(
		query,
		user.Name,
		user.DOB,
		time.Now(),
		time.Now(),
	).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err!=nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `
		SELECT id, name, dob, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user:=&models.User{}
	err:=r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err!=nil {
		if err==sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users from the database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query:=`
		SELECT id, name, dob, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err:=r.db.Query(query)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err:=rows.Scan(
			&user.ID,
			&user.Name,
			&user.DOB,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err!=nil {
			return nil, err
		}
		users=append(users, user)
	}

	return users, rows.Err()
}

// UpdateUser updates an existing user
func (r *UserRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	query:=`
		UPDATE users
		SET name = $1, dob = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, name, dob, created_at, updated_at
	`

	err:=r.db.QueryRow(
		query,
		user.Name,
		user.DOB,
		time.Now(),
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.DOB,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err!=nil {
		if err==sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id int) error {
	query:=`DELETE FROM users WHERE id = $1`
	_, err:=r.db.Exec(query, id)
	return err
}

// GetUserByName retrieves users by name (partial match)
func (r *UserRepository) GetUserByName(name string) ([]models.User, error) {
	query:=`
		SELECT id, name, dob, created_at, updated_at
		FROM users
		WHERE name ILIKE $1
		ORDER BY created_at DESC
	`

	rows, err:=r.db.Query(query, "%"+name+"%")
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err:=rows.Scan(
			&user.ID,
			&user.Name,
			&user.DOB,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err!=nil {
			return nil, err
		}
		users=append(users, user)
	}

	return users, rows.Err()
}
