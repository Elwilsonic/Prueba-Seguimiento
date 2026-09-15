# Informe de cobertura

Fecha: 2026-09-15
Módulo: `prueba-seguimiento` (Go 1.27.1)

## Resultado por paquete

Comando: `go test -coverprofile="cover.out" -covermode=atomic ./...`

```
ok      prueba-seguimiento/internal/project     0.353s  coverage: 100.0% of statements
```

Único paquete implementado hasta HU-01: `internal/project`.

## Detalle por función (`go tool cover -func=cover.out`)

```
prueba-seguimiento/internal/project/project.go:56:   NewRepository   100.0%
prueba-seguimiento/internal/project/project.go:61:   Create          100.0%
prueba-seguimiento/internal/project/project.go:88:   Update          100.0%
prueba-seguimiento/internal/project/project.go:120:  nameExists      100.0%
prueba-seguimiento/internal/project/project.go:131:  today           100.0%
total:                                               (statements)    100.0%
```

## Interpretación

- Cobertura de sentencias: **100%** en `internal/project`.
- Todas las reglas de negocio de HU-01 (RN-01..RN-05) tienen al menos un
  test que las ejercita, tanto en creación como en modificación.
- Este informe se irá actualizando a medida que se agreguen paquetes en
  las próximas historias.

## Evolución

| Historia | Paquete | Cobertura |
| --- | --- | --- |
| HU-01 | internal/project | 100.0% |
