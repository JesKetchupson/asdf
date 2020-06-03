package storage

//TODO
type DB interface {
	Get()
	Update()
	Insert()
	CreateNewTable()
	Delete()
	DeleteTable()
}
