# vite-go-react
Building a go backend and react frontend with vite. Connect both and serve as a docker container.

## Quickstart
```sh
docker build -t go-server .
```

```sh
docker run --detach --name vite-react-client -p 3000:8080 -d go-server
```

open browser: http://localhost:3000
