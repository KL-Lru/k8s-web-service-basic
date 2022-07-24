package repos

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*
	Fetch users info
*/
func (repo *Repository) UserList() (users []User, err error) {
	statement := "SELECT id, name FROM users;"
	users = []User{}
	rows, err := repo.Db.Query(statement)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = bindUser(rows, &user); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func bindUser(row Scanner, user *User) (err error) {
	err = row.Scan(&user.Id, &user.Name)
	return
}
