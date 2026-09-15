# HU-01 — Evidencia TDD (Red → Green → Refactor)

Historia: Como Agile Enabler, quiero crear y modificar un proyecto, para
poder gestionarlo desde el inicio.

Fecha de ejecución: 2026-09-15
Módulo: `prueba-seguimiento` (Go 1.27.1, paquete estándar `testing`)
Paquete: `internal/project`

## Fase RED — tests primero

Los tests se escribieron **antes** del código de implementación. El
repositorio solo contenía el módulo (`go.mod`), el archivo BDD y el
archivo de tests; el paquete `project` todavía no existía.

Comando:

```
go test ./...
```

Salida:

```
# prueba-seguimiento/internal/project [prueba-seguimiento/internal/project.test]
internal\project\project_test.go:14:10: undefined: NewRepository
internal\project\project_test.go:16:24: undefined: CreateInput
internal\project\project_test.go:28:17: undefined: StatusActive
internal\project\project_test.go:29:53: undefined: StatusActive
internal\project\project_test.go:44:10: undefined: NewRepository
internal\project\project_test.go:46:24: undefined: CreateInput
internal\project\project_test.go:54:17: undefined: StatusActive
internal\project\project_test.go:55:53: undefined: StatusActive
internal\project\project_test.go:68:10: undefined: NewRepository
internal\project\project_test.go:70:24: undefined: CreateInput
internal\project\project_test.go:70:24: too many errors
FAIL    prueba-seguimiento/internal/project [build failed]
FAIL
```

Interpretación: los 10 tests fallan porque la funcionalidad nueva aún no
existe (símbolos `undefined`), no por errores preexistentes del
repositorio.

Commit propuesto: `test: agregar tests para HU-01 (RED)`

## Fase GREEN — código mínimo

Se implementó `internal/project/project.go` con el mínimo necesario para
que todos los tests pasen: tipos `Project`, `CreateInput`, `UpdateInput`,
`Repository` (en memoria), funciones `Create` y `Update`, errores de
negocio (`ErrNameRequired`, `ErrDuplicateName`, `ErrProjectNotFound`) y
validaciones RN-01..RN-05.

Durante la fase se corrigieron dos detalles de compilación en los
propios tests (variable sin usar) y se agregaron dos tests que cierran
reglas ya especificadas (nombre vacío al modificar, actualización de
fecha de inicio).

Comando:

```
go test ./...
```

Salida:

```
ok      prueba-seguimiento/internal/project     0.298s
```

Commit propuesto: `feat: implementar HU-01 (GREEN)`

## Fase REFACTOR

No aplica: el código resultante es mínimo y legible; no se identificó una
mejora estructural que valga un cambio aparte.

## Validación final

Comandos y salidas:

```
gofmt -w internal\project\project.go internal\project\project_test.go
→ sin cambios pendientes de formato

go test ./...
ok      prueba-seguimiento/internal/project     0.280s

go test -coverprofile="cover.out" -covermode=atomic ./...
ok      prueba-seguimiento/internal/project     0.353s  coverage: 100.0% of statements
```

Detalle de cobertura (`go tool cover -func=cover.out`):

```
project.go:56   NewRepository   100.0%
project.go:61   Create          100.0%
project.go:88   Update          100.0%
project.go:120  nameExists      100.0%
project.go:131  today           100.0%
total:          (statements)    100.0%
```

## Trazabilidad: escenario BDD → test

| Escenario BDD (features/HU-01-gestion-proyectos.feature) | Test |
| --- | --- |
| Crear un proyecto con datos completos (happy path) | `TestCreate_ProyectoConDatosCompletos` |
| Crear un proyecto solo con nombre (alternativo) | `TestCreate_ProyectoSinDatosOpcionales` |
| Crear un proyecto con nombre con espacios en los extremos (límite) | `TestCreate_NombreConEspaciosSeRecorta` |
| Rechazar un proyecto sin nombre (error) | `TestCreate_NombreVacioDevuelveError` |
| Rechazar un proyecto con nombre duplicado (error) | `TestCreate_NombreDuplicadoDevuelveError` |
| Modificar la descripción de un proyecto existente (happy path) | `TestUpdate_ModificaProyectoExistente` / `TestUpdate_ModificaFechaInicio` |
| Modificar solo el nombre conservando los demás campos (límite) | `TestUpdate_ModificaSoloUnCampo` |
| Modificar un proyecto sin cambiar ningún campo (límite) | `TestUpdate_SinCambiosMantieneProyecto` |
| Modificar un proyecto inexistente (error) | `TestUpdate_IdInexistenteDevuelveError` |
| Rechazar una modificación hacia un nombre duplicado (error) | `TestUpdate_NombreDuplicadoDevuelveError` |

## Resumen de trazabilidad

HU-01: Historia → SDD → Criterios de Aceptación → BDD → Tests → Código Go
