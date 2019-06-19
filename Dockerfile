FROM golang:1.11.2

COPY . .

ENV DB_PATH=/home/user.db

RUN cd openshift-basic-identity-provider && go install --mod=vendor -v && cd /go && rm -rf /go/openshift-basic-identity-provider

EXPOSE 8080

ENTRYPOINT ["openshift-basic-identity-provider"]
