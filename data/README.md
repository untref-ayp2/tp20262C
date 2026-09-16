# Escenarios de ejemplo

Cada carpeta contiene un archivo `entrada.json` con la configuración de la red
logística y debajo la descripción del resultado esperado. Los grupos deben
generar **al menos dos escenarios propios** además de estos tres, cubriendo el
caso de éxito y el caso de falla controlada.

## `basico/`

Escenario mínimo de entrada: valida la carga de datos y el modelo (Entrega 1).

- 1 robopuerto (P1), 3 cofres y 1 robot.
- C1 provee 10 de hierro, C2 solicita 5 de hierro, C3 almacenamiento.
- Resultado esperado de la simulación completa: estado estable en 1 ciclo con
  batería de sobra. Antes de la Entrega 2, sirve para probar la carga, el hash
  y la cobertura.

## `estable/`

Red donde el sistema **alcanza el estado estable**.

- 2 robopuertos con zonas superpuestas (red conexa), 5 cofres y 2 robots.
- C1 provee 30 de hierro → C3 solicita 10; C2 provee 20 de circuitos → C4
  solicita 8; C5 es almacenamiento.
- Resultado esperado: todas las solicitudes satisfechas, sin ítems en
  tránsito, con registros de viajes y recargas en el archivo de salida.

## `no-estable/`

Red donde **no se alcanza** el estado estable: las solicitudes pendientes
deben reportarse con su causa.

- 1 robopuerto (P1, alcance 5).
- C1 provee solo 3 de hierro y C2 solicita 10 → falta de stock.
- C3 solicita circuitos pero está fuera de toda cobertura → cofre inaccesible.
- Resultado esperado: dos solicitudes no satisfechas, cada una con su causa
  (falta de stock / cofre inaccesible).