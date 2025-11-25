Go Cloud Scheduler



Финальный проект курса по Go. Это веб-сервер для планировщика задач с поддержкой повторяющихся событий, REST API, базой данных SQLite и аутентификацией.



Описание функционала



Приложение позволяет вести список задач. Основные возможности:



CRUD операций: Создание, чтение, обновление и удаление задач.



Отметка выполнения: При выполнении периодической задачи она автоматически переносится на следующую дату.



Правила повторения:



d <число> — каждые N дней.



y — каждый год.



w <дни> — по дням недели (например, w 1,5 — пн, пт).



m <дни> \[месяцы] — в определенные дни месяца (например, m 1,15, m -1 — последний день месяца).



Поиск: Фильтрация задач по строке поиска или дате.



Аутентификация: Доступ к API защищен JWT-токеном.



Выполненные задания со звёздочкой



В проекте реализованы все дополнительные задания повышенной сложности:



Расширенные правила повторения: Реализован расчет дат для правил w (дни недели) и m (дни месяца).



Поиск задач: Реализован обработчик параметра search (поиск по заголовку, комментарию и дате).



Аутентификация: Реализован вход по паролю (/api/signin), выдача JWT-токена и Middleware для защиты остальных ручек API.



Docker: Написан Dockerfile (multi-stage build) для сборки и запуска приложения в контейнере.



Запуск проекта локально

Требования



Go 1.22 или выше



SQLite3



Установка зависимостей

code

Bash

download

content\_copy

expand\_less

go mod download

Запуск сервера



Для работы аутентификации необходимо задать переменную окружения TODO\_PASSWORD.



Linux / macOS / Git Bash:



code

Bash

download

content\_copy

expand\_less

TODO\_PASSWORD=123 go run main.go



Windows (PowerShell):



code

Powershell

download

content\_copy

expand\_less

$env:TODO\_PASSWORD="123"; go run main.go



Windows (CMD):



code

Cmd

download

content\_copy

expand\_less

set TODO\_PASSWORD=123 \&\& go run main.go



После запуска откройте в браузере: http://localhost:7540



Пароль для входа: 123 (или тот, который вы указали в переменной окружения).



Запуск через Docker

1\. Сборка образа

code

Bash

download

content\_copy

expand\_less

docker build -t todo-app .

2\. Запуск контейнера



Команда ниже запускает сервер, пробрасывает порт 7540 и монтирует файл базы данных с хоста внутрь контейнера (чтобы данные сохранялись при перезапуске).



code

Bash

download

content\_copy

expand\_less

docker run -p 7540:7540 \\

&nbsp; -v "/$(pwd)/scheduler.db:/root/scheduler.db" \\

&nbsp; -e TODO\_PASSWORD=123 \\

&nbsp; todo-app



(Примечание: Для Windows, если $(pwd) не срабатывает, укажите полный путь к папке проекта).



Тестирование



Так как в проекте включена аутентификация, для запуска тестов необходимо получить токен.



Запустите сервер.



Выполните вход, чтобы получить токен. Это можно сделать через cURL:



code

Bash

download

content\_copy

expand\_less

curl -X POST http://localhost:7540/api/signin \\

&nbsp; -H "Content-Type: application/json" \\

&nbsp; -d '{"password": "123"}'



Скопируйте полученный токен (строка внутри поля "token").



Откройте файл tests/settings.go.



Вставьте токен в переменную Token:



code

Go

download

content\_copy

expand\_less

var Token = "ваш\_токен\_здесь"



Убедитесь, что поиск включен: var Search = true.



Запустите тесты:



code

Bash

download

content\_copy

expand\_less

go test ./tests

Переменные окружения

Переменная	Описание	Значение по умолчанию

TODO\_PORT	Порт веб-сервера	7540

TODO\_DBFILE	Путь к файлу базы данных	scheduler.db

TODO\_PASSWORD	Пароль для входа	(пусто, вход запрещен)

 
