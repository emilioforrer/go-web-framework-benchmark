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

##  Endpoints

All services return the data embedded from [`internal/data/data.json`](internal/data/data.json).

## Results

Results are ordered by **Score**.

Your find the formula used to calculate the score inside the `CalculateScore` function in the [`internal/benchmark/score.go`](internal/benchmark/score.go) file.

### Last benchmark

The last benchmark was run on **2024-08-25**.

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
Name: go-fiber
Requests Per Second: 39268.95
Aborted Due To Deadline: 51
Successful Requests: 1178221
P99.999 Latency: 26.668903ms
P99 Latency: 9.583607ms
P90 Latency: 4.579236ms
Average Latency: 2.543312ms
Min Latency: 136.542µs
Max Latency: 52.144212ms
Duration: 30.005181136s
Score: 10.88
---------------------------------
Name: go-bun
Requests Per Second: 44507.62
Aborted Due To Deadline: 14
Successful Requests: 1335311
P99.999 Latency: 29.846318ms
P99 Latency: 9.587716ms
P90 Latency: 4.345456ms
Average Latency: 2.24339ms
Min Latency: 102.427µs
Max Latency: 85.537194ms
Duration: 30.002168252s
Score: 9.80
---------------------------------
Name: go-fuego
Requests Per Second: 37873.19
Aborted Due To Deadline: 98
Successful Requests: 1136363
P99.999 Latency: 31.618332ms
P99 Latency: 11.84593ms
P90 Latency: 5.337309ms
Average Latency: 2.636905ms
Min Latency: 117.094µs
Max Latency: 47.438987ms
Duration: 30.007003945s
Score: 9.50
---------------------------------
Name: go-echo
Requests Per Second: 34262.74
Aborted Due To Deadline: 18
Successful Requests: 1027971
P99.999 Latency: 32.9082ms
P99 Latency: 12.845333ms
P90 Latency: 6.027738ms
Average Latency: 2.915814ms
Min Latency: 115.03µs
Max Latency: 54.112172ms
Duration: 30.003115383s
Score: 8.01
---------------------------------
Name: go-std
Requests Per Second: 31501.20
Aborted Due To Deadline: 84
Successful Requests: 945100
P99.999 Latency: 31.99724ms
P99 Latency: 13.926233ms
P90 Latency: 6.495333ms
Average Latency: 3.170954ms
Min Latency: 131.141µs
Max Latency: 51.428198ms
Duration: 30.004693918s
Score: 7.50
---------------------------------
Name: go-gin
Requests Per Second: 32367.91
Aborted Due To Deadline: 12
Successful Requests: 971260
P99.999 Latency: 33.665851ms
P99 Latency: 13.648238ms
P90 Latency: 6.47537ms
Average Latency: 3.085249ms
Min Latency: 125.921µs
Max Latency: 72.83891ms
Duration: 30.007246449s
Score: 6.80
---------------------------------
Name: go-gorilla-mux
Requests Per Second: 32486.46
Aborted Due To Deadline: 40
Successful Requests: 974671
P99.999 Latency: 32.271609ms
P99 Latency: 13.627165ms
P90 Latency: 6.510632ms
Average Latency: 3.073762ms
Min Latency: 114.389µs
Max Latency: 82.708647ms
Duration: 30.003604414s
Score: 6.68
---------------------------------
Name: go-fr
Requests Per Second: 17874.42
Aborted Due To Deadline: 73
Successful Requests: 536273
P99.999 Latency: 66.751515ms
P99 Latency: 31.553563ms
P90 Latency: 14.72117ms
Average Latency: 5.588963ms
Min Latency: 191.538µs
Max Latency: 106.436579ms
Duration: 30.006337946s
Score: 2.03
---------------------------------

```


---

#### Docker status after benchmark (`docker stats --format json --no-stream`)


```json
{"BlockIO":"0B / 0B","CPUPerc":"0.04%","Container":"bed5bc5ed8eb","ID":"bed5bc5ed8eb","MemPerc":"0.05%","MemUsage":"9.207MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-fiber-1","NetIO":"882MB / 18GB","PIDs":"23"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"cd4b1901d45c","ID":"cd4b1901d45c","MemPerc":"0.07%","MemUsage":"13.12MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-std-1","NetIO":"811MB / 16.5GB","PIDs":"24"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"4f226550230b","ID":"4f226550230b","MemPerc":"0.07%","MemUsage":"13.7MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-echo-1","NetIO":"748MB / 15.2GB","PIDs":"26"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"cf5d3d854d48","ID":"cf5d3d854d48","MemPerc":"0.06%","MemUsage":"12.68MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-fuego-1","NetIO":"819MB / 16.7GB","PIDs":"26"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"073e01f1fde0","ID":"073e01f1fde0","MemPerc":"0.07%","MemUsage":"13.25MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-gorilla-mux-1","NetIO":"719MB / 14.7GB","PIDs":"23"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"d73be7a1a268","ID":"d73be7a1a268","MemPerc":"0.10%","MemUsage":"20.02MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-fr-1","NetIO":"458MB / 9.78GB","PIDs":"23"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"bd198651a2f2","ID":"bd198651a2f2","MemPerc":"0.06%","MemUsage":"11.03MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-bun-1","NetIO":"883MB / 18.1GB","PIDs":"25"}
{"BlockIO":"0B / 0B","CPUPerc":"0.00%","Container":"50be4b1ad37a","ID":"50be4b1ad37a","MemPerc":"0.06%","MemUsage":"11.6MiB / 19.54GiB","Name":"go-web-framework-benchmark-go-gin-1","NetIO":"729MB / 14.9GB","PIDs":"24"}
{"BlockIO":"0B / 0B","CPUPerc":"0.01%","Container":"c5743de0d7d1","ID":"c5743de0d7d1","MemPerc":"0.08%","MemUsage":"16.21MiB / 19.54GiB","Name":"go-web-framework-benchmark-db-1","NetIO":"3.24kB / 0B","PIDs":"7"}
```