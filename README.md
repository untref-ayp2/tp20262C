# Trabajo Práctico Grupal — Red Logística Automatizada

## Algoritmos y Programación II — Segundo Cuatrimestre 2026

## 1. Objetivo

Diseñar e implementar, en Go, una simulación de una red logística
automatizada, aplicando estructuras de datos abstractas (TAD) y técnicas de
diseño de algoritmos de la cursada. El trabajo tiene **dos hitos** sobre un
mismo repositorio: un **checkpoint formativo** y la **entrega final**.

## 2. Modalidad de trabajo

- Grupos de 3 a 4 estudiantes.
- Cada grupo tiene asignado un docente o ayudante que guía y evalúa el
  desarrollo.
- Todos los integrantes tienen que registrar commits propios e identificables
  a lo largo del trabajo. La ausencia de participación registrada en el
  historial puede afectar la nota individual, independientemente de la nota
  grupal.
- Se usa Git desde el inicio: el historial debe reflejar el avance real.
  No se aceptan commits únicos de "entrega final".
- Se parte del repo template (asignación única en classroom50) y se trabaja
  sobre el repo clonado.

## 3. Fechas clave

| Hito                 | Fecha            | Referencia en el calendario                    |
| -------------------- | ---------------- | ---------------------------------------------- |
| Asignación del TP    | Clase 14         | 01/10 al 05/10 (según comisión)                |
| Checkpoint (Entrega 1) | Domingo 25/10/2026 | Antes de la Clase 21 (montículos)           |
| Segundo Parcial      | 10/11 (Ma-Ju) / 11/11 (Lu-Mi) | Clase 25                     |
| Entrega final        | Domingo 15/11/2026 | Después del 2º parcial, antes de las orales  |
| Presentaciones orales | 19/11 (Ma-Ju) / 20/11 (Lu-Mi) | Clase 28                   |

El **checkpoint** no se re-entrega: se evalúa el repositorio a esa fecha
(commits y estado del código) y se devuelve una corrección formativa. La nota
definitiva se obtiene en la corrección de la entrega final, de acuerdo con
`rubrica.md`.

## 4. Descripción del sistema

En una red logística automatizada, los ítems se transportan entre cofres
inteligentes mediante robots logísticos que operan dentro del alcance de
robopuertos.

### 4.1. Ítems

Cada ítem tiene un nombre único (por ejemplo, `hierro`, `circuito`, `motor`) y
existe en cantidades enteras dentro de los cofres.

### 4.2. Cofres

Cada cofre cumple un único rol:

- **Provisión:** ofrece ítems para abastecer otros cofres.
- **Solicitud:** solicita una cantidad de un ítem. La solicitud se considera
  satisfecha al alcanzar o superar esa cantidad.
- **Almacenamiento:** recibe ítems excedentes sin destino asignado; se usa como
  último recurso.

### 4.3. Robopuertos

Un robopuerto tiene una coordenada fija $(x, y)$ y un alcance. Los cofres
dentro de esa distancia están cubiertos por él. Dos robopuertos se **conectan**
si sus zonas de cobertura se superponen; agrupando robopuertos conectados se
obtienen las **redes conexas** de la red logística. Un cofre fuera de toda
cobertura es **inaccesible** y queda fuera de la simulación. Los robopuertos
también funcionan como estaciones de recarga.

### 4.4. Robots logísticos y energía

Cada robot tiene una capacidad de carga (cantidad de ítems por viaje) y una
batería máxima medida en una unidad abstracta llamada `Célula`. Cada
movimiento consume Células en proporción a la distancia euclídea recorrida:

```
célulasNecesarias = ceil(distanciaEuclídea(origen, destino) × factorDeConsumo)
```

El factor de consumo es configurable. Un robot puede viajar entre cofres de
una misma red conexa; si la batería no alcanza para el tramo, **recarga en su
robopuerto** (un ciclo de carga) y reintenta en el siguiente ciclo. Como el
objetivo es simular flujos de ítems y no la navegación física, no se planifican
rutas con paradas intermedias: la recarga siempre ocurre en el robopuerto de
origen y queda registrada como evento.

### 4.5. Objetivo de la simulación

Partiendo de un estado inicial con solicitudes sin satisfacer y proveedores con
ítems disponibles, el sistema ejecuta ciclos de transferencia hasta alcanzar un
**estado estable** (todas las solicitudes satisfechas y sin ítems en tránsito)
o, si eso no es posible (falta de stock, cofres inaccesibles, batería
insuficiente sin recarga viable), reporta con claridad qué solicitudes quedaron
sin satisfacer y por qué.

## 5. Glosario

- **Robot logístico:** unidad móvil que transporta ítems entre cofres, con
  batería limitada (Células) y capacidad de carga limitada.
- **Robopuerto:** infraestructura fija que define una zona de cobertura y
  funciona como estación de recarga.
- **Cofre logístico:** contenedor de ítems con un rol único (ver 4.2).
- **Zona de cobertura:** área circular definida por un robopuerto; puede haber
  zonas superpuestas.
- **Red conexa:** conjunto de robopuertos conectados directa o indirectamente
  por zonas superpuestas.
- **Célula:** unidad abstracta de energía de la batería de un robot.
- **Ciclo de simulación:** iteración en la que se evalúan solicitudes
  pendientes, se seleccionan proveedores y se ejecutan los movimientos
  posibles.
- **Estado estable:** situación en la que todas las solicitudes están
  satisfechas y no hay ítems en tránsito.

## 6. Carga de datos desde archivo

La configuración inicial se carga desde un archivo **JSON** cuyo formato está
documentado en `data/README.md` y en el informe. El parser viene resuelto en
`internal/carga`; el formato mínimo incluye:

- **La red:** robopuertos (posición y alcance) y cofres (posición, rol,
  inventario inicial y solicitudes).
- **Los robots:** robopuerto de origen, batería máxima y capacidad de carga.
- **El factor de consumo** de Células por unidad de distancia.

El archivo debe permitir reproducir cualquier escenario de prueba. Los tres
escenarios de ejemplo están en `data/basico`, `data/estable` y
`data/no-estable`.

## 7. Desarrollo por hitos

### 7.1. Checkpoint — Entrega 1: Modelo y estructuras lineales

**Temas involucrados:** TAD pilas y colas; TAD tablas de hash y diccionarios;
búsqueda lineal y binaria; ordenamiento por selección e inserción;
recursividad. Todos fueron vistos antes del primer parcial: esta entrega
refuerza el bloque de TAD lineales y búsquedas.

**Alcance:** se construye el modelo estático de la red logística, sin
movimiento de ítems.

Requisitos:

1. Completar el modelado del dominio (`internal/dominio`): cofres con sus tres
   roles, robopuertos y robots, tal como vienen esqueletizados en el template.
2. Implementar el TAD **tabla de hash** propio (`internal/tad/hash`) con su
   función de dispersión y su resolución de colisiones, y usarlo para indexar
   cofres por identificador (`red.CofrePorID`). El `map` de Go puede usarse
   como auxiliar interno, pero el TAD debe ser propio.
3. Implementar el TAD **cola** (`internal/tad/cola`) y justificar en el informe
   su aplicación (por ejemplo, solicitudes atendidas en orden de llegada y
   desempate FIFO en la simulación). Es válido, además, modelar el inventario
   de un cofre con una lista propia si el grupo lo justifica.
4. Implementar **búsqueda binaria** sobre el catálogo de ítems ordenado y usar
   **búsqueda lineal** donde el criterio no admita orden total.
5. Implementar el **ordenamiento por selección o inserción** (a elección,
   justificada) para ordenar el catálogo de ítems por nombre.
6. Implementar la **recursividad** de agrupamiento: dado el conjunto de
   robopuertos, determinar las redes conexas (dos robopuertos se conectan si
   sus zonas se superponen o si existe una cadena que los une) y construir la
   red (`red.ConstruirRed`) asignando cofres por cobertura; reportar los
   cofres inaccesibles.
7. _(Opcional, división y conquista)_ Implementar una función que determine el
   proveedor más cercano a un cofre mediante una estrategia de división y
   conquista sobre coordenadas, y comparar su complejidad contra la búsqueda
   lineal equivalente.

**Entregable:** código de las estructuras y consultas, con pruebas
automatizadas para cada TAD y para las búsquedas y el agrupamiento en redes;
el escenario `data/basico` cargado y verificado; un avance del informe con las
secciones correspondientes a esta etapa.

**Hito:** domingo 25 de octubre de 2026 (evaluación formativa del repo).

### 7.2. Entrega final: Robots, planificación y simulación

**Temas involucrados:** patrones de diseño Iterator y Composite; algoritmos
ávidos; montículos binarios y colas de prioridad.

**Alcance:** se incorporan los robots y la energía, se resuelve la selección de
proveedores con el algoritmo ávido y se cierra el sistema con el ciclo de
simulación completo.

Requisitos:

1. Completar el modelo de **robot** (batería, capacidad de carga, consumo y
   recarga), que viene esqueleto en `internal/dominio/robot`.
2. Implementar el **algoritmo ávido** de selección de proveedor
   (`internal/algoritmos/greedy`): para una solicitud, elige el proveedor con
   stock del ítem más conveniente según la regla (a definir y justificar en el
   informe): primero rol Provisión y, dentro del rol, la menor distancia al
   cofre solicitante.
3. Aplicar el patrón **Iterator** (`internal/patrones/iterator`) para recorrer
   de manera uniforme el inventario de un cofre y sus solicitudes pendientes,
   sin exponer la estructura interna.
4. Aplicar el patrón **Composite** (`internal/patrones/composite` +
   `internal/dominio/red`): listar cofres alcanzables y calcular el total
   disponible de un ítem de manera uniforme sobre un robopuerto o sobre la red
   completa.
5. Implementar el TAD **montículo binario / cola de prioridad**
   (`internal/tad/heap`) y usarlo en la simulación para decidir, en cada
   ciclo, qué solicitud atender primero (criterio a elección y justificado:
   cantidad faltante, cercanía al proveedor más conveniente, etc.).
6. Implementar el **ciclo de simulación** (`internal/simulacion`): construir
   la red, ejecutar ciclos de transferencia con los robots (viajes que
   respetan la capacidad de carga, recargas registradas como eventos, excedentes
   a cofres de almacenamiento) hasta alcanzar el estado estable o determinar
   que no es posible, reportando cada solicitud pendiente con su causa en un
   archivo de salida (`output/simulacion.txt`): movimientos, distancias,
   decisiones y recargas de cada ciclo.
7. Validar el sistema contra los tres escenarios de `data/` y generar al menos
   dos escenarios propios (uno de éxito y uno de falla controlada).

**Entregable (final):** sistema completo e integrado desde el archivo de
entrada hasta el archivo de salida; pruebas automatizadas de todos los
módulos; escenarios de ejemplo y propios; informe final según `template.md`.

**Fecha de entrega:** domingo 15 de noviembre de 2026.

## 8. Requisitos transversales

- Implementación en Go, organizada en paquetes según `estructura-proyecto.md`.
- Los TAD vistos en la cursada (cola, tabla de hash, montículo) deben
  implementarse como tipos propios del proyecto y no reemplazarse por
  `container/list`, `container/heap` ni bibliotecas externas. Sí pueden usarse
  slices y `map` de Go como bloques de construcción internos de esas
  implementaciones (por ejemplo, el inventario de un cofre).
- Cada algoritmo no trivial (búsquedas, ordenamientos, greedy) debe tener
  pruebas automatizadas (`_test.go`) y su complejidad temporal y espacial
  documentada en el informe.
- El sistema tiene que mostrar por log/archivo los eventos relevantes de cada
  ciclo (por ejemplo: `R1 recarga en P1`, `Solicitud de C2 satisfecha`,
  `C3 inaccesible`).
- Evitar variables globales mutables; usar errores de Go para condiciones
  esperables del dominio (sin stock, cofre inaccesible, batería insuficiente).

## 9. Formato de entrega

- Repositorio de GitHub con historial de commits de todos los integrantes.
- Código organizado según `estructura-proyecto.md`.
- Informe en Markdown, según `template.md`, dentro del repositorio.
- Archivos de entrada (`data/`) y de salida generados (`output/`).
- La entrega final evalúa el sistema completo; el checkpoint de la Entrega 1
  se corrige de forma formativa sobre el mismo repositorio.

## 10. Presentación oral

- Duración estimada: 15 a 20 minutos por grupo.
- Todos los integrantes participan y deben poder responder preguntas sobre
  cualquier parte del sistema, no solo sobre su rol asignado.
- Para acceder a la presentación oral, el trabajo debe estar aprobado por el
  docente a cargo.
- Fechas previstas: comisión Lunes-Miércoles, 20 de noviembre de 2026;
  comisión Martes-Jueves (y CUDI), 19 de noviembre de 2026.

## 11. Evaluación

La evaluación se rige por la rúbrica publicada en `rubrica.md` (100 puntos):
25 para el checkpoint de la Entrega 1, 50 para la entrega final y 25 para los
aspectos generales (código, pruebas, informe y presentación oral).

## 12. Observaciones

- Ante dudas durante el desarrollo, consultar con el docente/ayudante a cargo
  y revisar los apuntes de la materia.
- El diseño del ciclo de simulación (orden de atención, criterio de prioridad,
  manejo de viajes múltiples) es libre, pero debe quedar documentado en el
  informe.
- El template incluye esqueletos compilables (`go build ./...` y
  `go test ./...` pasan al clonar): lo didáctico está por completar y lo que
  ya viene resuelto (carga JSON, tipos de dominio) debe reutilizarse tal cual.