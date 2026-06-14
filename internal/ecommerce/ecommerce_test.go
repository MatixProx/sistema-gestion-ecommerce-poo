package ecommerce

import "testing"

func crearProductoPrueba(t *testing.T, id string, stock int) Producto {
	t.Helper()
	cat, err := NewCategoria("CAT", "General")
	if err != nil {
		t.Fatal(err)
	}
	producto, err := NewProducto(id, "Producto prueba", 10, stock, "Demo", *cat)
	if err != nil {
		t.Fatal(err)
	}
	return *producto
}

func TestProductoValidaPrecio(t *testing.T) {
	cat, _ := NewCategoria("CAT", "General")
	_, err := NewProducto("P1", "Producto", 0, 10, "", *cat)
	if err == nil {
		t.Fatal("se esperaba error por precio invalido")
	}
}

func TestCarritoCalculaTotal(t *testing.T) {
	producto := crearProductoPrueba(t, "P1", 10)
	carrito := NewCarrito("C1")
	if err := carrito.AgregarProducto(producto, 3); err != nil {
		t.Fatal(err)
	}
	if carrito.CalcularTotal() != 30 {
		t.Fatalf("total esperado 30, obtenido %.2f", carrito.CalcularTotal())
	}
}

func TestPedidoReduceStock(t *testing.T) {
	producto := crearProductoPrueba(t, "P1", 10)
	inventario := NewInventario()
	if err := inventario.AgregarProducto(&producto); err != nil {
		t.Fatal(err)
	}
	cliente, err := NewCliente("C1", "Ana", "ana@example.com", "clave123", "Quito", "099")
	if err != nil {
		t.Fatal(err)
	}
	carrito := NewCarrito("CAR1")
	if err := carrito.AgregarProducto(producto, 4); err != nil {
		t.Fatal(err)
	}
	_, err = NewPedido("PED1", *cliente, *carrito, inventario)
	if err != nil {
		t.Fatal(err)
	}
	actualizado, _ := inventario.BuscarProducto("P1")
	if actualizado.Stock() != 6 {
		t.Fatalf("stock esperado 6, obtenido %d", actualizado.Stock())
	}
}

func TestMetodoPagoPolimorfico(t *testing.T) {
	producto := crearProductoPrueba(t, "P1", 5)
	inventario := NewInventario()
	_ = inventario.AgregarProducto(&producto)
	cliente, _ := NewCliente("C1", "Luis", "luis@example.com", "clave123", "Loja", "098")
	carrito := NewCarrito("CAR1")
	_ = carrito.AgregarProducto(producto, 1)
	pedido, err := NewPedido("PED1", *cliente, *carrito, inventario)
	if err != nil {
		t.Fatal(err)
	}

	var metodo MetodoPago = NewTransferenciaPago("TRX-123")
	pago, err := metodo.Procesar(pedido)
	if err != nil {
		t.Fatal(err)
	}
	if pago.Estado() != "confirmado" || pedido.Estado() != EstadoPagado {
		t.Fatalf("pago o pedido no quedaron confirmados")
	}
}

func TestReporteConcurrente(t *testing.T) {
	producto := crearProductoPrueba(t, "P1", 10)
	inventario := NewInventario()
	_ = inventario.AgregarProducto(&producto)
	cliente, _ := NewCliente("C1", "Sofia", "sofia@example.com", "clave123", "Cuenca", "097")
	carrito := NewCarrito("CAR1")
	_ = carrito.AgregarProducto(producto, 2)
	pedido, _ := NewPedido("PED1", *cliente, *carrito, inventario)
	_ = pedido.CambiarEstado(EstadoPagado)

	reporte := NewReporteService(inventario).GenerarResumen([]Pedido{*pedido})
	if reporte.TotalVentas != 20 {
		t.Fatalf("ventas esperadas 20, obtenido %.2f", reporte.TotalVentas)
	}
}
