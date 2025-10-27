### Configuración kind

1. Crear un cluster de kind con configuración de puertos:
```bash
kind create cluster --name go-server --config kind-config.yaml
```

2. Construir la imagen de Docker:
```bash
docker build -t go-server:latest .
```

3. Cargar la imagen en kind:
```bash
kind load docker-image go-server:latest --name go-server
```

4. Instalar MetalLB para habilitar LoadBalancer en kind:
```bash
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
```

Esperar a que MetalLB esté listo:
```bash
kubectl wait --namespace metallb-system --for=condition=ready pod --selector=app=metallb --timeout=90s
```

5. Configurar el rango de IPs de MetalLB:
```bash
kubectl apply -f metallb-config.yaml
```

6. Obtener la IP del nodo kind y actualizar service.yaml:
```bash
docker inspect go-server-control-plane --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
```
Actualiza el valor de `externalIPs` en service.yaml con esta IP (por defecto: 172.18.0.2).

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

El servicio LoadBalancer estará accesible en:
- **Desde el host local**: http://localhost:8080/api o http://127.0.0.1:8080/api
- **Desde máquinas externas en la red**: http://192.168.56.10:8080/api (o la IP externa del host)

El servicio LoadBalancer distribuirá automáticamente las peticiones entre los 2 pods (replicas) configurados.

**Cómo funciona**:
1. MetalLB asigna una IP del rango Docker al LoadBalancer
2. `externalIPs` hace el servicio accesible en la IP del nodo kind (172.18.0.2)
3. `extraPortMappings` en kind-config.yaml reenvía el tráfico del host:8080 → nodo:8080
4. `listenAddress: 0.0.0.0` permite acceso desde todas las interfaces del host

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
