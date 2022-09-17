## Metrics Server 安裝

```js
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

```js
kubectl get pod -n kube-system
```
```js
export KUBE_EDITOR="nano"
```
```js
kubectl edit deploy metrics-server -n kube-system
```
```js
spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        - --kubelet-insecure-tls  <- 增加此行
        image: k8s.gcr.io/metrics-server/metrics-server:v0.6.1
```

## 驗證
```js
kubectl get pod -n kube-system
```
```js
kubectl top node
```


