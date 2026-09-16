// Package carga lee la configuración inicial de la simulación desde un
// archivo JSON y la convierte en tipos del dominio.
//
// Este paquete viene resuelto en el template: es el punto de partida para que
// el resto del sistema tenga datos con qué trabajar.
package carga

import (
	"encoding/json"
	"fmt"
	"os"

	"tp-red-logistica/internal/dominio/cofre"
	"tp-red-logistica/internal/dominio/robopuerto"
	"tp-red-logistica/internal/dominio/robot"
)

// Configuracion es la configuración completa de una simulación.
type Configuracion struct {
	// FactorConsumo multiplica la distancia para calcular Células por viaje.
	FactorConsumo float64             `json:"factor_consumo"`
	Robopuertos   []RobopuertoEntrada `json:"robopuertos"`
	Cofres        []CofreEntrada      `json:"cofres"`
	Robots        []robot.Robot       `json:"robots"`
}

// RobopuertoEntrada es la representación de un robopuerto en el JSON.
type RobopuertoEntrada struct {
	ID      string  `json:"id"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Alcance float64 `json:"alcance"`
}

// CofreEntrada es la representación de un cofre en el JSON, con el rol como
// string ("provision", "solicitud" o "almacenamiento").
type CofreEntrada struct {
	ID          string             `json:"id"`
	Rol         string             `json:"rol"`
	X           float64            `json:"x"`
	Y           float64            `json:"y"`
	Inventario  map[string]int     `json:"inventario"`
	Solicitudes []SolicitudEntrada `json:"solicitudes"`
}

// SolicitudEntrada es la representación de una solicitud en el JSON.
type SolicitudEntrada struct {
	Item     string `json:"item"`
	Cantidad int    `json:"cantidad"`
}

// Cargar lee y valida un archivo JSON de configuración.
func Cargar(ruta string) (Configuracion, error) {
	archivo, err := os.ReadFile(ruta)
	if err != nil {
		return Configuracion{}, fmt.Errorf("leer configuración: %w", err)
	}
	var cfg Configuracion
	if err := json.Unmarshal(archivo, &cfg); err != nil {
		return Configuracion{}, fmt.Errorf("parsear configuración: %w", err)
	}
	if cfg.FactorConsumo <= 0 {
		return Configuracion{}, fmt.Errorf("factor_consumo debe ser positivo")
	}
	if len(cfg.Robopuertos) == 0 || len(cfg.Cofres) == 0 {
		return Configuracion{}, fmt.Errorf("la configuración debe definir robopuertos y cofres")
	}
	for _, c := range cfg.Cofres {
		if _, err := rolDesdeString(c.Rol); err != nil {
			return Configuracion{}, err
		}
	}
	return cfg, nil
}

// ARobopuertos convierte las entradas de robopuertos en tipos del dominio.
func ARobopuertos(cfg Configuracion) []*robopuerto.Robopuerto {
	puertos := make([]*robopuerto.Robopuerto, 0, len(cfg.Robopuertos))
	for _, p := range cfg.Robopuertos {
		puertos = append(puertos, &robopuerto.Robopuerto{
			ID:      p.ID,
			X:       p.X,
			Y:       p.Y,
			Alcance: p.Alcance,
		})
	}
	return puertos
}

// ACofres convierte las entradas de cofres en tipos del dominio, con su
// inventario y sus solicitudes.
func ACofres(cfg Configuracion) ([]*cofre.Cofre, error) {
	cofres := make([]*cofre.Cofre, 0, len(cfg.Cofres))
	for _, c := range cfg.Cofres {
		rol, err := rolDesdeString(c.Rol)
		if err != nil {
			return nil, err
		}
		cofreNuevo := cofre.Nuevo(c.ID, rol, c.X, c.Y)
		for item, cantidad := range c.Inventario {
			cofreNuevo.AgregarStock(item, cantidad)
		}
		for _, s := range c.Solicitudes {
			cofreNuevo.Solicitudes = append(cofreNuevo.Solicitudes, cofre.Solicitud{
				Item:     s.Item,
				Cantidad: s.Cantidad,
			})
		}
		cofres = append(cofres, cofreNuevo)
	}
	return cofres, nil
}

// rolDesdeString convierte el nombre textual del rol en la enumeración del
// dominio.
func rolDesdeString(s string) (cofre.Rol, error) {
	switch s {
	case "provision":
		return cofre.RolProvision, nil
	case "solicitud":
		return cofre.RolSolicitud, nil
	case "almacenamiento":
		return cofre.RolAlmacenamiento, nil
	default:
		return 0, fmt.Errorf("rol desconocido %q (esperado: provision, solicitud o almacenamiento)", s)
	}
}
