package connection

//Connection is the interface for types related with external services
type Connection interface {
	Connect() error
	Close() error
}
