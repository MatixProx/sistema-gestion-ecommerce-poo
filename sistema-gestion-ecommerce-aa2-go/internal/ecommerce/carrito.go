package ecommerce

import "fmt"

type DetalleCarrito struct {
	producto Producto
	cantidad int
}

func NewDetalleCarrito(producto Producto, cantidad int) (DetalleCarrito, error) {
	if cantidad <= 0 {
		return DetalleCarrito{}, ErrCantidadInvalida
	}
	if !producto.Disponible(cantidad) {
		return DetalleCarrito{}, fmt.Errorf("producto %s: %w", producto.Nombre(), ErrStockInsuficiente)
	}
	return DetalleCarrito{producto: producto, cantidad: cantidad}, nil
}

func (d DetalleCarrito) Producto() Producto { return d.producto }
func (d DetalleCarrito) Cantidad() int      { return d.cantidad }
func (d DetalleCarrito) Subtotal() float64  { return d.producto.Precio() * float64(d.cantidad) }

// Carrito usa un slice porque los items pueden crecer o disminuir dinamicamente.
type Carrito struct {
	id    string
	items []DetalleCarrito
}

func NewCarrito(id string) *Carrito {
	return &Carrito{id: id, items: []DetalleCarrito{}}
}

func (c Carrito) ID() string              { return c.id }
func (c Carrito) Items() []DetalleCarrito { return append([]DetalleCarrito{}, c.items...) }

func (c *Carrito) AgregarProducto(producto Producto, cantidad int) error {
	if cantidad <= 0 {
		return ErrCantidadInvalida
	}
	if !producto.Disponible(cantidad) {
		return fmt.Errorf("producto %s: %w", producto.Nombre(), ErrStockInsuficiente)
	}
	for idx := range c.items {
		if c.items[idx].Producto().ID() == producto.ID() {
			nuevaCantidad := c.items[idx].cantidad + cantidad
			if !producto.Disponible(nuevaCantidad) {
				return fmt.Errorf("producto %s: %w", producto.Nombre(), ErrStockInsuficiente)
			}
			c.items[idx].cantidad = nuevaCantidad
			return nil
		}
	}
	detalle, err := NewDetalleCarrito(producto, cantidad)
	if err != nil {
		return err
	}
	c.items = append(c.items, detalle)
	return nil
}

func (c *Carrito) EliminarProducto(productoID string) error {
	for idx, item := range c.items {
		if item.Producto().ID() == productoID {
			c.items = append(c.items[:idx], c.items[idx+1:]...)
			return nil
		}
	}
	return ErrProductoNoExiste
}

func (c *Carrito) CambiarCantidad(productoID string, cantidad int) error {
	if cantidad <= 0 {
		return ErrCantidadInvalida
	}
	for idx := range c.items {
		if c.items[idx].Producto().ID() == productoID {
			if !c.items[idx].Producto().Disponible(cantidad) {
				return ErrStockInsuficiente
			}
			c.items[idx].cantidad = cantidad
			return nil
		}
	}
	return ErrProductoNoExiste
}

func (c Carrito) CalcularTotal() float64 {
	total := 0.0
	for _, item := range c.items {
		total += item.Subtotal()
	}
	return total
}

func (c Carrito) Vacio() bool { return len(c.items) == 0 }
