# Fibonacci

#### Сбор Docker образа:
 `docker build -t rauf-test .`

####Запуск сервиса fibonacci в отдельном контейнере:
`docker run rauf-test`

###REST API

`http://localhost:3000/fibonacci?arg=9`

####Пример ответа:
`[{"Number":1,"Value":1},{"Number":2,"Value":1},
{"Number":3,"Value":2},{"Number":4,"Value":3},
{"Number":5,"Value":5},{"Number":6,"Value":8}
,{"Number":7,"Value":13},{"Number":8,"Value":21},
{"Number":9,"Value":34}]`


###gRPC
Proto file (см. service.proto)
`./api/v1/service.proto`
