# Build section of the docker file
FROM golang:latest as build

WORKDIR /app
COPY . .

RUN go mod tidy && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/linux/notifier
RUN chmod +x ./build/linux/notifier

# App and server section of the docker file
FROM ubuntu:latest as serve

RUN apt-get update
RUN apt-get install -y ca-certificates

ARG NOTIFIER_DISCORDWEBHOOK=https://discord.com/api/webhooks/
ARG USERNAME=nonroot
ARG USER_UID=1000
ARG USER_GID=$USER_UID

WORKDIR /app
COPY --from=build /app/build/linux/notifier /app/notifier

EXPOSE 3038

RUN addgroup --gid $USER_GID $USERNAME
RUN adduser --uid $USER_UID --gid $USER_GID --disabled-password --gecos "" $USERNAME
RUN chown $USERNAME:$USERNAME ./notifier

USER $USERNAME:$USERNAME
ENV NOTIFIER_DISCORDWEBHOOK=$NOTIFIER_DISCORDWEBHOOK
ENTRYPOINT [ "./notifier", "server" ]