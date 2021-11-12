# Kustomize plugin showcase

Proves we could use containerized tools (like `konfigure`) to modify our
kustomizations.

`main.go`, containerized as `example.docker.io/functions/konfigure-plugin:latest`, adds a ConfigMap in
process of modifying existing app resources.

**Requires kustomize v4**

Usage:
```
docker build -t example.docker.io/functions/konfigure-plugin:latest .
kustomize build --enable-alpha-plugins app/
```

Expected output:
```
apiVersion: v1
data:
  value: important-data
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
