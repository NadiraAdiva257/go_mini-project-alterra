# pull image dari registry
FROM golang:1.20.1-alpine

# direktori u/ instruksi RUN, CMD, ENTRYPOINT, COPY, & ADD
WORKDIR /app

# copy file local ke workdir
COPY go.mod ./
COPY go.sum ./

# download go mod & dependencies saat building container
RUN go mod download

# copy semua file local ke workdir app
COPY . ./

# u/ build / compile aplikasi
RUN go build -o /mini-project

# port u/ container saat runtime
EXPOSE 8000

# command yg akan dieksekusi saat container dijalankan
CMD [ "/mini-project" ]