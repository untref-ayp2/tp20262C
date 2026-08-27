# Estructura de proyecto — Red Logística Automatizada

Estructura de referencia para organizar el código Go del trabajo práctico. Sigue
las convenciones habituales de un módulo Go (`cmd/`, `internal/`) y separa las
estructuras de datos abstractas (TAD) de la lógica de dominio y de los
algoritmos, de modo que cada entrega agregue paquetes concretos sin reescribir
los anteriores.

Puede adaptarse a las necesidades del grupo, pero se recomienda mantener la
separación entre TAD, dominio y algoritmos, ya que es la que evalúa la rúbrica.

```txt
red-logistica/
├── go.mod
├── go.sum
├── README.md
├── .gitignore
│
├── cmd/
│   └── simulador/
│       └── main.go             # Punto de entrada: Carga datos, corre la
│                                 simulación, escribe la salida
│
├── internal/
│   ├── tad/                    # Estructuras de datos abstractas propias
│   │   │                         (Entrega 1 y 3)
│   │   ├── list/               # Lista simple, doble y circular
│   │   ├── stack/              # Implementación de pila
│   │   ├── queue/              # Implementación propia de cola
│   │   ├── hash/               # Tabla de hash propia
│   │   ├── tree/               # ABB / árbol balanceado
│   │   └── heap/               # Montículo binario / cola de prioridad
│   │
│   ├── dominio/                # Tipos del dominio (Entrega 1 en adelante)
│   │   ├── item/
│   │   ├── cofre/              # Los 5 tipos de cofre
│   │   ├── robopuerto/
│   │   ├── robot/              # Entrega 2
│   │   └── red/                # Composite: Red de robopuertos (Entrega 2)
│   │
│   ├── patrones/               # Patrones de diseño (Entrega 2)
│   │   ├── iterator/
│   │   └── adapter/
│   │
│   ├── algoritmos/
│   │   ├── busqueda/           # Lineal y binaria (Entrega 1)
│   │   ├── ordenamiento/       # Selección/inserción (Entrega 1); quicksort
│   │   │                         /mergesort/heapsort, bucket/radix (Entrega 3)
│   │   ├── greedy/             # Selección de proveedor (Entrega 2)
│   │   ├── backtracking/       # Planificación de rutas con recarga (Entrega 2)
│   │   └── dp/                 # Selección de carga por programación dinámica
│   │                             (Entrega 2)
│   │
│   ├── carga/                  # Lectura de archivos de entrada (Entrega 1) y
│   │                             adapters de formato (Entrega 2)
│   │
│   └── simulacion/             # Ciclo de simulación, estado estable, registro
│                                 de salida (Entrega 3)
│
├── data/
│   ├── ejemplo-basico/         # Escenario mínimo de ejemplo
│   ├── ejemplo-estable/        # Escenario donde el sistema alcanza estado
│   │                             estable
│   └── ejemplo-no-estable/     # Escenario con solicitudes que no pueden
│                                 satisfacerse
│
├── output/                     # Archivos de salida generados por la simulación
│                                 (no versionar los generados en cada corrida)
│
└── docs/
    ├── informe.md              # Copia o enlace al informe (ver template.md)
    └── diagramas/              # Diagramas de arquitectura y de estructuras de
                                  datos
```

## Convenciones de paquetes

- Cada paquete bajo `internal/tad/` expone únicamente su tipo abstracto y sus
  operaciones (por ejemplo, `Push`, `Pop`, `Enqueue`, `Dequeue`, `Get`, `Set`,
  `Insertar`, `Recorrer`); no tiene que conocer nada del dominio logístico.
- Los paquetes bajo `internal/dominio/` dependen de `internal/tad/`, nunca al
  revés.
- Los paquetes bajo `internal/algoritmos/` operan sobre tipos del dominio o
  sobre tipos genéricos, según el caso, y no deben depender de
  `internal/simulacion/` ni de `cmd/`.
- `internal/simulacion/` es el único paquete que conoce y orquesta a todos los
  demás.
- `cmd/simulador/main.go` se limita a leer argumentos/configuración, invocar
  `internal/carga`, ejecutar `internal/simulacion` y escribir la salida; no
  debe contener lógica de negocio.

## Pruebas

Cada paquete de `internal/` tiene que incluir sus propios archivos `_test.go`
junto al código que prueban (convención estándar de Go), por ejemplo:

```txt
internal/tad/pila/
├── pila.go
└── pila_test.go
```

Se recomienda incluir, además, pruebas de integración en
`internal/simulacion/simulacion_test.go` que corran la simulación completa
sobre los escenarios de `data/` y verifiquen el resultado esperado (estado
estable alcanzado o solicitudes reportadas como no satisfechas).

## Correspondencia con las entregas

| Entrega   | Paquetes que se agregan o completan                                                                                                                                                                                                                                               |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Entrega 1 | `internal/tad/lista`, `internal/tad/pila`, `internal/tad/cola`, `internal/tad/hash`, `internal/dominio/item`, `internal/dominio/cofre`, `internal/dominio/robopuerto`, `internal/algoritmos/busqueda`, `internal/algoritmos/ordenamiento` (selección/inserción), `internal/carga` |
| Entrega 2 | `internal/dominio/robot`, `internal/dominio/red`, `internal/patrones/iterator`, `internal/patrones/adapter`, `internal/algoritmos/greedy`, `internal/algoritmos/backtracking`, `internal/algoritmos/dp`, ampliación de `internal/carga`                                           |
| Entrega 3 | `internal/tad/arbol`, `internal/tad/monticulo`, ampliación de `internal/algoritmos/ordenamiento` (recursivos y lineales), `internal/simulacion`                                                                                                                                   |

## Próximos pasos sugeridos

1. Inicializar el módulo Go (`go mod init`) y el repositorio Git.
2. Acordar el formato del archivo de datos de entrada antes de escribir
   `internal/carga`.
3. Configurar un flujo de pruebas (`go test ./...`) que corra en cada commit
   relevante.
