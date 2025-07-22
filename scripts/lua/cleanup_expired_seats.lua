-- cleanup_expired_seats.lua
-- KEYS[1]: key of hash (e.g., seats:trip:1234)
-- ARGV[1]: current timestamp (Unix time)

local hashKey = KEYS[1]
local currentTime = tonumber(ARGV[1])

-- Get all seats
local seats = redis.call("HGETALL", hashKey)

for i = 1, #seats, 2 do
    local seatNumber = seats[i]
    local seatValue = seats[i+1]

    -- Only process seats that are held
    if string.sub(seatValue, 1, 5) == "held:" then
        local parts = {}
        for str in string.gmatch(seatValue, "([^:]+)") do
            table.insert(parts, str)
        end

        local expireAt = tonumber(parts[3]) -- held:USER:expireAt

        if expireAt ~= nil and expireAt <= currentTime then
            -- Expired ➝ update status to available
            redis.call("HSET", hashKey, seatNumber, "available")
        end
    end
end

return "OK"
