## Local Path Provisioner

* Local Path Provisioner 為 Kubernetes 用戶提供了一種利用每個節點中的本地存儲的方法。

* 官方原始 yaml
```js
kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.22/deploy/local-path-storage.yaml
```
* 在後面能加入自己的 image
```js
kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.22/deploy/local-path-storage.yaml | sed's |rancher/local-path-provisioner|<quay.io/user/image>|g' |kubectl apply -f -
```
* 檢視成果
    * 非K8S物件為額外安裝，預設會自動刪除資料
```js
kubectl get storageclass
```

```js
kubectl apply -f pvc-nginx.yaml
```

```js
kubectl get pvc html-storage
```

* 建立 Pod 使用 PVC
```js
kubectl apply -f pod-nginx.yaml
```

* 查看執行在哪台主機
```js
kubectl get pod nginx -o wide
```

* 使用local-path的PVC產生POD，會自動產生對應的PV
```js
kubectl get pvc html-storage
```
```js
kubectl get pv 
```
```js
ls -al /opt/local-path-provision
```

## 刪除 Pod 檢查 PV PVC
* 刪除 pod、PVC
```js
kubectl delete -f pod-nginx.yaml
```
```js
kubectl delete -f pvc-nginx.yaml
```
* 檢查 PV 與執行主機路徑
```js
kubectl get pv
```
```js
ls -al /opt/local-path-provisioner
```
