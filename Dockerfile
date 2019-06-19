FROM golang:1.11.2

COPY . .

ENV DB_PATH=/home/user.db

RUN go build --mod=vendor && go install -v

EXPOSE 8080

ENTRYPOINT ["openshift-basic-identity-provider"]
