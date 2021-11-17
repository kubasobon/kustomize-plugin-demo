# Kustomize plugin showcase

[DockerHub](https://hub.docker.com/repository/docker/wttyf00l/konfigure-plugin)

Proves we could use containerized tools (like `konfigure`) to modify our
kustomizations.

`main.go`, containerized as `docker.io/wttyf00l/konfigure-plugin:latest`, adds a ConfigMap in
process of modifying existing app resources.

**Requires kustomize v4**

Usage:
```
docker build -t docker.io/wttyf00l/konfigure-plugin:latest .
kustomize build --enable-alpha-plugins app/
```

Expected output:
```
apiVersion: v1
data:
  registry: docker.io
  ttl: 24h
kind: ConfigMap
metadata:
  creationTimestamp: null
  name: from-plugin
---
apiVersion: v1
kind: Service
metadata:
  name: the-service
spec:
  ports:
  - port: 8666
    protocol: TCP
    targetPort: 8080
  type: LoadBalancer
```
