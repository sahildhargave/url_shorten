# multistage docker build . This redices the size of the final docker image
# stage 1 to build the app

FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

# stage 2 Deploys the app built in stage 1

# Alpine is a lightweight Linux distribution known for its 
# small size, making its a popular choice for containerized 
# applications.
FROM alpine

RUN adduser -S -D -H -h /app appuser

# Creating a new user named appuser
# - S create a system user
# - D disable password login
# - H does not create user home directory

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000 

CMD [ "./main" ]
