
# Тестовое задание:
Web-Сервис сотрудников, сделанный на Golang
Сервис должен уметь:

1. Добавлять сотрудников, в ответ должен приходить Id добавленного
   сотрудника.

2. Удалять сотрудников по Id.

3. Выводить список сотрудников для указанной компании. Все доступные поля.

4. Выводить список сотрудников для указанного отдела компании. Все доступные поля.

5. Изменять сотрудника по его Id. Изменения должно быть только тех
   полей, которые указаны в запросе.

Модель сотрудника:
```
{
    Id int
    Name string
    Surname string
    Phone string
    CompanyId int
    Passport {
        Type string
        Number string
    }
    Department {
    Name string
    Phone string
    }
}
```
Все методы должны быть реализованы в виде HTTP запросов в формате JSON. БД: любая.

--- 

# Пояснения к решению

При добавлении или изменении данных сотрудника проверяются следующие критерии:
- введены ли имя и фамилия с заглавной буквы;
- начинается ли телефон с цифры 8, а также состоит ли он из 11 знаков;
- состоит ли Passport Number из 10 цифр (как в паспорте РФ или ВУ);
- значение CompanyId > 0;
- добавляется/изменяется ли сотрудник с уже существующем в таблице ` employees` Passport Number.

1. Для того чтобы добавить нового сотрудника, необходимо заполнить тело запроса, с указанием обязательных параметров.
```go
type Employee struct {
	Id         int        `json:"id" db:"id"`
	Name       string     `json:"name" db:"name" binding:"required"`
	Surname    string     `json:"surname" db:"surname" binding:"required"`
	Phone      string     `json:"phone" db:"employee_phone" binding:"required"`
	CompanyId  int        `json:"companyId" db:"company_id" binding:"required"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type Passport struct {
	Type   string `json:"type" db:"passport_type" binding:"required"`
	Number string `json:"number" db:"passport_number" binding:"required"`
}

type Department struct {
	Name  string `json:"name" db:"department_name" binding:"required"`
	Phone string `json:"phone" db:"department_phone" binding:"required"`
}
```

2. Для удаления сотрудка достаточно указать его id в теле запроса.

```go
type deleteEmployeeInput struct {
	Id int `json:"id" binding:"required"`
}
```

3. Для получения списка сотрудников по `companyId` достаточно указать `companyId` в параметрах запроса (`localhost:8000/employee/company/1`). Для добавления фильтра по отделу - необходимо добавить тело запроса.
```go
type getAllEmployeesInput struct {
	DepartmentName string `json:"departmentName"`
}
```

4. При изменении данных меняться будут только те - которые были указаны в теле запроса. При попытке изменить нумер паспорта на уже существующий - выскочит сообщение об ошибке и изменений не произойдет. Обязательным полем является `id` пользователя

```go
type UpdateEmployee struct {
	Id        int    `json:"id" db:"id" binding:"required"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	Phone     string `json:"phone" db:"employee_phone"`
	CompanyId int    `json:"companyId" db:"company_id"`
	Passport  struct {
		Type   string `json:"type" db:"passport_type"`
		Number string `json:"number" db:"passport_number"`
	} `json:"passport"`
	Department struct {
		Name  string `json:"name" db:"department_name"`
		Phone string `json:"phone" db:"department_phone"`
	} `json:"department"`
}
```

---

# Установка и запуск

```
git clone https://github.com/Rpqshka/smartway-employee-app.git
```

```
cd smartway-employee-app
```

```
docker-compose up --build
```

Вы можете использовать тестовую конфигурацию, которая находится в файле ```.env```, либо настроить сервис под себя и добавить этот файл в ```.gitignore```

---

# Работа с сервисом

После загрузки БД и запуска сервиса становятся доступны эндпоинты:
- POST   /employee          (Добавление нового сотрудника)
- DELETE /employee          (Удаление сотрудника по id)
- PUT    /employee          (Изменение данных сотрудника)
- GET    /employee/company/:id     (Получение списка сотрудников по id компании и департаменту)

[![postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/24093475-89c5d5f3-8253-4982-ad1a-cfe5fe45cfdc?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D24093475-89c5d5f3-8253-4982-ad1a-cfe5fe45cfdc%26entityType%3Dcollection%26workspaceId%3Db59b32b4-2180-4904-8005-bad090c92886)


## Примеры

1. POST `localhost:8000/employee` (Добавление нового сотрудника)

TEST INPUT:
```
{
    "name": "Aleksandr",
    "surname": "Golikov",
    "phone": "89653779090",
    "companyId": 1,
    "passport":{
        "type": "id",
        "number": "0123456789"
    },
    "department":{
       "name": "IT",
        "phone": "89651234567"
    }
}
```
TEST OUTPUT:
```
{
    "id": 1
}
```


2. DELETE `localhost:8000/employee` (Удаление сотрудника по id)

TEST INPUT:
```
{
    "id": 3
}
```
TEST OUTPUT:
```
{
    "id": 3
}
```

3. GET `localhost:8000/employee/company/1` (Получение списка сотрудников по id компании и департаменту)

TEST INPUT:
```
{
    
}
```
TEST OUTPUT:
```
{
    "employees": [
        {
            "id": 1,
            "name": "Aleksandr",
            "surname": "Golikov",
            "phone": "89653779090",
            "companyId": 1,
            "passport": {
                "type": "id",
                "number": "0123456789"
            },
            "department": {
                "name": "IT",
                "phone": "89651234567"
            }
        },
        {
            "id": 2,
            "name": "Aleksandr",
            "surname": "Golikov",
            "phone": "89653779090",
            "companyId": 1,
            "passport": {
                "type": "id",
                "number": "0223456789"
            },
            "department": {
                "name": "Design",
                "phone": "89651234567"
            }
        }
    ]
}
```

TEST INPUT:
```
{
    "departmentName": "Design"
}
```
TEST OUTPUT:
```
{
    "employees": [
        {
            "id": 2,
            "name": "Aleksandr",
            "surname": "Golikov",
            "phone": "89653779090",
            "companyId": 1,
            "passport": {
                "type": "id",
                "number": "0223456789"
            },
            "department": {
                "name": "Design",
                "phone": "89651234567"
            }
        }
    ]
}
```

4. PUT `localhost:8000/employee` (Изменение данных сотрудника)
   TEST INPUT:
```
{
    "id": 2,
    "passport":{
        "number":"0223456789"
    }
}
```
TEST OUTPUT:
```
{
    "message": "employee with passport number 0223456789 already exists with id 2"
}

```
TEST INPUT:
```
{
    "id": 2,
    "surname": "Pupkin",
    "passport":{
        "type": "Driver License",
        "number":"0323456789"
    }
}
```
TEST OUTPUT:
```
{
    "id": 2
}
```


ТАБЛИЦА `employees`

[![2024-06-30-192953772.png](https://i.postimg.cc/L655SXvY/2024-06-30-192953772.png)](https://postimg.cc/ppNRYPwR)
---

# Контакты
Telegram : @rpqshka

email: rpqshka@gmail.com