#  Benchmarking the Transportation Service (Using WSL with Lua Scripts)

This document describes how to benchmark the `transportation-service` APIs using `wrk` with Lua scripting. The tests are executed in a WSL (Windows Subsystem for Linux) environment.

---

##  Requirements

- Run inside **WSL** or a Linux environment
- `wrk` installed with **Lua scripting support**
- The service must be running at: `http://172.26.16.1:8002`

---

##  `wrk` Parameters Explained

| Parameter | Description                            |
|----------|----------------------------------------|
| `-t8`    | Number of **threads** (8 threads)       |
| `-c200`  | Number of **concurrent connections** (200 connections) |
| `-d30s`  | **Duration** of the test (30 seconds)   |

---

##  Test: Get Trip Detail (Random ID from 1 to 3)

wrk -t8 -c200 -d30s --script=tests/benchmark/transportation/multiple_get_detail_trip_test.lua http://172.26.16.1:8002


---

##  Test: Get List Trip

wrk -t8 -c200 -d30s --script=tests/benchmark/transportation/multiple_get_list_trips_test.lua http://172.26.16.1:8002
