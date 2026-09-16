// Package composite implementa el patrón Composite aplicado a la red
// logística: un robopuerto individual y la red completa se tratan con la
// misma interfaz, permitiendo operaciones uniformes como "listar cofres" o
// "total de un ítem".
package composite

import "tp-red-logistica/internal/dominio/cofre"

// Componente es la interfaz común del patrón: las hojas (robopuertos con sus
// cofres) y el compuesto (la red) exponen las mismas operaciones.
type Componente interface {
	// Nombre identifica al componente.
	Nombre() string
	// ListarCofres devuelve los cofres asociados al componente.
	ListarCofres() []*cofre.Cofre
	// TotalItem suma las cantidades disponibles de un ítem entre sus cofres.
	TotalItem(item string) int
}
