# Hello world: an app showcasing load balancing on K8s

This project shows how to use load balancing with a K8s cluster.

Build this app as a Docker image using this command:
```shell
$ make build
```

Run this app on your host using this command:
```shell
$ make run
```

You should be able to access this app using your browser at
http://localhost:9000.

Deploy this app on K8s using this command:
```shell
$ kubectl apply -f helloworld.yml
```

All pods/services/pvc are created in the namespace `helloworld`.

A `LoadBalancer` object should be created with a public IP address, provided
this feature is implemented by your cloud provider.

This app is configurable. Set these environment variables in the K8s
descriptor (or your host when you're testing locally):

 - `SERVER_PORT`: set listening port (default is `9000`)
 - `MESSAGE`: set greetings message (default is `Hello world!`)
