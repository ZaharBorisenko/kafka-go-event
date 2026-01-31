# kafka-go-event

kafka-console-consumer --topic test-topic --bootstrap-server localhost:9092 --from-beginning

kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

kafka-topics --bootstrap-server localhost:9092 --list

kafka-console-consumer --topic test-topic --bootstrap-server localhost:9092 --group my-test-consumer-group1

localhost:8080 - kafka-ui

### topics account-service
1. `kafka-console-producer --topic OpenAccountEvent  --bootstrap-server localhost:9092`
2. `kafka-console-producer --topic DepositFundEvent  --bootstrap-server localhost:9092`
3. `kafka-console-producer --topic WithdrawFundEvent --bootstrap-server localhost:9092`
4. `kafka-console-producer --topic CloseAccountEvent --bootstrap-server localhost:9092`


### taskfile
# Запуск всех компонентов
task up

##### Только Docker
task docker

##### Только producer
task producer

##### Фоновый режим
task dev

##### Остановка
task down

