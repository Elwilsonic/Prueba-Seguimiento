# Prueba-Seguimiento
Repositorio de pruebas del equipo, usado para probar herramientas y flujos
de trabajo (agentes de IA, GitHub CLI, tableros Scrum) **antes** de aplicarlos
en el repositorio real del Trabajo Práctico Integrador.

⚠️ **Este repo no forma parte de la entrega.** Es solo un sandbox de pruebas.

## Qué se probó acá

- Instalación y configuración de OpenCode con modelo NVIDIA/GLM-5.3-Flash.
- Creación de Issues en GitHub vía `gh issue create` desde el agente de IA.
- Formato de especificaciones SDD + escenarios BDD generados automáticamente.
- Integración del tablero Scrum (GitHub Projects) con Issues del repositorio.

## Repositorio real del TP

👉 [seguimiento-medicion-GPC](https://github.com/Elwilsonic/seguimiento-medicion-GPC)

## Recordatorio: ciclo TDD (Red → Green → Refactor)

Antes de implementar cualquier función, hay que escribir primero el test.
El PDF pide dejar evidencia de este proceso en el historial de Git.

1. **RED** — Escribir el test primero. Correrlo con `go test` y confirmar
   que **falla** (porque el código todavía no existe o está incompleto).
2. **GREEN** — Escribir el código mínimo necesario para que ese test pase.
   Correr `go test` de nuevo y confirmar que ahora **pasa**.
3. **REFACTOR** — Mejorar el código (legibilidad, estructura) sin romper
   el test que ya pasaba.

**Sugerencia para dejar evidencia en Git:** hacer un commit separado por
cada fase, por ejemplo:

git commit -m "test: agregar test para CrearProyecto (RED)"
git commit -m "feat: implementar CrearProyecto (GREEN)"

Así, cualquiera que revise el historial del repositorio puede ver
claramente que el test se escribió antes que el código.