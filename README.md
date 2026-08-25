# Task Tracker (todosher)

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![Platforms](https://img.shields.io/badge/Platforms-Linux%20%7C%20macOS%20%7C%20Windows-blue?style=for-the-badge)](https://github.com/ВАШ_НИК/НАЗВАНИЕ_РЕПО/releases)
[![Release](https://img.shields.io/github/v/release/sanhasan/task-tracker?style=for-the-badge&color=green)](https://github.com/ВАШ_НИК/НАЗВАНИЕ_РЕПО/releases/latest)

Простой и удобный трекер задач с интерфейсом командной строки (CLI), который помогает отслеживать, что нужно сделать, что уже в процессе и что выполнено.

Пользователь может:

* Добавлять, обновлять и удалять задачи

* Отмечать прогресс по задачам

* Выводить список задач

* Выводить список сделанных задач

* Выводить список начатых задач

* Выводить список не начатых задач

## Скачать готовое приложение

Вам не нужно собирать проект самостоятельно. Скачайте готовый исполняемый файл для вашей системы из раздела [Releases](https://github.com/sanhasan/task-tracker/releases/latest) или по прямым ссылкам ниже:

| Операционная система | Архитектура | Прямая ссылка для скачивания                                                                                               |
| :--- | :--- |:---------------------------------------------------------------------------------------------------------------------------|
| **Linux** | AMD64 (x86_64) | [todosher-linux-amd64](https://github.com/sanhasan/task-tracker/releases/latest/download/todosher-linux-amd64)             |
| **Linux** | ARM64 | [todosher-linux-arm64](https://github.com/sanhasan/task-tracker/releases/latest/download/todosher-linux-arm64)             |
| **macOS** | Intel (x86_64) | [todosher-darwin-amd64](https://github.com/sanhasan/task-tracker/releases/latest/download/todosher-darwin-amd64)           |
| **macOS** | Apple Silicon (M1/M2/M3) | [todosher-darwin-arm64](https://github.com/sanhasan/task-tracker/releases/latest/download/todosher-darwin-arm64)           |
| **Windows** | AMD64 (x86_64) | [todosher-windows-amd64.exe](https://github.com/sanhasan/task-tracker/releases/latest/download/todosher-windows-amd64.exe) |

## Список технологий

* **Golang**
    * `os` (работа с файловой системой, аргументами и переменными окружения)
    * `encoding/json` (сериализация и десериализация данных)
* **Git / Github** (контроль версий и хостинг кода)
* **Clean Architecture** (четкое разделение ответственности на слои: `config`, `controller`, `usecase`, `repository`)
* **Bash** (скрипт для автоматической кросс-компиляции под разные ОС и архитектуры)

## Будущее развитие проекта

- [ ] Добавить логгер
- [ ] Добавить тесты
    - [ ] Unit-тесты
    - [ ] Integration-тесты
    - [ ] E2E-тесты
- [ ] Добавить документацию кода (Godoc)
- [ ] Расширить функционал и добавить поддержку различных флагов для операций
- [ ] Написать отчёт о собственном использовании в повседневности

## Как использовать

### 1. Установка и подготовка

1. Скачайте исполняемый файл для вашей операционной системы (из раздела Releases или соберите его самостоятельно с помощью скрипта `./build.sh`).
2. Сделайте файл исполняемым (для Linux/macOS):
   ```bash
   chmod +x todosher
   ```
3. **(Рекомендуется)** Переместите файл в директорию, которая находится в вашем системном `$PATH`, чтобы запускать команду `todosher` из любой папки в терминале:
   ```bash
   # Для Linux/macOS (в системную папку, потребуются права администратора)
   sudo mv todosher /usr/local/bin/todosher
   
   # ИЛИ в локальную директорию пользователя (без прав суперпользователя)
   mkdir -p ~/.local/bin
   mv todosher ~/.local/bin/todosher
   ```

### 2. Основные команды

Приложение управляется через аргументы командной строки. Вот список доступных команд:

| Команда                 | Описание | Пример |
|:------------------------| :--- | :--- |
| `help`                  | Показать справку по всем командам | `todosher help` |
| `add "<текст>"`         | Добавить новую задачу | `todosher add "Изучить Clean Architecture"` |
| `list`                  | Вывести список **всех** задач | `todosher list` |
| `list todo`             | Вывести только **не начатые** задачи | `todosher list todo` |
| `list in-progress`      | Вывести задачи, которые **в процессе** | `todosher list in_progress` |
| `list done`             | Вывести только **выполненные** задачи | `todosher list done` |
| `start <id>`            | Перевести задачу в статус "в процессе" | `todosher start 1` |
| `finish <id>`           | Перевести задачу в статус "выполнено" | `todosher finish 1` |
| `update <id> "<текст>"` | Обновить описание задачи | `todosher update 1 "Изучить Go и Clean Arch"` |
| `delete <id>`           | Удалить задачу по её ID | `todosher delete 1` |

> **Совет:** ID задачи всегда отображается при выводе команды `list`.

### 3. Пример рабочего процесса

```bash
# 1. Добавляем новую задачу
$ todosher add "Написать документацию для проекта"
✅ Задача успешно добавлена (ID: 1)

# 2. Смотрим список всех задач
$ todosher list
[1] [TODO] Написать документацию для проекта

# 3. Начинаем работу над задачей
$ todosher start 1
🔄 Задача ID 1 переведена в статус "в процессе"

# 4. Проверяем список начатых задач
$ todosher list in_progress
[1] [IN_PROGRESS] Написать документацию для проекта

# 5. Завершаем задачу
$ todosher finish 1
🎉 Задача ID 1 выполнена!

# 6. Проверяем список сделанных задач
$ todosher list done
[1] [DONE] Написать документацию для проекта
```

### 4. Где хранятся данные?

Вам не нужно настраивать пути вручную. Приложение автоматически определяет домашнюю директорию пользователя и создает там скрытую папку для хранения данных. Это гарантирует, что приложение будет работать корректно, из какой бы папки вы его ни запустили.

* **Linux / macOS**: `~/.todosher/tasks.json`
* **Windows**: `C:\Users\ВашеИмя\.todosher\tasks.json`

> 💡 Файл `tasks.json` имеет простую и читаемую структуру. Вы можете вручную сделать его резервную копию или перенести свои задачи на другой компьютер, просто скопировав этот файл.
