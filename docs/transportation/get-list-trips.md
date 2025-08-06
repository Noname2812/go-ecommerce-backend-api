Client
  │
  ▼
Check local cache
  │
  ├──► Cache HIT ────────► Return response
  │
  └──► Cache MISS
           │
           ▼
        Check redis cache
            │
            ├──► Cache HIT ────────► Save local cache ────────► Return response
            │
            └──► Cache MISS
                        │
                        ▼
                        Try Acquire Lock (e.g., Redis lock trip_id:321)
                            │
                            ├──► Lock FAIL ──► Wait & Retry until local cache available or Redis cache → Return OR return server error
                            │
                            └──► Lock SUCCESS
                                    │
                                    ▼
                                Double Check redis cache  ──► Cache HIT ────────► Return response
                                    │
                                    ▼
                                Cache MISS
                                    │
                                    ▼
                                Get data from db
                                        |
                                        ▼
                                Save local cache 
                                    |
                                    ▼
                                Return response