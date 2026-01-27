kindを利用してKubernetesクラスターを作成する。

```
kind create cluster --image=kindest/node:v1.29.0
```
![alt text](cluster.png)

クラスターの確認
``` bash
kubectl cluster-info --context kind-kind
```

クラスターの削除
``` bash
kind delete cluster
```