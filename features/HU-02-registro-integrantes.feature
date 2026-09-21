# language: es

Característica: Registro de integrantes de un proyecto
  Como Agile Enabler
  quiero registrar los integrantes de un proyecto
  para saber quién forma parte del equipo

  # HU-02 — Trazabilidad: docs/especificaciones/HU-02-registro-integrantes.md

  Escenario: Registrar un integrante con nombre y rol en un proyecto existente (happy path)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando registro en el proyecto el integrante "Ana Pérez" con rol "Product Builder"
    Entonces el integrante se registra con un ID único dentro del proyecto
    Y el nombre guardado es "Ana Pérez"
    Y el rol guardado es "Product Builder"

  Escenario: Registrar un integrante con el rol en mayúsculas o minúsculas (caso límite)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando registro en el proyecto el integrante "Ana Pérez" con rol "product architect"
    Entonces el integrante se registra con un ID único dentro del proyecto
    Y el rol guardado es "Product Architect"

  Escenario: Registrar un integrante con nombre y rol con espacios en los extremos (caso límite)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando registro en el proyecto el integrante "  Ana Pérez  " con rol "  agile enabler  "
    Entonces el integrante se registra con un ID único dentro del proyecto
    Y el nombre guardado es "Ana Pérez"
    Y el rol guardado es "Agile Enabler"

  Escenario: Rechazar un integrante sin rol (error)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando intento registrar en el proyecto el integrante "Ana Pérez" con rol vacío o compuesto solo por espacios
    Entonces recibo el error "rol obligatorio"
    Y no se registra ningún integrante

  Escenario: Rechazar un integrante sin nombre (error)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando intento registrar en el proyecto un integrante con nombre vacío o compuesto solo por espacios y rol "Product Builder"
    Entonces recibo el error "nombre del integrante obligatorio"
    Y no se registra ningún integrante

  Escenario: Rechazar un integrante con rol no válido (error)
    Dado que existe un proyecto con nombre "Sistema de Métricas"
    Cuando registro en el proyecto el integrante "Ana Pérez" con rol "Scrum Master"
    Entonces recibo el error "rol no válido"
    Y no se registra ningún integrante

  Escenario: Rechazar un integrante duplicado en el mismo proyecto (error)
    Dado que existe un proyecto con nombre "Sistema de Métricas" con el integrante "Ana Pérez" registrado
    Cuando intento registrar en el proyecto el integrante "Ana Pérez" con rol "Agile Enabler"
    Entonces recibo el error "integrante duplicado"

  Escenario: Registrar el mismo integrante en dos proyectos distintos (caso alternativo)
    Dado que existen los proyectos "Sistema de Métricas" y "Portal Web"
    Cuando registro el integrante "Ana Pérez" con rol "Product Builder" en cada proyecto
    Entonces el registro es exitoso en ambos proyectos

  Escenario: Listar los integrantes de un proyecto en orden de registro (happy path)
    Dado que existe un proyecto con nombre "Sistema de Métricas" con los integrantes "Ana Pérez", "Luis Gómez" y "Marta Díaz" registrados
    Cuando consulto los integrantes del proyecto
    Entonces la lista contiene "Ana Pérez", "Luis Gómez" y "Marta Díaz" en el orden en que fueron registrados

  Escenario: Listar los integrantes de un proyecto sin integrantes (caso límite)
    Dado que existe un proyecto con nombre "Sistema de Métricas" sin integrantes
    Cuando consulto los integrantes del proyecto
    Entonces la lista está vacía

  Escenario: Registrar un integrante en un proyecto inexistente (error)
    Cuando intento registrar en el proyecto con ID "P-999" el integrante "Ana Pérez" con rol "Product Builder"
    Entonces recibo el error "proyecto no encontrado"

  Escenario: Listar los integrantes de un proyecto inexistente (error)
    Cuando intento consultar los integrantes del proyecto con ID "P-999"
    Entonces recibo el error "proyecto no encontrado"
