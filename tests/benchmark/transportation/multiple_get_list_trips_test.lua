request = function()
    local path = string.format("/v1/transportation/search-trips?departure_date=2025-11-21&from_location=SG&to_location=HHZ")
    return wrk.format("GET", path)
  end
  