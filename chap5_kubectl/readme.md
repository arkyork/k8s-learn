# Kubernetes上にpodをデプロイする
kubectl apply --filename chap4_kube_nginx/myserver.yaml --namespace default

kubectl run myserver2 --image blux2/hello-server:2.4 --namespace default  

## 作成したリソースの確認

kubectl get pods -n default

kubectl describe pod myserver -n default

## outputオプション

outputオプションでは割り当てられれたipアドレスやnode名など詳細情報を確認できる。

kubectl get pod --output wide --namespace default

yamlも指定可能。

kubectl get pod myserver -n default -o yaml > chap5_kubectl/myserver_pod.yaml

diffで差分を確認。

diff chap5_kubectl/myserver_pod.yaml chap4_kube_nginx/myserver.yaml

## jsonpathオプション

kubectl get pod myserver --output jsonpath='{.spec.containers[].image}' -n default

## ログレベル

rest apiを確認できる。

kubectl get pod myserver -v=7 --namespace default

## debug用コンテナの立ち上げ

kubectl debug --stdin --tty myserver --image=curlimages/curl:8.4.0 --target=hello-server --namespace default -- sh

## port-forward
kubectl port-forward nginx 5555:80 --namespace default  
kubectl port-forward myserver 5555:8080 --namespace default  

## 