# Kafka Go Event Microservices

### 🛠 technologies
1. Golang
2. Kafka (Confluent Platform)
3. Primary DB: MySQL 8.0 (Account Data)
4. Analytics DB: ClickHouse (Event Logging)
5. Libraries: Sarama (Kafka), Viper (Config), GORM (MySQL), Fiber (Web)

### Quick Start
```go
docker-compose up --build
```

| Сервис           | Порт  | Описание                                     |
|------------------|-------|----------------------------------------------|
| Account Producer | 8000  | Accepts an HTTP request to create accounts |
| Kafka UI         | 8080  | Web interface for monitoring Kafka topics  |
| Kafka Broker     | 9092  | Internal port for Docker                   |
| Kafka Broker     | 29092 | External port (for local development)      |
| MySQL            | 3306  | Account storage                          |
| ClickHouse       | 9000  | Analytics warehouse                          |