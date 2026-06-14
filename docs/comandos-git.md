# Comandos para subir el avance a GitHub

Ejecutar desde la carpeta del proyecto:

```bash
git status
git add .
git commit -m "Avance AA2: desarrollo del sistema e-commerce en Go"
git branch -M main
git remote -v
```

Si el repositorio remoto ya existe:

```bash
git push origin main
```

Si no esta configurado el remoto:

```bash
git remote add origin https://github.com/MatixProx/sistema-gestion-ecommerce-poo.git
git push -u origin main
```

Antes de subir, ejecutar:

```bash
go test ./...
go run ./cmd/demo
```
