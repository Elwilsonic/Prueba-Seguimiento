# language: es

Característica: Gestión de proyectos
  Como Agile Enabler
  quiero crear y modificar un proyecto
  para poder gestionarlo desde el inicio

  # HU-01 — Trazabilidad: docs/especificaciones/HU-01-gestion-proyectos.md

  Escenario: Crear un proyecto con datos completos (happy path)
    Dado que no existe un proyecto con nombre "Sistema de Métricas"
    Cuando creo un proyecto con nombre "Sistema de Métricas", descripción "TP Integrador" y fecha de inicio "2026-09-15"
    Entonces el proyecto se crea con un ID único
    Y el estado del proyecto es "Activo"
    Y la descripción guardada es "TP Integrador"
    Y la fecha de inicio guardada es "2026-09-15"

  Escenario: Crear un proyecto solo con nombre (caso alternativo)
    Cuando creo un proyecto con nombre "Proyecto mínimo"
    Entonces el proyecto se crea con un ID único
    Y el estado del proyecto es "Activo"
    Y la descripción guardada está vacía
    Y la fecha de inicio guardada es la fecha actual

  Escenario: Crear un proyecto con nombre con espacios en los extremos (caso límite)
    Cuando creo un proyecto con nombre "  Proyecto con espacios  "
    Entonces el proyecto se guarda con el nombre "Proyecto con espacios"

  Escenario: Rechazar un proyecto sin nombre (error)
    Cuando intento crear un proyecto con nombre vacío o compuesto solo por espacios
    Entonces recibo el error "nombre obligatorio"
    Y no se crea ningún proyecto

  Escenario: Rechazar un proyecto con nombre duplicado (error)
    Dado que existe un proyecto con nombre "Duplicado"
    Cuando creo otro proyecto con nombre "Duplicado"
    Entonces recibo el error "nombre duplicado"

  Escenario: Modificar la descripción de un proyecto existente (happy path)
    Dado que existe un proyecto con nombre "Original" y descripción "Descripción inicial"
    Cuando modifico el proyecto cambiando la descripción a "Nueva descripción"
    Entonces la descripción guardada es "Nueva descripción"
    Y el nombre sigue siendo "Original"

  Escenario: Modificar solo el nombre conservando los demás campos (caso límite)
    Dado que existe un proyecto con nombre "Solo campo", descripción "Datos previos" y fecha de inicio "2026-09-15"
    Cuando modifico el proyecto cambiando solo el nombre a "Nuevo nombre"
    Entonces el nombre guardado es "Nuevo nombre"
    Y la descripción se conserva como "Datos previos"
    Y la fecha de inicio se conserva como "2026-09-15"

  Escenario: Modificar un proyecto sin cambiar ningún campo (caso límite)
    Dado que existe un proyecto con nombre "Sin cambios" y descripción "Estable"
    Cuando modifico el proyecto sin cambiar ningún campo
    Entonces la operación es exitosa
    Y el proyecto queda igual

  Escenario: Modificar un proyecto inexistente (error)
    Cuando intento modificar el proyecto con ID "P-999"
    Entonces recibo el error "proyecto no encontrado"

  Escenario: Rechazar una modificación hacia un nombre duplicado (error)
    Dado que existen los proyectos "Alpha" y "Beta"
    Cuando modifico el proyecto "Beta" cambiando su nombre a "Alpha"
    Entonces recibo el error "nombre duplicado"
    Y el nombre de "Beta" sigue siendo "Beta"
