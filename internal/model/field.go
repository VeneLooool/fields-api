package model

type Field struct {
	ID      uint64 `db:"id"`
	Name    string `db:"name"`
	Culture string `db:"culture"`
}
