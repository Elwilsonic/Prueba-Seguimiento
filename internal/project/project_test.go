package project

import (
	"errors"
	"testing"
	"time"
)

// fechaInicio representa la fecha usada en los escenarios BDD ("2026-09-15").
var fechaInicio = time.Date(2026, time.September, 15, 0, 0, 0, 0, time.UTC)

// Cubre escenario BDD: "Crear un proyecto con datos completos"
func TestCreate_ProyectoConDatosCompletos(t *testing.T) {
	repo := NewRepository()

	p, err := repo.Create(CreateInput{
		Name:        "Sistema de Métricas",
		Description: "TP Integrador",
		StartDate:   &fechaInicio,
	})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if p.ID == "" {
		t.Fatal("se esperaba un ID único no vacío")
	}
	if p.Status != StatusActive {
		t.Errorf("estado = %q, se esperaba %q", p.Status, StatusActive)
	}
	if p.Name != "Sistema de Métricas" {
		t.Errorf("nombre = %q, se esperaba %q", p.Name, "Sistema de Métricas")
	}
	if p.Description != "TP Integrador" {
		t.Errorf("descripción = %q, se esperaba %q", p.Description, "TP Integrador")
	}
	if !p.StartDate.Equal(fechaInicio) {
		t.Errorf("fecha de inicio = %v, se esperaba %v", p.StartDate, fechaInicio)
	}
}

// Cubre escenario BDD: "Crear un proyecto solo con nombre"
func TestCreate_ProyectoSinDatosOpcionales(t *testing.T) {
	repo := NewRepository()

	p, err := repo.Create(CreateInput{Name: "Proyecto mínimo"})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if p.ID == "" {
		t.Fatal("se esperaba un ID único no vacío")
	}
	if p.Status != StatusActive {
		t.Errorf("estado = %q, se esperaba %q", p.Status, StatusActive)
	}
	if p.Description != "" {
		t.Errorf("descripción = %q, se esperaba vacía", p.Description)
	}
	ahora := time.Now()
	if p.StartDate.Year() != ahora.Year() || p.StartDate.Month() != ahora.Month() || p.StartDate.Day() != ahora.Day() {
		t.Errorf("fecha de inicio = %v, se esperaba la fecha actual", p.StartDate)
	}
}

// Cubre escenario BDD: "Crear un proyecto con nombre con espacios en los extremos"
func TestCreate_NombreConEspaciosSeRecorta(t *testing.T) {
	repo := NewRepository()

	p, err := repo.Create(CreateInput{Name: "  Proyecto con espacios  "})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if p.Name != "Proyecto con espacios" {
		t.Errorf("nombre = %q, se esperaba %q", p.Name, "Proyecto con espacios")
	}
}

// Cubre escenario BDD: "Rechazar un proyecto sin nombre"
func TestCreate_NombreVacioDevuelveError(t *testing.T) {
	repo := NewRepository()
	casos := []string{"", "   "}

	for _, nombre := range casos {
		if _, err := repo.Create(CreateInput{Name: nombre}); !errors.Is(err, ErrNameRequired) {
			t.Errorf("nombre %q: se esperaba error %q, se obtuvo %v", nombre, ErrNameRequired, err)
		}
	}

	// No se crea ningún proyecto: el primer ID válido asignado sigue siendo P-001.
	p, err := repo.Create(CreateInput{Name: "Válido"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear un proyecto válido, se obtuvo: %v", err)
	}
	if p.ID != "P-001" {
		t.Errorf("ID = %q, se esperaba %q porque no debió crearse ningún proyecto antes", p.ID, "P-001")
	}
}

// Cubre escenario BDD: "Rechazar un proyecto con nombre duplicado"
func TestCreate_NombreDuplicadoDevuelveError(t *testing.T) {
	repo := NewRepository()

	if _, err := repo.Create(CreateInput{Name: "Duplicado"}); err != nil {
		t.Fatalf("no se esperaba error en la primera creación, se obtuvo: %v", err)
	}
	if _, err := repo.Create(CreateInput{Name: "Duplicado"}); !errors.Is(err, ErrDuplicateName) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrDuplicateName, err)
	}
}

// Cubre escenario BDD: "Modificar la descripción de un proyecto existente"
func TestUpdate_ModificaProyectoExistente(t *testing.T) {
	repo := NewRepository()
	creado, err := repo.Create(CreateInput{Name: "Original", Description: "Descripción inicial"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear, se obtuvo: %v", err)
	}

	nuevaDesc := "Nueva descripción"
	actualizado, err := repo.Update(creado.ID, UpdateInput{Description: &nuevaDesc})

	if err != nil {
		t.Fatalf("no se esperaba error al modificar, se obtuvo: %v", err)
	}
	if actualizado.Description != "Nueva descripción" {
		t.Errorf("descripción = %q, se esperaba %q", actualizado.Description, "Nueva descripción")
	}
	if actualizado.Name != "Original" {
		t.Errorf("nombre = %q, se esperaba %q sin cambios", actualizado.Name, "Original")
	}
}

// Cubre escenario BDD: "Modificar solo el nombre conservando los demás campos"
func TestUpdate_ModificaSoloUnCampo(t *testing.T) {
	repo := NewRepository()
	creado, err := repo.Create(CreateInput{Name: "Solo campo", Description: "Datos previos", StartDate: &fechaInicio})
	if err != nil {
		t.Fatalf("no se esperaba error al crear, se obtuvo: %v", err)
	}

	nuevoNombre := "Nuevo nombre"
	actualizado, err := repo.Update(creado.ID, UpdateInput{Name: &nuevoNombre})

	if err != nil {
		t.Fatalf("no se esperaba error al modificar, se obtuvo: %v", err)
	}
	if actualizado.Name != "Nuevo nombre" {
		t.Errorf("nombre = %q, se esperaba %q", actualizado.Name, "Nuevo nombre")
	}
	if actualizado.Description != "Datos previos" {
		t.Errorf("descripción = %q, se esperaba %q sin cambios", actualizado.Description, "Datos previos")
	}
	if !actualizado.StartDate.Equal(fechaInicio) {
		t.Errorf("fecha de inicio = %v, se esperaba %v sin cambios", actualizado.StartDate, fechaInicio)
	}
}

// Cubre escenario BDD: "Modificar un proyecto sin cambiar ningún campo"
func TestUpdate_SinCambiosMantieneProyecto(t *testing.T) {
	repo := NewRepository()
	creado, err := repo.Create(CreateInput{Name: "Sin cambios", Description: "Estable", StartDate: &fechaInicio})
	if err != nil {
		t.Fatalf("no se esperaba error al crear, se obtuvo: %v", err)
	}

	actualizado, err := repo.Update(creado.ID, UpdateInput{})

	if err != nil {
		t.Fatalf("no se esperaba error al modificar sin cambios, se obtuvo: %v", err)
	}
	if actualizado != creado {
		t.Errorf("proyecto = %+v, se esperaba %+v sin cambios", actualizado, creado)
	}
}

// Cubre escenario BDD: "Modificar un proyecto inexistente"
func TestUpdate_IdInexistenteDevuelveError(t *testing.T) {
	repo := NewRepository()

	if _, err := repo.Update("P-999", UpdateInput{}); !errors.Is(err, ErrProjectNotFound) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrProjectNotFound, err)
	}
}

// Cubre escenario BDD: "Modificar la descripción de un proyecto existente"
// (variante: modificar la fecha de inicio).
func TestUpdate_ModificaFechaInicio(t *testing.T) {
	repo := NewRepository()
	creado, err := repo.Create(CreateInput{Name: "Con fecha"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear, se obtuvo: %v", err)
	}

	nuevaFecha := time.Date(2026, time.December, 1, 0, 0, 0, 0, time.UTC)
	actualizado, err := repo.Update(creado.ID, UpdateInput{StartDate: &nuevaFecha})

	if err != nil {
		t.Fatalf("no se esperaba error al modificar, se obtuvo: %v", err)
	}
	if !actualizado.StartDate.Equal(nuevaFecha) {
		t.Errorf("fecha de inicio = %v, se esperaba %v", actualizado.StartDate, nuevaFecha)
	}
	if actualizado.Name != "Con fecha" {
		t.Errorf("nombre = %q, se esperaba %q sin cambios", actualizado.Name, "Con fecha")
	}
}

// Cubre RN-01 aplicada a la modificación: nombre vacío al actualizar.
func TestUpdate_NombreVacioDevuelveError(t *testing.T) {
	repo := NewRepository()
	creado, err := repo.Create(CreateInput{Name: "Existente"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear, se obtuvo: %v", err)
	}

	vacio := "   "
	if _, err := repo.Update(creado.ID, UpdateInput{Name: &vacio}); !errors.Is(err, ErrNameRequired) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrNameRequired, err)
	}
}

// Cubre escenario BDD: "Rechazar una modificación hacia un nombre duplicado"
func TestUpdate_NombreDuplicadoDevuelveError(t *testing.T) {
	repo := NewRepository()
	if _, err := repo.Create(CreateInput{Name: "Alpha"}); err != nil {
		t.Fatalf("no se esperaba error al crear Alpha, se obtuvo: %v", err)
	}
	beta, err := repo.Create(CreateInput{Name: "Beta"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear Beta, se obtuvo: %v", err)
	}

	nombre := "Alpha"
	if _, err := repo.Update(beta.ID, UpdateInput{Name: &nombre}); !errors.Is(err, ErrDuplicateName) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrDuplicateName, err)
	}

	// El nombre de Beta no debió cambiar tras el error.
	sinCambios, err := repo.Update(beta.ID, UpdateInput{})
	if err != nil {
		t.Fatalf("no se esperaba error al reconsultar, se obtuvo: %v", err)
	}
	if sinCambios.Name != "Beta" {
		t.Errorf("nombre de Beta = %q, se esperaba %q sin cambios", sinCambios.Name, "Beta")
	}
}
