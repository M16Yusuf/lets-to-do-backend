# Lets to do backend

![badge golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![badge postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

## 🗝️ Environment

```bash
# database
DBUSER=<your_database_user>
DBPASS=<your_database_password>
DBNAME=<your_database_name
DBHOST=<your_database_host>
DBPORT=<your_database_port>
```

## 🚧 API Documentation

| Method | Endpoint                | Body                                                                                                         | query / path                            | Description                                       |
| ------ | ----------------------- | ------------------------------------------------------------------------------------------------------------ | --------------------------------------- | ------------------------------------------------- |
| GET    | /api/todos              |                                                                                                              | query=string, page=integer, sort=string | get all todos with pagenation and optional filter |
| POST   | /api/todos              | json{"title":string, "description":string, "priority":string, "due_date":datetime, "category_id":integer }   | id=integer                              | "create new todo"                                 |
| GET    | /api/todos/:id          |                                                                                                              | id=integer                              | "view specific todo"                              |
| PUT    | /api/todos/:id          | json{ "title":"string", "description":string, "priority":string, "due_date":datetime, "category_id":integer} | id=integer                              | "update specific todo"                            |
| DELETE | /api/todos/:id          |                                                                                                              | id=integer                              | "delete specific todo"                            |
| PATCH  | /api/todos/:id/complete |                                                                                                              | id=integer                              | "toggle completion status"                        |
| GET    | /api/categories         |                                                                                                              |                                         | "list all categories"                             |
| POST   | /api/categories         | json{ "name":string, "color":string }                                                                        |                                         | "create new category"                             |
| PUT    | /api/categories/:id     | json{ "name":string, "color":string }                                                                        | id=integer                              | "update category"                                 |
| DELETE | /api/categories/:id     |                                                                                                              | id=integer                              | "delete category"                                 |
