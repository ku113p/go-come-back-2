# 🔁 Go Level-Up Plan: From Confident to Middle+

**Goal:** Подготовиться к собеседованиям уровня Middle+ Go Developer через интересные и практически значимые проекты, теоретическую проработку сложных тем и шлифовку GitHub-портфолио.

**Schedule:** 1.5–2.5 часа в будни. Выходные: Rust / отдых / pet-проекты по желанию.

**Duration:** 6–10 недель

---

## Phase 1: Deep Theory & Language Mastery (2 недели)

*Цель: Не просто "знать", а понимать — что происходит в рантайме и под капотом.*

- [ ] **Concurrency Deep Dive:**
    - [x] Race conditions (`go test -race`)
    - [x] Deadlocks (аналитика + примеры)
    - [x] Channel patterns (fan-in/out, worker pools, cancellation)
    - [ ] `context.Context` (timeouts, cancellation, values)
    - [x] Практика: Напиши минималистичный воркер пул с контекстом

- [ ] **Memory & GC:**
    - [ ] Stack vs Heap
    - [ ] Pointer escaping
    - [ ] Garbage Collector (как работает, когда триггерится)
    - [ ] Инструмент: `pprof`, `runtime`, `go tool trace`

- [ ] **Interface Internals:**
    - [ ] `interface{}` vs конкретные интерфейсы
    - [ ] Почему `nil` != `nil`?
    - [ ] Пустой интерфейс + type assertions vs generics
    - [ ] pointer-to-interface vs interface-to-pointer

- [ ] **Error Handling Advanced:**
    - [ ] `errors.Is`, `errors.As`, `errors.Unwrap`
    - [ ] Сравнение с Python exception-цепочками
    - [ ] Создай свою обёртку для ошибок (с метаданными)

**📘 Результат:** Пишешь на Go не просто "на автопилоте", а с чётким пониманием, как всё устроено.

---

## Phase 2: Прокачанные Pet-проекты (4–6 недель)

> Интересные, нестандартные проекты, связанные с ИИ, агентами, системными тулзами и параллелизмом.

### ✅ Проект 1: AI-Powered CLI Agent *(Medium)*

> Командный агент в терминале с несколькими режимами: shell-интерфейс, открытие файлов, поиск, вызов API, краткий вывод, чат с LLM.

- [ ] Ввод/вывод через CLI (`promptui`, `os.Stdin`)
- [ ] HTTP-интеграция с OpenAI / Llama API
- [ ] Архитектура: FSM (finite state machine) + context
- [ ] Конкурентные задачи (напр. параллельный поиск файлов + fetch документации)
- [ ] Генерация простых команд (обработка естественного языка в подсказках)

📦 Вариации:
- Bash-помощник
- Обработчик логов по запросу

---

### ✅ Проект 2: Task Orchestration Engine *(Medium+)*

> Мини-аналоги Temporal/Airflow с DSL или JSON-описанием задач.

- [ ] Чтение JSON конфигов
- [ ] Запуск цепочек задач с зависимостями
- [ ] Ретрай и таймауты (через `context`)
- [ ] Логгирование и персистентность
- [ ] Параллельное выполнение (`sync.WaitGroup`, worker pool)

Дополнительно:
- [ ] GUI (если хочется): web + websocket + progress

---

### ✅ Проект 3: Self-Hosting Tool *(Medium-Heavy)*

> Утилита для развёртывания проектов в Docker по описанию (аналог doppler/envkey + mini-docker-compose)

- [ ] Чтение `.env` + генерация Dockerfile
- [ ] HTTP UI или CLI
- [ ] Работа с файловой системой
- [ ] Статусы и прогресс в реальном времени (chan + goroutine)
- [ ] Хранение состояния (sqlite или flat file DB)

---

## Phase 3: Interview Battle Prep (Параллельно с проектами)

*Цель: Без паники решать задачи на собеседованиях и уверенно отвечать на вопросы про язык и архитектуру.*

- [ ] **50 Interview-style Tasks in Go:**
    - [ ] Алгоритмы и структуры данных (Leetcode Easy-Medium)
    - [ ] Конкурентность: safe map, воркеры, таймеры, race-detection
    - [ ] Структурирование проектов (чистая архитектура, слоями)

- [ ] **Top Interview Questions:**
    - [ ] Чем отличается goroutine от OS thread?
    - [ ] Как устроен GC в Go?
    - [ ] Почему slice может "протекать память"?
    - [ ] Чем interface отличается от struct?
    - [ ] Когда использовать pointer receiver?

- [ ] **Mock Interviews:**
    - [ ] 1-2 раза в неделю — с другом, ChatGPT, или платформа (pramp, interviewing.io)

---

## Phase 4: GitHub Portfolio & Документы

*Цель: Убедительный, чистый, “вкусный” профиль.*

- [ ] **Отбор 2 проектов:** самые яркие (желательно из фазы 2)
- [ ] **README.md:** скриншоты, GIF'ки, структура, tech stack
- [ ] **CI:** `go test`, `golangci-lint`, GitHub Actions
- [ ] **Dockerfile + Makefile:** чтобы "скачал и запустил"
- [ ] **Код стайл:** `go fmt`, `golint`, понятные имена
- [ ] **Cover letter / CV:** Подчеркнуть переход из Python в Go через реальный опыт

---

## Phase X: Inspiration & Growth

*Поддерживай интерес, не забывай развиваться и кайфовать от процесса.*

- [ ] Читать чужой код (стандартная библиотека, проект из Awesome Go)
- [ ] Пробовать generics в реальных задачах
- [ ] Участвовать в Go комьюнити (форумы, митапы)
- [ ] Rust vs Go: архитектурные сравнения (если интересно)

---

## Напоминания

- **Игровой подход > зубрёжка**
- **Делай проекты так, как будто ими будешь хвастаться на конференции**
- **Глубина важнее объёма**
- **Регулярность бьёт интенсивность**
- **Развлекай себя: экспериментируй, пробуй, твори**
