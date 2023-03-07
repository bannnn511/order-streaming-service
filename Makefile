dev:
	docker compose -f docker-compose-dev.yml up -d
down:
	docker compose down

kafka-create-topic:
	#make kafka-create-topic topic=test
	docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic $(topic)
