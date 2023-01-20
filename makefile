BINARY_NAME=restaurant-app

build:
	go build -o ${BINARY_NAME} main.go
 
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}

docker-build:
	docker build -t pramodshenkar/restaurant-app:1.0.0.RELEASE .

docker-run:
	docker run -p 5000:5000 -it pramodshenkar/restaurant-app:1.0.0.RELEASE

