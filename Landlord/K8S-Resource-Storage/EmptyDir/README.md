## EmptyDir 應用
* K8S 執行
```js
kubectl apply -f pod-emptydir.yaml
```
* 執行在哪台主機
```js
kubectl get pod meptydir -o wide
```

```js
kubectl logs emptydir -c c2
```
```js
sudo tree /var/lib/kubelet/pods/ | grep -B 10 cache
```

* 刪除 EmptyDir
```js
kubectl delete -f pod-emptydir.yaml
```
* 檢查刪除結果
```js
sudo tree /var/lib/kubelet/pods/ | grep -B 10 cache
```




