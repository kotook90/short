.golangci.yaml Файл конфигурации линтеров для github action workflows lint.yml Проводит проверку по линтерам:

govet # общий анализ возможных багов
goconst # нахождение строк, которые следует вынести в константы
funlen # детектирование слишком крупных функций
bodyclose # проверка на незакрытые body после чтения тела ответа
errcheck # проверка на обработку всех ошибок
deadcode # детектирование не использованного кода
exportloopref # детектирование экспорта указателя на переменную внутри цикла
dupl
gocritic
gocyclo
staticcheck
structcheck
stylecheck
typecheck
unconvert
unparam
unused
varcheck
.pre-commit-config.yaml Файл конфигурации линтеров для github action workflows pre-commit.yml В конфигурации заданы хуки для проверки:

Валидность всех yaml-файлов репозитория
Символ окончания файла
Лишние пробелы
Юнит тесты для го приложений
.github/workflows/lint.yml - workflow, срабатывает при пул-реквесте на ветку main и при любом пуше на любую ветку ( т.е. всегда)
.github/workflows/pre-commit.yaml - workflow, срабатывает при любоом пул-реквесте и при любом пуше на любую ветку ( т.е. всегда) Делает следующее:

Устанавливает ubuntu-20.04
Устанавливает Go 1.17
Билдит проверяемое приложение
Запускает go vet
Проводит анализ staticcheck ./...
Прогоняет юнит тесты, если они есть
Запускает проверки на валидность всех yaml-файлов репозитория
Запускает проверки на символ окончания файла
Запускает проверки на лишние пробелы Конфигурация выполняемых действий задана в файле .pre-commit-config.yaml
