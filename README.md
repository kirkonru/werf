# dapp [![Gem Version](https://badge.fury.io/rb/dapp.svg)](https://badge.fury.io/rb/dapp) [![Build Status](https://travis-ci.org/flant/dapp.svg)](https://travis-ci.org/flant/dapp) [![Code Climate](https://codeclimate.com/github/flant/dapp/badges/gpa.svg)](https://codeclimate.com/github/flant/dapp) [![Test Coverage](https://codeclimate.com/github/flant/dapp/badges/coverage.svg)](https://codeclimate.com/github/flant/dapp/coverage)

## Reference

### Dappfile

#### Основное

##### name \<name\>
Базовое имя для собираемых docker image`ей: <базовое имя>-dappstage:<signature>
Опционально, по умолчанию определяется исходя из имени директории, в которой находится Dappfile.

##### install\_depends\_on \<glob\>[,\<glob\>, \<glob\>, ...]
Список файлов зависимостей для стадии install.

* При изменении содержимого указанных файлов, произойдет пересборка стадии install.
* Учитывается лишь содержимое файлов и порядок в котором они указаны (имена файлов не учитываются).
* Поддерживаются glob-паттерны.
* Директории игнорируются.

##### setup\_depends\_on \<glob\>[,\<glob\>, \<glob\>, ...]
Список файлов зависимостей для стадии setup.

* При изменении содержимого указанных файлов, произойдет пересборка стадии setup.
* Учитывается лишь содержимое файлов и порядок в котором они указаны (имена файлов не учитываются).
* Поддерживаются glob-паттерны.
* Директории игнорируются.

##### builder \<builder\>
Тип сборки: :chef или :shell.
* Опционально, по умолчанию будет выбран тот builder, который будет использован первым (см. Chef, Shell).
* В одном Dappfile можно использовать только один builder.

##### app \<app\>[, &blk]
Определяет приложение <app> для сборки.

* Опционально, по умолчанию будет использоваться приложение с базовым именем (см. name \<name\>).
* Можно определять несколько приложений в одном Dappfile.
* При использовании блока создается новый контекст.
  * Наследуются все настройки родительского контекста.
  * Можно дополнять или переопределять настройки родительского контекста.
  * Можно использовать директиву app внутри нового контекста.
* При использовании вложенных вызовов директивы, будут использоваться только приложения указанные последними в иерархии. Другими словами, в описанном дереве приложений будут использованы только листья.
  * Правила именования вложенных приложений: <app>[-<subapp>-<subsubapp>...]
* Примеры:
  * Собирать приложения X и Y:```ruby
app 'X'
app 'Y'
```
  * Собирать приложения X, Y-Z и Y-K-M:```ruby
app 'X'
app 'Y' do
  app 'Z'

  app 'K' do
    app 'M'
  end
end
```

#### Артифакты
*TODO*

#### Docker
*TODO*

#### Shell
*TODO*

#### Chef

##### chef.module \<mod\>[, \<mod\>, \<mod\> ...]
Включить переданные модули для chef builder в данном контексте.

* Для каждого переданного модуля может существовать по одному рецепту на каждый из stage.
* Файл рецепта для \<stage\>: recipes/\<stage\>.rb
* Рецепт модуля будет добавлен в runlist для данного stage если существует файл рецепта.
* Порядок вызова рецептов модулей в runlist совпадает порядком их описания в конфиге.
* При сборке stage, для каждого из включенных модулей, при наличии файла рецепта, будут скопированы:
  * files/\<stage\>/ -> files/default/
  * templates/\<stage\>/ -> templates/default/
  * metadata.json

##### chef.skip_module \<mod\>[, \<mod\>, \<mod\> ...]
Выключить переданные модули для chef builder в данном контексте.

##### chef.reset_modules
Выключить все модули для chef builder в данном контексте.

##### chef.recipe \<recipe\>[, \<recipe\>, \<recipe\> ...]
Включить переданные рецепты из проекта для chef builder в данном контексте.

* Для каждого преданного рецепта может существовать файл рецепта в проекте на каждый из stage.
* Файл рецепта для \<stage\>: recipes/\<stage\>/\<recipe\>.rb
* Рецепт будет добавлен в runlist для данного stage если существует файл рецепта.
* Порядок вызова рецептов в runlist совпадает порядком их описания в конфиге.
* При сборке stage, при наличии хотя бы одного файла рецепта из включенных, будут скопированы:
  * files/\<stage\> -> files/default/
  * templates/\<stage\>/ -> templates/default/
  * metadata.json

##### chef.remove_recipe \<recipe\>[, \<recipe\>, \<recipe\> ...]
Выключить переданные рецепты из проекта для chef builder в данном контексте.

##### chef.reset_recipes
Выключить все рецепты из проекта для chef builder в данном контексте.

##### chef.reset_all
Выключить все рецепты из проекта и все модули для chef builder в данном контексте.

##### Примеры
* [Dappfile](doc/example/Dappfile.chef.1)

### Команды

#### dapp build
*TODO*

#### dapp push
*TODO*

#### dapp smartpush
*TODO*

## Architecture

### Стадии
*TODO*

### Хранение данных

#### Кэш стадий
*TODO*

#### Временное
*TODO*

#### Метаданные
*TODO*

#### Кэш сборки
*TODO*
