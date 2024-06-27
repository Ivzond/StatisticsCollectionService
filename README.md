**1. Общие сведения**

   * Название проекта: Сервис сбора статистики
   * Описание проекта: Микросервис на golang для сбора статистики
   * Цель проекта: Создать сервис с возможностью хранения статистики в базе данных

**2. Требования к функциональности**

   Должно быть 4 api ручки:
   * GetOrderBook(exchange_name, pair string) ([]*DepthOrder, error)
   * SaveOrderBook(exchange_name, pair string, orderBook []*DepthOrder) error
   * GetOrderHistory(client *Client)  ([]*HistoryOrder, error)
   * SaveOrder(client *Client, order *HistoryOrder) error

**3. Технические требования**
   * Язык программирования: Go 1.22
   * Библиотеки и фреймворки: Любые на ваш выбор
   * Архитектура: REST API
   * База данных:  Лучше реализовать ClickHouse, но можно и  Postgres
   
**4. Нефункциональные требования**
   * Производительность: Время отклика сервера не более 200мс
   * Rps: до 200 на запись, до 100 на чтение

**5. Тестирование**
   * Тестирование: Unit-тесты для всех основных функций

**6. Требования к документации**
   * Документация кода: Комментарии к основным модулям и функциям
   * Пользовательская документация: Руководство пользователя

**7. Сроки выполнения**
   * Ожидаемые сроки: 1 неделя

**8. Детали реализации:**

*Таблица OrderBook*
```
id                  int64
exchange            string
pair                string
asks                []depthOrder
bids                []depthOrder

type DepthOrder struct{
Price       float64
BaseQty     float64
}
```

*Таблица Order_History*
```
client_name                 string
exchange_name   	    string
label		            string
pair  		            string
side    		    string
type                        string
base_qty                    float64
price                       float64
algorithm_name_placed       string
lowest_sell_prc             float64
highest_buy_prc             float64
commission_quote_qty        float64
time_placed                 time.time

type HistoryOrder struct{
client_name                 string
exchange_name   	    string
label		            string
pair  		            string
side    		    string
type                        string
base_qty                    float64
price                       float64
algorithm_name_placed       string
lowest_sell_prc             float64
highest_buy_prc             float64
commission_quote_qty        float64
time_placed                 time.time
}

type Client struct{
client_name                     string
exchange_name   	        string
label		                string
pair  		                string
}
```
**9. Отправка на проверку**

После написания сервиса, нужно загрузить его в гитхаб в публичный репо и скинуть ссылку на почту:  andrew@vortex.foundation

