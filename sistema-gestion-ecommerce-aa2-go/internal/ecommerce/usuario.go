package ecommerce

import (
	"fmt"
	"strings"
)

// Usuario representa la informacion comun de cualquier actor del sistema.
// Los campos son privados para aplicar encapsulacion: solo se modifican con setters.
type Usuario struct {
	id         string
	nombre     string
	correo     string
	contrasena string
	rol        string
}

func NewUsuario(id, nombre, correo, contrasena, rol string) (*Usuario, error) {
	u := &Usuario{}
	if err := u.SetID(id); err != nil {
		return nil, err
	}
	if err := u.SetNombre(nombre); err != nil {
		return nil, err
	}
	if err := u.SetCorreo(correo); err != nil {
		return nil, err
	}
	if err := u.SetContrasena(contrasena); err != nil {
		return nil, err
	}
	if err := u.SetRol(rol); err != nil {
		return nil, err
	}
	return u, nil
}

func (u Usuario) ID() string     { return u.id }
func (u Usuario) Nombre() string { return u.nombre }
func (u Usuario) Correo() string { return u.correo }
func (u Usuario) Rol() string    { return u.rol }

func (u *Usuario) SetID(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return fmt.Errorf("id: %w", ErrCampoVacio)
	}
	u.id = id
	return nil
}

func (u *Usuario) SetNombre(nombre string) error {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return fmt.Errorf("nombre: %w", ErrCampoVacio)
	}
	u.nombre = nombre
	return nil
}

func (u *Usuario) SetCorreo(correo string) error {
	correo = strings.TrimSpace(strings.ToLower(correo))
	if correo == "" {
		return fmt.Errorf("correo: %w", ErrCampoVacio)
	}
	if !strings.Contains(correo, "@") || !strings.Contains(correo, ".") {
		return ErrCorreoInvalido
	}
	u.correo = correo
	return nil
}

func (u *Usuario) SetContrasena(contrasena string) error {
	if len(strings.TrimSpace(contrasena)) < 6 {
		return ErrContrasenaInvalida
	}
	u.contrasena = contrasena
	return nil
}

func (u *Usuario) SetRol(rol string) error {
	rol = strings.TrimSpace(strings.ToLower(rol))
	if rol == "" {
		return fmt.Errorf("rol: %w", ErrCampoVacio)
	}
	u.rol = rol
	return nil
}

func (u Usuario) Autenticar(correo, contrasena string) bool {
	return u.correo == strings.ToLower(strings.TrimSpace(correo)) && u.contrasena == contrasena
}

// Cliente compone Usuario. En Go se usa composicion en lugar de herencia clasica.
type Cliente struct {
	Usuario
	direccion string
	telefono  string
}

func NewCliente(id, nombre, correo, contrasena, direccion, telefono string) (*Cliente, error) {
	u, err := NewUsuario(id, nombre, correo, contrasena, "cliente")
	if err != nil {
		return nil, err
	}
	c := &Cliente{Usuario: *u}
	if err := c.SetDireccion(direccion); err != nil {
		return nil, err
	}
	if err := c.SetTelefono(telefono); err != nil {
		return nil, err
	}
	return c, nil
}

func (c Cliente) Direccion() string { return c.direccion }
func (c Cliente) Telefono() string  { return c.telefono }

func (c *Cliente) SetDireccion(direccion string) error {
	direccion = strings.TrimSpace(direccion)
	if direccion == "" {
		return fmt.Errorf("direccion: %w", ErrCampoVacio)
	}
	c.direccion = direccion
	return nil
}

func (c *Cliente) SetTelefono(telefono string) error {
	telefono = strings.TrimSpace(telefono)
	if telefono == "" {
		return fmt.Errorf("telefono: %w", ErrCampoVacio)
	}
	c.telefono = telefono
	return nil
}

// Administrador representa un usuario con permisos administrativos.
type Administrador struct {
	Usuario
	cargo    string
	permisos []string
}

func NewAdministrador(id, nombre, correo, contrasena, cargo string, permisos []string) (*Administrador, error) {
	u, err := NewUsuario(id, nombre, correo, contrasena, "administrador")
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(cargo) == "" {
		return nil, fmt.Errorf("cargo: %w", ErrCampoVacio)
	}
	return &Administrador{Usuario: *u, cargo: cargo, permisos: permisos}, nil
}

func (a Administrador) Cargo() string      { return a.cargo }
func (a Administrador) Permisos() []string { return append([]string{}, a.permisos...) }
