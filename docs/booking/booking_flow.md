Client (Request Seat Booking)
   │
   ▼
Check seat status in Redis (trip_seat_status:321)
   │
   ├──► Seat NOT available ──► Return "Already booked"
   │
   └──► Seat available
           │
           ▼
   Try Acquire Lock (Redis lock seat_id:xyz)
           │
           ├──► Lock FAIL ──► Return "Already locked by another user"
           │
           └──► Lock SUCCESS
                   │
                   ▼
        Mark seat as "reserved" in Redis 
        (set TTL = booking_hold_time, e.g., 5 min)
                   │
                   ▼
        Publish Kafka Event: SeatReserved
                   │
                   ▼
        Return lock seat success