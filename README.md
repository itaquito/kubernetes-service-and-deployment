### Configuración kind

1. Crear un cluster de kind:
```bash
kind create cluster --name go-server
```

2. Construir la imagen de Docker:
```bash
docker build -t go-server:latest .
```

3. Cargar la imagen en kind:
```bash
kind load docker-image go-server:latest --name go-server
```

### Desplegar en Kubernetes

1. Crear el namespace:
```bash
kubectl apply -f namespace.yaml
```

2. Desplegar la aplicación:
```bash
kubectl apply -f deployment.yaml
```

3. Crear el servicio de load balancer:
```bash
kubectl apply -f service.yaml
```

Desplegar todo de una vez:
```bash
kubectl apply -f namespace.yaml -f deployment.yaml -f service.yaml
```

### Verificar el despliegue

Revisar todos los recursos en el namespace:
```bash
kubectl get all -n go-server
```

Revisar el estado de los pods:
```bash
kubectl get pods -n go-server
```

Revisar los detalles del servicio:
```bash
kubectl get service -n go-server
```

### Acceder al servicio

Como kind no tiene un proveedor de load balancer, usa port forward:
```bash
kubectl port-forward --address 0.0.0.0 -n go-server service/go-server-service 8080:8080
```

## Limpieza

Eliminar todos los recursos:
```bash
kubectl delete -f service.yaml -f deployment.yaml -f namespace.yaml
```

O eliminar todo el namespace:
```bash
kubectl delete namespace go-server
```

Eliminar el cluster de kind:
```bash
kind delete cluster --name go-server
```
