# HU-02 — Registro de integrantes de un proyecto

## Historia de usuario

Como Agile Enabler, quiero registrar los integrantes de un proyecto, para
saber quién forma parte del equipo.

## Especificación SDD

### Objetivo

Permitir al rol Agile Enabler dar de alta integrantes en un proyecto
existente y consultar la lista de integrantes de ese proyecto, de modo que
quede registrado quién forma parte del equipo de cada proyecto.

### Entradas

- **Registrar integrante:** ID del proyecto existente, nombre del
  integrante (obligatorio), rol dentro del equipo (obligatorio; debe
  corresponder a uno de los roles válidos del proyecto: Product
  Architect, Agile Enabler, Product Builder).
- **Listar integrantes:** ID del proyecto existente.

### Salidas esperadas

- Integrante registrado con ID único dentro del proyecto, nombre, rol en
  su forma canónica y datos guardados.
- Lista de integrantes del proyecto consultado, en el orden en que fueron
  registrados.
- Si la operación no es válida (nombre vacío, rol vacío o no válido,
  integrante duplicado, proyecto inexistente), un error descriptivo en
  lugar de un integrante.

### Reglas de negocio

- **RN-06:** el nombre del integrante es obligatorio; no puede quedar
  vacío ni contener solo espacios.
- **RN-07:** no se puede registrar dos veces el mismo integrante (mismo
  nombre) dentro del mismo proyecto; en proyectos distintos sí se permite.
- **RN-08:** solo se pueden registrar y listar integrantes de proyectos
  existentes, identificados por su ID.
- **RN-09:** los espacios al inicio y fin del nombre y del rol se recortan
  antes de validar y guardar.
- **RN-10:** el rol es obligatorio y debe corresponder a uno de los roles
  válidos del proyecto: Product Architect, Agile Enabler, Product
  Builder. La validación ignora mayúsculas/minúsculas (por ejemplo,
  "product architect" es válido) y el valor se guarda con su forma
  canónica de la lista (por ejemplo, "Product Architect").

> Nota: los campos del integrante (nombre y rol obligatorios, el rol
> validado contra los 3 roles válidos con comparación insensible a
> mayúsculas/minúsculas y forma canónica), el alcance (solo registrar y
> listar, sin quitar) y la regla de duplicados por proyecto fueron
> supuestos confirmados o ajustados con el usuario; el resto de detalles
> queda sujeto a ajuste si el dominio real define otros.

### Restricciones

- Código en Go, con el paquete estándar `testing` (sin librerías
  externas).
- Almacenamiento en memoria (repositorio en memoria); no se pide
  persistencia en disco ni interfaz de usuario en esta historia.
- El control de permisos por rol (quién puede registrar) no forma parte
  de esta historia.
- No se pide quitar ni modificar integrantes en esta historia.

### Casos límite

- Nombre con solo espacios → rechazado (RN-06).
- Nombre con espacios en los extremos → se recorta y se guarda recortado
  (RN-09).
- Rol con solo espacios → rechazado (RN-10).
- Rol con espacios en los extremos → se recorta y se valida (RN-09,
  RN-10).
- Rol en cualquier combinación de mayúsculas/minúsculas (por ejemplo
  "product architect" o "AGILE ENABLER") → aceptado; se guarda con su
  forma canónica ("Product Architect" / "Agile Enabler") (RN-10).
- Rol no correspondiente a los roles válidos (por ejemplo "Scrum Master"
  o "DevOps") → rechazado (RN-10).
- Listar integrantes de un proyecto sin integrantes → aceptado; devuelve
  una lista vacía.
- Registrar el mismo nombre en dos proyectos distintos → aceptado en
  ambos (RN-07).

### Condiciones de error

- Nombre del integrante vacío o solo espacios → error: nombre del
  integrante obligatorio.
- Rol vacío o solo espacios → error: rol obligatorio.
- Rol que no corresponde a los roles válidos → error: rol no válido.
- Integrante ya registrado en el mismo proyecto → error: integrante
  duplicado.
- ID de proyecto inexistente al registrar o listar → error: proyecto no
  encontrado.

### Criterios de aceptación

1. Puedo registrar un integrante con nombre y rol en un proyecto
   existente; el sistema le asigna un ID único dentro del proyecto.
2. No puedo registrar un integrante sin nombre o con nombre compuesto
   solo por espacios.
3. No puedo registrar un integrante sin rol o con rol compuesto solo por
   espacios.
4. No puedo registrar dos veces el mismo integrante (mismo nombre) en el
   mismo proyecto; sí puedo registrarlo en proyectos distintos.
5. El nombre y el rol con espacios en los extremos se guardan recortados.
6. Puedo registrar un integrante con el rol escrito en cualquier
   combinación de mayúsculas/minúsculas (por ejemplo "product architect")
   y el sistema lo guarda con su forma canónica ("Product Architect").
7. No puedo registrar un integrante con un rol que no sea uno de los
   roles válidos: Product Architect, Agile Enabler, Product Builder.
8. Puedo listar los integrantes de un proyecto y la lista refleja todos
   los registrados, en el orden en que fueron agregados; un proyecto sin
   integrantes devuelve una lista vacía.
9. Si intento registrar o listar integrantes de un proyecto con ID
   inexistente, recibo el error "proyecto no encontrado".
10. Los escenarios BDD de esta historia están cubiertos por tests
    automatizados en Go y todos pasan (`go test ./...`).
