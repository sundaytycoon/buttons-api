package model

const (
	GetUser = "SELECT id, name, state FROM profileme_users WHERE id = ? ORDER BY id DESC"
)