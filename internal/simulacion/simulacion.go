// Package simulacion contiene el ciclo de simulación de la red logística:
// a partir de la configuración cargada, ejecuta ciclos de transferencia hasta
// alcanzar el estado estable o hasta reportar, con causa, las solicitudes que
// no pueden satisfacerse.
//
// Es el paquete que orquesta a los demás: usa el montículo (cola de
// prioridad), el algoritmo ávido, los robots y la red.
package simulacion

import (
	"errors"

	"tp-red-logistica/internal/carga"
)

// ErrNoImplementado se devuelve mientras la función está pendiente de
// implementación en el trabajo práctico.
var ErrNoImplementado = errors.New("no implementado: completar en el TP")

// SolicitudPendiente describe una solicitud que no pudo satisfacerse al
// finalizar la simulación, con la causa (falta de stock, cofre inaccesible,
// batería insuficiente, etc.).
type SolicitudPendiente struct {
	Cofre string
	Item  string
	Causa string
}

// Resultado resume la corrida de la simulación.
type Resultado struct {
	// EstadoEstable es true si todas las solicitudes quedaron satisfechas y no
	// hay ítems en tránsito.
	EstadoEstable bool
	// Ciclos es la cantidad de ciclos ejecutados.
	Ciclos int
	// Eventos es el registro cronológico de los movimientos y decisiones.
	Eventos []string
	// NoSatisfechas son las solicitudes que quedaron pendientes.
	NoSatisfechas []SolicitudPendiente
}

// Simular ejecuta el ciclo de simulación completo sobre la configuración
// cargada y devuelve el resultado.
//
// Lineamientos (el detalle de diseño queda a criterio del grupo y se
// documenta en el informe):
//
//  1. Construir la red (red.ConstruirRed) y detectar cofres inaccesibles.
//  2. En cada ciclo, encolar en el montículo las solicitudes pendientes con
//     su prioridad (a elección justificada: cantidad faltante, cercanía al
//     proveedor más conveniente, etc.).
//  3. Extraer la solicitud más prioritaria y seleccionar proveedor con el
//     algoritmo ávido (greedy); si la batería del robot no alcanza, registrar
//     el evento "recarga" y reintentar en el siguiente ciclo.
//  4. Transferir los ítems en viajes que respetan la capacidad de carga de
//     los robots; los excedentes sin destino van a cofres de almacenamiento.
//  5. Terminar cuando no quedan solicitudes (estado estable) o cuando una
//     iteración completa no produjo ningún avance (reportar las solicitudes
//     pendientes con su causa).
//
// TODO(implementar): el ciclo completo, integrando los demás paquetes.
func Simular(cfg carga.Configuracion) (Resultado, error) {
	return Resultado{}, ErrNoImplementado
}
