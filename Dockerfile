# Use a imagem oficial do Golang como imagem base
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho no contêiner
COPY go.mod .

COPY go.sum .

# Baixe as dependências do módulo
RUN go mod download

# Copie o código-fonte Go para o diretório de trabalho no contêiner
COPY cmd/serverGrpc/main.go .
COPY pb/ /app/pb

# Compile o código Go
RUN go build -o app

# Exponha a porta em que o servidor gRPC está escutando (porta 9000)
EXPOSE 9000

# Comando para executar o aplicativo quando o contêiner for iniciado
CMD ["./app"]