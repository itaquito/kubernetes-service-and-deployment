## Despliegue en Kubernetes

### Para Usuarios de Kind

1. Crear un clúster de kind:
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

**O desplegar todo de una vez:**
```bash
kubectl apply -f namespace.yaml -f deployment.yaml -f service.yaml
```

### Verificar el Despliegue

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

### Acceder al Servicio

Como kind no tiene un proveedor de LoadBalancer, usa port-forward:
```bash
kubectl port-forward -n go-server service/go-server-service 8080:8080
```
Luego accede en: http://localhost:8080/api

## Limpieza

Eliminar todos los recursos:
```bash
kubectl delete -f service.yaml -f deployment.yaml -f namespace.yaml
```

O eliminar todo el namespace:
```bash
kubectl delete namespace go-server
```

Eliminar el clúster de kind:
```bash
kind delete cluster --name go-server
```
