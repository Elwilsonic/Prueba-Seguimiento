package project

import (
	"errors"
	"testing"
)

// Cubre escenario BDD: "Registrar un integrante con nombre y rol en un
// proyecto existente" (happy path).
func TestAddMember_RegistraIntegranteConNombreYRol(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	m, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: "Product Builder"})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if m.ID == "" {
		t.Fatal("se esperaba un ID único dentro del proyecto, no vacío")
	}
	if m.Name != "Ana Pérez" {
		t.Errorf("nombre = %q, se esperaba %q", m.Name, "Ana Pérez")
	}
	if m.Role != "Product Builder" {
		t.Errorf("rol = %q, se esperaba %q", m.Role, "Product Builder")
	}
}

// Cubre escenario BDD: "Registrar un integrante con el rol en mayúsculas
// o minúsculas" (caso límite): la validación ignora mayúsculas/minúsculas
// y guarda la forma canónica del rol (RN-10).
func TestAddMember_RolEnMinusculasGuardaCanonico(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	m, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: "product architect"})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if m.Role != "Product Architect" {
		t.Errorf("rol = %q, se esperaba la forma canónica %q", m.Role, "Product Architect")
	}
}

// Cubre escenario BDD: "Registrar un integrante con nombre y rol con
// espacios en los extremos" (caso límite): se recortan antes de validar
// y guardar (RN-09).
func TestAddMember_NombreYRolConEspaciosSeRecortan(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	m, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "  Ana Pérez  ", Role: "  agile enabler  "})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if m.Name != "Ana Pérez" {
		t.Errorf("nombre = %q, se esperaba %q", m.Name, "Ana Pérez")
	}
	if m.Role != "Agile Enabler" {
		t.Errorf("rol = %q, se esperaba la forma canónica %q", m.Role, "Agile Enabler")
	}
}

// Cubre escenario BDD: "Rechazar un integrante sin rol" (error):
// rol vacío o compuesto solo por espacios (RN-10).
func TestAddMember_SinRolDevuelveError(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	for _, rol := range []string{"", "   "} {
		if _, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: rol}); !errors.Is(err, ErrMemberRoleRequired) {
			t.Errorf("rol %q: se esperaba error %q, se obtuvo %v", rol, ErrMemberRoleRequired, err)
		}
	}
}

// Cubre escenario BDD: "Rechazar un integrante sin nombre" (error):
// nombre vacío o compuesto solo por espacios (RN-06).
func TestAddMember_NombreVacioDevuelveError(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	for _, nombre := range []string{"", "   "} {
		if _, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: nombre, Role: "Product Builder"}); !errors.Is(err, ErrMemberNameRequired) {
			t.Errorf("nombre %q: se esperaba error %q, se obtuvo %v", nombre, ErrMemberNameRequired, err)
		}
	}
}

// Cubre escenario BDD: "Rechazar un integrante con rol no válido" (error):
// el rol no corresponde a los roles válidos del proyecto (RN-10).
func TestAddMember_RolNoValidoDevuelveError(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	for _, rol := range []string{"Scrum Master", "DevOps"} {
		if _, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: rol}); !errors.Is(err, ErrInvalidRole) {
			t.Errorf("rol %q: se esperaba error %q, se obtuvo %v", rol, ErrInvalidRole, err)
		}
	}

	// No se registra ningún integrante tras los errores.
	integrantes, err := repo.ListMembers(proyecto.ID)
	if err != nil {
		t.Fatalf("no se esperaba error al listar, se obtuvo: %v", err)
	}
	if len(integrantes) != 0 {
		t.Errorf("cantidad de integrantes = %d, se esperaba 0", len(integrantes))
	}
}

// Cubre escenario BDD: "Rechazar un integrante duplicado en el mismo
// proyecto" (error): mismo nombre dos veces en el mismo proyecto (RN-07).
func TestAddMember_IntegranteDuplicadoDevuelveError(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}
	if _, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: "Product Builder"}); err != nil {
		t.Fatalf("no se esperaba error en el primer registro, se obtuvo: %v", err)
	}

	if _, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: "Ana Pérez", Role: "Agile Enabler"}); !errors.Is(err, ErrMemberDuplicate) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrMemberDuplicate, err)
	}
}

// Cubre escenario BDD: "Registrar el mismo integrante en dos proyectos
// distintos" (caso alternativo): la regla de duplicados es por proyecto
// (RN-07).
func TestAddMember_MismoIntegranteEnDistintosProyectos(t *testing.T) {
	repo := NewRepository()
	p1, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el primer proyecto, se obtuvo: %v", err)
	}
	p2, err := repo.Create(CreateInput{Name: "Portal Web"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el segundo proyecto, se obtuvo: %v", err)
	}

	m1, err := repo.AddMember(p1.ID, AddMemberInput{Name: "Ana Pérez", Role: "Product Builder"})
	if err != nil {
		t.Fatalf("no se esperaba error al registrar en el primer proyecto, se obtuvo: %v", err)
	}
	m2, err := repo.AddMember(p2.ID, AddMemberInput{Name: "Ana Pérez", Role: "Product Builder"})
	if err != nil {
		t.Fatalf("no se esperaba error al registrar en el segundo proyecto, se obtuvo: %v", err)
	}

	if m1.Name != "Ana Pérez" || m2.Name != "Ana Pérez" {
		t.Errorf("nombres = %q, %q; se esperaba %q en ambos proyectos", m1.Name, m2.Name, "Ana Pérez")
	}
}

// Cubre escenario BDD: "Listar los integrantes de un proyecto en orden de
// registro" (happy path).
func TestListMembers_ListaEnOrdenDeRegistro(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}
	nombres := []string{"Ana Pérez", "Luis Gómez", "Marta Díaz"}
	roles := []string{"Product Builder", "Product Architect", "Agile Enabler"}
	ids := make([]string, 0, len(nombres))
	for i, nombre := range nombres {
		m, err := repo.AddMember(proyecto.ID, AddMemberInput{Name: nombre, Role: roles[i]})
		if err != nil {
			t.Fatalf("no se esperaba error al registrar %q, se obtuvo: %v", nombre, err)
		}
		ids = append(ids, m.ID)
	}

	integrantes, err := repo.ListMembers(proyecto.ID)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if len(integrantes) != len(nombres) {
		t.Fatalf("cantidad de integrantes = %d, se esperaba %d", len(integrantes), len(nombres))
	}
	for i, nombre := range nombres {
		if integrantes[i].Name != nombre {
			t.Errorf("integrante %d: nombre = %q, se esperaba %q en el orden de registro", i, integrantes[i].Name, nombre)
		}
		if integrantes[i].Role != roles[i] {
			t.Errorf("integrante %d: rol = %q, se esperaba %q", i, integrantes[i].Role, roles[i])
		}
	}
	// IDs únicos dentro del proyecto.
	vistos := make(map[string]bool, len(ids))
	for i, id := range ids {
		if vistos[id] {
			t.Errorf("ID %q del integrante %d repetido dentro del proyecto", id, i)
		}
		vistos[id] = true
	}
	if vistos[integrantes[0].ID] == false || vistos[integrantes[1].ID] == false {
		t.Error("los IDs de los integrantes listados no coinciden con los asignados al registrar")
	}
}

// Cubre escenario BDD: "Listar los integrantes de un proyecto sin
// integrantes" (caso límite): devuelve una lista vacía.
func TestListMembers_ProyectoSinIntegrantesDevuelveListaVacia(t *testing.T) {
	repo := NewRepository()
	proyecto, err := repo.Create(CreateInput{Name: "Sistema de Métricas"})
	if err != nil {
		t.Fatalf("no se esperaba error al crear el proyecto, se obtuvo: %v", err)
	}

	integrantes, err := repo.ListMembers(proyecto.ID)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if len(integrantes) != 0 {
		t.Errorf("cantidad de integrantes = %d, se esperaba 0", len(integrantes))
	}
}

// Cubre escenario BDD: "Registrar un integrante en un proyecto
// inexistente" (error) (RN-08).
func TestAddMember_ProyectoInexistenteDevuelveError(t *testing.T) {
	repo := NewRepository()

	if _, err := repo.AddMember("P-999", AddMemberInput{Name: "Ana Pérez", Role: "Product Builder"}); !errors.Is(err, ErrProjectNotFound) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrProjectNotFound, err)
	}
}

// Cubre escenario BDD: "Listar los integrantes de un proyecto
// inexistente" (error) (RN-08).
func TestListMembers_ProyectoInexistenteDevuelveError(t *testing.T) {
	repo := NewRepository()

	if _, err := repo.ListMembers("P-999"); !errors.Is(err, ErrProjectNotFound) {
		t.Errorf("se esperaba error %q, se obtuvo %v", ErrProjectNotFound, err)
	}
}
