IMAGE_NAME=pow
NETWORK_NAME=pow-network

build:
	docker build -t $(IMAGE_NAME) .

create_network:
	docker network create $(NETWORK_NAME)

run_server:
	docker run --rm --network $(NETWORK_NAME) --name server $(IMAGE_NAME)

run_client:
	docker run --rm --network $(NETWORK_NAME) --name client $(IMAGE_NAME) /app/client

clean:
	docker rmi $(IMAGE_NAME)
	docker network rm $(NETWORK_NAME)

