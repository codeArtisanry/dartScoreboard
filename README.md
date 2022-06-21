# darts-scoreboard

Dart Scoring application with Nuxt js and Golang To record Score When Playing Dart.

## Run Using Docker

Clone Project

```bash
  git clone https://github.com/Improwised/darts-scoreboard.git
```

Go to the project directory

```bash
  cd darts-scoreboard
```

### Start Backend (_Make sure you have docker installed locally_)

go to backend API directory

```bash
  cd api
```

start backend Docker container first time with following commands

```bash
  docker build --tag dart-api .
```

```bash
  docker run --name=dartapi -it --mount type=bind,source="$(pwd)"/dart.db,target=/app/dart.db --publish 8585:8080 dart-api
```

After first time use below command to run backend

```bash
  docker container start dartapi
```

## Run Locally

```bash
  go run main.go
```
