package infrastructure

type IDatabase interface {
	Exec() error
}
