# Use a imagem oficial do Golang como imagem base
FROM golang:latest

# Instale o compilador protobuf
RUN apt-get update && apt-get -y install protobuf-compiler

# Instale o plugin protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho no contêiner
COPY go.mod .
COPY go.sum .

# Baixe as dependências do módulo
RUN go mod download

# Copie o arquivo .proto para o diretório de trabalho no contêiner
COPY helloworld.proto .

# Gere os arquivos Go a partir do arquivo .proto
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto

# Copie o código-fonte Go para o diretório de trabalho no contêiner
COPY grpc-server.go .

# Compile o código Go
RUN go build -o app

# Exponha a porta em que o servidor gRPC está escutando (porta 50051)
EXPOSE 50051

# Comando para executar o aplicativo quando o contêiner for iniciado
CMD ["./app"]