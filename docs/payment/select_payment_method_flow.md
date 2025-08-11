    Redis TTL Watcher
           │
           ├──► Payment method selected before expire
           │         │
           │         ▼
           │     Call external service (SEPAY, VNPAY, ...)
           │         │
           │         ▼
           │     Set expired payment
           │         │
           │         ▼
           │    Return QR or url payment
           │
           └──► No payment method selected OR expired
                     │
                     ▼
              Publish Kafka Event: SeatReleased
                     │
                     ▼
             Unlock seat in Redis
             (Delete reservation key)
                     │
                     ▼
             Update DB if needed