-- ARGV[1]: Booking ID
local lockValue = ARGV[1]

-- Release each key
for i = 1, #KEYS do
    -- Only delete if value == lockValue
    local v = redis.call("GET", KEYS[i])
    if v == lockValue then
        redis.call("DEL", KEYS[i])
    end
end

return true
