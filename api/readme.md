# Feed Time Calculation

Сервис расчёта времени подачи

`POST /feed-time`

Пример запроса:
```json
{
	"lat": 55.752992,
	"lng": 37.618333,
	"limit": 3
}
```

Пример ответа:
```json
{
  "error": "",
  "result": [
    {
      "car_id": 313,
      "time": 31
    },
    {
      "car_id": 114,
      "time": 31
    },
    {
      "car_id": 210,
      "time": 59
    }
  ]
}
```