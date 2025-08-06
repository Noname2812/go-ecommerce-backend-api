-- KEYS[1]: redis hash key (example: trip_seat_status:trip_id)
-- ARGV[1]: JSON array of seat list
-- ARGV[2]: TTL (seconds)

local seat_list = cjson.decode(ARGV[1])
local hash_key = KEYS[1]
local ttl = tonumber(ARGV[2])

for i, seat in ipairs(seat_list) do
    local seat_id = tostring(seat.seat_id)
    local seat_json = cjson.encode(seat)
    redis.call("HSET", hash_key, seat_id, seat_json)
end

if ttl and ttl > 0 then
    redis.call("EXPIRE", hash_key, ttl)
end

return true
