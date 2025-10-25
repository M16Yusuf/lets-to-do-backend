# Lets to do backend

![badge golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![badge postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

## Entity-Relationship Diagram

Database ERD that used in this project is look like this. The project is simple to-do app for tracking to-do list a user where every todo item can only have 1 category, like shown in ERD that both table has relation one to many relationship.

On table categories is use for store data categories list, with data name and color for display on frontend. On Table This table stores individual to-do items created by users. Each todo can optionally belong to a category.

This design is provides a clear, normalized, and scalable schema for managing tasks and their categories. It ensures data integrity, supports flexible querying, and aligns well with common real-world use cases in task management systems.

![erd_database](/demo/erd-project.png)

## Code Splitting and Struckture Folder

<img src="/demo/Structure_folder.png" loading="lazy" width="370px" align="right">

**This project use monolithic architecture** where the application run on one server. Why using monolithc? because the application scale and complexity is small, monolithic is best option then using microservice and it will make application overengineering.

The code splited like shown beside, the code splited following **Separation of concerns** principle where code divided by their function, every section asign for one responsibilty. This approach makes software more maintainable, understandable, and easier to test by minimizing the overlap of responsibilities between different parts of the code.

Every directory/section have function like this:

- handlers, for code that responsibilty with handling request and response
- config, for code that responsibity with configuration that connect project with database or something similiar.
- middleware, for store middleware code
- models, for store struck/object database, body, or response.
- repositories, for code that responsibilty do something on database.
- roouters, for code that responsibity make API route.

## ⚙️ Installation From Clone Repository

1. Clone the project

```sh
$ https://github.com/M16Yusuf/lets-to-do-backend.git
```

2. Navigate to project directory

```sh
$ cd lets-to-do-backend
```

3. Install dependencies

```sh
$ go mod tidy
```

4. Setup your [environment](##-environment)

```bash
# database
DBUSER=<your_database_user>
DBPASS=<your_database_password>
DBNAME=<your_database_name
DBHOST=<your_database_host>
DBPORT=<your_database_port>
```

5. Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) for DB migration

6. Do the DB Migration

```sh
$ migrate -database YOUR_DATABASE_URL -path ./db/migrations up
```

or if u install Makefile run command

```sh
$ make migrate-createUp
```

7. Run the project

```sh
$ go run ./cmd/main.go
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

## Screenshot Demo API response

<div align="center">

<img src="/demo/demo_geteallorfiltertodo.png" loading="lazy" width="370px">
<img src="/demo/demo_create_todo.png" loading="lazy" width="370px"> </br>
<img src="/demo/demo_detailtodo.png" loading="lazy" width="370px">
<img src="/demo/demo_updatetodo.png" loading="lazy" width="370px"> </br>

<img src="/demo/demo_deletetodo.png" loading="lazy" width="370px">
<img src="/demo/demo_togglecomplete.png" loading="lazy" width="370px"></br>
<img src="/demo/demo_getCategory.png" loading="lazy" width="370px"> 
<img src="/demo/demo_createCategory.png" loading="lazy" width="370px"></br>

<img src="/demo/demo_updateCategory.png" loading="lazy" width="370px">
<img src="/demo/demo_deleteCategory.png" loading="lazy" width="370px">

</div>

## 📧 Contact Info

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/m16yusuf/)

## 🎯 Related Project

[https://github.com/M16Yusuf/lets-to-do-frontend](https://github.com/M16Yusuf/lets-to-do-frontend)
