FROM golang:1.11.2

COPY . /go/openshift-basic-identity-provider

ENV DB_PATH=/home/user.db

RUN cd /go/openshift-basic-identity-provider && go install --mod=vendor -v && cd /go && rm -rf /go/openshift-basic-identity-provider

EXPOSE 8080

CMD ["openshift-basic-identity-provider"]
