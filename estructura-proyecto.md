# Estructura de proyecto — Red Logística Automatizada

Estructura de referencia del código Go del trabajo práctico. Sigue las
convenciones habituales de un módulo Go (`cmd/`, `internal/`) y separa los TAD
de la lógica de dominio y de los algoritmos, de modo que cada hito agregue
paquetes concretos sin reescribir los anteriores.

Puede adaptarse a las necesidades del grupo, pero se recomienda mantener la
separación entre TAD, dominio, patrones y algoritmos, ya que es la que evalúa
la rúbrica.

```txt
red-logistica/
├── go.mod
├── README.md
├── Makefile
├── .gitignore
│
├── cmd/
│   └── simulador/
│       └── main.go             # Punto de entrada: carga datos y corre la
│                                 simulación
│
├── internal/
│   ├── tad/                    # TAD propios del proyecto
│   │   ├── cola/               # Cola FIFO (Entrega 1)
│   │   ├── hash/               # Tabla de hash propia (Entrega 1)
│   │   └── heap/               # Montículo binario / cola de prioridad
│   │                             (Entrega final)
│   │
│   ├── dominio/                # Tipos del dominio (Entrega 1 en adelante)
│   │   ├── cofre/              # Los 3 roles de cofre
│   │   ├── robopuerto/         # Cobertura y conexión entre robopuertos
│   │   ├── robot/              # Batería, consumo y recarga (Entrega final)
│   │   └── red/                # Red construida y compuesta (Entrega 1 / final)
│   │
│   ├── patrones/               # Patrones de diseño (Entrega final)
│   │   ├── iterator/           # Interfaz de iteración uniforme
│   │   └── composite/          # Interfaz Componente del Composite
│   │
│   ├── algoritmos/
│   │   ├── busqueda/           # Lineal y binaria (Entrega 1)
│   │   ├── ordenamiento/       # Selección o inserción (Entrega 1)
│   │   └── greedy/             # Selección de proveedor (Entrega final)
│   │
│   ├── carga/                  # Lectura del JSON de entrada (resuelto)
│   │
│   └── simulacion/             # Ciclo de simulación y estado estable
│                                 (Entrega final)
│
├── data/
│   ├── basico/                 # Escenario mínimo (Entrega 1)
│   ├── estable/                # Estado estable alcanzable
│   └── no-estable/             # Solicitudes no satisfacibles
│
├── output/                     # Archivos de salida generados por la
│                                 simulación (no versionar corridas)
│
└── docs/
    ├── informe.md              # Copia o enlace al informe (ver template.md)
    └── diagramas/              # Diagramas de arquitectura y de estructuras
```

## Convenciones de paquetes

- Cada paquete bajo `internal/tad/` expone únicamente su tipo abstracto y sus
  operaciones (`Encolar`, `Desencolar`, `Insertar`, `Obtener`,
  `ExtraerMinimo`, etc.); no conoce nada del dominio logístico.
- El dominio (`internal/dominio/`) modela los conceptos del problema y **no**
  depende de los TAD; estos se aplican desde `red`, `algoritmos` y
  `simulacion`.
- Los paquetes bajo `internal/patrones/` contienen las interfaces de los
  patrones; las implementaciones concretas viven en `dominio/red` (Composite)
  o se aplican en `simulacion` (Iterator).
- Los paquetes bajo `internal/algoritmos/` operan sobre tipos del dominio o
  sobre tipos genéricos, y no dependen de `simulacion` ni de `cmd/`.
- `internal/simulacion/` es el único paquete que conoce y orquesta a los
  demás.
- `cmd/simulador/main.go` se limita a leer la ruta del archivo, invocar
  `internal/carga`, ejecutar `internal/simulacion` y escribir la salida; no
  contiene lógica de negocio.

## Pruebas

Cada paquete de `internal/` debe incluir sus propios archivos `_test.go`
junto al código que prueban (convención estándar de Go), por ejemplo:

```txt
internal/tad/hash/
├── hash.go
└── hash_test.go
```

Se recomienda incluir, además, pruebas de integración en
`internal/simulacion/simulacion_test.go` que corran la simulación completa
sobre los escenarios de `data/` y verifiquen el resultado esperado (estado
estable alcanzado o solicitudes reportadas con su causa).

## Correspondencia con los hitos

| Hito                     | Paquetes que se agregan o completan                                                                                                             |
| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| Checkpoint (Entrega 1)   | `internal/tad/cola`, `internal/tad/hash`, `internal/dominio/cofre`, `internal/dominio/robopuerto`, `internal/dominio/red`, `internal/algoritmos/busqueda`, `internal/algoritmos/ordenamiento` |
| Entrega final            | `internal/dominio/robot`, `internal/patrones/iterator`, `internal/patrones/composite`, `internal/algoritmos/greedy`, `internal/tad/heap`, `internal/simulacion` |

## Próximos pasos sugeridos

1. Familiarizarse con el template: correr `go build ./...` y
   `go test ./...`, y luego `go run ./cmd/simulador data/basico/entrada.json`.
2. Acordar en el grupo el diseño de cada TAD (función de dispersión del hash,
   criterio de prioridad del montículo) antes de escribir código.
3. Completar los TAD y algoritmos con sus pruebas módulo por módulo, y correr
   `make test` en cada avance relevante.