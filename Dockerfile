FROM golang:1.16.0-buster as builder

WORKDIR /api

COPY ./api/ .

RUN ls -alF && pwd

RUN go build main.go auth.go error_checking.go func_agenda.go func_notes.go

EXPOSE 8000

CMD ["./main"]