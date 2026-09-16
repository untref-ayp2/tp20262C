# Informe del Trabajo Práctico Grupal — Red Logística Automatizada

## Algoritmos y Programación II — Segundo cuatrimestre 2026

**Integrantes del Grupo:**

- [Apellido y nombre del estudiante 1 — Legajo]
- [Apellido y nombre del estudiante 2 — Legajo]
- [Apellido y nombre del estudiante 3 — Legajo]

**Docente/Ayudante a cargo:** [Nombre del docente o ayudante]

**Comisión:** [Lunes-Miércoles / Martes-Jueves / CUDI]

**Fecha de entrega:** [Fecha]

> Este informe se completa de manera incremental. En el checkpoint (Entrega 1)
> se completan las secciones correspondientes al modelo y las estructuras
> lineales; en la Entrega final se integran las secciones de robots, patrones,
> montículo y simulación.

---

## 1. Introducción

- Breve descripción del problema: simulación de una red logística automatizada.
- Objetivos del trabajo práctico.
- Resumen de la solución implementada al momento de esta entrega.
- Organización del informe.

## 2. Diseño del sistema

### 2.1. Arquitectura General

- Diagrama de paquetes del sistema (ver `estructura-proyecto.md` como
  referencia), mostrando los módulos principales y sus dependencias.
- Descripción de las decisiones de diseño y justificación de la arquitectura.

### 2.2. Estructuras de datos

Para cada estructura, indicar en qué hito se incorporó:

- **[Checkpoint]** Tabla de hash y cola propias: descripción, justificación de
  la elección frente a alternativas (por ejemplo, `map` de Go o
  `container/list`), diagrama de buckets o estructura interna si se considera
  relevante.
- **[Entrega final]** Montículo binario: descripción del montículo mínimo
  (mínimo = extrae el de menor prioridad), justificación y diagrama si
  corresponde.

### 2.3. Algoritmos implementados

Para cada algoritmo, indicar en qué hito se incorporó:

- **[Checkpoint]** Búsqueda lineal y binaria; ordenamiento por selección o
  inserción; detección recursiva de redes conexas de robopuertos y
  construcción de la red.
- **[Entrega final]** Selección ávida de proveedor; ciclo de simulación con
  montículo de prioridad, robots, recargas y reporte de solicitudes
  insatisfechas.

Para cada uno: explicación del funcionamiento, pseudocódigo o diagrama de
flujo si es necesario, y referencia a los conceptos teóricos de la cursada
aplicados.

### 2.4. Patrones de diseño aplicados

**[Entrega final]** Descripción de la aplicación de los patrones Iterator y
Composite en el sistema: qué problema de diseño resuelve cada uno en este
dominio y cómo se implementó en Go (para Composite: interfaz Componente en
`internal/patrones/composite`, hojas y compuesto en `internal/dominio/red`).

### 2.5. Gestión de datos de entrada y salida

- Formato del archivo JSON de entrada (robopuertos, cofres, robots, factor de
  consumo), con ejemplo corto.
- Formato del archivo de salida de la simulación: movimientos, distancias,
  recargas y decisiones de cada ciclo.

## 3. Implementación

- Descripción general de la implementación en Go.
- Organización del código en paquetes (correspondencia con
  `estructura-proyecto.md`).
- Explicación de las interfaces definidas y las decisiones de modularización.
- Consideraciones sobre buenas prácticas de programación aplicadas: manejo de
  errores, pruebas, nombres, ausencia de variables globales mutables.

## 4. Análisis de complejidad

Para cada algoritmo y estructura relevante:

- Complejidad temporal (notación Big O), en el mejor, peor y caso promedio
  cuando corresponda.
- Complejidad espacial (notación Big O).
- Justificación del análisis.

## 5. Pruebas y resultados

- Descripción de los escenarios de prueba utilizados:
  `data/basico`, `data/estable`, `data/no-estable`, y al menos dos escenarios
  propios (éxito y falla controlada).
- Metodología de las pruebas automatizadas (`_test.go`) por módulo.
- Resultados obtenidos: tablas o descripciones de cada corrida (cantidad de
  ciclos, solicitudes satisfechas, solicitudes no satisfechas con su causa).
- Análisis de los resultados y su interpretación.

## 6. Conclusiones

- Resumen de los logros del trabajo práctico.
- Evaluación de la eficiencia del sistema implementado.
- Identificación de limitaciones y posibles mejoras.
- Reflexiones sobre la experiencia de trabajo en grupo.

## 7. Bibliografía

- Listado de las fuentes consultadas: apuntes de la cátedra, libros, documentación
  de Go, material de la cursada.

## 8. Apéndice

- Enlace al repositorio con el código fuente completo.
- Datos de prueba detallados (archivos de entrada y salida utilizados).
- Cualquier otra información relevante que complemente el informe.