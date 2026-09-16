// Package red modela la red logística con el patrón Composite (la interfaz
// Componente vive en internal/patrones/composite) y mantiene un índice de
// cofres por identificador usando el TAD tabla de hash.
package red

import (
	"errors"

	"tp-red-logistica/internal/dominio/cofre"
	"tp-red-logistica/internal/dominio/robopuerto"
	"tp-red-logistica/internal/patrones/composite"
	"tp-red-logistica/internal/tad/hash"
)

// ErrNoImplementado se devuelve mientras las funciones están pendientes de
// implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// RobopuertoNodo es la hoja del Composite: un robopuerto con los cofres que
// caen dentro de su zona de cobertura.
type RobopuertoNodo struct {
	puertoID string
	cofres   []*cofre.Cofre
}

// NuevoRobopuertoNodo crea la hoja para un robopuerto y sus cofres.
func NuevoRobopuertoNodo(puertoID string, cofres []*cofre.Cofre) *RobopuertoNodo {
	return &RobopuertoNodo{puertoID: puertoID, cofres: cofres}
}

// Nombre devuelve el identificador del robopuerto.
func (n *RobopuertoNodo) Nombre() string {
	return n.puertoID
}

// ListarCofres devuelve los cofres de la hoja.
func (n *RobopuertoNodo) ListarCofres() []*cofre.Cofre {
	return n.cofres
}

// TotalItem suma las cantidades del ítem entre los cofres de la hoja.
func (n *RobopuertoNodo) TotalItem(item string) int {
	total := 0
	for _, c := range n.cofres {
		total += c.StockDe(item)
	}
	return total
}

// RedLogistica es el compuesto del Composite: agrupa componentes y delega las
// operaciones. Además mantiene un índice de cofres por ID con el TAD hash.
type RedLogistica struct {
	componentes  []composite.Componente
	indiceCofres *hash.TablaHash[*cofre.Cofre]
}

// NuevaRed crea una red vacía con su índice de cofres.
func NuevaRed() *RedLogistica {
	return &RedLogistica{
		indiceCofres: hash.Nueva[*cofre.Cofre](16),
	}
}

// Agregar incorpora un componente a la red.
func (r *RedLogistica) Agregar(c composite.Componente) {
	r.componentes = append(r.componentes, c)
}

// ListarCofres delega el recorrido en todos los componentes.
func (r *RedLogistica) ListarCofres() []*cofre.Cofre {
	var todos []*cofre.Cofre
	for _, c := range r.componentes {
		todos = append(todos, c.ListarCofres()...)
	}
	return todos
}

// TotalItem delega el cálculo en todos los componentes.
func (r *RedLogistica) TotalItem(item string) int {
	total := 0
	for _, c := range r.componentes {
		total += c.TotalItem(item)
	}
	return total
}

// CofrePorID busca un cofre en el índice por su identificador.
//
// TODO(implementar): usar indiceCofres (TAD hash) para la búsqueda.
func (r *RedLogistica) CofrePorID(id string) (*cofre.Cofre, error) {
	return nil, ErrNoImplementado
}

// ConstruirRed arma la red a partir de los robopuertos y los cofres cargados:
// asigna cada cofre al/los robopuertos que lo cubren (distancia euclídea
// dentro del alcance) y detecta los cofres sin cobertura, que quedan fuera.
//
// TODO(implementar): crear los nodos, asignar cofres por cobertura e indexar
// los cofres por ID en indiceCofres.
func ConstruirRed(puertos []*robopuerto.Robopuerto, cofres []*cofre.Cofre) (*RedLogistica, error) {
	return nil, ErrNoImplementado
}
