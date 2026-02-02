# gss — Go Service Scaffold

Генератор базовой структуры Go-проекта из шаблона. Поддерживает выбор CI/CD: **GitHub Actions** или **GitLab CI**.

## Установка

```bash
go build -o gss .
# или
go install
```

## Конфигурация

При первом запуске создаётся `~/.config/gss/config.yaml`. В нём задаются:

- `template_repository_url` — URL репозитория с шаблоном (или переменная окружения `GSS_TEMPLATE_REPO`)
- `template_repository_path` — каталог, куда клонируется шаблон (по умолчанию `~/.config/gss/tpl`)
- `template_path` — каталог используемого шаблона (при вызове с аргументом `v2` — `tpl/v2`)

## Команды

| Команда | Описание |
|--------|----------|
| `gss install` | Клонирует репозиторий шаблона в `~/.config/gss/tpl` |
| `gss update` | Обновляет шаблон (`git pull`) |
| `gss init` | Создаёт новый проект в текущей директории по шаблону |

### init

Интерактивно (или через флаги) запрашивает:

- **group** — Git group/org (например, `myorg`)
- **project** — имя проекта (например, `my-service`)
- **kuber** — namespace в Kubernetes
- **ci** — `github` или `gitlab` (какой CI/CD генерировать)

Флаги: `-g`, `-p`, `-k`, `--ci`.

## Шаблон

Шаблон — репозиторий с файлами `.tmpl` и конфигом **scaffold.json** в корне (или в `v2/`).

Формат **scaffold.json**:

- **mkdir** — список каталогов (пути могут содержать `{{.Project.K}}`, `{{.Project.S}}`, `{{.Group}}`, `{{.CI}}` и т.д.)
- **gitkeep** — каталоги, в которые положить `.gitkeep`
- **copy** — пары "источник → назначение" (копирование как есть)
- **render** — пары `from` (шаблон) → `to` (файл). Опционально поле **ci**: `"github"` или `"gitlab"` — правило применяется только при выбранном CI
- **command** — команды, выполняемые в корне нового проекта (например, `git init`, `go mod tidy`)

Данные для шаблонов: `Group`, `Project.C` / `Project.K` / `Project.S`, `Kuber`, `CI`.

Пример конфига и структуры шаблона — в `templates/scaffold.json.example`.

## Разработка

См. [hero-rewrite-guide.md](hero-rewrite-guide.md) — как устроен перенос логики и выбор GitHub/GitLab для CI/CD.
