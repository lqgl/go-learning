# RocketMQ Demo
This repository contains a small demo application showing how to use rocketmq with Go. 

## Installation and setup
First install the rocketmq with docker:
```
$ docker pull xuchengen/rocketmq:latest
$ docker volume create rocketmq_data
$ docker run -itd \
 --name=rocketmq \
 --hostname rocketmq \
 --restart=always \
 -p 8080:8080 \
 -p 9876:9876 \
 -p 10909:10909 \
 -p 10911:10911 \
 -p 10912:10912 \
 -v rocketmq_data:/home/app/data \
 -v /etc/localtime:/etc/localtime \
 -v /var/run/docker.sock:/var/run/docker.sock \
 --net=host \
 xuchengen/rocketmq:latest
```
[More in Here](https://hub.docker.com/r/xuchengen/rocketmq)

## Using the demo
To run the demo, you need to have Go installed on your machine. You can then clone the repository and run the following commands:
```
$ cd rocketmq-demo
$ go mod tidy
$ go run main.go
```

Then view the console output or view the [rocketmq web panel](http://127.0.0.1:9876) 

## License
This code is released under the [MIT License](https://opensource.org/license/mit/).

## Contributing
If you find a bug or have a suggestion for improvement, feel free to open an issue or submit a pull request. We welcome contributions from anyone!