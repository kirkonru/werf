---
title: Артефакты Stapel
sidebar: documentation
permalink: ru/documentation/configuration/stapel_artifact.html
ref: documentation_configuration_stapel_artifact
lang: ru
author: Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vRD-K_z7KEoliEVT4GpTekCkeaFMbSPWZpZkyTDms4XLeJAWEnnj4EeAxsdwnU3OtSW_vuKxDaaFLgD/pub?w=1800&amp;h=850" data-featherlight="image">
  <img src="https://docs.google.com/drawings/d/e/2PACX-1vRD-K_z7KEoliEVT4GpTekCkeaFMbSPWZpZkyTDms4XLeJAWEnnj4EeAxsdwnU3OtSW_vuKxDaaFLgD/pub?w=640&amp;h=301">
  </a>
---

## Что такое артефакты?

***Артефакты*** — это специальный образ, используемый в других артефактах или отдельных образах, описанных в конфигурации, предназначенный преимущественно для отделения ресурсов инструментов сборки от процесса сборки образа приложения. Примерами таких ресурсов могут быть — данные или программное обеспечение, которые необходимы для сборки, но не нужны для запуска приложения, и т.п.

Образ _артефакта_ нельзя [протэгировать]({{ site.baseurl }}/documentation/reference/publish_process.html) как обычный образ, и использовать как отдельное приложение.

Используя артефакты, вы можете собирать неограниченное количество компонентов, что позволяет решать, например, следующие задачи:
- Если приложение состоит из набора компонент, каждый со своими зависимостями от других компонент, то обычно вам приходится пересобирать все компоненты каждый раз. Вам бы хотелось пересобирать только те компоненты, которым это действительно нужно.
- Компоненты должны быть собраны в разных окружениях.

Importing _resources_ from _artifacts_ are described in [import directive]({{ site.baseurl }}/documentation/configuration/stapel_image/import_directive.html) in _destination image_ config section ([_image_]({{ site.baseurl }}/documentation/configuration/introduction.html#image-config-section) or [_artifact_]({{ site.baseurl }}/documentation/configuration/introduction.html#artifact-config-section)).

## Configuration

The configuration of the _artifact_ is not much different from the configuration of _image_. Each _artifact_ should be described in a separate [artifact config section]({{ site.baseurl }}/documentation/configuration/introduction.html#artifact-config-section).

The instructions associated with the _from stage_, namely the [_base image_]({{ site.baseurl }}/documentation/configuration/stapel_image/base_image.html) and [mounts]({{ site.baseurl }}/documentation/configuration/stapel_image/mount_directive.html), and also [imports]({{ site.baseurl }}/documentation/configuration/stapel_image/import_directive.html) remain unchanged.

The _docker_instructions stage_ and the [corresponding instructions]({{ site.baseurl }}/documentation/configuration/stapel_image/docker_directive.html) are not supported for the _artifact_. An _artifact_ is an assembly tool and only the data stored in it is required.

The remaining _stages_ and instructions are considered further separately.

### Naming

<div class="summary" markdown="1">
```yaml
artifact: <artifact name>
```
</div>

_Artifact images_ are declared with `artifact` directive: `artifact: <artifact name>`. Unlike the [naming of the _image_]({{ site.baseurl }}/documentation/configuration/stapel_image/naming.html), the artifact has no limitations associated with docker naming convention, as used only internal.

```yaml
artifact: "application assets"
```

### Adding source code from git repositories

<div class="summary">

<a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vQQiyUM9P3-_A6O5tLms_y1UCny9X6lxQSxtMtBalcyjcHhYV4hnPnISmTVY09c-ANOFqwHeOxYPs63/pub?w=2031&amp;h=144" data-featherlight="image">
<img src="https://docs.google.com/drawings/d/e/2PACX-1vQQiyUM9P3-_A6O5tLms_y1UCny9X6lxQSxtMtBalcyjcHhYV4hnPnISmTVY09c-ANOFqwHeOxYPs63/pub?w=1016&amp;h=72">
</a>

</div>

Unlike with _image_, _artifact stage conveyor_ has no _gitCache_ and _gitLatestPatch_ stages.

> Werf implements optional dependence on changes in git repositories for _artifacts_. Thus, by default werf ignores them and _artifact image_ is cached after the first assembly, but you can specify any dependencies for assembly instructions.

Read about working with _git repositories_ in the corresponding [article]({{ site.baseurl }}/documentation/configuration/stapel_image/git_directive.html).

### Running assembly instructions

<div class="summary">

<a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vTlpKbAr6wQCE4bSxVB5Kr6uxzbCGu_ncsviT2Ax6_qLL3zAVLWIsYElAi9_LMuVeFiDi1lo97HNvD2/pub?w=1428&h=649" data-featherlight="image">
      <img src="https://docs.google.com/drawings/d/e/2PACX-1vTlpKbAr6wQCE4bSxVB5Kr6uxzbCGu_ncsviT2Ax6_qLL3zAVLWIsYElAi9_LMuVeFiDi1lo97HNvD2/pub?w=426&h=216">
</a>

</div>

Directives and _user stages_ remain unchanged: _beforeInstall_, _install_, _beforeSetup_ and _setup_.

If there are no dependencies on files specified in git `stageDependencies` directive for _user stages_, the image is cached after the first build and will no longer be reassembled while the related _stages_ exist in _stages storage_.

> If the artifact should be rebuilt on any change in the related git repository, you should specify the _stageDependency_ `**/*` for any _user stage_, e.g., for _install stage_:
```yaml
git:
- to: /
  stageDependencies:
    install: "**/*"
```

Read about working with _assembly instructions_ in the corresponding [article]({{ site.baseurl }}/documentation/configuration/stapel_image/assembly_instructions.html).

## All directives
```yaml
artifact: <artifact_name>
from: <image>
fromLatest: <bool>
fromCacheVersion: <version>
fromImage: <image_name>
fromImageArtifact: <artifact_name>
git:
# local git
- add: <absolute path in git repository>
  to: <absolute path inside image>
  owner: <owner>
  group: <group>
  includePaths:
  - <path or glob relative to path in add>
  excludePaths:
  - <path or glob relative to path in add>
  stageDependencies:
    install:
    - <path or glob relative to path in add>
    beforeSetup:
    - <path or glob relative to path in add>
    setup:
    - <path or glob relative to path in add>
# remote git
- url: <git repo url>
  branch: <branch name>
  commit: <commit>
  tag: <tag>
  add: <absolute path in git repository>
  to: <absolute path inside image>
  owner: <owner>
  group: <group>
  includePaths:
  - <path or glob relative to path in add>
  excludePaths:
  - <path or glob relative to path in add>
  stageDependencies:
    install:
    - <path or glob relative to path in add>
    beforeSetup:
    - <path or glob relative to path in add>
    setup:
    - <path or glob relative to path in add>
shell:
  beforeInstall:
  - <cmd>
  install:
  - <cmd>
  beforeSetup:
  - <cmd>
  setup:
  - <cmd>
  cacheVersion: <version>
  beforeInstallCacheVersion: <version>
  installCacheVersion: <version>
  beforeSetupCacheVersion: <version>
  setupCacheVersion: <version>
ansible:
  beforeInstall:
  - <task>
  install:
  - <task>
  beforeSetup:
  - <task>
  setup:
  - <task>
  cacheVersion: <version>
  beforeInstallCacheVersion: <version>
  installCacheVersion: <version>
  beforeSetupCacheVersion: <version>
  setupCacheVersion: <version>
mount:
- from: build_dir
  to: <absolute_path>
- from: tmp_dir
  to: <absolute_path>
- fromPath: <absolute_or_relative_path>
  to: <absolute_path>
import:
- artifact: <artifact name>
  image: <image name>
  before: <install || setup>
  after: <install || setup>
  add: <absolute path>
  to: <absolute path>
  owner: <owner>
  group: <group>
  includePaths:
  - <relative path or glob>
  excludePaths:
  - <relative path or glob>
asLayers: <bool>
```

## Using artifacts

Unlike [*stapel image*]({{ site.baseurl }}/documentation/configuration/stapel_image/assembly_instructions.html), *stapel artifact* does not have a git latest patch stage.

Git latest patch stage is supposed to be updated on every commit, which brings new changes to files. *Stapel artifact* though is recommended to be used as a deeply cached image, which will be updated in rare cases, when some special files changed.

For example: import git into *stapel artifact* and rebuild assets in this artifact only when dependent assets files in git has changes. For every other change in git where non-dependent files has been changed assets will not be rebuilt.

However in the case when there is a need to bring changes of any git files into *stapel artifact* (to build golang application for example) user should define `git.stageDependencies` of some stage that needs these files explicitly as `*` pattern:

```
git:
- add: /
  to: /app
  stageDependencies:
    setup:
    - "*"
```

In this case every change in git files will result in artifact rebuild, all *stapel images* that import this artifact will also be rebuilt.

**NOTE** User should employ multiple separate `git.add` directive invocations in every [*stapel image*]({{ site.baseurl }}/documentation/configuration/stapel_image/assembly_instructions.html) and *stapel artifact* that needs git files — it is an optimal way to add git files into any image. Adding git files to artifact and then importing it into image using `import` directive is not recommended.
