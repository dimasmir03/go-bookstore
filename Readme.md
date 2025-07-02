# 📚 Go Bookstore API

🔧 Учебный проект — простое REST API на Go для управления книгами с использованием Gin и Gorm.  

Проект разработан по статье: [Building a REST API with Golang using Gin and Gorm](https://blog.logrocket.com/rest-api-golang-gin-gorm/)

---

## 🧩 Стек технологий

- **Язык:** Go
- **Фреймворк:** [Gin](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/)
- **База данных:** SQLite

---

## 📁 Установка и запуск

```bash
go mod tidy
go run main.go
```
📘 Эндпоинты API
📖 Получить список книг
GET /books
Ответ:

```json
{
  "data": [
    { "id": 1, "title": "Start with Why", "author": "Simon Sinek" }
  ]
}
```
➕ Создать книгу
POST /books
Тело запроса:

```json
{
  "title": "Start with Why",
  "author": "Simon Sinek"
}
```
📘 Получить книгу по ID
GET /books/:id
Ответ:

```json
{
  "data": {
    "id": 1,
    "title": "Start with Why",
    "author": "Simon Sinek"
  }
}
```
✏️ Обновить книгу
PATCH /books/:id
Тело запроса (одно из):

```json
{ "title": "The Infinite Game" }
{ "author": "Charls Edwin" }
```
❌ Удалить книгу
DELETE /books/:id
Ответ:

```json
{ "data": true }
```
🧠 Чему научился
* Создавать REST API на Go

* Использовать маршруты и middleware в Gin

* Работать с SQLite через GORM

* Структурировать модель и обработчики

* Обрабатывать JSON-запросы и ответы

📦 Зависимости
```bash
github.com/gin-gonic/gin  
gorm.io/gorm  
gorm.io/driver/sqlite
```
📜 Лицензия
Проект учебный, предназначен для демонстрации базовых навыков разработки REST API на Go.
