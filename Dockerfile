# Stage 1: Build Tailwind CSS
FROM node:latest AS build-tailwind
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npx tailwindcss -i ./internal/assets/tailwind.css -o ./internal/assets/dist/styles.css

# Stage 2: Build Go application
FROM golang:latest AS build-go
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=build-tailwind /app/internal/assets/dist ./internal/assets/dist
RUN go build -o /app/out cmd/web/main.go

# Stage 3: Final image
FROM golang:1.21-alpine
WORKDIR /app
COPY --from=build-go /app/out ./out
COPY --from=build-go /app/internal/assets/dist ./internal/assets/dist
COPY --from=build-go /app/data ./data
COPY --from=build-go /app/cmd/web/main.go ./cmd/web/main.go
COPY --from=build-go /app/internal ./internal
COPY --from=build-go /app/cmd ./cmd
COPY --from=build-go /app/go.mod /app/go.sum ./
COPY --from=build-go /app/.env ./

# Expose port
EXPOSE 3000

# Run the application
CMD ["./out"]