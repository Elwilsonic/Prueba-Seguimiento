# HU-02 — Evidencia TDD (Red → Green → Refactor)

Historia: Como Agile Enabler, quiero registrar los integrantes de un
proyecto, para saber quién forma parte del equipo.

Fecha de ejecución: 2026-09-20
Módulo: `prueba-seguimiento` (Go 1.27.1, paquete estándar `testing`)
Paquete: `internal/project`

## Fase RED — tests primero

Los tests (`internal/project/members_test.go`) se escribieron **antes**
del código de implementación. El working tree contenía además el struct
`Repository` pre-cableado sin commit (mapa `members` por proyecto) que
referenciaba el tipo `Member`, aún inexistente, por lo que el build ya
estaba roto antes de agregar los tests.

Comando:

```
go test ./...
```

Salida:

```
# prueba-seguimiento/internal/project [prueba-seguimiento/internal/project.test]
internal\project\project.go:54:23: undefined: Member
internal\project\project.go:61:30: undefined: Member
internal\project\members_test.go:17:17: repo.AddMember undefined (type *Repository has no field or method AddMember)
internal\project\members_test.go:17:40: undefined: AddMemberInput
internal\project\members_test.go:43:17: repo.AddMember undefined (type *Repository has no field or method AddMember)
internal\project\members_test.go:43:40: undefined: AddMemberInput
internal\project\members_test.go:63:17: repo.AddMember undefined (type *Repository has no field or method AddMember)
internal\project\members_test.go:63:40: undefined: AddMemberInput
internal\project\members_test.go:86:21: repo.AddMember undefined (type *Repository has no field or method AddMember)
internal\project\members_test.go:86:44: undefined: AddMemberInput
internal\project\members_test.go:86:44: too many errors
FAIL    prueba-seguimiento/internal/project [build failed]
FAIL
```

Interpretación: los 12 tests fallan porque la funcionalidad nueva aún no
existe (símbolos `undefined`: `Member`, `AddMemberInput`, `AddMember`,
`ListMembers` y los errores de negocio de integrantes). El fallo
preexistente del working tree (`undefined: Member` en `project.go:54,61`)
también corresponde a la funcionalidad de HU-02, por lo que el RED es
atribuible a esta historia y no a un error ajeno del repositorio.

Commit propuesto: `test: agregar tests para HU-02 (RED)`

## Fase GREEN — código mínimo

Se completó `internal/project/project.go` con el mínimo necesario para
que todos los tests pasen: tipo `Member`, entrada `AddMemberInput`, lista
`rolesValidos`, errores de negocio (`ErrMemberNameRequired`,
`ErrMemberRoleRequired`, `ErrInvalidRole`, `ErrMemberDuplicate`),
funciones `AddMember` y `ListMembers`, contador de IDs por proyecto
(`nextMember`) y validación de rol con comparación insensible a
mayúsculas/minúsculas y forma canónica (`canonicalRole`, RN-06..RN-10).

Comando:

```
go test ./...
```

Salida:

```
ok      prueba-seguimiento/internal/project     0.300s
```

Commit propuesto: `feat: implementar HU-02 (GREEN)`

## Fase REFACTOR

No aplica: el código resultante es mínimo y legible; no se identificó una
mejora estructural que valga un cambio aparte.

## Validación final

Comandos y salidas:

```
gofmt -w internal\project\project.go internal\project\members_test.go
→ sin cambios pendientes de formato

go test ./...
ok      prueba-seguimiento/internal/project     0.506s

go test -coverprofile="cover.out" -covermode=atomic ./...
ok      prueba-seguimiento/internal/project     0.366s  coverage: 100.0% of statements
```

Detalle de cobertura (`go tool cover -func=cover.out`):

```
project.go:84   NewRepository   100.0%
project.go:93   Create          100.0%
project.go:120  Update          100.0%
project.go:152  nameExists      100.0%
project.go:163  AddMember       100.0%
project.go:196  ListMembers     100.0%
project.go:205  canonicalRole   100.0%
project.go:220  today           100.0%
total:          (statements)    100.0%
```

## Trazabilidad: escenario BDD → test

| Escenario BDD (features/HU-02-registro-integrantes.feature) | Test |
| --- | --- |
| Registrar un integrante con nombre y rol en un proyecto existente (happy path) | `TestAddMember_RegistraIntegranteConNombreYRol` |
| Registrar un integrante con el rol en mayúsculas o minúsculas (límite) | `TestAddMember_RolEnMinusculasGuardaCanonico` |
| Registrar un integrante con nombre y rol con espacios en los extremos (límite) | `TestAddMember_NombreYRolConEspaciosSeRecortan` |
| Rechazar un integrante sin rol (error) | `TestAddMember_SinRolDevuelveError` |
| Rechazar un integrante sin nombre (error) | `TestAddMember_NombreVacioDevuelveError` |
| Rechazar un integrante con rol no válido (error) | `TestAddMember_RolNoValidoDevuelveError` |
| Rechazar un integrante duplicado en el mismo proyecto (error) | `TestAddMember_IntegranteDuplicadoDevuelveError` |
| Registrar el mismo integrante en dos proyectos distintos (alternativo) | `TestAddMember_MismoIntegranteEnDistintosProyectos` |
| Listar los integrantes de un proyecto en orden de registro (happy path) | `TestListMembers_ListaEnOrdenDeRegistro` |
| Listar los integrantes de un proyecto sin integrantes (límite) | `TestListMembers_ProyectoSinIntegrantesDevuelveListaVacia` |
| Registrar un integrante en un proyecto inexistente (error) | `TestAddMember_ProyectoInexistenteDevuelveError` |
| Listar los integrantes de un proyecto inexistente (error) | `TestListMembers_ProyectoInexistenteDevuelveError` |

## Resumen de trazabilidad

HU-02: Historia → SDD → Criterios de Aceptación → BDD → Tests → Código Go
