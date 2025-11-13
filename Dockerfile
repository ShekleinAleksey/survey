
# Этап 1: Сборка приложения
FROM golang:alpine as builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

# Собираем приложение (флаг -ldflags уменьшает размер бинарника)
RUN go build -ldflags="-s -w" -o survey ./cmd/main.go

# Этап 2: Запуск приложения
FROM alpine

# Устанавливаем клиент PostgreSQL 
RUN apk add --no-cache postgresql-client

# Копируем только собранный бинарник
COPY --from=builder /go/survey /app/survey

WORKDIR /app

CMD [ "./survey" ]
