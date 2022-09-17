## MetalLB 安裝

```js
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml
```

```js
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml
```

* 建立在 NS metallb-system 的 ConfigMap
```js
echo '
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: mlb1
      protocol: layer2
      addresses:
      - 192.168.61.220-192.168.61.230' | kubectl apply -f -
```

## MetalLB 檢查
```js
kubectl get pod -n metallb-system -o wide
```

## MetalLB 測試
* 建立 Pod
```js
echo '
apiVersion: apps/v1
kind: Deployment
metadata:
  name: d1.dep
spec:
  replicas: 2
  selector:
    matchLabels:
      app: d1.dep
  template:
    metadata:
      labels:
        app: d1.dep
    spec:
      containers:
      - name: app
        image: quay.io/flysangel/image:app.golang' | kubectl apply -f -
```

* 建立 service
```js
echo '
apiVersion: v1
kind: Service
metadata:
  name: d1
  annotations:
    metallb.universe.tf/address-pool: mlb1
spec:
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: d1.dep
  type: LoadBalancer' | kubectl apply -f -
```
* 檢視 pod、service
```js
kubectl get pod
```
```js
kubectl get service
```
* 測試
```js
curl -w "\n" http://<EXTERNAL-IP>
```
```js
curl -w "\n" http://<EXTERNAL-IP>/hostname
```
```js
curl -w "\n" http://<EXTERNAL-IP>/hostname
```

## 驗證
```js
sudo arping -c 4 <EXTERNAL-IP>
```

## 移除測試
```js
kubectl delete service s1
```
```js
kubectl delete deployment s1.dep
```

