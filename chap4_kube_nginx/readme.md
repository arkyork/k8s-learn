# Kubernetes上にNginxをデプロイする

kubernetes上にマニフェストで定義した内容を適用する

kubectl apply --filename chap4_kube_nginx/nginx.yaml --namespace default

kubectl runもあるが、一時的なPodを作成する場合に利用することが多い。

## 作成したPodの確認

port-forwardでローカルからアクセスする

kubectl port-forward -n default pod/nginx 9080:80
curl http://localhost:9080