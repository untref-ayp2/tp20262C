// Package ordenamiento implementa los ordenamientos cuadráticos de la cursada
// (selección e inserción), aplicados al catálogo de ítems por nombre.
//
// El grupo elige UNO de los dos y justifica la elección en el informe según el
// escenario (por ejemplo, pocos elementos casi ordenados favorecen inserción).
package ordenamiento

import "errors"

// ErrNoImplementado se devuelve mientras la función está pendiente de
// implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// Seleccion ordena ascendentemente un slice de nombres in-place (por
// selección).
//
// TODO(implementar): en cada pasada, buscar el menor y ubicarlo en su
// posición definitiva.
func Seleccion(items []string) error {
	return ErrNoImplementado
}

// Insercion ordena ascendentemente un slice de nombres in-place (por
// inserción).
//
// TODO(implementar): insertar cada elemento en la posición que le
// corresponde en el prefijo ya ordenado.
func Insercion(items []string) error {
	return ErrNoImplementado
}
