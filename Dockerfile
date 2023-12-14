# Используйте официальный образ Golang
FROM golang:1.16

# Установите рабочую директорию в GOPATH
WORKDIR /lanin/GolandProjects/TestCase

# Скопируйте исходный код приложения в контейнер
COPY . .

# Соберите Go-приложение
RUN go build -o calculator calculator.go

# Укажите команду, которая будет запускать ваше приложение
CMD ["calculator"]

# Убедитесь, что ваше приложение слушает порт, если это необходимо
EXPOSE 8080
