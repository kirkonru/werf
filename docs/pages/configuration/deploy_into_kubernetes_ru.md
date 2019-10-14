---
title: Деплой в Kubernetes
sidebar: documentation
permalink: ru/documentation/configuration/deploy_into_kubernetes.html
author: Timofey Kirillov <timofey.kirillov@flant.com>
ref: documentation_configuration_deploy_into_kubernetes
lang: ru
---

## Имя релиза

Werf позволяет определять пользовательский шаблон имени helm-релиза, который используется во время [процесса деплоя]({{ site.baseurl }}/ru/documentation/reference/deploy_process/deploy_into_kubernetes.html#release-name) для генерации имени релиза.

Пользовательский шаблон имени релиза определяется в [секции мета-информации]({{ site.baseurl }}/ru/documentation/configuration/introduction.html#meta-config-section) в файле `werf.yaml`:

```yaml
project: PROJECT_NAME
configVersion: 1
deploy:
  helmRelease: TEMPLATE
  helmReleaseSlug: false
```

`deploy.helmReleaseSlug` включает или отключает [slug'ификацию]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#release-name-slug) имени helm-релиза. Включен по умолчанию.

В качестве значения для `deploy.helmRelease` указывается Go-шаблон с разделителям `[[` и `]]`. Поддерживаюся функции `project` и `env`. Значение шаблона имени релиза по умолчанию: `[[ project ]]-[[ env ]]`.

Т.к. в качестве значения для `deploy.helmRelease` указывается Go-шаблон, то возможно использование соответствующих [функций Werf]({{ site.baseurl }}/ru/documentation/configuration/introduction.html#go-templates-1) (впрочем, как и для любых других параметров в конфигурации). Например, вы можете комбинировать переменные доступные в Go-шаблоне с переменными окружения следующим образом:
{% raw %}
```yaml
deploy:
  helmRelease: >-
    [[ project ]]-{{ env "HELM_RELEASE_EXTRA" }}-[[ env ]]
```
{% endraw %}

## Kubernetes namespace

Werf позволяет определять пользовательский шаблон namespace в Kubernetes, который будет использоваться во время [процесса деплоя]({{ site.baseurl }}/ru/documentation/reference/deploy_process/deploy_into_kubernetes.html#kubernetes-namespace) для генерации имени namespace.

Пользовательский шаблон namespace Kubernetes определяется в [секции мета-информации]({{ site.baseurl }}/ru/documentation/configuration/introduction.html#meta-config-section) в файле `werf.yaml`:


```yaml
project: PROJECT_NAME
configVersion: 1
deploy:
  namespace: TEMPLATE
  namespaceSlug: true|false
```

В качестве значения для `deploy.namespace` указывается Go-шаблон с разделителям `[[` и `]]`. Поддерживаюся функции `project` и `env`. Значение шаблона имени namespace по умолчанию: `[[ project ]]-[[ env ]]`.

`deploy.namespaceSlug` включает или отключает [slug'ификацию]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#kubernetes-namespace-slug) имени namespace Kubernetes. Включен по умолчанию.
