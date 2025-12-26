- マルチステージ
- Http server

```bash
docker build .\chap1_docker\goServer --tag go-server:1.0
```

```bash
docker run --rm --detach --publish 9000:8080 --name go-server go-server:1.0
```