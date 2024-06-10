# Используем официальный образ PostgreSQL
FROM postgres:16

# Устанавливаем переменные окружения для PostgreSQL
ENV POSTGRES_DB=mydatabase
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword

# Создаем volume для хранения данных
VOLUME ["/var/lib/postgresql/data"]
