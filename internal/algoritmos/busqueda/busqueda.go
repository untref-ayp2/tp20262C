// Package busqueda implementa búsqueda lineal y búsqueda binaria sobre el
// catálogo de ítems (nombres) de la red logística.
package busqueda

import "errors"

// ErrNoImplementado se devuelve mientras la función está pendiente de
// implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// Binaria busca un nombre en un slice ordenado ascendentemente y devuelve la
// posición de la primera aparición, o -1 si no existe.
//
// TODO(implementar): búsqueda binaria con partición sucesiva por la mitad;
// precondición: items está ordenado (ver internal/algoritmos/ordenamiento).
func Binaria(items []string, objetivo string) int {
	return -1
}

// Lineal recorre items y devuelve la primera posición cuyo elemento cumple el
// criterio. Se usa cuando el criterio de búsqueda no admite un orden total
// (por ejemplo, buscar cofres por rol).
func Lineal[T any](items []T, coincide func(T) bool) (int, bool) {
	for i, item := range items {
		if coincide(item) {
			return i, true
		}
	}
	return 0, false
}
