<!-- run on wsl with lua-->
<!-- detail trip route with random id 1 -> 3 -->
wrk -t8 -c200 -d30s --script=tests/benchmark/transportation/multiple_get_detail_trip_test.lua http://172.26.16.1:8002

<!-- get list trip random data in file lua -->
wrk -t8 -c200 -d30s --script=tests/benchmark/transportation/multiple_get_list_trips_test.lua http://172.26.16.1:8002