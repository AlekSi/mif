# mif

Неофициальное REST API для книг издательства [МИФ](http://mann-ivanov-ferber.ru).

Текущая версия является пробной и нестабильной, в будущем возможны изменения, не являющиеся обратно совместимыми.
Если вы хотите застразовать себя от таких изменений, пожалуйста, используйте фиксированные [теги](releases).

Вклад в проект посредством создания issues и pull-requests приветствуется 🤓

## БД

Используется БД PostgreSQL.

При сборе данных выяснилось, что часть данных является неконситентной 
(например, некоторые книги относятся к несуществующим категориям). 
Поэтому некоторые внешние и уникальные ключи между таблицами на текущий момент не предполагаются.
При желании можно восстановить целостность данных и добавить внешние ключи.

### Структура данных

#### База данных

    createdb  mifbooks --encoding='utf-8' --locale=en_US.utf8 --template=template0;
    
#### Необходимые таблицы

    CREATE TABLE IF NOT EXISTS categories (
      id SERIAL PRIMARY KEY,
      title VARCHAR(256) NOT NULL,
      created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
      updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );
    
    CREATE TABLE IF NOT EXISTS books (
      id SERIAL PRIMARY KEY,
      mif_id INTEGER,
      category_id INTEGER,
      title VARCHAR(512),
      isbn VARCHAR(64),
      authors VARCHAR(512),
      url VARCHAR(512) UNIQUE NOT NULL,
      created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
      updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );
    
    CREATE TYPE volume_type AS ENUM ('paperbook', 'ebook', 'audiobook');
    CREATE TABLE IF NOT EXISTS volumes (
      id SERIAL PRIMARY KEY,
      book_id INTEGER NOT NULL REFERENCES books(id),
      type volume_type NOT NULL,
      created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
      updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
      
      UNIQUE (book_id, type)
    );
    
     CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      email VARCHAR(128) UNIQUE NOT NULL,
      token VARCHAR(128) UNIQUE NOT NULL,
      created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
      updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
     );
     
     CREATE TABLE IF NOT EXISTS library (
       id SERIAL PRIMARY KEY,
       user_id INTEGER NOT NULL REFERENCES users(id),
       volume_id INTEGER NOT NULL REFERENCES volumes(id),
       created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
       updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
       
       UNIQUE (user_id, volume_id)
     );
    
## Импорт данных

Импорт данных на текущий момент предполагается в полуручном режиме: все предлагаемые команды можно запускать или вручную,
или добавить их в расписание планировщика задач (cron) для регулярного запуска.

Подробности - в [data](data)

## Запуск REST API

Пример запуска сервера API:

    ENV host=127.0.0.1 port=80 \
    db_host=localhost db_port=5432 db_user=postgres db_pass=mysecretpassword db=mifbooks \
    go run main.go --debug
    
    
Параметры `host` и `port` отвечают за то, по какому адресу и порту слушать запросы.
Параметры с префиксом `db_` задают конфигурацию доступа к БД PostgreSQL: 
хост, порт, пользователь, пароль, название БД.

## Документация по REST API

### Все книги издательства
##### Список книг

`GET /api/v1/books` - получить список книг издательства.
В выдаче книги будут отсортированы по `id` в обратном порядке: от более новых к более старым.

__Опциональные параметры запроса:__

- `search` - поиск книги по части названия
- `author` - поиск книги по автору
- `mif_id` - поиск книги по идентификатору в МИФ
- `type` - тип "носителя" книги: `paperbook`, `ebook`, `audiobook`.

- `limit` - "пагинация": кол-во записей на странице (по умолчанию 10)
- `offset` - "пагинация": сдвиг записей (по умолчанию 0)

__Пример запроса:__ 

    curl -X GET http://mif.grahovac.me/api/v1/books?search=маркетинг&limit=50


### Информация о книге 

`GET /api/v1/books/:id` - получить конкретную книгу по ее идентификатору в БД

__Пример запроса:__ 

    curl -X GET http://mif.grahovac.me/api/v1/books/33
   
#### Личная библиотека

Запросы к личной библиотеке подразумевают аутентификацию пользователя.
Для аутентификации используется секретный токен, передаваемый через параметр запроса `token`.

В случае, если запрос предполагает наличие тела запроса, оно передается в формате JSON.

#### Список книг

`GET /api/v1/library` - получить список книг, имеющихся в собственной библиотеке.
В выдаче книги будут отсортированы по `id` в обратном порядке: добавленные последними будут в начале выдачи.

__Обязательные параметры запроса:__

__Опциональные параметры запроса:__

- `limit` - "пагинация": кол-во записей на странице (по умолчанию 10)
- `offset` - "пагинация": сдвиг записей (по умолчанию 0)

__Пример запроса:__ 


#### Добавить книгу в свою библиотеку

`POST /api/v1/library`

__Обязательные параметры запроса:__

- `token` - уникальный токен пользователя для аутентификации

__Обязательные параметры тела запроса:__

- `book_id` - идентификатор добавляемой книги
- `type` - "носитель" добавляемой книги: `paperbook`, `ebook`, `audiobook`.

__Пример запроса:__ 

    curl -X GET http://mif.grahovac.me/api/v1/library?token=6f9c1a78-36c7-4703-adb7-e0101e480c64&limit=50

#### Удалить книгу из своей библиотеки

`DELETE /api/v1/library`

__Обязательные параметры запроса:__

- `token` - уникальный токен пользователя для аутентификации

__Обязательные параметры тела запроса:__

- `book_id` - идентификатор добавляемой книги
- `type` - "носитель" добавляемой книги: `paperbook`, `ebook`, `audiobook`.

__Пример запроса:__ 

    curl -X DELETE http://mif.grahovac.me/api/v1/library?token=6f9c1a78-36c7-4703-adb7-e0101e480c64 \
    -H 'Content-Type: application/json' \
    -d '{"book_id": 33, "type": "paperbook"}'


## Road map

- Добавить клиента для API и консольное приложение для работы с ним
- Добавить swagger-документацию для API
- Рефакторинг кишков middleware (запросы к БД)
- Поддержка HTTPS
- Запуск импорта данных по расписанию
- Утилита для управления пользователями
