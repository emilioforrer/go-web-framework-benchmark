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
- [Go Chi](https://github.com/go-chi/chi)
- [Go Hertz](https://github.com/cloudwego/hertz)

## Run the benchmark

Run the services:

```bash	
docker-compose up
```
In a separate terminal run the benchmark:

```bash	
task becnhmark
```

###  Docker services endpoint responses

All services return the data embedded from [`internal/data/data.json`](internal/data/data.json).

### Command used to benchmark all services

```bash
oha -j --no-tui -n 120000 -c 1000 -p 500 -z 1m http://localhost:{port}
```

## Results

Results are ordered by **Score**.

Your find the formula used to calculate the score inside the `CalculateScore` function in the [`internal/benchmark/score.go`](internal/benchmark/score.go) file.

### Last benchmark

The last benchmark was run on **2024-10-5**.

---
Specs used (from `neofetch --backend off`):

```bash
OS: Ubuntu 22.04.1 LTS on Windows 10.0.22631 x86_64
Kernel: 5.15.153.1-microsoft-standard-WSL2
Uptime: 8 hours, 25 mins
Packages: 912 (dpkg), 2 (nix-user), 46 (nix-default), 6 (snap)
Shell: bash 5.2.26
Theme: Adwaita [GTK3]
Icons: Adwaita [GTK3]
Terminal: Windows Terminal
CPU: AMD Ryzen 9 6900HX with Radeon Graphics (8) @ 3.293GHz
GPU: Microsoft Corporation Basic Render Driver
Memory: 4619MiB / 20006MiB
```

---

#### Results (output)

```bash
Name: go-hertz
Requests Per Second: 36815.85
Aborted Due To Deadline: 124
Successful Requests: 2210226
P99.999 Latency: 173.921821ms
P99 Latency: 57.535842ms
P90 Latency: 39.602673ms
Average Latency: 27.065555ms
Min Latency: 310.199µs
Max Latency: 300.462612ms
Duration: 1m0.03800698s
Score: 3.09
---------------------------------
Name: go-std-env-test-off
Requests Per Second: 29602.54
Aborted Due To Deadline: 473
Successful Requests: 1777513
P99.999 Latency: 127.64361ms
P99 Latency: 68.25369ms
P90 Latency: 48.39304ms
Average Latency: 33.679369ms
Min Latency: 249.03µs
Max Latency: 241.037054ms
Duration: 1m0.061942831s
Score: 3.03
---------------------------------
Name: go-fiber
Requests Per Second: 34690.32
Aborted Due To Deadline: 328
Successful Requests: 2082268
P99.999 Latency: 167.591851ms
P99 Latency: 81.84998ms
P90 Latency: 47.831477ms
Average Latency: 28.739445ms
Min Latency: 211.323µs
Max Latency: 289.791513ms
Duration: 1m0.033916396s
Score: 2.88
---------------------------------
Name: go-std-env-test-auto
Requests Per Second: 29068.08
Aborted Due To Deadline: 487
Successful Requests: 1744510
P99.999 Latency: 157.70361ms
P99 Latency: 71.674126ms
P90 Latency: 49.592156ms
Average Latency: 34.294376ms
Min Latency: 234.242µs
Max Latency: 277.407669ms
Duration: 1m0.031384048s
Score: 2.54
---------------------------------
Name: go-std-env-test-on
Requests Per Second: 28612.80
Aborted Due To Deadline: 415
Successful Requests: 1717211
P99.999 Latency: 171.527517ms
P99 Latency: 73.99696ms
P90 Latency: 50.16427ms
Average Latency: 34.846574ms
Min Latency: 203.993µs
Max Latency: 368.126729ms
Duration: 1m0.029988797s
Score: 2.19
---------------------------------
Name: go-bun
Requests Per Second: 29474.78
Aborted Due To Deadline: 548
Successful Requests: 1770188
P99.999 Latency: 280.822164ms
P99 Latency: 138.798028ms
P90 Latency: 62.323629ms
Average Latency: 33.761665ms
Min Latency: 155.759µs
Max Latency: 356.394652ms
Duration: 1m0.076307863s
Score: 1.60
---------------------------------
Name: go-fuego
Requests Per Second: 28760.44
Aborted Due To Deadline: 204
Successful Requests: 1726180
P99.999 Latency: 321.594741ms
P99 Latency: 157.845282ms
P90 Latency: 62.845269ms
Average Latency: 34.674647ms
Min Latency: 200.242µs
Max Latency: 400.081363ms
Duration: 1m0.026338696s
Score: 1.38
---------------------------------
Name: go-chi
Requests Per Second: 28683.87
Aborted Due To Deadline: 387
Successful Requests: 1721937
P99.999 Latency: 322.075113ms
P99 Latency: 165.476555ms
P90 Latency: 65.842095ms
Average Latency: 34.752861ms
Min Latency: 157.874µs
Max Latency: 455.784047ms
Duration: 1m0.04504078s
Score: 1.33
---------------------------------
Name: go-gorilla-mux
Requests Per Second: 27802.30
Aborted Due To Deadline: 140
Successful Requests: 1669176
P99.999 Latency: 308.93981ms
P99 Latency: 177.078075ms
P90 Latency: 75.240385ms
Average Latency: 35.834482ms
Min Latency: 181.77µs
Max Latency: 570.148108ms
Duration: 1m0.042360913s
Score: 1.23
---------------------------------
Name: go-gin
Requests Per Second: 28212.43
Aborted Due To Deadline: 572
Successful Requests: 1693222
P99.999 Latency: 312.379736ms
P99 Latency: 161.177572ms
P90 Latency: 68.514333ms
Average Latency: 35.316033ms
Min Latency: 180.199µs
Max Latency: 679.429998ms
Duration: 1m0.037149827s
Score: 1.20
---------------------------------
Name: go-echo
Requests Per Second: 29620.06
Aborted Due To Deadline: 511
Successful Requests: 1777948
P99.999 Latency: 414.400695ms
P99 Latency: 156.677121ms
P90 Latency: 66.417089ms
Average Latency: 33.654044ms
Min Latency: 131.972µs
Max Latency: 587.87331ms
Duration: 1m0.042376055s
Score: 1.11
---------------------------------
Name: go-std
Requests Per Second: 27229.28
Aborted Due To Deadline: 626
Successful Requests: 1634805
P99.999 Latency: 441.443551ms
P99 Latency: 179.643816ms
P90 Latency: 72.825259ms
Average Latency: 36.567183ms
Min Latency: 166.652µs
Max Latency: 595.80428ms
Duration: 1m0.06147945s
Score: 0.96
---------------------------------
Name: go-fr
Requests Per Second: 13037.00
Aborted Due To Deadline: 919
Successful Requests: 782440
P99.999 Latency: 1.20804671s
P99 Latency: 604.45533ms
P90 Latency: 290.61296ms
Average Latency: 76.174679ms
Min Latency: 192.071µs
Max Latency: 1.653890207s
Duration: 1m0.087359437s
Score: 0.16
---------------------------------

```


---

#### Docker status after benchmark (`docker stats --format json --no-stream`)


```json
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"2b61dd9ef787","ID":"2b61dd9ef787","MemPerc":"0.66%","MemUsage":"26.88MiB / 4GiB","Name":"go-web-framework-benchmark-go-std-env-test-auto-1","NetIO":"2.77GB / 56.3GB","PIDs":"8"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"eaca13fb9bae","ID":"eaca13fb9bae","MemPerc":"0.75%","MemUsage":"30.54MiB / 4GiB","Name":"go-web-framework-benchmark-go-std-env-test-on-1","NetIO":"3.34GB / 67.9GB","PIDs":"8"}
{"BlockIO":"0B / 0B","CPUPerc":"0.01%","Container":"65a8642d6817","ID":"65a8642d6817","MemPerc":"95.30%","MemUsage":"3.812GiB / 4GiB","Name":"go-web-framework-benchmark-go-std-env-test-off-1","NetIO":"3.43GB / 69.8GB","PIDs":"9"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"577a4c8cd444","ID":"577a4c8cd444","MemPerc":"1.11%","MemUsage":"45.39MiB / 4GiB","Name":"go-web-framework-benchmark-go-gorilla-mux-1","NetIO":"3.77GB / 77GB","PIDs":"28"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"85a9324c73f7","ID":"85a9324c73f7","MemPerc":"2.68%","MemUsage":"109.8MiB / 4GiB","Name":"go-web-framework-benchmark-go-fr-1","NetIO":"1.55GB / 33.3GB","PIDs":"25"}
{"BlockIO":"0B / 0B","CPUPerc":"0.19%","Container":"ccdc3ab0fbfe","ID":"ccdc3ab0fbfe","MemPerc":"1.03%","MemUsage":"42.14MiB / 4GiB","Name":"go-web-framework-benchmark-go-chi-1","NetIO":"2.87GB / 58.7GB","PIDs":"30"}
{"BlockIO":"0B / 0B","CPUPerc":"0.01%","Container":"dccae7f8071d","ID":"dccae7f8071d","MemPerc":"1.47%","MemUsage":"60.39MiB / 4GiB","Name":"go-web-framework-benchmark-go-hertz-1","NetIO":"4.27GB / 85.9GB","PIDs":"19"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"63c99e69f142","ID":"63c99e69f142","MemPerc":"1.02%","MemUsage":"41.85MiB / 4GiB","Name":"go-web-framework-benchmark-go-bun-1","NetIO":"3.42GB / 70GB","PIDs":"28"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"cd0194d6976c","ID":"cd0194d6976c","MemPerc":"1.07%","MemUsage":"44MiB / 4GiB","Name":"go-web-framework-benchmark-go-echo-1","NetIO":"2.85GB / 58.2GB","PIDs":"28"}
{"BlockIO":"0B / 0B","CPUPerc":"0.02%","Container":"d08e350e031c","ID":"d08e350e031c","MemPerc":"0.63%","MemUsage":"25.97MiB / 4GiB","Name":"go-web-framework-benchmark-go-fiber-1","NetIO":"4.02GB / 81.6GB","PIDs":"31"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"6b57d4c58caa","ID":"6b57d4c58caa","MemPerc":"1.02%","MemUsage":"41.69MiB / 4GiB","Name":"go-web-framework-benchmark-go-fuego-1","NetIO":"3.28GB / 67.1GB","PIDs":"28"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"136ddc98ae7e","ID":"136ddc98ae7e","MemPerc":"1.12%","MemUsage":"45.98MiB / 4GiB","Name":"go-web-framework-benchmark-go-std-1","NetIO":"3.9GB / 79.5GB","PIDs":"29"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"f2b686330000","ID":"f2b686330000","MemPerc":"1.13%","MemUsage":"46.11MiB / 4GiB","Name":"go-web-framework-benchmark-go-gin-1","NetIO":"3.28GB / 67GB","PIDs":"30"}
```