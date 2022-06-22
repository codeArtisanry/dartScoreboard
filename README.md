
# Dart-Scoreboard

Dart-Scoreboard is use for calculating and managing your dart game.

### Used technology
- **Front-end** : vue (v4.5.x), nuxt(v2.15.x)
- **Back-end** : golang (v1.17.x)
- **Database** : sqlite (v3)


### Run Locally (UI)

Clone the project

```bash
  git clone https://github.com/Improwised/darts-scoreboard.git
```
Go to the project web directory (after clone project)

```bash
  cd dart-scoreboard/web
```

Run project UI locally using nuxt command

```bash
# install dependencies
  npm install

# serve with hot reload at localhost:3000
  npm run dev

# for build app
  npm run build

# for build app run at localhost:3000
  npm start
```


Run project UI using docker command

```bash
# build docker images for app UI
  docker build . -t dart-app

# run docker build images
  docker run --add-host=host.docker.internal:host-gateway -p 5000:3000 dart-app
```
### Run Locally (Api)

Go to the project api directory (after clone project)

```bash
  cd darts-scoreboard/api
```

Run project api locally using go command

```bash
  go run main.go
```

Run project api using docker command

```bash
# build docker images for app api
  docker build --tag dart-api .

# run docker build images
  docker run --name=dartapi -it --mount type=bind,source="$(pwd)"/dart.db,target=/app/dart.db --publish 8585:8080 dart-api

# after first time use below command to run backend
  docker container start dartapi
```


_You will see the UI running on (with nuxt command)_ :

http://localhost:3000/

_You will see the UI running on (with docker command)_ :

http://127.0.0.1:5000/

_You will see the api running on (with go command)_ :

http://localhost:8080/

_You will see the UI running on (with docker command)_ :

http://localhost:8585/
