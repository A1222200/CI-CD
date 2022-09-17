## VMware 新增硬碟
> player -> Manage -> Virtual Machine Setting
>> Add -> Hard Disk -> SCSI
>>> Create a new virtual disk
>>>> Store virtual disk as a single file # 100GB


## 安裝 LVM
```js
sudo modprobe dm-mod
```
```js
echo dm-mod | sudo tee /etc/modules-load.d/dm.conf
```

```js
sudo apk update
```
```js
sudo apk add lvm2
```

## Rook Ceph 安裝
* Rook 將分佈式存儲系統轉變為自我管理、自我擴展、自我修復的存儲服務。它可以自動執行存儲管理員的任務：部署、引導、配置、供應、擴展、升級、遷移、災難恢復、監控和資源管理。

* Rook 使用 Kubernetes 平台的強大功能通過 Kubernetes Operator 為每個存儲提供商提供服務。

*  Ceph 是一種高度可擴展的分佈式存儲解決方案，提供對象、文件和塊存儲。在每個存儲節點上，您將找到Ceph 存儲對象的文件系統和Ceph OSD（對象存儲守護程序）進程。

#### 環境需求
* K8S 新增100GB 硬碟
* K8S 完成安裝 LVM
* K8S 重新啟動
    * 使用 lsblk -l 檢查硬碟是否新增完成
    * 使用 lvm version 檢查是否安裝 LVM
* CMD 視窗登入 "主機"

### 下載 rook-ceph
```js
mkdir rook-ceph
```
```js
git clone --single-branch --branch v1.9.2 https://github.com/rook/rook.git
```
```js
cp rook/deploy/examples/crds.yaml rook-ceph
```
```js
cp rook/deploy/examples/common.yaml rook-ceph
```
```js
cp rook/deploy/examples/operator.yaml rook-ceph
```
```js
cp rook/deploy/examples/cluster.yaml rook-ceph
```
```js
cp rook/deploy/examples/toolbox.yaml rook-ceph
```
```js
cp rook/deploy/examples/direct-mount.yaml rook-ceph
```
```js
rm -rf rook
```
```js
cd rook-ceph
```

* 安裝 rook-ceph operator
```js
sed -i 's|rook/ceph|quay.io/flysangel/ceph|g' operator.yaml
```
```js
kubectl create -f crds.yaml -f common.yaml -f operator.yaml
```

* 檢查 operator running
```js
kubectl get pod -n rook-ceph
```

* 設定 rook-ceph 只使用 sdb 這顆硬碟
```js
sed -i 's|useAllDevices: true|useAllDevices: false\n    deviceFilter: sdb|g' cluster.yaml
```
```js
kubectl create -f cluster.yaml
```

## Rook Ceph檢查
```js
kubectl apply -f toolbox.yaml
```
```js
export tools=$(kubectl get pod -n rook-ceph | grep tools | awk '{print $1}’)
```
```js
kubectl exec -it -n rook-ceph $tools -- bash
```
* 在 rook-reph 測試環境
```js
ceph status
```
    * cluster:
        id:     8d9db0fc-6c24-4df5-8d4f-5637548c4e48
        health: HEALTH_OK -> 代表有成功
```js
ceph osd status
```
```js
ceph df
```


