// Package greedy implementa el algoritmo ávido de selección de proveedor:
// para una solicitud, elige el cofre más conveniente entre los que pueden
// satisfacerla.
package greedy

import (
	"errors"

	"tp-red-logistica/internal/dominio/cofre"
)

// ErrNoImplementado se devuelve mientras la función está pendiente de
// implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// SeleccionarProveedor elige, para la solicitud de un cofre, el proveedor más
// conveniente entre los candidatos: cofres con stock del ítem solicitado.
// La regla ávida propuesta: primero rol Provisión, y dentro del mismo rol, la
// menor distancia euclídea al cofre solicitante.
//
// TODO(implementar): aplicar la regla ávida sobre proveedores y devolver un
// error si ningún candidato puede satisfacer la solicitud (sin stock o sin
// cobertura alcanzable).
func SeleccionarProveedor(solicitante *cofre.Cofre, solicitud cofre.Solicitud, proveedores []*cofre.Cofre) (*cofre.Cofre, error) {
	return nil, ErrNoImplementado
}
