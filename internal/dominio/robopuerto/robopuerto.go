// Package robopuerto define los robopuertos y su zona de cobertura.
package robopuerto

import "math"

// Robopuerto es una infraestructura fija con una zona de cobertura circular y
// que funciona como estación de recarga de robots.
type Robopuerto struct {
	ID      string
	X, Y    float64
	Alcance float64
}

// Cubre indica si el punto (x, y) queda dentro de la zona de cobertura.
func (r *Robopuerto) Cubre(x, y float64) bool {
	return Distancia(r.X, r.Y, x, y) <= r.Alcance
}

// ConectaCon indica si dos robopuertos tienen zonas superpuestas (sus centros
// distan menos que la suma de los alcances).
func (r *Robopuerto) ConectaCon(otro *Robopuerto) bool {
	return Distancia(r.X, r.Y, otro.X, otro.Y) <= r.Alcance+otro.Alcance
}

// Distancia calcula la distancia euclídea entre dos puntos.
func Distancia(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Hypot(dx, dy)
}
