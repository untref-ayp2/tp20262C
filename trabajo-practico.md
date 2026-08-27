# Trabajo Práctico Grupal — Red Logística Automatizada

## Algoritmos y Programación II — Segundo Cuatrimestre 2026

## 1. Objetivo

Diseñar e implementar, en Go, un sistema de simulación de una red logística
automatizada, aplicando las estructuras de datos abstractas (TAD) y las técnicas
de diseño de algoritmos (DA) desarrolladas a lo largo de la cursada. El trabajo
se organiza en tres entregas incrementales, cada una asociada a un bloque de
temas del programa, de modo que el sistema construido en una entrega sea la base
sobre la que se apoya la siguiente.

## 2. Modalidad de trabajo

- Grupos de 4 estudiantes.
- Cada grupo tiene asignado un docente o ayudante que guía y evalúa el
  desarrollo.
- Todos los integrantes tienen que registrar commits propios e identificables
  en el repositorio a lo largo de las tres entregas. La ausencia de
  participación registrada en el historial de commits puede afectar la nota
  individual, independientemente de la nota grupal.
- Se tiene que usar Git desde el inicio del trabajo, con un historial de
  commits que refleje el avance real del desarrollo. No se aceptan commits
  únicos de "entrega final".

## 3. Fechas clave

| Hito                  | Fecha                                                    | Referencia en el calendario                                  |
| --------------------- | -------------------------------------------------------- | ------------------------------------------------------------ |
| Entrega 1             | Domingo 27/09/2026                                       | Antes de la Clase 12 (repaso) y la Clase 13 (Primer Parcial) |
| Primer Parcial        | Martes 29/09/2026 (Ma-Ju) / Miércoles 30/09/2026 (Lu-Mi) | Clase 13                                                     |
| Entrega 2             | Domingo 18/10/2026                                       | Antes de la Clase 18 (Árboles binarios)                      |
| Entrega 3 (final)     | Domingo 08/11/2026                                       | Antes de la Clase 25 (Segundo Parcial)                       |
| Segundo Parcial       | Lunes 09/11/2026 (Lu-Mi) / Martes 10/11/2026 (Ma-Ju)     | Clase 25                                                     |
| Presentaciones orales | Miércoles 18/11/2026 (Lu-Mi) / Jueves 19/11/2026 (Ma-Ju) | Clase 28                                                     |

## 4. Descripción del sistema

En una red logística automatizada, los ítems se transportan entre cofres
inteligentes mediante robots logísticos que operan dentro del alcance de
robopuertos.

### 4.1. Ítems

Cada ítem tiene un nombre único (por ejemplo, "hierro", "circuito", "motor") y
existe en cantidades enteras dentro de los cofres.

### 4.2. Cofres

Cada cofre cumple un único rol respecto de un ítem dado. Los tipos de cofre son:

- **Provisión activa:** Ofrece ítems y los empuja proactivamente hacia cofres
  que los solicitan.
- **Provisión pasiva:** Ofrece ítems disponibles, pero no los ofrece
  activamente; tienen que ser retirados por un robot.
- **Solicitud:** Solicita una cantidad de un ítem. La solicitud se considera
  satisfecha al alcanzar o superar esa cantidad. Se abastece, en orden de
  prioridad, desde cofres de _Provisión activa_, luego _Búfer_, luego
  _Provisión pasiva_.
- **Intermedio / Búfer:** Solicita ítems igual que un cofre de _Solicitud_, pero
  también puede actuar como proveedor pasivo de otros cofres. Una solicitud
  imposible de satisfacer en un cofre _Búfer_ no impide que el sistema alcance
  el estado estable.
- **Almacenamiento:** Recibe ítems excedentes sin destino asignado; se usa como
  último recurso.

Un mismo cofre no puede tener roles distintos para ítems diferentes: por cofre
y por ítem, el rol es único. Por ejemplo, un cofre puede solicitar "circuito"
(20 unidades), u ofrecer "hierro" (50 unidades), o almacenar cualquier otro ítem
restante, pero no combinar roles sobre ítems distintos.

### 4.3. Robopuertos

Un robopuerto tiene una coordenada fija $(x, y)$ y una distancia máxima de
alcance; los cofres dentro de esa distancia están cubiertos por él. Los robots
sólo pueden moverse entre cofres cubiertos por al menos un mismo robopuerto, o
entre robopuertos cuyas zonas de cobertura formen una red conexa. Un cofre fuera
de toda cobertura es inaccesible y queda fuera de la simulación. Los robopuertos
también funcionan como estaciones de recarga.

### 4.4. Robots logísticos y energía

Cada robot tiene una capacidad de carga (cantidad de ítems por viaje,
configurable globalmente) y una batería máxima medida en una unidad abstracta
llamada `Célula`. Cada movimiento consume Células en proporción a la distancia
euclídea recorrida:

```
célulasNecesarias = distanciaEuclídea(cofreOrigen, cofreDestino) × factorDeConsumo
```

El factor de consumo es configurable (por ejemplo, 1 Célula por unidad de
distancia). Un robot puede recargar su batería completamente al pasar por un
robopuerto (requiere un ciclo de carga). Antes de cada entrega, el robot tiene
que verificar si su batería alcanza para el trayecto; si no alcanza, tiene que
planificar una ruta alternativa con paradas de recarga en robopuertos
intermedios. Si no existe una ruta viable, la entrega se cancela o pospone.

### 4.5. Objetivo de la simulación

Partiendo de un estado inicial con solicitudes sin satisfacer y proveedores con
ítems disponibles, el sistema debe ejecutar ciclos de transferencia hasta
alcanzar un **estado estable** (todas las solicitudes satisfechas y sin ítems en
tránsito) o, si eso no es posible (por falta de ítems, cobertura insuficiente,
rutas inviables por batería, etc.), reportar con claridad qué solicitudes
quedaron sin satisfacer y por qué.

## 5. Glosario

- **Robot logístico:** Unidad móvil que transporta ítems entre cofres, con
  batería limitada (Células) y capacidad de carga limitada. Parte de un
  robopuerto y planifica sus rutas considerando cobertura, energía y
  prioridades.
- **Robopuerto:** Infraestructura fija que define una zona de cobertura y
  funciona como estación de recarga.
- **Cofre logístico:** Contenedor inteligente de ítems, con un rol específico
  (ver 4.2).
- **Zona de cobertura:** Área definida por un robopuerto, dentro de la cual sus
  cofres pueden enviar y recibir ítems. Puede haber zonas superpuestas.
- **Red logística:** Conjunto de robopuertos cuyas zonas de cobertura están
  directa o indirectamente conectadas.
- **Célula:** Unidad abstracta de energía de la batería de un robot.
- **Ciclo de simulación:** Iteración en la que el sistema evalúa solicitudes
  pendientes, selecciona proveedores y ejecuta los movimientos posibles.
- **Estado estable:** Situación en la que todas las solicitudes están
  satisfechas y no hay ítems en tránsito.

## 6. Carga de datos desde archivo

La configuración inicial de la simulación debe cargarse desde archivo. El
formato queda a criterio de cada grupo (JSON, CSV, YAML, texto estructurado,
etc.), pero tiene que estar documentado en el informe y contener, como mínimo:

- **La red logística:** Ubicación y alcance de cada robopuerto; posición y tipo
  de cada cofre.
- **El contenido inicial de cada cofre:** Ítems ofrecidos, ítems solicitados y
  cantidades, ítems almacenados.
- **Los robots:** Robopuerto de origen, capacidad de batería y capacidad de
  carga (si varían entre robots).

El diseño del archivo es libre, pero debe permitir reproducir fácilmente
cualquier escenario de prueba.

## 7. Desarrollo incremental

### 7.1. Entrega 1: Modelo de datos y estructuras lineales

**Temas de la cursada involucrados:** Búsqueda lineal y binaria; ordenamiento
por selección e inserción; TAD pilas y colas; TAD listas simples, dobles y
circulares; TAD mapas de bits y tablas de hash; TAD conjuntos y diccionarios;
recursividad; división y conquista.

**Alcance:** En esta entrega se construye el modelo estático de la red
logística, sin robots ni movimiento de ítems. El objetivo es representar
correctamente el dominio y responder consultas sobre él.

Requisitos:

1. Definir los tipos (structs) e interfaces necesarias para representar Ítems,
   Cofres (con sus cinco variantes) y Robopuertos, según la descripción del
   dominio (sección 4).
2. Implementar la carga de la configuración inicial desde archivo (sección 6).
3. Indexar los cofres por identificador y los ítems por nombre utilizando una
   tabla de hash propia (TAD tabla de hash), evitando delegar la búsqueda en el
   `map` nativo de Go como única solución: el `map` puede usarse como estructura
   auxiliar interna, pero la tabla de hash debe ser un TAD del proyecto, con su
   propia función de hashing y su propia resolución de colisiones.
4. Representar el inventario de cada cofre con una lista propia (simple, doble
   o circular, a elección justificada del grupo) y las solicitudes pendientes
   de cada cofre con una pila o una cola propia, según corresponda a la
   semántica del caso (por ejemplo, las solicitudes se atienden en orden de
   llegada (cola), el historial de movimientos de un cofre puede modelarse como
   una pila, para auditar en orden inverso).
5. Determinar, para cada robopuerto, qué cofres están dentro de su zona de
   cobertura (por distancia euclídea) y, aplicando recursividad, agrupar los
   robopuertos en redes conexas (dos robopuertos pertenecen a la misma red si
   sus zonas se superponen o si existe una cadena de robopuertos que los
   conecta). Un cofre fuera de toda cobertura debe reportarse como inaccesible.
6. Implementar un catálogo ordenado de ítems (por nombre) y resolver sobre él
   las búsquedas por nombre exacto usando búsqueda binaria, usar búsqueda
   lineal donde el criterio de búsqueda no admita orden total (por ejemplo,
   búsqueda de cofres por tipo).
7. Ordenar, para cada robopuerto, las solicitudes pendientes por prioridad y
   después por distancia al proveedor más cercano, utilizando un algoritmo de
   ordenamiento por selección o por inserción implementado por el grupo (no
   `sort.Slice`). Justificar la elección entre selección e inserción según el
   escenario.
8. _(Opcional, división y conquista)_ Implementar una función que, dado un
   conjunto de cofres proveedores, determine el más cercano a un cofre de
   solicitud mediante una estrategia de división y conquista sobre las
   coordenadas, y comparar su complejidad contra la búsqueda lineal equivalente.

**Entregable:** Código fuente del modelo y las estructuras descriptas, con
pruebas automatizadas para cada estructura (pilas, colas, listas, tabla de
hash) y para las búsquedas y el agrupamiento en redes; un archivo de datos de
ejemplo; un informe parcial breve (a integrar después en el informe final) que
documente las estructuras elegidas y su justificación.

**Fecha de entrega:** domingo 27 de septiembre de 2026.

### 7.2. Entrega 2: Robots, planificación y patrones de diseño

**Temas de la cursada involucrados:** Patrones de diseño Iterator, Adapter y
Composite; algoritmos ávidos; backtracking / fuerza bruta; programación
dinámica.

**Alcance:** Se incorporan los robots logísticos, su energía y su capacidad de
carga, y se resuelve cómo seleccionar proveedores, planificar rutas con recarga
y decidir qué ítems transportar en cada viaje.

Requisitos:

1. Definir el tipo Robot logístico, con batería máxima (en Células), batería
   actual, capacidad de carga y robopuerto de origen. Implementar el cálculo de
   consumo de energía por movimiento (fórmula de la sección 4.4, con factor de
   consumo configurable).
2. Modelar la red logística (robopuertos, sus zonas de cobertura y los cofres
   que contienen) con el patrón **Composite**, de manera que operaciones como
   "listar todos los cofres alcanzables" o "calcular el total disponible de un
   ítem" puedan aplicarse de manera uniforme sobre un robopuerto individual o
   sobre la red completa.
3. Definir una interfaz **Iterator** (o utilizar iteradores nativos de Go
   mediante `range-over-func` / `iter.Seq`) que permita recorrer de manera
   uniforme el inventario de un cofre y las solicitudes pendientes,
   independientemente de si están implementados como lista simple, doble o
   circular.
4. Implementar un **Adapter** que permita incorporar un segundo formato de
   archivo de entrada, distinto del usado en la Entrega 1 (por ejemplo, JSON si
   en la Entrega 1 se usó CSV), sin modificar la lógica interna de carga ya
   construida.
5. Implementar, con un algoritmo **ávido (greedy)**, la selección del cofre
   proveedor para una solicitud, respetando el orden de prioridad (Provisión
   activa > Búfer > Provisión pasiva) y, dentro de cada nivel de prioridad, la
   menor distancia al cofre solicitante.
6. Implementar, con **backtracking**, la planificación de una ruta desde el
   robopuerto de origen de un robot hasta el cofre destino, insertando paradas
   de recarga en robopuertos intermedios cuando la batería no alcance para el
   tramo directo. Si no existe ninguna ruta viable con recargas, la entrega
   debe cancelarse o posponerse, y esto debe quedar registrado.
7. Implementar, con **programación dinámica**, la selección de qué solicitudes
   (o qué unidades de ítems) transportar en un viaje de un robot, maximizando
   la cantidad de unidades entregadas sin superar la capacidad de carga del
   robot (problema de la mochila, en su variante 0/1 o fraccionaria, según se
   justifique).

**Entregable:** Código de las secciones anteriores integrado con el de la
Entrega 1; pruebas automatizadas para el algoritmo ávido, el backtracking y la
programación dinámica, incluyendo al menos un caso sin solución viable para
cada uno; informe parcial con la justificación de cada patrón y algoritmo,
incluyendo su complejidad.

**Fecha de entrega:** domingo 18 de octubre de 2026.

### 7.3. Entrega 3: Estructuras jerárquicas, prioridad y cierre del sistema

**Temas de la cursada involucrados:** Árboles binarios y recorridos; árboles
binarios de búsqueda; árboles balanceados e iteradores; montículos binarios y
colas de prioridad; ordenamientos recursivos (quicksort, mergesort, heapsort);
ordenamientos en tiempo lineal (urnas/bucket sort y radix sort).

**Alcance:** Se completa el sistema con las estructuras jerárquicas y de
prioridad, se reemplazan los algoritmos provisorios de la Entrega 1 por
versiones más eficientes, y se ejecuta la simulación completa hasta alcanzar un
estado estable o reportar los pedidos no satisfechos.

Requisitos:

1. Indexar los robopuertos (o los cofres, a elección del grupo) en un **árbol
   binario de búsqueda** propio, ordenado por coordenada, que permita resolver
   de forma eficiente consultas como "cofres dentro de un radio de una
   coordenada dada", reemplazando el recorrido exhaustivo usado en la Entrega 1.
   Documentar si el árbol se balancea y con qué estrategia.
2. Implementar un **iterador in-order** sobre el árbol anterior, reutilizando
   la interfaz Iterator definida en la Entrega 2.
3. Implementar una **cola de prioridad** basada en un **montículo binario**
   propio, y utilizarla para decidir, en cada ciclo de simulación, qué solicitud
   atender primero (por urgencia, cantidad faltante o cercanía al proveedor más
   conveniente).
4. Reemplazar el ordenamiento por selección/inserción de la Entrega 1 por un
   **ordenamiento recursivo** (quicksort, mergesort o heapsort, a elección)
   para ordenar el registro de movimientos o el catálogo de ítems, y comparar
   empíricamente su desempeño contra la versión anterior sobre un mismo
   conjunto de datos de tamaño creciente.
5. Identificar al menos un caso dentro del sistema donde las claves a ordenar
   sean enteros acotados (por ejemplo, distancias discretizadas o cantidades de
   ítems) e implementar ahí un **ordenamiento en tiempo lineal** (bucketsort o
   radixsort), justificando por qué es aplicable.
6. Integrar todo lo anterior en un ciclo de simulación que, a partir del estado
   inicial cargado, ejecute ciclos de transferencia hasta que el sistema alcance
   un estado estable (todas las solicitudes satisfechas y sin ítems en
   tránsito) o hasta determinar que no es posible, informando claramente qué
   solicitudes quedaron sin satisfacer y por qué (falta de ítems, cofre
   inaccesible, batería insuficiente sin ruta de recarga viable, etc.).
7. Registrar en un archivo de salida los movimientos realizados, las distancias
   recorridas y las decisiones tomadas en cada ciclo.

**Entregable (final):** Sistema completo e integrado, con pruebas automatizadas
de todos los módulos; al menos dos escenarios de datos de ejemplo (uno donde el
sistema alcanza estado estable y otro donde no lo alcanza); archivo de salida
de la simulación; informe final completo según `template.md`.

**Fecha de entrega:** domingo 8 de noviembre de 2026.

## 8. Requisitos transversales

- Implementación en Go, organizada en paquetes según la estructura de proyecto
  de referencia (ver `estructura-proyecto.md`).
- Las estructuras correspondientes a TAD vistos en la cursada (pilas, colas,
  listas, tablas de hash, árboles, montículos) deben implementarse como tipos
  propios del proyecto; no se aceptan como reemplazo `container/list`,
  `container/heap` ni otras bibliotecas externas de estructuras de datos. Sí
  pueden usarse slices y maps de Go como bloques de construcción internos de
  esas implementaciones.
- Cada algoritmo no trivial (búsquedas, ordenamientos, greedy, backtracking,
  programación dinámica) debe tener pruebas automatizadas (`_test.go`) y su
  complejidad temporal y espacial documentada en el informe.
- El sistema tiene que mostrar por log los eventos relevantes de cada ciclo de
  simulación (por ejemplo: "Robot R1 recarga en Robopuerto P2", "Solicitud del
  Cofre C4 satisfecha", "Cofre C7 inaccesible").
- Se debe evitar el uso de variables globales mutables y el manejo de
  condiciones esperables del dominio mediante errores de Go.

## 9. Formato de entrega

- Repositorio de GitHub con historial de commits de todos los integrantes.
- Código fuente organizado según `estructura-proyecto.md`.
- Informe en Markdown, según `template.md`, dentro del repositorio.
- Archivos de datos de entrada de ejemplo y archivos de salida generados por el
  sistema.
- Cada entrega parcial (Entrega 1 y Entrega 2) se evalúa de forma independiente
  y acumulativa; la Entrega 3 integra y cierra el trabajo.

## 10. Presentación oral

- Duración estimada: 15 a 20 minutos por grupo.
- Todos los integrantes deben participar y ser capaces de responder preguntas
  sobre cualquier parte del sistema, no sólo sobre su rol asignado.
- Para acceder a la presentación oral, el trabajo práctico debe estar aprobado
  por el docente a cargo.
- Fechas previstas: comisión Lunes-Miércoles, 18 de noviembre de 2026; comisión
  Martes-Jueves, 19 de noviembre de 2026.

## 11. Evaluación

La evaluación de este trabajo práctico se rige por la rúbrica publicada en
`rubrica.md`, que pondera cada entrega, la integración final del sistema y los
aspectos generales de código, informe y presentación oral.

## 12. Observaciones

- Ante dudas durante el desarrollo, consultar con el docente/ayudante a cargo.
- El diseño del archivo de datos de entrada es libre, pero debe quedar
  documentado en el informe y permitir reproducir cualquier escenario de prueba.
- Se espera que cada grupo genere al menos dos escenarios de prueba propios,
  además de los provistos como ejemplo, cubriendo tanto el caso de éxito
  (estado estable alcanzable) como el caso de falla controlada.
