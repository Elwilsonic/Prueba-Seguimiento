# HU-01 — Gestión de proyectos: crear y modificar

## Historia de usuario

Como Agile Enabler, quiero crear y modificar un proyecto, para poder
gestionarlo desde el inicio.

## Especificación SDD

### Objetivo

Permitir al rol Agile Enabler dar de alta un nuevo proyecto en el sistema y
modificar sus datos después de creado, de modo que las demás
funcionalidades (Product Backlog, Sprints, métricas) operen siempre sobre
un proyecto existente y con datos correctos.

### Entradas

- **Crear proyecto:** nombre (obligatorio), descripción (opcional),
  fecha de inicio (opcional; por defecto la fecha actual).
- **Modificar proyecto:** ID del proyecto existente más los campos a
  actualizar (mismos campos que la creación: nombre, descripción,
  fecha de inicio).

### Salidas esperadas

- Proyecto creado con ID único, datos guardados y estado inicial
  "Activo".
- Proyecto modificado con los nuevos valores reflejados al consultarlo.
- Si la operación no es válida, un error descriptivo en lugar de un
  proyecto.

### Reglas de negocio

- **RN-01:** el nombre es obligatorio; no puede quedar vacío ni contener
  solo espacios.
- **RN-02:** el nombre debe ser único entre los proyectos existentes.
- **RN-03:** solo se pueden modificar proyectos existentes,
  identificados por su ID.
- **RN-04:** un proyecto recién creado nace en estado "Activo".
- **RN-05:** los espacios al inicio y fin del nombre se recortan antes de
  validar y guardar.

> Nota: los campos de un proyecto (nombre, descripción, fecha de inicio,
> estado) y las reglas RN-01..RN-05 son supuestos de especificación
> pendientes de confirmación; ajustar si el dominio real define otros.

### Restricciones

- Código en Go, con el paquete estándar `testing` (sin librerías
  externas).
- Almacenamiento en memoria (repositorio en memoria); no se pide
  persistencia en disco ni interfaz de usuario en esta historia.
- El control de permisos por rol (quién puede crear/modificar) no forma
  parte de esta historia.

### Casos límite

- Nombre con solo espacios → rechazado (RN-01).
- Nombre con espacios en los extremos → se recorta y se guarda recortado
  (RN-05).
- Crear proyecto sin descripción ni fecha de inicio → aceptado; quedan
  vacíos / fecha actual.
- Modificar un proyecto sin cambiar ningún campo → operación válida, el
  proyecto queda igual.
- Modificar solo un campo (por ejemplo, solo la descripción) → los demás
  campos se conservan sin cambios.

### Condiciones de error

- Nombre vacío o solo espacios → error: nombre obligatorio.
- Nombre igual al de otro proyecto existente → error: nombre duplicado.
- ID inexistente al modificar → error: proyecto no encontrado.

### Criterios de aceptación

1. Puedo crear un proyecto con nombre, descripción y fecha de inicio; el
   sistema le asigna un ID único y estado "Activo".
2. No puedo crear un proyecto sin nombre o con nombre compuesto solo por
   espacios.
3. No puedo crear dos proyectos con el mismo nombre.
4. Un nombre con espacios en los extremos se guarda recortado.
5. Puedo modificar nombre, descripción y/o fecha de inicio de un proyecto
   existente por su ID, y los cambios quedan reflejados; los campos no
   modificados se conservan.
6. Si intento modificar un proyecto con ID inexistente, recibo el error
   "proyecto no encontrado".
7. Los escenarios BDD de esta historia están cubiertos por tests
   automatizados en Go y todos pasan (`go test ./...`).
