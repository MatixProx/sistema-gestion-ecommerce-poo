package ecommerce

import (
	"fmt"
	"strings"
)

type Categoria struct {
	id     string
	nombre string
}

func NewCategoria(id, nombre string) (*Categoria, error) {
	cat := &Categoria{}
	if err := cat.SetID(id); err != nil {
		return nil, err
	}
	if err := cat.SetNombre(nombre); err != nil {
		return nil, err
	}
	return cat, nil
}

func (c Categoria) ID() string     { return c.id }
func (c Categoria) Nombre() string { return c.nombre }

func (c *Categoria) SetID(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return fmt.Errorf("id categoria: %w", ErrCampoVacio)
	}
	c.id = id
	return nil
}

func (c *Categoria) SetNombre(nombre string) error {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return fmt.Errorf("nombre categoria: %w", ErrCampoVacio)
	}
	c.nombre = nombre
	return nil
}

// Producto mantiene sus atributos privados y valida sus cambios con setters.
type Producto struct {
	id          string
	nombre      string
	precio      float64
	stock       int
	descripcion string
	categoria   Categoria
}

func NewProducto(id, nombre string, precio float64, stock int, descripcion string, categoria Categoria) (*Producto, error) {
	p := &Producto{categoria: categoria}
	if err := p.SetID(id); err != nil {
		return nil, err
	}
	if err := p.SetNombre(nombre); err != nil {
		return nil, err
	}
	if err := p.SetPrecio(precio); err != nil {
		return nil, err
	}
	if err := p.SetStock(stock); err != nil {
		return nil, err
	}
	p.descripcion = strings.TrimSpace(descripcion)
	return p, nil
}

func (p Producto) ID() string           { return p.id }
func (p Producto) Nombre() string       { return p.nombre }
func (p Producto) Precio() float64      { return p.precio }
func (p Producto) Stock() int           { return p.stock }
func (p Producto) Descripcion() string  { return p.descripcion }
func (p Producto) Categoria() Categoria { return p.categoria }

func (p *Producto) SetID(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return fmt.Errorf("id producto: %w", ErrCampoVacio)
	}
	p.id = id
	return nil
}

func (p *Producto) SetNombre(nombre string) error {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return fmt.Errorf("nombre producto: %w", ErrCampoVacio)
	}
	p.nombre = nombre
	return nil
}

func (p *Producto) SetPrecio(precio float64) error {
	if precio <= 0 {
		return fmt.Errorf("precio %.2f: %w", precio, ErrPrecioInvalido)
	}
	p.precio = precio
	return nil
}

func (p *Producto) SetStock(stock int) error {
	if stock < 0 {
		return fmt.Errorf("stock %d: %w", stock, ErrStockInvalido)
	}
	p.stock = stock
	return nil
}

func (p Producto) Disponible(cantidad int) bool {
	return cantidad > 0 && p.stock >= cantidad
}

func (p *Producto) ReducirStock(cantidad int) error {
	if cantidad <= 0 {
		return ErrCantidadInvalida
	}
	if p.stock < cantidad {
		return fmt.Errorf("producto %s: %w", p.nombre, ErrStockInsuficiente)
	}
	p.stock -= cantidad
	return nil
}

func (p *Producto) AumentarStock(cantidad int) error {
	if cantidad <= 0 {
		return ErrCantidadInvalida
	}
	p.stock += cantidad
	return nil
}
