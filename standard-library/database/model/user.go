package model

import "database/sql"

type User struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func Create(db *sql.DB, u *User) error {
	err := db.QueryRow(`INSERT INTO sql_users (name, address, age) VALUES ($1, $2, $3)`, u.Name, u.Address, u.Age).
		Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func QueryById(db *sql.DB, userId int) (User, error) {
	var u User
	rows, err := db.Query("SELECT * FROM sql_users WHERE id=$1", userId)
	if err != nil {
		return u, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Address, &u.Age); err != nil {
			return u, err
		}
	}

	return u, nil
}

func UpdateById(db *sql.DB, u *User) error {
	result, err := db.Exec(`UPDATE sql_users SET name=$1, address=$2, age=$3 WHERE id = $4`,
		u.Name, u.Address, u.Age, u.ID)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		return err
	}

	return nil
}

func DeleteById(db *sql.DB, userId int) error {
	result, err := db.Exec(`DELETE FROM sql_users WHERE id=$1`, userId)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		return err
	}

	return nil
}
