curl -X POST http://localhost:8080/budget \
-H "Content-Type: application/json" \
-d '{
      "month": "2025-01-01T00:00:00Z",
      "categoryId": 1,
      "budget": 450.00
    }'
