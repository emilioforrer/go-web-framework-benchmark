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
Requests Per Second: 49014.10
Aborted Due To Deadline: 427
Successful Requests: 2942476
P99.999 Latency: 116.943425ms
P99 Latency: 43.23109ms
P90 Latency: 29.826613ms
Average Latency: 20.330005ms
Min Latency: 187.9µs
Max Latency: 280.246975ms
Duration: 1m0.041964679s
Memory (MiB): 62.69
CPU (%): 0.02
Network (NetIO - MiB): 5.02
Disk (BlockIO - MiB): 0.00
Performance Score: 5.45
Resource Utilization Score: 6.91
Total Score: 5.98
---------------------------------
Name: go-fiber
Requests Per Second: 46920.83
Aborted Due To Deadline: 285
Successful Requests: 2816494
P99.999 Latency: 201.180092ms
P99 Latency: 61.370151ms
P90 Latency: 34.983864ms
Average Latency: 21.233266ms
Min Latency: 142.389µs
Max Latency: 290.552598ms
Duration: 1m0.032587759s
Memory (MiB): 26.07
CPU (%): 0.02
Network (NetIO - MiB): 4.73
Disk (BlockIO - MiB): 0.00
Performance Score: 3.65
Resource Utilization Score: 8.74
Total Score: 5.50
---------------------------------
Name: go-std-env-test-auto
Requests Per Second: 35757.28
Aborted Due To Deadline: 529
Successful Requests: 2146625
P99.999 Latency: 128.217075ms
P99 Latency: 58.142922ms
P90 Latency: 40.474456ms
Average Latency: 27.845254ms
Min Latency: 172.25µs
Max Latency: 315.070319ms
Duration: 1m0.048024835s
Memory (MiB): 32.72
CPU (%): 0.00
Network (NetIO - MiB): 3.31
Disk (BlockIO - MiB): 0.00
Performance Score: 3.48
Resource Utilization Score: 8.40
Total Score: 5.27
---------------------------------
Name: go-std-env-test-on
Requests Per Second: 34277.61
Aborted Due To Deadline: 703
Successful Requests: 2057279
P99.999 Latency: 138.281197ms
P99 Latency: 61.710264ms
P90 Latency: 42.70273ms
Average Latency: 29.091307ms
Min Latency: 131.688µs
Max Latency: 261.873368ms
Duration: 1m0.038675133s
Memory (MiB): 28.84
CPU (%): 0.00
Network (NetIO - MiB): 3.86
Disk (BlockIO - MiB): 0.00
Performance Score: 3.36
Resource Utilization Score: 8.60
Total Score: 5.26
---------------------------------
Name: go-chi
Requests Per Second: 38337.13
Aborted Due To Deadline: 537
Successful Requests: 2301049
P99.999 Latency: 257.28865ms
P99 Latency: 110.09385ms
P90 Latency: 49.175117ms
Average Latency: 25.994839ms
Min Latency: 152.392µs
Max Latency: 387.965259ms
Duration: 1m0.03542742s
Memory (MiB): 42.50
CPU (%): 0.00
Network (NetIO - MiB): 3.45
Disk (BlockIO - MiB): 0.00
Performance Score: 2.24
Resource Utilization Score: 7.91
Total Score: 4.30
---------------------------------
Name: go-fuego
Requests Per Second: 36514.60
Aborted Due To Deadline: 630
Successful Requests: 2191287
P99.999 Latency: 235.250977ms
P99 Latency: 116.018291ms
P90 Latency: 50.401752ms
Average Latency: 27.29842ms
Min Latency: 143.615µs
Max Latency: 360.454081ms
Duration: 1m0.028506417s
Memory (MiB): 44.37
CPU (%): 0.00
Network (NetIO - MiB): 3.84
Disk (BlockIO - MiB): 0.00
Performance Score: 2.27
Resource Utilization Score: 7.82
Total Score: 4.29
---------------------------------
Name: go-gorilla-mux
Requests Per Second: 39234.58
Aborted Due To Deadline: 602
Successful Requests: 2355264
P99.999 Latency: 256.701907ms
P99 Latency: 121.064154ms
P90 Latency: 53.625942ms
Average Latency: 25.390759ms
Min Latency: 110.832µs
Max Latency: 380.518715ms
Duration: 1m0.045652135s
Memory (MiB): 46.45
CPU (%): 0.00
Network (NetIO - MiB): 4.37
Disk (BlockIO - MiB): 0.00
Performance Score: 2.28
Resource Utilization Score: 7.72
Total Score: 4.26
---------------------------------
Name: go-gin
Requests Per Second: 35383.15
Aborted Due To Deadline: 584
Successful Requests: 2124452
P99.999 Latency: 250.541738ms
P99 Latency: 124.230529ms
P90 Latency: 55.375655ms
Average Latency: 28.134796ms
Min Latency: 145.499µs
Max Latency: 395.676116ms
Duration: 1m0.057846172s
Memory (MiB): 44.41
CPU (%): 0.00
Network (NetIO - MiB): 3.82
Disk (BlockIO - MiB): 0.00
Performance Score: 2.05
Resource Utilization Score: 7.82
Total Score: 4.15
---------------------------------
Name: go-echo
Requests Per Second: 35965.54
Aborted Due To Deadline: 362
Successful Requests: 2160219
P99.999 Latency: 245.936605ms
P99 Latency: 125.794806ms
P90 Latency: 53.637771ms
Average Latency: 27.699933ms
Min Latency: 134.368µs
Max Latency: 367.694084ms
Duration: 1m0.073646261s
Memory (MiB): 49.60
CPU (%): 0.00
Network (NetIO - MiB): 3.40
Disk (BlockIO - MiB): 0.00
Performance Score: 2.15
Resource Utilization Score: 7.55
Total Score: 4.11
---------------------------------
Name: go-std
Requests Per Second: 35647.39
Aborted Due To Deadline: 134
Successful Requests: 2140200
P99.999 Latency: 250.322655ms
P99 Latency: 128.82052ms
P90 Latency: 55.02267ms
Average Latency: 27.977318ms
Min Latency: 122.824µs
Max Latency: 430.874717ms
Duration: 1m0.041820449s
Memory (MiB): 46.45
CPU (%): 0.00
Network (NetIO - MiB): 4.44
Disk (BlockIO - MiB): 0.00
Performance Score: 2.02
Resource Utilization Score: 7.72
Total Score: 4.09
---------------------------------
Name: go-bun
Requests Per Second: 36622.09
Aborted Due To Deadline: 439
Successful Requests: 2198714
P99.999 Latency: 234.378226ms
P99 Latency: 115.308332ms
P90 Latency: 50.258369ms
Average Latency: 27.201824ms
Min Latency: 141.701µs
Max Latency: 295.011929ms
Duration: 1m0.049900626s
Memory (MiB): 73.16
CPU (%): 0.00
Network (NetIO - MiB): 6.44
Disk (BlockIO - MiB): 0.00
Performance Score: 2.39
Resource Utilization Score: 6.41
Total Score: 3.85
---------------------------------
Name: go-fr
Requests Per Second: 15903.35
Aborted Due To Deadline: 749
Successful Requests: 954588
P99.999 Latency: 1.013386898s
P99 Latency: 488.766112ms
P90 Latency: 220.487259ms
Average Latency: 62.549239ms
Min Latency: 137.052µs
Max Latency: 1.690303397s
Duration: 1m0.07143065s
Memory (MiB): 84.17
CPU (%): 0.00
Network (NetIO - MiB): 3.97
Disk (BlockIO - MiB): 0.00
Performance Score: 0.23
Resource Utilization Score: 5.83
Total Score: 2.27
---------------------------------
Name: go-std-env-test-off
Requests Per Second: 31355.46
Aborted Due To Deadline: 500
Successful Requests: 1882761
P99.999 Latency: 179.009328ms
P99 Latency: 71.846076ms
P90 Latency: 45.398098ms
Average Latency: 31.76856ms
Min Latency: 225.642µs
Max Latency: 309.115502ms
Duration: 1m0.061652014s
Memory (MiB): 3903.49
CPU (%): 0.00
Network (NetIO - MiB): 3.91
Disk (BlockIO - MiB): 0.00
Performance Score: 2.49
Resource Utilization Score: -185.14
Total Score: -65.74
---------------------------------

```

---
