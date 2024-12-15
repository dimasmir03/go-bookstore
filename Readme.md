Building a REST API with Golang using Gin and Gorm
Это небольшой проект для обучения.

Статья по которой делался проект: 
https://blog.logrocket.com/rest-api-golang-gin-gorm/

Golang+REST-API+GIN+Gorm+SQLite

Для проекта использовались пакеты 
github.com/gin-gonic/gin 
gorm.io/gorm
gorm.io/driver/sqlite

Проект реализует CRUD взаимодействие с базой данных SQlite

Прослушиваемые Endpoits с помощью фраемворка Gin:
GET "/books" - получение всех книг из бд
POST "/books" - создание новой книги
GET "/books/:id"  - получение книги по id
PATCH "/books/:id" - обновление книги по id
DELETE "/books/:id" - удаление книги по id

GET "/books" 
получения списка всех книг из бд books.db
формат ответа JSON
{
  "data": []
}


POST "/books"
создание новой книги в бд books.db
формат запроса JSON
{
  "title": "Start with Why",
  "author": "Simon Sinek"
}
формат ответа JSON
{
  "data": {
    "id": 1,
    "title": "Start with Why",
    "author": "Simon Sinek"
  }
}


GET "/books/:id"
получения по id книги из бд books.db
формат ответа JSON
{
  "data": {
    "id": 1,
    "title": "Start with Why",
    "author": "Simon Sinek"
  }
}


PATCH "/books/:id"
обновление по id книги в бд books.db
формат запроса JSON
{
  "title": "The Infinite Game"
}
{
  "author": "Charls Edwin"
}
формат ответа JSON
{
  "data": {
    "id": 1,
    "title": "The Infinite Game",
    "author": "Simon Sinek"
  }
}
{
  "data": {
    "id": 1,
    "title": "The Infinite Game",
    "author": "Charls Edwin"
  }
}


DELETE "/books/:id"
удаление по id книги в бд books.db
формат ответа JSON
{
  "data": true
}
