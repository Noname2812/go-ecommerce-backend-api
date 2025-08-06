request = function()
  local id = math.random(1, 3)
  local path = string.format("/v1/transportation/trip-detail/%d", id)
  return wrk.format("GET", path)
end
