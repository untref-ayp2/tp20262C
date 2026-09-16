// Package iterator define una interfaz de iteración uniforme, de modo que las
// colecciones del dominio (inventarios, solicitudes) puedan recorrerse sin
// exponer su estructura interna.
//
// Incluye una implementación de referencia sobre slices; el grupo debe aplicar
// el patrón a sus propias colecciones (por ejemplo, recorrer las solicitudes
// pendientes de un cofre o el inventario).
package iterator

// Iterador recorre una colección elemento por elemento.
type Iterador[T any] interface {
	// Siguiente avanza a la siguiente posición; devuelve false al terminar.
	Siguiente() bool
	// Valor devuelve el elemento en la posición actual.
	Valor() (T, bool)
}

// IteradorSlice es una implementación de referencia del patrón sobre un
// slice (colección concreta: un slice de Go).
type IteradorSlice[T any] struct {
	items []T
	pos   int
}

// DesdeSlice crea un iterador sobre un slice.
func DesdeSlice[T any](items []T) *IteradorSlice[T] {
	return &IteradorSlice[T]{items: items, pos: -1}
}

// Siguiente avanza una posición en el slice.
func (it *IteradorSlice[T]) Siguiente() bool {
	if it.pos+1 >= len(it.items) {
		return false
	}
	it.pos++
	return true
}

// Valor devuelve el elemento actual.
func (it *IteradorSlice[T]) Valor() (T, bool) {
	if it.pos < 0 || it.pos >= len(it.items) {
		var cero T
		return cero, false
	}
	return it.items[it.pos], true
}
