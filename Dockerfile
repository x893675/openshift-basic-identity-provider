FROM golang:1.11.2

COPY . /go/openshift-basic-identity-provider


ENV DB_PATH=/home/user.db
ENV SALT_KEY=1234567887654321

RUN cd /go/openshift-basic-identity-provider && mkdir -pv /etc/origin/master/custom_auth/ && cp certs/* /etc/origin/master/custom_auth/ && go install --mod=vendor -v && cd /go && rm -rf /go/openshift-basic-identity-provider

EXPOSE 8080

CMD ["openshift-basic-identity-provider"]
