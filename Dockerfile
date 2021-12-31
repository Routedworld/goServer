# syntax=docker/dockerfile:1
FROM golang
WORKDIR /user/agallo/Code/kubeAcademy
COPY . .
RUN go build -o /server
CMD ["/server"]
