-- ARGV[1]: TTL (milliseconds)
-- ARGV[2]: Booking ID

local ttl = tonumber(ARGV[1])
local lockValue = ARGV[2]

-- Lock each key
for i = 1, #KEYS do
    local ok = redis.call("SET", KEYS[i], lockValue, "PX", ttl, "NX")
    if not ok then
        -- Rollback previous locks
        for j = 1, i - 1 do
            redis.call("DEL", KEYS[j])
        end
        return false
    end
end

return true