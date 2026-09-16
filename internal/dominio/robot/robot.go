// Package robot define los robots logísticos y su consumo de energía.
package robot

import "math"

// Robot es una unidad móvil que transporta ítems entre cofres, con batería
// medida en Células y capacidad de carga limitada.
type Robot struct {
	ID               string `json:"id"`
	BateriaMaxima    int    `json:"bateria_maxima"`
	BateriaActual    int    `json:"bateria_actual"`
	CapacidadCarga   int    `json:"capacidad_carga"`
	RobopuertoOrigen string `json:"robopuerto_origen"`
}

// ConsumoPara devuelve las Células necesarias para recorrer una distancia con
// el factor de consumo indicado (se redondea hacia arriba).
func (r *Robot) ConsumoPara(distancia, factor float64) int {
	return int(math.Ceil(distancia * factor))
}

// PuedeViajar indica si la batería actual alcanza para recorrer la distancia.
func (r *Robot) PuedeViajar(distancia, factor float64) bool {
	return r.BateriaActual >= r.ConsumoPara(distancia, factor)
}

// Recargar restaura la batería al máximo (requiere un ciclo de carga en un
// robopuerto).
func (r *Robot) Recargar() {
	r.BateriaActual = r.BateriaMaxima
}

// Consumir descuenta las Células de la batería tras un viaje.
func (r *Robot) Consumir(cels int) {
	r.BateriaActual -= cels
}
