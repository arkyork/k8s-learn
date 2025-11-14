docker run --rm --detach --publish 9000:80 --name nginxserver nginx:1.25.3

# --rm コンテナを自動的に削除
# --detach -d バックグラウンドで実行
# --publish ホストのPCの9000番 -> 80番に転送
# 　http://localhost:9000/ に移動したらWelcome to nginx!
# dockerでは一つ一つレイヤーをDL