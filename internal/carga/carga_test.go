package carga

import (
	"testing"

	"tp-red-logistica/internal/dominio/cofre"
)

// TestCargarBasico es un ejemplo de prueba sobre el escenario mínimo: valida
// que el archivo data/basico/entrada.json se carga y convierte correctamente.
// Cada grupo debe escribir sus propias pruebas para los demás módulos.
func TestCargarBasico(t *testing.T) {
	cfg, err := Cargar("../../data/basico/entrada.json")
	if err != nil {
		t.Fatalf("Cargar: %v", err)
	}

	if got := len(cfg.Robopuertos); got != 1 {
		t.Errorf("robopuertos: se esperaba 1, se obtuvo %d", got)
	}
	if got := len(cfg.Cofres); got != 3 {
		t.Errorf("cofres: se esperaban 3, se obtuvieron %d", got)
	}
	if got := len(cfg.Robots); got != 1 {
		t.Errorf("robots: se esperaba 1, se obtuvo %d", got)
	}
	if got := cfg.FactorConsumo; got != 1 {
		t.Errorf("factor_consumo: se esperaba 1, se obtuvo %v", got)
	}

	cofres, err := ACofres(cfg)
	if err != nil {
		t.Fatalf("ACofres: %v", err)
	}
	for _, c := range cofres {
		if c.Rol == cofre.RolProvision && c.StockDe("hierro") != 10 {
			t.Errorf("cofre %s: se esperaban 10 de hierro, se obtuvieron %d", c.ID, c.StockDe("hierro"))
		}
		if c.Rol == cofre.RolSolicitud && len(c.Solicitudes) != 1 {
			t.Errorf("cofre %s: se esperaba 1 solicitud, se obtuvieron %d", c.ID, len(c.Solicitudes))
		}
	}
}

// TestRolDesdeString cubre el mapeo textual de roles y el caso inválido.
func TestRolDesdeString(t *testing.T) {
	casos := []struct {
		texto    string
		rol      cofre.Rol
		invalido bool
	}{
		{"provision", cofre.RolProvision, false},
		{"solicitud", cofre.RolSolicitud, false},
		{"almacenamiento", cofre.RolAlmacenamiento, false},
		{"buzon", 0, true},
	}
	for _, c := range casos {
		rol, err := rolDesdeString(c.texto)
		if c.invalido && err == nil {
			t.Errorf("%q: se esperaba error", c.texto)
		}
		if !c.invalido && rol != c.rol {
			t.Errorf("%q: se esperaba %v, se obtuvo %v", c.texto, c.rol, rol)
		}
	}
}
