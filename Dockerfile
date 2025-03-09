# Gunakan image dasar Golang versi 1.21.6
FROM golang:1.21.6

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode
COPY . .

# Buat file .env dengan variabel environment yang dibutuhkan
RUN echo "DB_HOST={your env}" >> .env && \
    echo "DB_USER={your env}" >> .env && \
    echo "DB_PASSWORD={your env}" >> .env && \
    echo "DB_PORT={your env}" >> .env && \
    echo "DB_NAME={your env}" >> .env && \
    echo "HOST_ADDRESS = {your env}" >> .env && \
    echo "HOST_PORT = {your env}" >> .env && \
    echo "SALT={your env}" >> .env && \
    echo "LOG_PATH = {your env}" >> .env
# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
