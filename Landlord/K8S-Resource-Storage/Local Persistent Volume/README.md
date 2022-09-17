## 建立 Local PV
```js
kubectl apply -f pv-rwo-5m.yaml
```
```js
kubectl get pv pv-rwo-5m
```


## 建立 local PVC
```js
kubectl apply -f pvc-rwo-3m.yaml
```

## 檢查 PV 與 PVC 有無綁定
```js
kubectl get pvc pvc-rwo-3m
```

```js
kubectl get pv pv-rwo-5m
```


## 建立 Pod 使用 PVC
* K8S 執行
```js
kubectl apply -f pod-nginx-1.yaml
```
* K8S pod
```js
kubectl get pod nginx-1 -o wide
```

* 若發現建置失敗
```js
kubectl describe pod nginx-1
```
```js
kubectl taint nodes m1 node-role.kubernetes.io/master:NoSchedule-
```
```js
kubectl taint nodes m1 node-role.kubernetes.io/control-plane
```

* 到指定主機
```js
sudo mkdir -p /opt/pc/5m
```


# 建立 nginx 首頁
```js
kubectl exec -it nginx-1 -- bash
```
* 進入 執行環境 root@nginx-1:/#
```js
echo 'Hello github !' > /usr/share/nginx/html/index.html
```
```js
exit
```

## 驗證 nginx 首頁
```js
ssh w1 cat /opt/pv/5m/index.html
```

## 移除 Pod 與 PVC 重新建立
```js
kubectl delete -f pod-nginx-1.yaml
```
```js
kubectl delete -f pvc-rwo-3m.yaml
```
* 檢視 STATUS 是否處於 Released
```js
kubectl get pv pv-rwo-5m
```
* 重新建立
```js
kubectl apply -f pvc-rwo-3m.yaml
```
# 在執行狀態時修改 yaml 內容
* 將 kubernetes 修改器改為 nano
```js
export KUBE_EDITOR="nano"
```
```js
kubectl edit pv pv-rwo-5m
```
* 將以下內容刪除
'''
claimRef:
    apiVersion: v1
    kind: PersistentVolumeClaim
    name: pvc-rwo-3m
    namespace: default
    resourceVersion: "5082487"
    uid: 57bda012-bacb-4798-912f-c890c3b3496e
'''

* 刪除 Pod PV PVC
```js
kubectl delete -f pod-nginx-1.yaml
```
```js
kubectl delete -f pvc-rwo-3m.yaml
```
```js
kubectl delete -f pv-rwo-5m.yaml
```
* 在執行主機刪除路徑
```js
sudo rm -r /opt/pv/5m
```
