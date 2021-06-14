package repositories

const (
	SqlCreateUser = `INSERT INTO user(name, nick, email, password, created_at) VALUES (
		?,?,?,?,?	
	); `

	SelectAllUsers = `SELECT * FROM user`
)
