# gss-template

Шаблон для [gss](https://github.com/psds-microservice/gss). Не используется напрямую — клонируется через `gss install`, затем `gss init` генерирует из него новый Go-проект.

## Содержимое

- **GORM** + PostgreSQL: config, model, database (миграции, сиды), service. Отдельный слой repository не используется — сервис работает с GORM напрямую; при необходимости его можно добавить позже.
- **HTTP**: health/ready, `cmd/<project>/main.go`.
- **CI**: при `gss init` можно выбрать GitHub Actions или GitLab CI — рендерятся соответствующие файлы.

## Как использовать

1. Установить gss и настроить `template_repository_url` на этот репозиторий.
2. Выполнить `gss install` (один раз).
3. В каталоге будущего проекта: `gss init` (указать group, project, kuber, ci).

После генерации: скопировать `.env.example` в `.env`, выполнить `go mod tidy` (или `make tidy`), при необходимости `go get gorm.io/gorm gorm.io/driver/postgres`.
