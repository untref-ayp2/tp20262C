# Informe del Trabajo Práctico Grupal - Red Logística Automatizada

## Algoritmos y Programación II - Segundo cuatrimestre 2026

**Integrantes del Grupo:**

- [Apellido y nombre del estudiante 1 — Legajo]
- [Apellido y nombre del estudiante 2 — Legajo]
- [Apellido y nombre del estudiante 3 — Legajo]
- [Apellido y nombre del estudiante 4 — Legajo]

**Docente/Ayudante a cargo:** [Nombre del tutor]

**Comisión:** [Lunes-Miércoles / Martes-Jueves]

**Fecha de entrega:** [Fecha]

> Este informe se completa de manera incremental. En la Entrega 1 y la Entrega
> 2 se completan únicamente las secciones correspondientes a lo desarrollado
> hasta el momento (indicado entre corchetes en cada apartado); en la Entrega 3
> el informe queda completo e integrado.

## 1. Introducción

- Breve descripción del problema a resolver (sistema de red logística
  automatizada).
- Objetivos del trabajo práctico.
- Resumen de la solución implementada al momento de esta entrega.
- Organización del informe.

## 2. Diseño del sistema

### 2.1. Arquitectura General

- Diagrama de paquetes del sistema (ver `estructura-proyecto.md` como
  referencia), mostrando los módulos principales y su interacción.
- Descripción de la arquitectura adoptada y justificación de las decisiones de
  diseño.

### 2.2. Estructuras de datos

Para cada estructura utilizada, indicar en qué entrega se incorporó:

- **[Entrega 1]** Tabla de hash, listas (simple/doble/circular), pilas y colas:
  descripción, justificación de la elección frente a alternativas, y diagrama
  si corresponde.
- **[Entrega 3]** Árbol binario de búsqueda y montículo binario: descripción,
  justificación, y diagrama si corresponde.

### 2.3. Algoritmos implementados

Para cada algoritmo, indicar en qué entrega se incorporó:

- **[Entrega 1]** Búsqueda lineal y binaria; ordenamiento por selección o
  inserción; detección recursiva de redes de robopuertos.
- **[Entrega 2]** Selección ávida de proveedor; backtracking para planificación
  de rutas con recarga; programación dinámica para selección de carga.
- **[Entrega 3]** Ordenamiento recursivo (quicksort/mergesort/heapsort);
  ordenamiento en tiempo lineal (bucket/radix); ciclo de simulación y detección
  de estado estable.

Para cada uno: explicación de su funcionamiento, pseudocódigo o diagrama de
flujo si es necesario, y referencia a los conceptos teóricos de la cursada
aplicados.

### 2.4. Patrones de diseño aplicados

**[Entrega 2]** Descripción de la aplicación de los patrones Composite,
Iterator y Adapter en el sistema: qué problema de diseño resuelve cada uno en
este dominio y cómo se implementó en Go.

### 2.5. Gestión de datos de entrada y salida

- Formato y contenido de los archivos de entrada (configuración de la red,
  cofres, robots).
- **[Entrega 2]** Formato del segundo archivo de entrada incorporado mediante
  el patrón Adapter.
- **[Entrega 3]** Formato del archivo de salida con el registro de movimientos,
  distancias y decisiones de cada ciclo.

## 3. Implementación

- Descripción general de la implementación en Go.
- Organización del código en paquetes (correspondencia con
  `estructura-proyecto.md`).
- Explicación de las interfaces definidas y las decisiones de modularización.
- Consideraciones sobre las buenas prácticas de programación aplicadas (manejo
  de errores, pruebas, nombres).

## 4. Análisis de complejidad

Para cada algoritmo y estructura relevante:

- Complejidad temporal (notación Big O), en el mejor, peor y caso promedio
  cuando corresponda.
- Complejidad espacial (notación Big O).
- Justificación del análisis.
- **[Entrega 3]** Comparación empírica entre el ordenamiento por
  selección/inserción de la Entrega 1 y el ordenamiento recursivo que lo
  reemplaza, sobre conjuntos de datos de tamaño creciente.

## 5. Pruebas y resultados

- Descripción de los escenarios de prueba utilizados, incluyendo al menos un
  escenario donde el sistema alcanza estado estable y uno donde no lo alcanza.
- Metodología de las pruebas automatizadas (`_test.go`) por módulo.
- Resultados obtenidos: tablas o gráficos con métricas relevantes (por ejemplo,
  cantidad de ciclos hasta el estado estable, cantidad de solicitudes no
  satisfechas, tiempo de ejecución para distintos tamaños de red).
- Análisis de los resultados y su interpretación en relación con la eficiencia
  del sistema.

## 6. Conclusiones

- Resumen de los logros del trabajo práctico.
- Evaluación de la eficiencia del sistema implementado.
- Identificación de limitaciones y posibles cuellos de botella.
- Propuestas de mejoras o extensiones futuras.
- Reflexiones sobre la experiencia de trabajo en grupo y los aprendizajes
  obtenidos.

## 7. Bibliografía

- Listado de las fuentes consultadas (apuntes de la cátedra, libros, artículos,
  documentación de Go).

## 8. Apéndice

- Enlace al repositorio con el código fuente completo.
- Datos de prueba detallados (archivos de entrada y salida utilizados).
- Cualquier otra información relevante que complemente el informe.
