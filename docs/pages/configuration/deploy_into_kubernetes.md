---
title: Deploy into kubernetes
sidebar: documentation
permalink: documentation/configuration/deploy_into_kubernetes.html
ref: documentation_configuration_deploy_into_kubernetes
author: Timofey Kirillov <timofey.kirillov@flant.com>
---

## Release name

Werf allows to define a custom release name template, which [used during deploy process]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#release-name) to generate a release name.

Custom release name template is defined in the [meta configuration section]({{ site.baseurl }}/documentation/configuration/introduction.html#meta-config-section) of `werf.yaml`:

```yaml
project: PROJECT_NAME
configVersion: 1
deploy:
  helmRelease: TEMPLATE
  helmReleaseSlug: false
```

`deploy.helmRelease` is a Go template with `[[` and `]]` delimiters. There are `[[ project ]]`, `[[ env ]]` functions support. Default: `[[ project ]]-[[ env ]]`.

`deploy.helmReleaseSlug` defines whether to apply or not [slug]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#release-name-slug) to generated helm release name. Default: `true`.

`TEMPLATE` as well as any value of the config can include [Werf Go templates functions]({{ site.baseurl }}/documentation/configuration/introduction.html#go-templates-1). E.g. you can mix the value with an environment variable:

{% raw %}
```yaml
deploy:
  helmRelease: >-
    [[ project ]]-{{ env "HELM_RELEASE_EXTRA" }}-[[ env ]]
```
{% endraw %}

## Kubernetes namespace

Werf allows to define a custom kubernetes namespace template, which [used during deploy process]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#kubernetes-namespace) to generate a Kubernetes Namespace.

Custom Kubernetes Namespace template is defined in the [meta configuration section]({{ site.baseurl }}/documentation/configuration/introduction.html#meta-config-section) of `werf.yaml`:

```yaml
project: PROJECT_NAME
configVersion: 1
deploy:
  namespace: TEMPLATE
  namespaceSlug: true|false
```

`deploy.namespace` is a Go template with `[[` and `]]` delimiters. There are `[[ project ]]`, `[[ env ]]` functions support. Default: `[[ project ]]-[[ env ]]`.

`deploy.namespaceSlug` defines whether to apply or not [slug]({{ site.baseurl }}/documentation/reference/deploy_process/deploy_into_kubernetes.html#kubernetes-namespace-slug) to generated kubernetes namespace. Default: `true`.
