# Rúbrica del Trabajo Práctico Grupal — Red Logística Automatizada

Puntaje total: **100 puntos** (80 puntos por las tres entregas y la integración
final del sistema, 20 puntos por aspectos generales).

Nota de aplicación: la rúbrica se usa de forma acumulativa. En la corrección de
la Entrega 1 y la Entrega 2 se evalúan únicamente los puntos correspondientes a
esa entrega, con devolución formativa; la nota final integra las tres entregas
más la integración y los aspectos generales, evaluados sobre el sistema
completo.

## Entrega 1: Modelo de datos y estructuras lineales (20 puntos)

1. **Modelado de Ítems, Cofres y Robopuertos (Informe — 2 puntos):**
   Descripción clara y justificada de los tipos e interfaces utilizados para
   representar el dominio, coherente con lo efectivamente implementado.
2. **Carga de datos desde archivo (Implementación — 3 puntos):** Lectura
   correcta de la red logística, el contenido inicial de los cofres y la
   configuración de robots, con formato documentado.
3. **Tabla de hash propia (Implementación — 3 puntos):** Implementación
   funcional de una tabla de hash propia (función de hashing y resolución de
   colisiones) usada para indexar cofres e ítems.
4. **Listas, pilas y colas propias (Implementación — 5 puntos):** Implementación
   correcta de los TAD lista (simple, doble o circular), pila y cola, aplicados
   con una semántica adecuada al inventario y a las solicitudes pendientes de
   cada cofre.
5. **Búsqueda lineal y binaria (Implementación — 3 puntos):** Catálogo ordenado
   de ítems con búsqueda binaria correcta; uso justificado de búsqueda lineal
   donde corresponda.
6. **Ordenamiento por selección o inserción (Implementación — 2 puntos):**
   Implementación propia y correcta, aplicada para ordenar solicitudes por
   prioridad y distancia, con justificación de la elección del algoritmo.
7. **Recursividad para detección de redes (Implementación — 2 puntos):**
   Agrupamiento recursivo y correcto de robopuertos en redes conexas, con
   detección adecuada de cofres inaccesibles.

## Entrega 2: Robots, planificación y patrones de diseño (25 puntos)

1. **Modelado de robots y energía (Informe — 2 puntos):** Descripción clara del
   modelo de batería, capacidad de carga y fórmula de consumo energético.
2. **Patrón Composite (Implementación — 4 puntos):** La red logística se modela
   de forma composite, permitiendo operaciones uniformes sobre robopuertos
   individuales y sobre la red completa.
3. **Patrón Iterator (Implementación — 4 puntos):** Interfaz de iteración
   uniforme sobre inventarios y solicitudes, independiente de la estructura
   lineal subyacente.
4. **Patrón Adapter (Implementación — 3 puntos):** Incorporación de un segundo
   formato de entrada sin modificar la lógica de carga existente.
5. **Algoritmo ávido de selección de proveedor (Implementación — 4 puntos):**
   Selección correcta del proveedor respetando el orden de prioridad y, adentro
   de cada nivel, la menor distancia.
6. **Backtracking para planificación de rutas con recarga (Implementación — 5
   puntos):** Planificación correcta de rutas con paradas de recarga cuando la
   batería no alcanza para el tramo directo, incluyendo el caso sin ruta viable.
7. **Programación dinámica para selección de carga (Implementación — 3
   puntos):** Selección de solicitudes o unidades a transportar que maximiza la
   cantidad entregada respetando la capacidad de carga, con justificación de la
   formulación elegida.

## Entrega 3: Estructuras jerárquicas, prioridad y cierre del sistema (25 puntos)

1. **Árbol de indexación con iterador (Informe — 2 puntos / Implementación — 4
   puntos):** Árbol binario de búsqueda propio, ordenado por coordenada, con
   consultas de cobertura eficientes y un iterador in-order funcional.
2. **Montículo y cola de prioridad (Implementación — 6 puntos):** Montículo
   binario propio utilizado correctamente para priorizar solicitudes en cada
   ciclo de simulación.
3. **Ordenamiento recursivo (Implementación — 4 puntos):** Quicksort, mergesort
   o heapsort propio, aplicado en reemplazo del ordenamiento de la Entrega 1,
   con comparación empírica de desempeño.
4. **Ordenamiento en tiempo lineal (Implementación — 4 puntos):** Bucketsort o
   radixsort propio, aplicado a un caso del dominio con claves enteras acotadas,
   con justificación de su aplicabilidad.
5. **Ciclo de simulación y detección de estado estable (Implementación — 5
   puntos):** Ejecución correcta de ciclos hasta alcanzar estado estable o hasta
   reportar, con causa, las solicitudes no satisfechas.

## Integración final del sistema (10 puntos)

1. **Funcionamiento extremo a extremo (5 puntos):** El sistema, ejecutado desde
   el archivo de entrada hasta el archivo de salida, funciona de forma
   consistente sobre los escenarios de prueba propios y los de ejemplo.
2. **Registro y trazabilidad (3 puntos):** El archivo de salida registra
   movimientos, distancias y decisiones tomadas en cada ciclo, de forma clara y
   verificable.
3. **Reporte de solicitudes no satisfechas (2 puntos):** Cuando el sistema no
   alcanza estado estable, informa con precisión qué solicitudes quedaron
   pendientes y la causa.

## Aspectos generales (20 puntos)

1. **Organización y calidad del código (Implementación — 6 puntos):** Código
   modularizado en paquetes según `estructura-proyecto.md`, siguiendo buenas
   prácticas de Go (nombres, manejo de errores, uso apropiado de interfaces).
2. **Pruebas automatizadas (Implementación — 6 puntos):** Cobertura de pruebas
   (`_test.go`) sobre las estructuras y los algoritmos de las tres entregas,
   incluyendo casos límite y casos sin solución.
3. **Calidad del informe escrito (General — 4 puntos):** Informe estructurado
   según `template.md`, claro, conciso, completo, con lenguaje técnico adecuado
   y coherencia entre las secciones de las tres entregas.
4. **Calidad de la presentación oral (General — 4 puntos):** Presentación clara
   y bien organizada; todos los integrantes demuestran conocimiento del trabajo
   completo y responden con precisión.

## Distribución de puntos

| Categoría          | Subcategoría                                  | Puntos  |
| ------------------ | --------------------------------------------- | ------- |
| Entrega 1          | Modelado de Ítems, Cofres y Robopuertos       | 2       |
|                    | Carga de datos desde archivo                  | 3       |
|                    | Tabla de hash propia                          | 3       |
|                    | Listas, pilas y colas propias                 | 5       |
|                    | Búsqueda lineal y binaria                     | 3       |
|                    | Ordenamiento por selección o inserción        | 2       |
|                    | Recursividad para detección de redes          | 2       |
| Entrega 2          | Modelado de Robots y Energía                  | 2       |
|                    | Patrón Composite                              | 4       |
|                    | Patrón Iterator                               | 4       |
|                    | Patrón Adapter                                | 3       |
|                    | Algoritmo ávido de selección de proveedor     | 4       |
|                    | Backtracking para planificación de rutas      | 5       |
|                    | Programación dinámica para selección de carga | 3       |
| Entrega 3          | Árbol de indexación con iterador              | 6       |
|                    | Montículo y cola de prioridad                 | 6       |
|                    | Ordenamiento recursivo                        | 4       |
|                    | Ordenamiento en tiempo lineal                 | 4       |
|                    | Ciclo de simulación y estado estable          | 5       |
| Integración Final  | Funcionamiento extremo a extremo              | 5       |
|                    | Registro y trazabilidad                       | 3       |
|                    | Reporte de solicitudes no satisfechas         | 2       |
| Aspectos Generales | Organización y calidad del código             | 6       |
|                    | Pruebas automatizadas                         | 6       |
|                    | Calidad del informe escrito                   | 4       |
|                    | Calidad de la presentación oral               | 4       |
| **Total**          |                                               | **100** |

## Criterios de devolución en entregas parciales

- **Entrega 1:** Se corrige contra los criterios de la sección "Entrega 1" (20
  puntos posibles) más el criterio de pruebas automatizadas correspondiente a
  lo desarrollado hasta el momento. El resultado es formativo y no promediable
  de manera independiente; se integra en la nota final.
- **Entrega 2:** Se corrige contra los criterios de la sección "Entrega 2" (25
  puntos posibles), asumiendo como base lo entregado en la Entrega 1. Un
  rediseño posterior de estructuras de la Entrega 1, si está justificado en el
  informe, no penaliza.
- **Entrega 3:** Corrección final sobre la totalidad de la rúbrica (100 puntos),
  incluyendo integración final y aspectos generales.

Una entrega parcial no realizada (Entrega 1 o Entrega 2) se considera
incumplimiento de proceso y debe informarse al docente a cargo antes de la
fecha límite correspondiente; su ausencia no impide corregir la Entrega 3, pero
el docente puede solicitar la implementación faltante como condición para
aprobar el trabajo.
