package directorio

type DB interface {
	BuscarPorNombre(nombre string) string
	BuscarPorTelefono(telefono string) string
	AgregarEntrada(nombre, telefono string) error
}

type database struct{}

func NewDB() DB {
	return &database{}
}

func (db database) BuscarPorNombre(nombre string) string         { return "" }
func (db database) BuscarPorTelefono(telefono string) string     { return "" }
func (db database) AgregarEntrada(nombre, telefono string) error { return nil }
