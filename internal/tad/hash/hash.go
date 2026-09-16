// Package hash implementa el TAD tabla de hash con resolución de colisiones
// por encadenamiento (cada bucket es un slice de pares clave-valor).
//
// Es un TAD genérico: la clave es siempre string y el valor puede ser
// cualquier tipo (por ejemplo, *cofre.Cofre o la cantidad de un ítem).
package hash

import "errors"

// ErrNoImplementado se devuelve mientras los métodos del TAD están pendientes
// de implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// parClaveValor es la unidad de almacenamiento dentro de un bucket.
type parClaveValor[V any] struct {
	clave string
	valor V
}

// TablaHash es una tabla de hash con capacidad fija y encadenamiento.
type TablaHash[V any] struct {
	buckets  [][]parClaveValor[V]
	cantidad int
}

// Nueva crea una tabla de hash con la capacidad inicial indicada.
func Nueva[V any](capacidad int) *TablaHash[V] {
	return &TablaHash[V]{
		buckets: make([][]parClaveValor[V], capacidad),
	}
}

// Insertar agrega o reemplaza el valor asociado a una clave.
//
// TODO(implementar): usar funcionHash para ubicar el bucket y la resolución
// de colisiones por encadenamiento; mantener actualizada la cantidad.
func (t *TablaHash[V]) Insertar(clave string, valor V) error {
	return ErrNoImplementado
}

// Obtener devuelve el valor asociado a una clave.
//
// TODO(implementar): recorrer el bucket correspondiente y devolver un error
// si la clave no existe.
func (t *TablaHash[V]) Obtener(clave string) (V, error) {
	var cero V
	return cero, ErrNoImplementado
}

// Contiene indica si la clave existe en la tabla.
func (t *TablaHash[V]) Contiene(clave string) bool {
	return false // TODO(implementar)
}

// Eliminar quita la clave de la tabla y devuelve un error si no existía.
func (t *TablaHash[V]) Eliminar(clave string) error {
	return ErrNoImplementado
}

// Cantidad devuelve la cantidad de claves almacenadas.
func (t *TablaHash[V]) Cantidad() int {
	return t.cantidad
}

// funcionHash convierte una clave en un índice de bucket válido.
//
// TODO(implementar): una función de dispersión propia (por ejemplo, sumar o
// ponderar los bytes de la clave y reducir módulo el tamaño de buckets).
func (t *TablaHash[V]) funcionHash(clave string) int {
	return 0
}
