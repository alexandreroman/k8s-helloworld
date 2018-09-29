all: run

build:
	docker build -t alexandreroman/helloworld .

run: build
	docker run --rm -p 9000:9000 alexandreroman/helloworld /root/helloworld
