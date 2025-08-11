Client
  │
  ▼
Check local cache (full response by trip_id)
  │
  ├──► Cache HIT ────────► Return response
  │
  └──► Cache MISS
           │
           ▼
    Try Acquire Lock (e.g., Redis lock trip_id:321)
           │
           ├──► Lock FAIL ──► Wait & Retry until local cache available → Return OR return server error
           │
           └──► Lock SUCCESS
                   │
                   ▼
              Double Check local cache  ──► Cache HIT ────────► Return response
                   │
                   ▼
              Cache MISS
                   │
                   ▼
        Run 2 Goroutines in parallel:
          ├─ Get Static Data (layout, bus info…) 
          │      ├─ Redis GET
          │      └─ If miss → DB query
          └─ Get Dynamic Data (seat status)
                 ├─ Redis GET (trip_seat_status:321)
                 └─ If miss → DB query
                   (seat locks)
                    |
                    ▼
              Merge static + dynamic
                    |
                    ▼
     Save local cache (TTL 3–5s, full merged response)
                  |
                  ▼
              Return merged response