# Informe de cobertura

Fecha: 2026-09-20
Módulo: `prueba-seguimiento` (Go 1.27.1)

## Resultado por paquete

Comando: `go test -coverprofile="cover.out" -covermode=atomic ./...`

```
ok      prueba-seguimiento/internal/project     0.366s  coverage: 100.0% of statements
```

Único paquete implementado hasta HU-02: `internal/project`.

## Detalle por función (`go tool cover -func cover.out`)

```
prueba-seguimiento/internal/project/project.go:84:	NewRepository	100.0%
prueba-seguimiento/internal/project/project.go:93:	Create		100.0%
prueba-seguimiento/internal/project/project.go:120:	Update		100.0%
prueba-seguimiento/internal/project/project.go:152:	nameExists	100.0%
prueba-seguimiento/internal/project/project.go:163:	AddMember	100.0%
prueba-seguimiento/internal/project/project.go:196:	ListMembers	100.0%
prueba-seguimiento/internal/project/project.go:205:	canonicalRole	100.0%
prueba-seguimiento/internal/project/project.go:220:	today		100.0%
total:							(statements)	100.0%
```

## Interpretación

- Cobertura de sentencias: **100%** en `internal/project`.
- Todas las reglas de negocio de HU-01 (RN-01..RN-05) y de HU-02
  (RN-06..RN-10) tienen al menos un test que las ejercita.
- HU-02 queda cubierta en sus casos normales, alternativos, límites y de
  error: alta de integrantes con nombre y rol, validación de roles contra
  los 3 roles válidos con cualquier capitalización y forma canónica,
  recorte de espacios, duplicados por nombre dentro del mismo proyecto,
  proyectos distintos con el mismo integrante, proyecto inexistente y
  consulta de integrantes.
- Este informe se irá actualizando a medida que se agreguen paquetes en
  las próximas historias.

## Evolución

| Historia | Paquete | Cobertura |
| --- | --- | --- |
| HU-01 | internal/project | 100.0% |
| HU-02 | internal/project | 100.0% |
