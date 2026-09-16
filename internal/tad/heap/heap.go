// Package heap implementa el TAD montículo binario mínimo (cola de
// prioridad): el elemento con la prioridad más baja (número más chico) es el
// primero en extraerse.
//
// En el TP se usa en el ciclo de simulación para decidir qué solicitud
// atender primero en cada ciclo.
package heap

import "errors"

// ErrNoImplementado se devuelve mientras los métodos del TAD están pendientes
// de implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// Elemento asocia un valor con una prioridad entera.
type Elemento[T any] struct {
	Prioridad int
	Valor     T
}

// Monticulo es un montículo binario mínimo de elementos de tipo T.
type Monticulo[T any] struct {
	datos []Elemento[T]
}

// Nuevo crea un montículo vacío.
func Nuevo[T any]() *Monticulo[T] {
	return &Monticulo[T]{}
}

// Insertar agrega un elemento con su prioridad.
//
// TODO(implementar): agregar al final y restaurar la propiedad del montículo
// (flotar hacia arriba).
func (m *Monticulo[T]) Insertar(prioridad int, valor T) error {
	return ErrNoImplementado
}

// ExtraerMinimo quita y devuelve el elemento de menor prioridad.
//
// TODO(implementar): devolver la raíz, reemplazarla por el último elemento y
// hundir hacia abajo; error si el montículo está vacío.
func (m *Monticulo[T]) ExtraerMinimo() (Elemento[T], error) {
	var cero Elemento[T]
	return cero, ErrNoImplementado
}

// EstaVacio indica si el montículo no tiene elementos.
func (m *Monticulo[T]) EstaVacio() bool {
	return len(m.datos) == 0
}

// Cantidad devuelve la cantidad de elementos en el montículo.
func (m *Monticulo[T]) Cantidad() int {
	return len(m.datos)
}
