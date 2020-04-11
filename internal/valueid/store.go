package valueid

// Maybe some kind of store interface?

type Store interface {
	GetByID(id int64) (value string, ok bool)
	GetOrCreateValueID(value string) (ID int64, err error)
}
