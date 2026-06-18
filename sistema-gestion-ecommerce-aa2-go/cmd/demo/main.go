package main

import (
	"fmt"
	"log"

	"github.com/MatixProx/sistema-gestion-ecommerce-poo/internal/ecommerce"
)

func main() {
	categoria, err := ecommerce.NewCategoria("CAT-TEC", "Tecnologia")
	if err != nil {
		log.Fatal(err)
	}

	mouse, err := ecommerce.NewProducto("PROD-001", "Mouse gamer", 24.50, 8, "Mouse ergonomico RGB", *categoria)
	if err != nil {
		log.Fatal(err)
	}
	teclado, err := ecommerce.NewProducto("PROD-002", "Teclado mecanico", 59.90, 2, "Teclado switch azul", *categoria)
	if err != nil {
		log.Fatal(err)
	}

	inventario := ecommerce.NewInventario()
	_ = inventario.AgregarProducto(mouse)
	_ = inventario.AgregarProducto(teclado)

	cliente, err := ecommerce.NewCliente("CLI-001", "Matias Castro", "matias@example.com", "clave123", "Quito", "0999999999")
	if err != nil {
		log.Fatal(err)
	}

	carrito := ecommerce.NewCarrito("CAR-001")
	if err := carrito.AgregarProducto(*mouse, 2); err != nil {
		log.Fatal(err)
	}
	if err := carrito.AgregarProducto(*teclado, 1); err != nil {
		log.Fatal(err)
	}

	pedido, err := ecommerce.NewPedido("PED-001", *cliente, *carrito, inventario)
	if err != nil {
		log.Fatal(err)
	}

	metodo := ecommerce.NewTarjetaPago("4111111111111111", cliente.Nombre())
	pago, err := metodo.Procesar(pedido)
	if err != nil {
		log.Fatal(err)
	}

	reporteService := ecommerce.NewReporteService(inventario)
	resumen := reporteService.GenerarResumen([]ecommerce.Pedido{*pedido})

	fmt.Println("=== DEMOSTRACION DEL SISTEMA E-COMMERCE ===")
	fmt.Printf("Cliente: %s (%s)\n", cliente.Nombre(), cliente.Rol())
	fmt.Printf("Total carrito: $%.2f\n", carrito.CalcularTotal())
	fmt.Printf("Pedido: %s | Estado: %s | Total: $%.2f\n", pedido.ID(), pedido.Estado(), pedido.Total())
	fmt.Printf("Pago: %s | Metodo: %s | Estado: %s\n", pago.ID(), pago.Metodo(), pago.Estado())
	fmt.Printf("Ventas confirmadas: $%.2f\n", resumen.TotalVentas)
	fmt.Printf("Pedidos por estado: %v\n", resumen.PedidosPorEstado)
	fmt.Printf("Productos mas vendidos: %v\n", resumen.ProductosMasVendidos)
	fmt.Println("Productos bajo stock:")
	for _, producto := range resumen.ProductosBajoStock {
		fmt.Printf("- %s: %d unidades\n", producto.Nombre(), producto.Stock())
	}
}
