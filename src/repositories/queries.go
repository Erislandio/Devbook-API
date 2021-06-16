package repositories

const (
	SqlCreateUser = `INSERT INTO user(name, nick, email, password, created_at) VALUES (
		?,?,?,?,?	
	); `

	SelectAllUsers = `select id, name, nick, email, created_at from user;`
)
