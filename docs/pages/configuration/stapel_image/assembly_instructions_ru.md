---
title: Запуск инструкций сборки
sidebar: documentation
permalink: ru/documentation/configuration/stapel_image/assembly_instructions.html
ref: documentation_configuration_stapel_image_assembly_instructions
lang: ru
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vQcjW39mf0TUxI7yqNzKPq4_9ffzg2IsMxQxu1Uk1-M0V_Wq5HxZCQJ6x-iD-33u2LN25F1nbk_1Yx5/pub?w=2031&amp;h=144" data-featherlight="image">
      <img src="https://docs.google.com/drawings/d/e/2PACX-1vQcjW39mf0TUxI7yqNzKPq4_9ffzg2IsMxQxu1Uk1-M0V_Wq5HxZCQJ6x-iD-33u2LN25F1nbk_1Yx5/pub?w=1016&amp;h=72">
  </a>

  <div class="tabs">
    <a href="javascript:void(0)" class="tabs__btn active" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'shell')">Shell</a>
    <a href="javascript:void(0)" class="tabs__btn" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'ansible')">Ansible</a>
  </div>

  <div id="shell" class="tabs__content active">
    <div class="language-yaml highlighter-rouge"><pre class="highlight"><code><span class="na">shell</span><span class="pi">:</span>
    <span class="na">beforeInstall</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;bash command&gt;</span>
    <span class="na">install</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;bash command&gt;</span>
    <span class="na">beforeSetup</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;bash command&gt;</span>
    <span class="na">setup</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;bash command&gt;</span>
    <span class="na">cacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">beforeInstallCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">installCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">beforeSetupCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">setupCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span></code></pre>
    </div>
  </div>

  <div id="ansible" class="tabs__content">
    <div class="language-yaml highlighter-rouge"><pre class="highlight"><code><span class="na">ansible</span><span class="pi">:</span>
    <span class="na">beforeInstall</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;task&gt;</span>
    <span class="na">install</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;task&gt;</span>
    <span class="na">beforeSetup</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;task&gt;</span>
    <span class="na">setup</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;task&gt;</span>
    <span class="na">cacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">beforeInstallCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">installCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">beforeSetupCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span>
    <span class="na">setupCacheVersion</span><span class="pi">:</span> <span class="s">&lt;arbitrary string&gt;</span></code></pre>
      </div>
  </div>

  <br/>
  <b>Running assembly instructions with git</b>

  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vRv56S-dpoTSzLC_24ifLqJHQoHdmJ30l1HuAS4dgqBgUzZdNQyA1balT-FwK16pBbbXqlLE3JznYDk/pub?w=1956&amp;h=648" data-featherlight="image">
    <img src="https://docs.google.com/drawings/d/e/2PACX-1vRv56S-dpoTSzLC_24ifLqJHQoHdmJ30l1HuAS4dgqBgUzZdNQyA1balT-FwK16pBbbXqlLE3JznYDk/pub?w=622&amp;h=206">
  </a>
---

## Пользовательские стадии

***Пользовательские стадии*** — это [_стадии_]({{ site.baseurl }}/ru/documentation/reference/stages_and_images.html) со сборочными инструкциями из [конфигурации]({{ site.baseurl}}/ru/documentation/configuration/introduction.html#%D1%87%D1%82%D0%BE-%D1%82%D0%B0%D0%BA%D0%BE%D0%B5-%D0%BA%D0%BE%D0%BD%D1%84%D0%B8%D0%B3%D1%83%D1%80%D0%B0%D1%86%D0%B8%D1%8F-werf). Другими словами, — это стадии, которые конфигурирует пользователь (существуют также служебные стадии, корторые пользователь конфигурировать не может). В настоящее время существует два вида сборочных инструкций: _shell_ и _ansible_.

В Werf существуют 4 _пользовательские стадии_, которые исполняются последовательно в следующем порядке: _beforeInstall_, _install_, _beforeSetup_ и _setup_. В результате исполнения инструкций какой-либо стадии создается один Docker-слой. Т.е. по одному слою на все стадию, в независимости от количества инструкций в ней.

## Причины использовать стадии

### Своя концепция структуры сборки (opinionated software)

Шаблон и механизм работы _пользовательских стадий_ основан на анализе сборки реальных приложений. В результате анализа мы пришли к выводу, что для того, чтобы качественно улучшить сборку большинства приложений достаточно разбить инструкции сборки на 4 группы (эти группы и есть — _пользовательские стадии_) подчиняющиеся определенным правилам. Такая группировка уменьшает количество слоев и ускоряет сборку.

### Четкая структура сборочного процесса

Наличие _пользовательских стадий_ определяет структуру процесса сборки и, таким образом, устанавливает некоторые рамки для разработчика. Не смотря на дополнительное ограничение по сравнению с неструктурированными инструкциями Docker-файла, это наоборот дает выигрыш в скорости, т.к. разработчик знает, какие инструкции на каком этапе должны быть.

### Запуск инструкций сборки при изменениях в git-репозитории

Werf может использовать как локальные так и удаленные git-репозитории при сборке. Любой _пользовательской стадии_ можно определить зависимость от изменения конкретных файлов или папок в одном или нескольких git-репозиториях. Это дает возможность принудительно пересобирать _пользовательскую стадию_ если в локальном или удаленном репозитории (или репозиториях) изменяются какие-либо файлы.

### Больше инструментов сборки: shell, ansible, ...

_Shell_ — знакомый и хорошо известный инструмент сборки. _Ansible_ — более молодой инструмент, требующий чуть больше времени на изучение.

Если вам нужен быстрый результат с минимальными затратами времени и как можно быстрее, то использования _shell_ может быть вполне достаточно — все работает аналогично директиве `RUN` в Dockerfile.

В случае с _Ansible_ применяется декларативный подход и подразумевается идемпотентность операций. Такой подход дает более предсказуемый результат, особенно в случае проектов большого жизненного цикла.

Архитектура Werf позволяет добавлять в будущем поддержку и других инструментов сборки.

## Использование пользовательских стадий

Werf позволяет определять до четырех _пользовательских стадий_ с инструкциями сборки. На содержание самих инструкций сборки Werf не накладывает каких-либо ограничений, т.е. вы можете указывать все те же инструкции, которые указывали в Dockerfile в директиве `RUN`. Однако важно не просто перенести инструкции из Dockerfile, а правильно разбить их на _пользовательские стадии_. Мы предлагаем такое разбиение исходя из опыта работы с реальными приложениями, и вся суть тут в том, что большинство сборок приложений проходят следующие этапы:
- установка системных пакетов
- установка системных зависимостей
- установка зависимостей приложения
- настройка системных пакетов
- настройка приложения

Какая может быть наилучшая стратегия выполнения этих этапов? Может показаться очевидным, что лучше всего выполнять эти этапы последовательно, кешируя промежуточные результаты. Либо, как вариант, — не смешивать инструкции этапов, из-за разных файловых зависимостей. Шаблон _пользовательских стадий_ предлагает следующую стратегию:
- использовать стадию _beforeInstall_ для инсталляции системных пакетов;
- использовать стадию _install_ для инсталляции системных зависимостей и зависимостей приложения;
- использовать стадию _beforeSetup_ для настройки системных параметров и установки приложения;
- использовать стадию _setup_ для настройки приложения.

### beforeInstall

Данная стадия предназначена для выполнения инструкций перед установкой приложения. Сюда следует относить установку системных приложений которые редко изменяются, но установка которых занимает много времени. Примером таких приложений могут быть — языковые пакеты, инструменты сборки, такие как composer, java, gradle и т. д. Также сюда правильно относить инструкции настройки системы, которые меняются редко, например — языковые настройки, настройки часового пояса, добавление пользователей и групп.

Поскольку эти компоненты меняются редко, они будут кэшироваться в рамках стадии _beforeInstall_ на длительный период.

### install

Данная стадия предназначена для установки приложения и его зависимостей, а также выполнения базовых настроек.

На данной стадии появляется доступ к исходному коду (директива git), и возможна установка зависимостей на основе manifest-файлов с использованием таких инструментов как composer, gradle, npm и т.д. Поскольку сборка стадии зависит от manifest-файла, для достижения наилучшего результата важно связать изменение в manifest-файлах репозитория с данной стадией. Например, если в проекте используется composer, то установление зависимости от файла composer.lock позволит пересобирать стадию _beforeInstall_, в случае изменения файла composer.lock.

### beforeSetup

Данная стадия предназначена для подготовки приложения перед настройкой.

На данной стадии рекомендуется выполнять разного рода компиляцию и обработку. Например — компиляция jar-файлов, бинарных файлов, файлов библиотек, создание ассетов web-приложений, минификация, шифрование и т.п. Перечисленные операции как правило зависят от изменений в исходном коде, и на данной стадии также важно определить достаточные зависимости изменений в репозитории. Логично, что зависимость данной стадии от изменений в репозитории будет покрывать уже большую область файлов в репозитории, чем на предыдущей стадии, и, соответственно, ее пересборка будет выполняться чаще.

При правильно определенных зависимостях, изменение в коде приложения дожно вызывать пересборку стадии _beforeSetup_, в случае если не изменялся manifest-файл. А в случае если manifest-файл изменился, должна уже вызываться пересборка стадии _install_ и следующих за ней стадий.

### setup

Данная стадия предназначена для настройки приложения.

Обычно на данной стадии выполняется копирование файлов конфигурации (например в папку `/etc`), создание файлов текущей версии приложения и т.д.  Такого рода операции не должны быть затратными по времени, т.к. они скорее всего будут выполняться при каждом новом коммите в репозитории.

### Пользовательская стратегия

Не смотря на изложенную четкую стратегию шаблона _пользовательских стадий_ и назначения каждой стадии, по сути нет никаких ограничений. Предложенные назначения каждой стадии являются лишь рекомендацией, которые основаны на нашем анализе работы реальных приложений. Вы можете использовать только одну пользовательскую стадию, либо определить свою стратегию группировки инструкций, чтобы получить преимущества кэширования и зависимостей от изменений в git-репозиториях с учетом особенностей сборки вашего приложения.

## Синтаксис

Пользовательские стадии и инструкции сборки определяются внутри двух взаимоисключающих директив вида сборочных инструкций — `shell` и `ansible`. Вы можете собирать образ используя внутри всех стадий либо сборочные инструкции ***shell***, либо сборочные инструкции ***ansible***.

Внутри директив вида сборочных инструкций можно указывать четыре директивы сборочных инструкций каждой _пользовательской стадии_, соответственно:
- `beforeInstall`
- `install`
- `beforeSetup`
- `setup`

Внутри директив вида сборочных инструкций также можно указывать директивы версий кэша (***cacheVersion***), которые по сути являются частью сигнатуры каждой _пользовательской стадии_. Более подробно об этом читай в [соответствующем разделе](#dependency-on-cacheversion-values) section.

## Shell

Синтаксис описания _пользовательских стадий_ при использовании сборочных инструкций _shell_:

```yaml
shell:
  beforeInstall:
  - <bash_command 1>
  - <bash_command 2>
  ...
  - <bash_command N>
  install:
  - bash command
  ...
  beforeSetup:
  - bash command
  ...
  setup:
  - bash command
  ...
  cacheVersion: <version>
  beforeInstallCacheVersion: <version>
  installCacheVersion: <version>
  beforeSetupCacheVersion: <version>
  setupCacheVersion: <version>
```

Сборочные инструкции _shell_ — это массив bash-комманд для соответствующей _пользовательской стадии_. Все команды одной стадии выполняются как одна инструкция `RUN`  в Dockerfile, т.е. в результате создается один слой на каждую _пользовательскую стадию_.

Werf при сборке использует собственный исполняемый файл bash и вам не нужно отдельно добавлять его в образ (или [базовый образ]({{ site.baseurl }}/documentation/configuration/stapel_image/base_image.html)) при сборке. Все команды одной стадии объединяются с помощью выражения `&&` bash и кодируются алгоритмом base64 перед передачей в _сборочный контейнер_. _Сборочный контейнер_ пользовательской стадии декодирует команды и запускает их.

Пример описания стадии _beforeInstall_ содержащей команды `apt-get update` и `apt-get install`:

```yaml
beforeInstall:
- apt-get update
- apt-get install -y build-essential g++ libcurl4
```

Werf выполнит команды стадии следующим образом::
- на хост-машине сгенерируется временный скрипт:

    ```bash
    #!/.werf/stapel/embedded/bin/bash -e

    apt-get update
    apt-get install -y build-essential g++ libcurl4
    ```

- скрипт смонтируется в _сборочный контейнер_ как `/.werf/shell/script.sh`
- скрипт выполнится.

> Исполняемый фаил `bash` находится внутри Docker-тома _stapel_. Подробнее про эту концепцию можно узнать в этой [статье](https://habr.com/company/flant/blog/352432/) (упоминаемый в статье `dappdeps` был переименован в `stapel`, но принцип сохранился)

## Ansible

Синтаксис описания _пользовательских стадий_ при использовании сборочных инструкций _ansible_:

```yaml
ansible:
  beforeInstall:
  - <ansible task 1>
  - <ansible task 2>
  ...
  - <ansible task N>
  install:
  - ansible task
  ...
  beforeSetup:
  - ansible task
  ...
  setup:
  - ansible task
  ...
  cacheVersion: <version>
  beforeInstallCacheVersion: <version>
  installCacheVersion: <version>
  beforeSetupCacheVersion: <version>
  setupCacheVersion: <version>
```

### Ansible config and stage playbook

Сборочные инструкции _ansible_ —  это массив Ansible-заданий для соответствующей _пользовательской стадии_. Для запуска этих заданий с помощью `ansible-playbook` Werf монтирует следующую структуру папок в _сборочнй контейнер_:

```bash
/.werf/ansible-workdir
├── ansible.cfg
├── hosts
└── playbook.yml
```

`ansible.cfg` содержит настройки для Ansible:
- использование локального транспорта (transport = local)
- подключение callback плагина werf для удобного логирования (stdout_callback = werf)
- включение режима цвета (force_color = 1)
- установка использования `sudo` для повышения привелегий (чтобы небыло необходимости использовать `become` в ansible-заданиях)

`hosts` — inventory-файл, содержит только localhost и некоторые `ansible_*` параметры.

`playbook.yml` — playbook, содержащий все задания соответствующей _пользовательской стадии_. Пример `werf.yaml` с описанием стадии _install_:

```yaml
ansible:
  install:
  - debug: msg='Start install'
  - file: path=/etc mode=0777
  - copy:
      src: /bin/sh
      dest: /bin/sh.orig
  - apk:
      name: curl
      update_cache: yes
  ...
```

В приведенном примере, Werf сгенерирует следующий `playbook.yml` для стадии _install_:
```yaml
- hosts: all
  gather_facts: 'no'
  tasks:
  - debug: msg='Start install'  \
  - file: path=/etc mode=0777   |
  - copy:                        > эти строки будут скопированы из werf.yaml
      src: /bin/sh              |
      dest: /bin/sh.orig        |
  - apk:                        |
      name: curl                |
      update_cache: yes         /
  ...
```

Werf выполняет playbook _пользовательской стадии_ в сборочном контейнере стадии с помощью команды `playbook-ansible`:

```bash
$ export ANSIBLE_CONFIG="/.werf/ansible-workdir/ansible.cfg"
$ ansible-playbook /.werf/ansible-workdir/playbook.yml
```

> Исполняемые файлы и библиотеки `ansible` и `python` находятся внутри Docker-тома _stapel_. Подробнее про эту концепцию можно узнать в этой [статье](https://habr.com/company/flant/blog/352432/) (упоминаемый в статье `dappdeps` был переименован в `stapel`, но принцип сохранился)

### Поддерживаемые модули

Одной из концепций, которую использует Werf, является идемпотентность сборки. Это значит что если "ничего не изменилось", то Werf при повторном и последующих запусках сборки должен создавать бинарно идентичные образы. В Werf эта задача решается с помощью подсчета _сигнатур стадий_.

Многие модули Ansible не являются идемпотентными, т.е. они могут давать разный результат запусков при неизменных входных параметрах. Это, конечно, не дает возможность корректно высчитывать _сигнатуру_ стадии, чтобы определять реальную необходимость её пересборки из-за изменений. Это привело к тому, что список поддерживаемых модулей был ограничен.

На текущий момент, список поддерживаемых модулей Ansible следующий:

- [Commands modules](https://docs.ansible.com/ansible/2.5/modules/list_of_commands_modules.html): command, shell, raw, script.
- [Crypto modules](https://docs.ansible.com/ansible/2.5/modules/list_of_crypto_modules.html): openssl_certificate и другие.
- [Files modules](https://docs.ansible.com/ansible/2.5/modules/list_of_files_modules.html): acl, archive, copy, stat, tempfile и другие.
- [Net Tools Modules](https://docs.ansible.com/ansible/2.5/modules/list_of_net_tools_modules.html): get_url, slurp, uri.
- [Packaging/Language modules](https://docs.ansible.com/ansible/2.5/modules/list_of_packaging_modules.html#language): composer, gem, npm, pip и другие.
- [Packaging/OS modules](https://docs.ansible.com/ansible/2.5/modules/list_of_packaging_modules.html#os): apt, apk, yum и другие.
- [System modules](https://docs.ansible.com/ansible/2.5/modules/list_of_system_modules.html): user, group, getent, locale_gen, timezone, cron и другие.
- [Utilities modules](https://docs.ansible.com/ansible/2.5/modules/list_of_utilities_modules.html): assert, debug, set_fact, wait_for.

При указании в _конфигурации сборки_ модуля отсутствущего в приведенном списке, сборка прервется с ошибкой. Не стесняйтесь [сообщать](https://github.com/flant/werf/issues/new) нам, если вы считаете что какой-либо модуль должен быть включен в список поддерживаемых.

### Копирование файлов

Предпочтительный способ копирования файлов в образ ­— использование [_git-маппинга_]({{ site.baseurl }}/ru/documentation/configuration/stapel_image/git_directive.html). Werf не может определять изменения в копируемых файлах при использовании модуля `copy`. Единственный вариант копирования внешнего файла в образ на такущий момент — использовать метод `.Files.Get` Go-шаблона. Данный метод возвращает содержимое файла как строку, что дает возможность использовать содержимое как часть _пользовательской стадии_. Таким образом, при изменении содержимого файла изменится сигнатура соответствующей стадии, что приведет к пересборке всей стадии.

Пример копирования файла `nginx.conf` в образ:

{% raw %}
```yaml
ansible:
  install:
  - copy:
      content: |
{{ .Files.Get "/conf/etc/nginx.conf" | indent 8}}
      dest: /etc/nginx/nginx.conf
```
{% endraw %}

Werf применит Go-шаблонизатор и в результате получится подобный `playbook.yml`:

```yaml
- hosts: all
  gather_facts: no
  tasks:
    install:
    - copy:
        content: |
          http {
            sendfile on;
            tcp_nopush on;
            tcp_nodelay on;
            ...
```

### Шаблоны Jinja

В Ansible реализована поддержка шаблонов [Jinja](https://docs.ansible.com/ansible/2.5/user_guide/playbooks_templating.html) в playbook'ах. Однако, у Go-шаблонов и Jinja-шаблонов одинаковый разделитель: {% raw %}`{{` и `}}`{% endraw %}. Чтобы использовать Jinja-шаблоны в конфигурации Werf, их нужно экранировать. Для этого есть два варианта: экранировать только {% raw %}`{{`{% endraw %}, либо экранировать все выражение шаблона Jinja.

Например, у вас есть следующая задача Ansible:

{% raw %}
```yaml
- copy:
    src: {{item}}
    dest: /etc/nginx
    with_files:
    - /app/conf/etc/nginx.conf
    - /app/conf/etc/server.conf
```
{% endraw %}

{% raw %}
Тогда, выражение Jinja-шаблона `{{item}}` должно быть экранировано:
{% endraw %}

{% raw %}
```yaml
# экранируем только {{
src: {{"{{"}} item }}
```
либо
```yaml
# экранируем все выражение
src: {{`{{item}}`}}
```
{% endraw %}

### Проблемы с Ansible

- Live-вывод реализован только для модулей `raw` и `command`. Остальные модули отображают вывод каналов `stdout` и `stderr` после выполнения, что приводит к задержкам, скачкообразному выводу.
- Большой вывод в `stderr` может подвесить выполнение Ansible-задачи ([issue #784](https://github.com/flant/werf/issues/784)).
- Модуль `apt` подвисает на некоторых версиях Debian и Ubuntu. Проявляется также на наследуемых образах([issue #645](https://github.com/flant/werf/issues/645)).

## Зависимости пользовательских стадий

Одна из особенностей Werf — возможность определять зависимости при которых происходит пересборка _стадии_.
Как указано в [справочнике]({{ site.baseurl }}/ru/documentation/reference/stages_and_images.html), сборка _стадий_ выполняется последовательно, одна за другой, и для каждой _стадии_ высчитывается _сигнатура стадии_. У _сигнатур_ есть ряд зависимостей, при изменении которых _сигнатура стадии_ меняется, что служит для Werf сигналом для пересборки стадии с измененной _сигнатурой_. Поскольку каждая следующая _стадия_ имеет зависимость в том числе и от предыдущей _стадии_ согласно _конвейеру стадий_, при изменении сигнатуры какой-либо _стадии_, произойдет пересборка и _стадии_ с измененной сигнатурой и всех последующих _стадий_.

_Сигнатура пользовательских стадий_ и соответственно пересборка _пользовательских стадий_ зависит от изменений:
- в инструкциях сборки
- в директивах семейства _cacheVersion_
- в git-репозитории (или git-репозиториях)
- в файлах, импортируемых из [артефактов]({{ site.baseurl }}/ru/documentation/configuration/stapel_artifact.html)

Первые три описанных варианта зависимостей, рассматриваются подробно далее.

## Зависимость от изменений в инструкциях сборки

_Сигнатура пользовательской стадии_ зависит от итогового текста инструкций, т.е. после применения шаблонизатора. Любые изменения в тексте инструкций с учетмо применения шаблонизатора Go или Jinja (в случае Ansible) в _пользовательской стадии_ приводят к пересборке _стадии_. Например, вы используете следующие _shell-инструкции_ :

```yaml
shell:
  beforeInstall:
  - echo "Commands on the Before Install stage"
  install:
  - echo "Commands on the Install stage"
  beforeSetup:
  - echo "Commands on the Before Setup stage"
  setup:
  - echo "Commands on the Setup stage"
```

При первой сборке этого образа буду выполнены инструкции всех четырех _пользовательских стадий_. В данной конфигурации нет _git-маппинга_, так что последующие сборки не приведут к повторному выполнению инструкций — _сигнатура пользовательских стадий_ не изменялась, сборочный кэш содержит актуальную информацию (валиден).

Изменим инструкцию сборки для стадии _install_:

```yaml
shell:
  beforeInstall:
  - echo "Commands on the Before Install stage"
  install:
  - echo "Commands on the Install stage"
  - echo "Installing ..."
  beforeSetup:
  - echo "Commands on the Before Setup stage"
  setup:
  - echo "Commands on the Setup stage"
```

Запуск Werf для сборки приведет к выплонению всех инструкций стадии _install_ и инструкций последующих _стадий_.

Go-templating and using environment variables can changes assembly instructions
and lead to unforeseen rebuilds. For example:

{% raw %}
```yaml
shell:
  beforeInstall:
  - echo "Commands on the Before Install stage for {{ env "CI_COMMIT_SHA” }}"
  install:
  - echo "Commands on the Install stage"
  ...
```
{% endraw %}

First build renders _beforeInstall command_ into:
```bash
echo "Commands on the Before Install stage for 0a8463e2ed7e7f1aa015f55a8e8730752206311b"
```

Build for the next commit renders _beforeInstall command_ into:

```bash
echo "Commands on the Before Install stage for 36e907f8b6a639bd99b4ea812dae7a290e84df27"
```

Using `CI_COMMIT_SHA` assembly instructions text changes every commit.
So this configuration rebuilds _beforeInstall_ user stage on every commit.

## Dependency on git repo changes

<a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vRv56S-dpoTSzLC_24ifLqJHQoHdmJ30l1HuAS4dgqBgUzZdNQyA1balT-FwK16pBbbXqlLE3JznYDk/pub?w=1956&amp;h=648" data-featherlight="image">
    <img src="https://docs.google.com/drawings/d/e/2PACX-1vRv56S-dpoTSzLC_24ifLqJHQoHdmJ30l1HuAS4dgqBgUzZdNQyA1balT-FwK16pBbbXqlLE3JznYDk/pub?w=622&amp;h=206">
  </a>

As stated in a _git mapping_ reference, there are _gitArchive_ and _gitLatestPatch_ stages. _gitArchive_ is executed after _beforeInstall_ user stage, and _gitLatestPatch_ is executed after _setup_ user stage if a local git repository has changes. So, to execute assembly instructions with the latest version of source codes, you may rebuild _gitArchive_ with [special commit]({{site.baseurl}}/documentation/configuration/stapel_image/git_directive.html#rebuild-of-git_archive-stage) or rebuild _beforeInstall_ (change _cacheVersion_ or instructions for _beforeInstall_ stage).

_install_, _beforeSetup_ and _setup_ user stages are also dependant on git repository changes. A git patch is applied at the beginning of _user stage_ to execute assembly instructions with the latest version of source codes.

> During image build process source codes are updated **only within one stage**, subsequent stages are based on this stage and use actualized files. First build adds sources on _gitArchive_ stage. Any other build updates sources on _gitCache_, _gitLatestPatch_ or on one of the following user stages: _install_, _beforeSetup_ and _setup_.
<br />
<br />
This stage is shown in _Calculation signature phase_
![git files actualized on specific stage]({{ site.baseurl }}/images/build/git_mapping_updated_on_stage.png)

_User stage_ dependency on git repository changes is defined with `git.stageDependencies` parameter. Syntax is:

```yaml
git:
- ...
  stageDependencies:
    install:
    - <mask 1>
    ...
    - <mask N>
    beforeSetup:
    - <mask>
    ...
    setup:
    - <mask>
```

`git.stageDependencies` parameter has 3 keys: `install`, `beforeSetup` and `setup`. Each key defines an array of masks for one user stage. User stage is rebuilt if a git repository has changes in files that match with one of the masks defined for _user stage_.

For each _user stage_ werf creates a list of matched files and calculates a checksum over each file attributes and content. This checksum is a part of _stage signature_. So signature is changed with every change in a repository: getting new attributes for the file, changing file's content, adding a new matched file, deleting a matched file, etc.

`git.stageDependencies` masks work together with `git.includePaths` and `git.excludePaths` masks. werf considers only files matched with `includePaths` filter and `stageDependencies` masks. Likewise, werf considers only files not matched with `excludePaths` filter and matched with `stageDependencies` masks.

`stageDependencies` masks works like `includePaths` and `excludePaths` filters. Masks are matched with files paths and may contain the following glob patterns:

- `*` — matches any file. This pattern includes `.` and excludes `/`
- `**` — matches directories recursively or files expansively
- `?` — matches any one character. Equivalent to /.{1}/ in regexp
- `[set]` — matches any one character in the set. Behaves exactly like character sets in regexp, including set negation ([^a-z])
- `\` — escapes the next metacharacter

Mask that starts with `*` is treated as anchor name by yaml parser. So mask with `*` or `**` patterns at the beginning should be quoted:

```
# * at the beginning of mask, so use double quotes
- "*.rb"
# single quotes also work
- '**/*'
# no star at the beggining, no quoting needed
- src/**/*.js
```

Werf determines whether the files changes in the git repository with use of checksums. For _user stage_ and for each mask, the following algorithm is applied:

- werf creates a list of all files from `add` path and apply `excludePaths` and `includePaths` filters:
- each file path from the list compared to the mask with the use of glob patterns;
- if mask matches a directory then this directory content is matched recursively;
- werf calculates checksum of attributes and content of all matched files.

These checksums are calculated in the beginning of the build process before any stage container is ran.

Example:

```yaml
---
image: app
git:
- add: /src
  to: /app
  stageDependencies:
    beforeSetup:
    - "*"
shell:
  install:
  - echo "install stage"
  beforeSetup:
  - echo "beforeSetup stage"
  setup:
  - echo "setup stage"
```

This `werf.yaml` has a git mapping configuration to transfer `/src` content from local git repository into `/app` directory in the image. During the first build, files are cached in _gitArchive_ stage and assembly instructions for _install_ and _beforeSetup_ are executed. The next builds of commits that have only changes outside of the `/src` do not execute assembly instructions. If a commit has changes inside `/src` directory, then checksums of matched files are changed, werf will apply git patch, rebuild all existing stages since _beforeSetup_: _beforeSetup_ and _setup_. Werf will apply patch on the _beforeSetup_ stage itself.

## Dependency on CacheVersion values

There are situations when a user wants to rebuild all or one of _user stages_. This
can be accomplished by changing `cacheVersion` or `<user stage name>CacheVersion` values.

Signature of the _install_ user stage depends on the value of the
`installCacheVersion` parameter. To rebuild the _install_ user stage (and
subsequent stages), you need to change the value of the `installCacheVersion` parameter.

> Note that `cacheVersion` and `beforeInstallCacheVersion` directives have the same effect. When these values are changed, then the _beforeInstall_ stage and subsequent stages rebuilt.

### Example: common image for multiple applications

You can define an image with common packages in separated `werf.yaml`. `cacheVersion` value can be used to rebuild this image to refresh packages versions.

```yaml
image: ~
from: ubuntu:latest
shell:
  beforeInstallCacheVersion: 2
  beforeInstall:
  - apt update
  - apt install ...
```

This image can be used as base image for multiple applications if images from hub.docker.com doesn't suite your needs.

### External dependency example

_CacheVersion directives_ can be used with [go templates]({{ site.baseurl }}/documentation/configuration/introduction.html#go-templates) to define _user stage_ dependency on files, not in the git tree.

{% raw %}
```yaml
image: ~
from: ubuntu:latest
shell:
  installCacheVersion: {{.Files.Get "some-library-latest.tar.gz" | sha256sum}}
  install:
  - tar zxf some-library-latest.tar.gz
  - <build application>
```
{% endraw %}

Build script can be used to download `some-library-latest.tar.gz` archive and then execute `werf build` command. If the file is changed then werf rebuilds _install user stage_ and subsequent stages.
