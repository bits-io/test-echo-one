# Gunakan image resmi Golang berbasis Alpine untuk ukuran lebih kecil
FROM golang:1.21.5-alpine

# Set working directory dalam container
WORKDIR /app

# Copy file go.mod dan go.sum terlebih dahulu (agar caching build lebih efisien)
COPY go.mod go.sum ./

# Download dependencies sebelum meng-copy semua file (mempercepat build)
RUN go mod download

# Copy semua source code ke dalam container
COPY . .

# Build aplikasi dengan output 'main'
RUN go build -o main ./cmd/main.go

# Expose port yang akan digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
