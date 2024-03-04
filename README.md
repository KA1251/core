USAGE: create a txt file with your conf 

the struct of txt:

REDIS_ENABLED:

REDIS_HOST:

REDIS_PORT:

REDIS_PASSWORD: 

RABBITMQ_ENABLED:

RABBITMQ_HOST: 

RABBITMQ_PORT:

RABBITMQ_USERNAME: 

RABBITMQ_PASSWORD:

PROMETHEUS_ENABLED:

PROMETHEUS_HOST:

PROMETHEUS_PORT: 

KAFKA_ENABLED:

KAFKA_PORT:

KAFKA_HOST:

KAFKA_USERNAME:

KAFKA_PASSWORD:

SQL_ENABLED:

SQL_PORT:

SQL_USERNAME:

SQL_PASSWORD:

SQL_HOST:

SQL_DB:

SQL_DRIVER:

Если вы используете брокер/бд, то устанавливаем значение enabled "T", в остальных случаях игнорируем все ненужные поля.
Далее используем функцю: 

loadConfig(yourconf.txt)//загружаем ваш конфиг в переменные окружения

Создаем: 

var yourconfig core.Struct

var yourconnection core.ConnectionHandler

И юзаем функцию инициализации:

core.Initiallizing(&yourConfig,&yourConnection)

Должно работать)




