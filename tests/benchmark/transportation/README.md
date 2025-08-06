<!-- run on wsl -->
wrk -t8 -c200 -d30s --script=tests/benchmark/transportation/multiple_detail_test.lua http://172.26.16.1:8002