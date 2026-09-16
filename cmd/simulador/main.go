// Command simulador es el punto de entrada del trabajo práctico: carga la
// configuración desde un archivo JSON y corre la simulación.
//
// Uso:
//
//	go run ./cmd/simulador data/basico/entrada.json
package main

import (
	"fmt"
	"os"

	"tp-red-logistica/internal/carga"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: go run ./cmd/simulador <entrada.json>")
		os.Exit(1)
	}

	cfg, err := carga.Cargar(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error al cargar la configuración: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Red cargada: %d robopuertos, %d cofres, %d robots, factor de consumo %.2f\n",
		len(cfg.Robopuertos), len(cfg.Cofres), len(cfg.Robots), cfg.FactorConsumo)
	for _, c := range cfg.Cofres {
		fmt.Printf("  cofre %s (%s) en (%.0f, %.0f)\n", c.ID, c.Rol, c.X, c.Y)
	}

	// TODO: una vez implementado internal/simulacion, conectar el ciclo:
	//   resultado, err := simulacion.Simular(cfg)
	//   ...escribir el resultado en output/simulacion.txt
	fmt.Println("(la simulación todavía no está implementada: completar internal/simulacion)")
}
