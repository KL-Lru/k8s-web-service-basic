package repos

type Scanner interface {
	Scan(dist ...interface{}) error
}
