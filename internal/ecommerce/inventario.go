package ecommerce

import "fmt"

// Inventario usa un map para localizar productos por ID en tiempo eficiente.
type Inventario struct {
	productos map[string]*Producto
}

func NewInventario() *Inventario {
	return &Inventario{productos: make(map[string]*Producto)}
}

func (i *Inventario) AgregarProducto(producto *Producto) error {
	if producto == nil {
		return fmt.Errorf("producto: %w", ErrCampoVacio)
	}
	i.productos[producto.ID()] = producto
	return nil
}

func (i Inventario) BuscarProducto(id string) (*Producto, error) {
	producto, ok := i.productos[id]
	if !ok {
		return nil, fmt.Errorf("id %s: %w", id, ErrProductoNoExiste)
	}
	return producto, nil
}

func (i Inventario) ListarProductos() []Producto {
	productos := make([]Producto, 0, len(i.productos))
	for _, producto := range i.productos {
		productos = append(productos, *producto)
	}
	return productos
}

func (i Inventario) ProductosBajoStock(minimo int) []Producto {
	bajoStock := []Producto{}
	for _, producto := range i.productos {
		if producto.Stock() <= minimo {
			bajoStock = append(bajoStock, *producto)
		}
	}
	return bajoStock
}
