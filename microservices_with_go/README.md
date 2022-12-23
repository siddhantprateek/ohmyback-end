Cross Compilation

```shell
env GOOS=linux GOARCH=386 go build main.go
```
It will not generate `.exe` file, but instead it will generate a binary file that is support for Linux environment


To build the docker image
```shell
docker-compose build 
```

To run the container
```shell
docker-compose up -d
```


```shell
kubectl app -f kubernetes/
kubectl get po,svc
kubectl delete deploy --all
```