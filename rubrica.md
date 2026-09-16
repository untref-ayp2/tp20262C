# Rúbrica del Trabajo Práctico Grupal — Red Logística Automatizada

Puntaje total: **100 puntos** (25 para el checkpoint de la Entrega 1, 50 para
la entrega final y 25 para los aspectos generales).

Nota de aplicación: el **checkpoint** de la Entrega 1 se corrige sobre el
repositorio a la fecha del hito (domingo 25/10/2026) con devolución formativa;
sus puntos se acreditan en la corrección final, evaluando lo efectivamente
desarrollado. La **entrega final** (domingo 15/11/2026) se evalúa sobre el
sistema completo contra la totalidad de la rúbrica.

## Checkpoint — Entrega 1: Modelo y estructuras lineales (25 puntos)

1. **Modelado del dominio y carga de datos (Informe — 5 puntos):** Descripción
   clara del modelo de cofres (tres roles), robopuertos y robots; carga correcta
   de los escenarios desde el JSON documentado, reutilizando el parser provisto.
2. **Tabla de hash propia (Implementación — 5 puntos):** Implementación
   funcional de la tabla de hash (función de dispersión y resolución de
   colisiones), usada para indexar cofres por identificador.
3. **Cola propia (Implementación — 4 puntos):** Implementación correcta del
   TAD cola, con pruebas; aplicación justificada (solicitudes por orden de
   llegada y desempate FIFO en la simulación).
4. **Búsqueda lineal y binaria (Implementación — 4 puntos):** Catálogo
   ordenado de ítems con búsqueda binaria correcta; uso justificado de
   búsqueda lineal donde corresponda.
5. **Ordenamiento por selección o inserción (Implementación — 3 puntos):**
   Implementación propia y correcta del algoritmo elegido, con justificación de
   la elección.
6. **Recursividad para redes conexas (Implementación — 4 puntos):**
   Agrupamiento recursivo y correcto de robopuertos en redes conexas,
   construcción de la red por cobertura y detección de cofres inaccesibles.

## Entrega final: Robots, planificación y simulación (50 puntos)

1. **Robots y energía (Implementación — 4 puntos):** Modelo de batería,
   capacidad de carga, consumo por distancia y recarga, correcto y bien
   integrado.
2. **Algoritmo ávido de selección de proveedor (Implementación — 8 puntos):**
   Selección correcta del proveedor según la regla definida (rol Provisión y
   menor distancia), con manejo claro del caso sin proveedor viable.
3. **Patrón Iterator (Implementación — 5 puntos):** Recorrido uniforme del
   inventario y de las solicitudes de un cofre sin exponer la estructura
   interna.
4. **Patrón Composite (Implementación — 5 puntos):** Operaciones uniformes
   (listar cofres, total de un ítem) sobre un robopuerto individual y sobre la
   red completa.
5. **Montículo binario / cola de prioridad (Implementación — 10 puntos):**
   Montículo propio correcto, usado para priorizar las solicitudes en cada
   ciclo, con criterio justificado.
6. **Ciclo de simulación y estado estable (Implementación — 12 puntos):**
   Ejecución correcta de los ciclos hasta alcanzar estado estable o hasta
   reportar, con causa, las solicitudes no satisfechas (falta de stock, cofre
   inaccesible, batería insuficiente).
7. **Escenarios y salida verificable (Implementación — 6 puntos):** Los tres
   escenarios de ejemplo y al menos dos propios (éxito y falla controlada)
   funcionan de extremo a extremo; el archivo de salida registra movimientos,
   distancias, recargas y decisiones de cada ciclo.

## Aspectos generales (25 puntos)

1. **Organización y calidad del código (Implementación — 6 puntos):** Código
   modularizado en paquetes según `estructura-proyecto.md`, con buenas
   prácticas de Go (nombres, manejo de errores, uso apropiado de interfaces).
2. **Pruebas automatizadas (Implementación — 7 puntos):** Cobertura de pruebas
   (`_test.go`) sobre los TAD, los algoritmos y la simulación, incluyendo
   casos límite y casos sin solución.
3. **Calidad del informe escrito (General — 6 puntos):** Informe estructurado
   según `template.md`, claro y completo, con análisis de complejidad y
   coherencia entre lo documentado y lo implementado.
4. **Calidad de la presentación oral (General — 6 puntos):** Presentación
   clara y bien organizada; todos los integrantes responden con precisión
   sobre el trabajo completo.

## Distribución de puntos

| Categoría                       | Subcategoría                                 | Puntos |
| ------------------------------- | -------------------------------------------- | ------ |
| Checkpoint (Entrega 1)          | Modelado y carga                              | 5      |
|                                 | Tabla de hash propia                          | 5      |
|                                 | Cola propia                                   | 4      |
|                                 | Búsqueda lineal y binaria                     | 4      |
|                                 | Ordenamiento selección/inserción              | 3      |
|                                 | Recursividad para redes conexas               | 4      |
| Entrega final                   | Robots y energía                              | 4      |
|                                 | Algoritmo ávido de selección de proveedor     | 8      |
|                                 | Patrón Iterator                               | 5      |
|                                 | Patrón Composite                              | 5      |
|                                 | Montículo / cola de prioridad                 | 10     |
|                                 | Ciclo de simulación y estado estable          | 12     |
|                                 | Escenarios y salida verificable               | 6      |
| Aspectos generales              | Organización y calidad del código             | 6      |
|                                 | Pruebas automatizadas                         | 7      |
|                                 | Calidad del informe escrito                   | 6      |
|                                 | Calidad de la presentación oral               | 6      |
| **Total**                       |                                              | **100** |

## Criterios de devolución

- **Checkpoint (Entrega 1):** se corrige contra los criterios de la sección
  "Checkpoint" (25 puntos posibles) más el estado de las pruebas de lo
  desarrollado. El resultado es formativo y no promediable de manera
  independiente; se acredita en la corrección final.
- **Entrega final:** corrección final sobre la totalidad de la rúbrica (100
  puntos), incluyendo la integración del sistema y los aspectos generales.

El incumplimiento del checkpoint (sin avance registrado en el repositorio a la
fecha del hito) se informa al docente a cargo antes de la fecha límite; no
impide corregir la entrega final, pero puede solicitarse la implementación
pendiente como condición para aprobar el trabajo.