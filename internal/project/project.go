// Package project implementa la gestión de proyectos (HU-01):
// crear y modificar proyectos del sistema.
package project

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	// ErrNameRequired corresponde a la regla de negocio RN-01.
	ErrNameRequired = errors.New("nombre obligatorio")
	// ErrDuplicateName corresponde a la regla de negocio RN-02.
	ErrDuplicateName = errors.New("nombre duplicado")
	// ErrProjectNotFound corresponde a la regla de negocio RN-03 y RN-08.
	ErrProjectNotFound = errors.New("proyecto no encontrado")
	// ErrMemberNameRequired corresponde a la regla de negocio RN-06.
	ErrMemberNameRequired = errors.New("nombre del integrante obligatorio")
	// ErrMemberRoleRequired corresponde a la regla de negocio RN-10.
	ErrMemberRoleRequired = errors.New("rol obligatorio")
	// ErrInvalidRole corresponde a la regla de negocio RN-10.
	ErrInvalidRole = errors.New("rol no válido")
	// ErrMemberDuplicate corresponde a la regla de negocio RN-07.
	ErrMemberDuplicate = errors.New("integrante duplicado")
)

// rolesValidos son los roles válidos del proyecto (RN-10).
var rolesValidos = []string{"Product Architect", "Agile Enabler", "Product Builder"}

// StatusActive es el estado inicial de todo proyecto nuevo (RN-04).
const StatusActive = "Activo"

// Project representa un proyecto de software gestionado por el sistema.
type Project struct {
	ID          string
	Name        string
	Description string
	StartDate   time.Time
	Status      string
}

// CreateInput agrupa los datos de alta de un proyecto.
// Description y StartDate son opcionales.
type CreateInput struct {
	Name        string
	Description string
	StartDate   *time.Time
}

// UpdateInput agrupa los datos de modificación de un proyecto.
// Los campos en nil no se modifican.
type UpdateInput struct {
	Name        *string
	Description *string
	StartDate   *time.Time
}

// AddMemberInput agrupa los datos de alta de un integrante en un
// proyecto. Ambos campos son obligatorios (RN-06, RN-10).
type AddMemberInput struct {
	Name string
	Role string
}

// Member representa un integrante registrado en un proyecto.
type Member struct {
	ID   string
	Name string
	Role string
}

// Repository es el almacenamiento en memoria de los proyectos (RN-03)
// y de los integrantes de cada proyecto (RN-08).
type Repository struct {
	next       int
	byID       map[string]Project
	members    map[string][]Member
	nextMember map[string]int
}

// NewRepository devuelve un repositorio vacío.
func NewRepository() *Repository {
	return &Repository{
		byID:       make(map[string]Project),
		members:    make(map[string][]Member),
		nextMember: make(map[string]int),
	}
}

// Create da de alta un proyecto (RN-01, RN-02, RN-04, RN-05).
func (r *Repository) Create(in CreateInput) (Project, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return Project{}, ErrNameRequired
	}
	if r.nameExists(name, "") {
		return Project{}, ErrDuplicateName
	}

	r.next++
	start := today()
	if in.StartDate != nil {
		start = *in.StartDate
	}
	p := Project{
		ID:          fmt.Sprintf("P-%03d", r.next),
		Name:        name,
		Description: in.Description,
		StartDate:   start,
		Status:      StatusActive,
	}
	r.byID[p.ID] = p
	return p, nil
}

// Update modifica un proyecto existente por su ID (RN-03, RN-05).
// Los campos de in en nil se conservan sin cambios.
func (r *Repository) Update(id string, in UpdateInput) (Project, error) {
	p, ok := r.byID[id]
	if !ok {
		return Project{}, ErrProjectNotFound
	}

	name := p.Name
	if in.Name != nil {
		name = strings.TrimSpace(*in.Name)
		if name == "" {
			return Project{}, ErrNameRequired
		}
	}
	if r.nameExists(name, id) {
		return Project{}, ErrDuplicateName
	}

	if in.Name != nil {
		p.Name = name
	}
	if in.Description != nil {
		p.Description = *in.Description
	}
	if in.StartDate != nil {
		p.StartDate = *in.StartDate
	}
	r.byID[id] = p
	return p, nil
}

// nameExists indica si ya existe otro proyecto (distinto de excludeID)
// con el mismo nombre (RN-02).
func (r *Repository) nameExists(name, excludeID string) bool {
	for id, p := range r.byID {
		if id != excludeID && p.Name == name {
			return true
		}
	}
	return false
}

// AddMember registra un integrante en un proyecto existente
// (RN-06, RN-07, RN-08, RN-09, RN-10).
func (r *Repository) AddMember(projectID string, in AddMemberInput) (Member, error) {
	if _, ok := r.byID[projectID]; !ok {
		return Member{}, ErrProjectNotFound
	}

	name := strings.TrimSpace(in.Name)
	if name == "" {
		return Member{}, ErrMemberNameRequired
	}

	role, err := canonicalRole(in.Role)
	if err != nil {
		return Member{}, err
	}

	for _, m := range r.members[projectID] {
		if m.Name == name {
			return Member{}, ErrMemberDuplicate
		}
	}

	r.nextMember[projectID]++
	m := Member{
		ID:   fmt.Sprintf("M-%03d", r.nextMember[projectID]),
		Name: name,
		Role: role,
	}
	r.members[projectID] = append(r.members[projectID], m)
	return m, nil
}

// ListMembers devuelve los integrantes de un proyecto existente (RN-08),
// en el orden en que fueron registrados.
func (r *Repository) ListMembers(projectID string) ([]Member, error) {
	if _, ok := r.byID[projectID]; !ok {
		return nil, ErrProjectNotFound
	}
	return r.members[projectID], nil
}

// canonicalRole valida el rol contra los roles válidos del proyecto
// ignorando mayúsculas/minúsculas (RN-10) y devuelve su forma canónica.
func canonicalRole(role string) (string, error) {
	recortado := strings.TrimSpace(role)
	if recortado == "" {
		return "", ErrMemberRoleRequired
	}
	for _, valido := range rolesValidos {
		if strings.EqualFold(recortado, valido) {
			return valido, nil
		}
	}
	return "", ErrInvalidRole
}

// today devuelve la fecha actual truncada al día (valor por defecto de
// la fecha de inicio cuando no se indica).
func today() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location())
}
