// Package cofre define el modelo de cofres logísticos y sus roles.
package cofre

// Rol indica la función que cumple un cofre respecto de los ítems.
type Rol int

const (
	// RolProvision ofrece ítems a otros cofres.
	RolProvision Rol = iota
	// RolSolicitud solicita una cantidad de un ítem.
	RolSolicitud
	// RolAlmacenamiento recibe excedentes sin destino asignado.
	RolAlmacenamiento
)

// String devuelve el nombre del rol.
func (r Rol) String() string {
	switch r {
	case RolProvision:
		return "provision"
	case RolSolicitud:
		return "solicitud"
	case RolAlmacenamiento:
		return "almacenamiento"
	default:
		return "desconocido"
	}
}

// Solicitud es un pedido de una cantidad de un ítem.
type Solicitud struct {
	Item     string
	Cantidad int
}

// Cofre es un contenedor inteligente de ítems con un rol único.
type Cofre struct {
	ID          string
	Rol         Rol
	X, Y        float64
	Inventario  map[string]int
	Solicitudes []Solicitud
}

// Nuevo crea un cofre vacío con el rol indicado.
func Nuevo(id string, rol Rol, x, y float64) *Cofre {
	return &Cofre{
		ID:         id,
		Rol:        rol,
		X:          x,
		Y:          y,
		Inventario: make(map[string]int),
	}
}

// AgregarStock suma una cantidad de un ítem al inventario.
func (c *Cofre) AgregarStock(item string, cantidad int) {
	c.Inventario[item] += cantidad
}

// StockDe devuelve la cantidad disponible de un ítem.
func (c *Cofre) StockDe(item string) int {
	return c.Inventario[item]
}

// TieneSolicitudes indica si el cofre tiene pedidos sin satisfacer.
func (c *Cofre) TieneSolicitudes() bool {
	return len(c.Solicitudes) > 0
}
