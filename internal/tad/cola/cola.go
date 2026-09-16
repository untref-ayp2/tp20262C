// Package cola implementa el TAD cola (FIFO: primero en entrar, primero en
// salir), genérico sobre el tipo de los elementos.
//
// En el TP se usa para modelar las solicitudes pendientes de un cofre (se
// atienden en orden de llegada) y como desempate FIFO en el ciclo de
// simulación.
package cola

import "errors"

// ErrNoImplementado se devuelve mientras los métodos del TAD están pendientes
// de implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// Cola es una cola FIFO de elementos de tipo T.
type Cola[T any] struct {
	elementos []T
}

// Nueva crea una cola vacía.
func Nueva[T any]() *Cola[T] {
	return &Cola[T]{}
}

// Encolar agrega un elemento al final de la cola.
//
// TODO(implementar): agregar el elemento al final y devolver nil.
func (c *Cola[T]) Encolar(elemento T) error {
	return ErrNoImplementado
}

// Desencolar quita y devuelve el primer elemento de la cola.
//
// TODO(implementar): devolver el primer elemento y quitarlo, o un error si la
// cola está vacía.
func (c *Cola[T]) Desencolar() (T, error) {
	var cero T
	return cero, ErrNoImplementado
}

// EstaVacia indica si la cola no tiene elementos.
func (c *Cola[T]) EstaVacia() bool {
	return len(c.elementos) == 0
}

// Cantidad devuelve la cantidad de elementos encolados.
func (c *Cola[T]) Cantidad() int {
	return len(c.elementos)
}
