## Ingress Nginx 安裝
```js
wget -O - https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.1/deploy/static/provider/cloud/deploy.yaml | sed 's|name: nginx|name: ig1|g' | kubectl apply -f -
```
```js
kubectl get service -n ingress-nginx
```
```js
kubectl get ingressclass
```
## 測試
* 環境建置
```js
kubectl apply -f ingress.yaml
```

```js
kubectl get pod
```
```js
kubectl get service
```
* 環境建置
```js
kubectl apply -f ingress.yaml
```

```js
kubectl get ingress
```
```js
curl -w '\n' --resolve test.k8s.org:80:192.168.61.220 http://test.k8s.org/app1
```
```js
curl -w '\n' --resolve test.k8s.org:80:192.168.61.220 http://test.k8s.org/app2
```

## 移除測試
```js
kubectl delete ingress ig1
```
```js
kubectl delete service app1 app2
```
```js
kubectl delete pod app1 app2
```


