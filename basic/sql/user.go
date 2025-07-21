package sql

type User struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func (u *User) TableName() string {
	return "sql_users"
}

func (u *User) Create() error {
	result, err := pgDB.Exec(`INSERT INTO sql_users (name, address, age) VALUES ($1, $2, $3)`, u.Name, u.Address, u.Age)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		return err
	}

	return nil
}

func (u *User) CreateReturnId() error {
	var id int
	err := pgDB.QueryRow(`INSERT INTO sql_users (name, address, age) VALUES ($1, $2, $3) RETURNING id`, u.Name, u.Address, u.Age).
		Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) QueryById() error {
	rows, err := pgDB.Query("SELECT * FROM sql_users WHERE id=$1", u.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Address, &u.Age); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) UpdateById() error {
	result, err := pgDB.Exec(`UPDATE sql_users SET name=$1, address=$2, age=$3 WHERE id = $4`,
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

func (u *User) DeleteById() error {
	result, err := pgDB.Exec(`DELETE FROM sql_users WHERE id=$1`, u.ID)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		return err
	}

	return nil
}
