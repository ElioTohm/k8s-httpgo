- For testing the readiness of the pod we have:
  the test will run on build of the container and on deployment.
  - the unit tests inside the dockerfile
  - and the integration test which is a pod that test connection on startup
  - readiness prob

- build container using ```docker-compose up --build``` to test it locally 
- building a deployment should be ran as follows
  ```
  docker build -t <repo>/k8s-httpgo:0.0.1 .
  ```
- use `<repo>/k8s-httpgo:0.0.1` in the `values.yaml` file

- using helm v3 
  - in case you would want to see the k8s manifest simply run and apply the file 
  ```
  helm template charts/k8s-httpgo/ -f charts/k8s-httpgo/values.yaml > k8s.yaml
  ```
  or simply install the project using helm 3
  ```
  helm upgrade --install prometheus  charts/k8s-httpgo -f charts/k8s-httpgo/values.yaml
  ```
Docker image size is slim
```
REPOSITORY                                            TAG       IMAGE ID       CREATED          SIZE
k8s-httpgo_k8s-httpgo                                 latest    e86e90007c4d   48 seconds ago   36.7MB
```