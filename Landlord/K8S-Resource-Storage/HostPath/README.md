# HostPath 應用
* K8S 執行
```js
kubectl apply -f pod-hostpath.yaml
```
* 執行在哪台主機
```js
kubectl get pod hostpath -o wide
```

* 檢視 HostPath
```js
sudo tree /opt/pv/cache
```
* 在執行主機上
```js
cat /opt/pv/cache/time
```

* 刪除 HostPath
```js
kubectl delete -f pod-hostpath.yaml
```
* 在執行主機上
```js
sudo tree /opt/pv/cache
```

