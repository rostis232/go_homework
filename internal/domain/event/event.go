package event

type Event struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
