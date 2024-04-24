# Тестовое задание

Оценка выполнения:

Уровень 1: Оценивается полнота реализации API, обработка ошибок и безопасность.

Уровень 2: Оценивается функциональность веб-приложения, включая интерфейс и взаимодействие с API.

Уровень 3: Оценивается создание десктоп-приложения, его интерфейс, работа с API и обработка ошибок.

## Уровень 1 Серверная часть

Необходимо разработать серверное приложение на языке программирования Golang. При запуске сервер должен загружать в структуру список пациентов из файла list_patients.json, который находится в каталоге data в корне приложения. Реализовать API с использованием пакета Gin.
API запросы:

GET запрос: Получить список пациентов (GetListPatients)

Описание: Передает клиенту JSON-список пациентов, загруженный из файла data/list_patients.json.

POST запрос: Добавить пациента (NewPatient)

Описание: Получает от клиента в теле запроса JSON одного пациента и добавляет его в структуру списка пациентов. Сервер генерирует новый GUID для пациента и сохраняет весь список на диск.

POST запрос: Редактировать пациента (EditPatient)

Описание: Получает от клиента в теле запроса JSON одного пациента с измененными данными по GUID. Находит нужного пациента по GUID и заменяет в структуре все значения на новые (кроме GUID) и сохраняет весь список на диск.

POST запрос: Удалить пациента (DelPatient)

Описание: Удаляет пациента по GUID и сохраняет весь список на диск.

Защита операций:
Для операций добавления удаления и редактирования рекомендуется использовать мьютексы или стек операций, чтобы обеспечить корректное выполнение операций, даже если несколько клиентов выполняют их одновременно для одного пользователя.


## Уровень 2 Клиентская Web часть

Реализовать веб-приложение с использованием React, HTML, CSS и JavaScript (или jQuery DataTables).

Таблица пациентов:

Отображает список пациентов.

Использует jQuery DataTables для удобного отображения и фильтрации данных.

Кнопки:

Добавить: Открывает окно с полями для добавления нового пациента.

Редактировать: Открывает окно с полями для редактирования выбранного пациента.

Удалить: Открывает окно подтверждения удаления.

Обновить:

Кнопка для обновления данных пациента.

Обработка ошибок:

Окно ошибок с кнопкой "OK" для отображения сообщений об ошибках.

## Уровень 3 Клиентская C# Avalonia десктоп приложение

Интерфейс:

Реализовать десктопное приложение на Framework Avalonia на аналогии с веб-версией.

Кнопки: Добавить, Редактировать, Удалить, Обновить.

Окна:

Окно с полями для добавления/редактирования.

Окно подтверждения удаления.

Окно ошибок.

Отправка запросов:

Взаимодействие с сервером через API запросы для добавления, редактирования и удаления пациентов.

## Пример файла list_patients.json

[
  {"fullname": "Иванов Петр Сергеевич", "birthday": "1980-03-15", "gender": 0, "guid": "e23fcee5-5424-4e20-b17d-51b6c4b5d328"},
  {"fullname": "Смирнова Анна Ивановна", "birthday": "1995-08-21", "gender": 1, "guid": "93af21d4-6a61-4a53-834c-7b4ee7b1b63e"},
  {"fullname": "Петров Алексей Владимирович", "birthday": "1989-11-02", "gender": 0, "guid": "f80e2f39-3d4d-49a4-8c3e-23b94a71dbb1"},
  {"fullname": "Козлов Владимир Николаевич", "birthday": "1982-07-10", "gender": 0, "guid": "bc92591e-7aaf-4f71-bb94-6c01e18c994c"},
  {"fullname": "Федорова Ольга Игоревна", "birthday": "1973-06-28", "gender": 1, "guid": "b2e754f3-8e4e-4fc2-b010-d4d5adca5380"},
  {"fullname": "Сидоров Игорь Андреевич", "birthday": "1997-09-05", "gender": 0, "guid": "d3d1bc24-e5cb-4870-aed1-0f1f3b3d1d5a"},
  {"fullname": "Григорьева Татьяна Васильевна", "birthday": "1984-12-30", "gender": 1, "guid": "c2952086-7ea2-4122-bc95-8e9f3c9f858e"},
  {"fullname": "Павлов Артем Михайлович", "birthday": "1991-04-18", "gender": 0, "guid": "a8c874f4-6bf0-4f4c-a9a5-71dbb3d0d77e"},
  {"fullname": "Медведева Наталья Петровна", "birthday": "1987-02-25", "gender": 1, "guid": "e5d2d083-184f-40a0-925e-5d1c0c0471ee"},
  {"fullname": "Новиков Сергей Дмитриевич", "birthday": "1993-10-08", "gender": 0, "guid": "e6a1ff4c-0325-44f3-a06e-784b6112da22"},
  {"fullname": "Кузнецова Елена Викторовна", "birthday": "1981-01-14", "gender": 1, "guid": "4a7c7c0a-72bf-4c8b-871d-38eb9d39ed13"},
  {"fullname": "Андреев Андрей Александрович", "birthday": "1998-06-03", "gender": 0, "guid": "bc5f4ea6-e76a-4ad3-82b4-2a01c6da4132"},
  {"fullname": "Борисова Ирина Валерьевна", "birthday": "1986-09-22", "gender": 1, "guid": "bcb0a5e2-eaca-4d8d-a004-cab48ab6da10"},
  {"fullname": "Сергеев Денис Анатольевич", "birthday": "1979-08-17", "gender": 0, "guid": "19c2d5b2-1600-4e7f-8dd9-870d6d8e9f04"},
  {"fullname": "Захарова Екатерина Игоревна", "birthday": "1990-05-11", "gender": 1, "guid": "b123af7e-e55c-46b8-8fcd-2c4c2da0e07f"},
  {"fullname": "Королев Владислав Сергеевич", "birthday": "1983-03-07", "gender": 0, "guid": "3c5d93e8-7fc1-4ad3-bb7e-7c24e4efcf66"},
  {"fullname": "Герасимова Ольга Алексеевна", "birthday": "1996-11-20", "gender": 1, "guid": "6348eef0-04d8-4803-bf86-d3f13ef9e4da"},
  {"fullname": "Константинов Максим Владимирович", "birthday": "1988-04-02", "gender": 0, "guid": "dd1d87b1-6bd8-4c62-9a2e-8dd5a077150d"},
  {"fullname": "Тарасова Анна Александровна", "birthday": "1994-01-19", "gender": 1, "guid": "4163c530-ff85-45ed-877e-b3a8bb8a78a2"},
  {"fullname": "Марков Александр Степанович", "birthday": "1975-07-25", "gender": 0, "guid": "aa50c991-653a-4db3-bbf5-e3e8d9cb4771"},
  {"fullname": "Савельева Вера Сергеевна", "birthday": "1986-08-12", "gender": 1, "guid": "25fb62a2-909d-427d-8aeb-3d7dcb75824d"},
  {"fullname": "Алексеев Игорь Дмитриевич", "birthday": "1992-12-05", "gender": 0, "guid": "96c906b4-731e-4a10-9382-16db4db443f0"},
  {"fullname": "Полякова Мария Андреевна", "birthday": "1985-03-18", "gender": 1, "guid": "2083e461-ae65-4b1f-8a7d-62e3e576d498"},
  {"fullname": "Никитин Дмитрий Сергеевич", "birthday": "1997-06-30", "gender": 0, "guid": "3da6f19c-062d-4815-879e-d78632dd3bb5"},
  {"fullname": "Миронова Екатерина Валерьевна", "birthday": "1989-09-14", "gender": 1, "guid": "f2a7627f-3e9c-4e05-99e2-82eb16e23e44"},
  {"fullname": "Соколов Александр Николаевич", "birthday": "1978-11-28", "gender": 0, "guid": "46bea979-6b07-4cd1-b3c8-03ac5ab1b8e7"},
  {"fullname": "Титова Анна Владимировна", "birthday": "1996-02-07", "gender": 1, "guid": "d352a0a1-b2a9-4971-9d91-8ed14e5c5c25"},
  {"fullname": "Белов Артем Станиславович", "birthday": "1984-05-23", "gender": 0, "guid": "7ab6f5b7-9be3-4ee0-94d4-9e1bc2e4c8ab"},
  {"fullname": "Полякова Елена Игоревна", "birthday": "1990-04-01", "gender": 1, "guid": "8897a174-e0a8-4bc9-839c-bb0f5c48c3f1"},
  {"fullname": "Кудрявцев Валентин Викторович", "birthday": "1983-10-15", "gender": 0, "guid": "23c49918-1014-4b20-b65b-8bfe4e7de6f0"},
  {"fullname": "Андреева Оксана Александровна", "birthday": "1992-08-19", "gender": 1, "guid": "0361784a-3c11-429a-bb86-cd48ab68bc65"}
]

