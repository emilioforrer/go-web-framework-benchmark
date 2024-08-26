# Becnhmarking Go Web Frameworks

## External dependencies

- [Go Task](https://taskfile.dev/)
- [Oha](https://github.com/hatoo/oha)

## Install dependencies

If you are using [Mise](https://mise.jdx.dev/) to manage your dependencies, you can run the following commands:

```bash
# on the project root directory
mise install
mise shell
```

## Frameworks and libraries

- [Go Standard Library (go-std)](https://github.com/golang/go)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Echo](https://github.com/labstack/echo)
- [GoFiber](https://github.com/gofiber/fiber)
- [GoFuego](https://github.com/go-fuego/fuego)
- [GoFr](https://github.com/gofr-dev/gofr)
- [Go BunRouter](https://github.com/uptrace/bunrouter)

## Run the benchmark

Run the services:

```bash	
docker-compose up
```
In a separate terminal run the benchmark:

```bash	
task becnhmark
```

## Results

Results are ordered by **Score**.

Your find the formula used to calculate the score inside the `CalculateScore` function in the [`internal/benchmark/score.go`](internal/benchmark/score.go) file.

### Last benchmark

The last benchmark was run on **2024-08-25**.


---


