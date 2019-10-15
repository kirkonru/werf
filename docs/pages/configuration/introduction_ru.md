---
title: Общие сведения
sidebar: documentation
permalink: ru/documentation/configuration/introduction.html
ref: documentation_configuration_introduction
lang: ru
author: Alexey Igrychev <alexey.igrychev@flant.com>, Timofey Kirillov <timofey.kirillov@flant.com>
---

## Что такое конфигурация Werf?

Для использования Werf в приложении должна быть описана конфигурация, которая включает в себя:

1. Определение мета-информации проекта, например — имени проекта, которое будет впоследствии влиять на результат сборки, деплоя и другие команды.
2. Определение списка образов проекта и инструкций их сборки.

Конфигурация Werf хранится в YAML-файле `werf.yaml` в корневой папке проекта (приложения), и представляет собой набор секций конфигурации -- частей YAML-файла, разделенных тремя дефисами (http://yaml.org/spec/1.2/spec.html#id2800132):

```yaml
CONFIG_SECTION
---
CONFIG_SECTION
---
CONFIG_SECTION
```

Каждая секция конфигурации может быть одного типа из трех:

1. Секция для описания мета-информации проекта, далее — *секция мета-информации*.
2. Секция описания инструкций сборки образа, далее — *секция образа*.
3. Секция описания инструкций сборки артефакта, далее — *секция артефакта*.

В будущем, возможно, количество типов увеличится.

### Секция мета-информации

```
project: PROJECT_NAME
configVersion: CONFIG_VERSION
OTHER_FIELDS
---
```

Секция мета-информации, это **обязательная** секция конфигурации, содержащая ключи `project: PROJECT_NAME` и `configVersion: CONFIG_VERSION`. В каждом файле конфигурации `werf.yaml` должна быть только одна секция мета-информации.

#### Имя проекта

Ключ `project` определяет уникальное имя проекта вашего приложения. Имя проекта влияет на имена образов в сборочном кэше, namespace в Kubernetes, имя Helm-релиза и зависящие от него имена (смотри подробнее про [развертывание в Kubernetes]({{ site.baseurl }}/ru/documentation/reference/deploy_process/deploy_into_kubernetes.html)). Ключ `project` — единственное обязательное поле секции мета-информации.

Имя проекта должно быть уникальным в пределах группы проектов, собираемых на одном сборочном узле и развертываемых на один и тот-же кластер Kubernetes (например уникальным в пределах всех групп одного GitLab).

Имя проекта должно быть не более 50 символов, содержать только строчные буквы латинского алфавита, цифры и знак дефиса.

**ВНИМАНИЕ**. Никогда не меняйте имя проекта в процессе работы, если вы не осознаете всех последствий.

Смена имени проекта приводит к следующим проблемам:
1. Инвалидация сборочного кэша. Все образы должны быть собранные повторно, а старые — удалены из локального хранилища или Docker registry вручную.
2. Создание совершенно нового helm-релиза. Если вы уже развернули ваше приложение в кластере Kubernetes, смена имени проекта и повторное его развертывание приведет к созданию еще одного экземпляра приложения.

Werf не поддерживает изменение имени проекта и все возникающие проблемы должны быть решены вручную.

#### Версии конфигурации

Директива `configVersion` определяет формат файла `werf.yaml`. В настоящее время, это всегда — `1`.

### Image config section

Each image config section defines instructions to build one independent docker image. There may be multiple image config sections defined in the same `werf.yaml` config to build multiple images.

Config section with the key `image: IMAGE_NAME` is the image config section. `image` defines short name of the docker image to be built. This name must be unique in a single `werf.yaml` config.

```
image: IMAGE_NAME_1
OTHER_FIELDS
---
image: IMAGE_NAME_2
OTHER_FIELDS
---
...
---
image: IMAGE_NAME_N
OTHER_FIELDS
```

### Artifact config section

Artifact config section also defines instructions to build one independent artifact docker image. Arifact is a secondary image aimed to isolate a build process and build tools resources (environments, software, data, see [artifacts article for the details]({{ site.baseurl }}/documentation/configuration/stapel_artifact.html)). There may be multiple artifact config sections for multiple artifact config sections defined in the same `werf.yaml` config.

Config section with the key `artifact: IMAGE_NAME` is the artifact config section. `artifact` defines short name of the artifact to be referred to from another config sections. This name must be unique in a single `werf.yaml` config.

### Minimal config example

Currently Werf requires to define meta config section and at least one image config section. Image config sections will be fully optional soon.

Example of minimal werf config:

```yaml
project: my-project
configVersion: 1
---
image: ~
from: alpine:latest
```

## Organizing configuration

Part of the configuration can be moved in ***separate template files*** and then included into __werf.yaml__. _Template files_ should live in the ***.werf*** directory with **.tmpl** extension (any nesting is supported).

> **Tip:** templates can be generated or downloaded before running werf. For example, for sharing common logic between projects.

Werf parses all files in one environment, thus described [define](#include) of one _template file_ becomes available in other files, including _werf.yaml_.

<div class="details active">
<a href="javascript:void(0)" class="details__summary">werf.yaml</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
{{ $_ := set . "RubyVersion" "2.3.4" }}
{{ $_ := set . "BaseImage" "alpine" }}

project: my-project
configVersion: 1
---

image: rails
from: {{ .BaseImage }}
ansible:
  beforeInstall:
  {{- include "(component) mysql client" . }}
  {{- include "(component) ruby" . }}
  install:
  {{- include "(component) Gemfile dependencies" . }}
```
{% endraw %}

</div>
</div>

<div class="details">
<a href="javascript:void(0)" class="details__summary">.werf/ansible/components.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
{{- define "(component) Gemfile dependencies" }}
  - file:
      path: /root/.ssh
      state: directory
      owner: root
      group: root
      recurse: yes
  - name: "Setup ssh known_hosts used in Gemfile"
    shell: |
      set -e
      ssh-keyscan github.com >> /root/.ssh/known_hosts
      ssh-keyscan mygitlab.myorg.com >> /root/.ssh/known_hosts
    args:
      executable: /bin/bash
  - name: "Install Gemfile dependencies with bundler"
    shell: |
      set -e
      source /etc/profile.d/rvm.sh
      cd /app
      bundle install --without development test --path vendor/bundle
    args:
      executable: /bin/bash
{{- end }}

{{- define "(component) mysql client" }}
  - name: "Install mysql client"
    apt:
      name: "{{`{{ item }}`}}"
      update_cache: yes
    with_items:
    - libmysqlclient-dev
    - mysql-client
    - g++
{{- end }}

{{- define "(component) ruby" }}
  - command: gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
  - get_url:
      url: https://raw.githubusercontent.com/rvm/rvm/master/binscripts/rvm-installer
      dest: /tmp/rvm-installer
  - name: "Install rvm"
    command: bash -e /tmp/rvm-installer
  - name: "Install ruby {{ .RubyVersion }}"
    raw: bash -lec {{`{{ item | quote }}`}}
    with_items:
    - rvm install {{ .RubyVersion }}
    - rvm use --default {{ .RubyVersion }}
    - gem install bundler --no-ri --no-rdoc
    - rvm cleanup all
{{- end }}
```
{% endraw %}

</div>
</div>

> If there are templates with the same name werf will use template defined in _werf.yaml_ or the latest described in _templates files_.

If need to use the whole _template file_, use template file path relative to _.werf_ directory as a template name in [include](#include) function.

<div class="details active">
<a href="javascript:void(0)" class="details__summary">werf.yaml</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
project: my-project
configVersion: 1
---
image: app
from: java:8-jdk-alpine
shell:
  beforeInstall:
  - mkdir /app
  - adduser -Dh /home/gordon gordon
import:
- artifact: appserver
  add: '/usr/src/atsea/target/AtSea-0.0.1-SNAPSHOT.jar'
  to: '/app/AtSea-0.0.1-SNAPSHOT.jar'
  after: install
- artifact: storefront
  add: /usr/src/atsea/app/react-app/build
  to: /static
  after: install
docker:
  ENTRYPOINT: ["java", "-jar", "/app/AtSea-0.0.1-SNAPSHOT.jar"]
  CMD: ["--spring.profiles.active=postgres"]
---
{{ include "artifact/appserver.tmpl" . }}
---
{{ include "artifact/storefront.tmpl" . }}
```
{% endraw %}

</div>
</div>

<div class="details">
<a href="javascript:void(0)" class="details__summary">.werf/artifact/appserver.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
artifact: appserver
from: maven:latest
git:
- add: '/app'
  to: '/usr/src/atsea'
shell:
  install:
  - cd /usr/src/atsea
  - mvn -B -f pom.xml -s /usr/share/maven/ref/settings-docker.xml dependency:resolve
  - mvn -B -s /usr/share/maven/ref/settings-docker.xml package -DskipTests
```
{% endraw %}

</div>
</div>

<div class="details active">
<a href="javascript:void(0)" class="details__summary">.werf/artifact/storefront.tmpl</a>
<div class="details__content" markdown="1">

{% raw %}
```yaml
artifact: storefront
from: node:latest
git:
- add: /app/react-app
  to: /usr/src/atsea/app/react-app
shell:
  install:
  - cd /usr/src/atsea/app/react-app
  - npm install
  - npm run build
```
{% endraw %}

</div>
</div>

## Processing of config

The following steps could describe the processing of a YAML configuration file:
1. Reading `werf.yaml` and extra templates from `.werf` directory;
2. Executing Go templates;
3. Saving dump into `.werf.render.yaml` (that file will remain after build and will be available until next render);
4. Splitting rendered YAML file into separate config sections (part of YAML stream separated by three hyphens, https://yaml.org/spec/1.2/spec.html#id2800132);
5. Validating each config section:
  * Validating YAML syntax (you could read YAML reference [here](http://yaml.org/refcard.html)).
  * Validating werf syntax.
6. Generating a set of images.

### Go templates

Go templates are available within YAML configuration. The following functions are supported:

* [Built-in Go template functions](https://golang.org/pkg/text/template/#hdr-Functions) and other language features. E.g. using common variable:<a id="go-templates" href="#go-templates" class="anchorjs-link " aria-label="Anchor link for: go templates" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  {{ $base_image := "golang:1.11-alpine" }}

  project: my-project
  configVersion: 1
  ---

  image: gogomonia
  from: {{ $base_image }}
  ---
  image: saml-authenticator
  from: {{ $base_image }}
  ```
  {% endraw %}

* [Sprig functions](http://masterminds.github.io/sprig/). E.g. using environment variable:<a id="sprig-functions" href="#sprig-functions" class="anchorjs-link " aria-label="Anchor link for: sprig functions" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  {{ $_ := env "SPECIFIC_ENV_HERE" | set . "GitBranch" }}

  image: ~
  from: alpine
  git:
  - url: https://github.com/company/project1.git
    branch: {{ .GitBranch }}
    add: /
    to: /app/project1
  - url: https://github.com/company/project2.git
    branch: {{ .GitBranch }}
    add: /
    to: /app/project2
  ```
  {% endraw %}

* `include` function with `define` for reusing configs:<a id="include" href="#include" class="anchorjs-link " aria-label="Anchor link for: include" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  image: app1
  from: alpine
  ansible:
    beforeInstall:
    {{- include "(component) ruby" . }}
  ---
  image: app2
  from: alpine
  ansible:
    beforeInstall:
    {{- include "(component) ruby" . }}

  {{- define "(component) ruby" }}
    - command: gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
    - get_url:
        url: https://raw.githubusercontent.com/rvm/rvm/master/binscripts/rvm-installer
        dest: /tmp/rvm-installer
    - name: "Install rvm"
      command: bash -e /tmp/rvm-installer
    - name: "Install ruby 2.3.4"
      raw: bash -lec {{`{{ item | quote }}`}}
      with_items:
      - rvm install 2.3.4
      - rvm use --default 2.3.4
      - gem install bundler --no-ri --no-rdoc
      - rvm cleanup all
  {{- end }}
  ```
  {% endraw %}

* `.Files.Get` function for getting project file content:<a id="files-get" href="#files-get" class="anchorjs-link " aria-label="Anchor link for: .Files.Get" data-anchorjs-icon=""></a>

  {% raw %}
  ```yaml
  project: my-project
  configVersion: 1
  ---

  image: app
  from: alpine
  ansible:
    setup:
    - name: "Setup /etc/nginx/nginx.conf"
      copy:
        content: |
  {{ .Files.Get ".werf/nginx.conf" | indent 8 }}
        dest: /etc/nginx/nginx.conf
  ```
  {% endraw %}
