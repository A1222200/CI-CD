## 部署 Landlord on K8S

* 環境準備
    * CNT.2022.v4 K8S 建置
    * 以下 K8S 完成部署
        1. MetalLB
        2. Ingress NGINX Controller
        3. Local Path Provisioner
* 用 CMD 視窗登入

## 部署前置作業
```js
kubectl create ns landlord
```
```js
kubectl create -n landlord configmap kuser-conf --from-file /home/bigred/.kube/config
```
* 使用課程準備的環境
```js
kubectl apply -f https://web.flymks.com/landlord/v1/landlord.yaml
```
* 確認部署狀態
```js
kubectl get all -n landlord
```

# 測試環境
```js
ssh -p 20000 tenant-0@192.168.61.221
```
* tenant-0:~/wk$
```js
kubectl get pod
```
```js
kubectl run a1 --image=quay.io/flysangel/nginx
```
```js
kubectl get pod -o wide
```
```js
curl 10.244.2.38
```
```js
kubectl delete pod a1
```
```js
kubectl get ns
```
```js
exit
```

## 檢查 Mariadb
```js
kubectl get pod -n landlord
```
---
NAME                       READY   STATUS    RESTARTS   AGE
gateway-759f9676cd-bk5n8   1/1     Running   0          56m
kuser-995f699bf-59stj      1/1     Running   0          56m
logger-7c68f4cf59-wdk7d    1/1     Running   0          56m
**mariadb-b8fdcc6c4-z9m7d**    1/1     Running   0          56m
tenant-0                   1/1     Running   0          56m
tenant-1                   1/1     Running   0          55m
tenant-2                   1/1     Running   0          55m

---
```js
kubectl exec -it -n landlord mariadb-b8fdcc6c4-z9m7d -- bash
```
* root@mariadb-b8fdcc6c4-z9m7d:/# 
```js
mysql -u root -p landlord
```
```js
show tables;
```
```js
select * from history;
```
```js
select * from inotify;
```

## Landlord 移除
```js
kubectl delete -f https://web.flymks.com/landlord/v1/landlord.yaml
```
```js
kubectl delete ns tenant-0 tenant-1 tenant-2
```

```js
kubectl delete csr tenant-0-csr tenant-1-csr tenant-2-csr
```

* 使用課程提供的實作包
```js
wget https://web.flymks.com/landlord/v1/sre_landlord.zip
```
```js
unzip sre_landlord.zip
```
```js
cd sre_landlord
```

## 設計需要的 Dockerfile
* 設計 Dockerfile.landlord.gateway
* 設計 Dockerfile.landlord.logger
* 設計 Dockerfile.landlord.tenant

## 重新部署
```js
kubectl create ns landlord
```
```js
kubectl create -n landlord configmap kuser-conf --from-file /home/bigred/.kube/config
```
```js
kubectl apply -f landlord.yaml
```
```js
kubectl get all -n landlord
```

