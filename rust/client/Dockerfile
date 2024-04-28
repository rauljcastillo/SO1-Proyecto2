# Usa la imagen oficial de Rust como base
FROM rust:latest as builder

# Crea un directorio de trabajo dentro del contenedor
WORKDIR /cliente

# Copia los archivos de tu aplicación al contenedor
COPY . .

# Compila la aplicación
RUN cargo build --release

# Crea una nueva imagen ligera
FROM ubuntu:latest

# Copia el binario compilado desde la imagen de compilación a la imagen final
COPY --from=builder /cliente/target/release/client /usr/local/bin/client

# Exponer el puerto que utiliza tu aplicación
EXPOSE 5000

# Comando para ejecutar tu aplicación cuando se inicie el contenedor
CMD ["client"]

# docker build -t jhonathantocay/app_server .
# docker run -p 8080:8080 jhonathantocay/app_server