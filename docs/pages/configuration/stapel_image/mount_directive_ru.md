---
title: Использование монтирования для уменьшения размера образа и ускорения сборки
sidebar: documentation
permalink: ru/documentation/configuration/stapel_image/mount_directive.html
ref: documentation_configuration_stapel_image_mount_directive
lang: ru
author: Artem Kladov <artem.kladov@flant.com>, Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vReDSY8s7mMtxuxwDTwtPLFYjEXePaoIB-XbEZcunJGNEHrLbrb9aFxyOoj_WeQe0XKQVhq7RWnG3Eq/pub?w=2031&amp;h=144" data-featherlight="image">
      <img src="https://docs.google.com/drawings/d/e/2PACX-1vReDSY8s7mMtxuxwDTwtPLFYjEXePaoIB-XbEZcunJGNEHrLbrb9aFxyOoj_WeQe0XKQVhq7RWnG3Eq/pub?w=1016&amp;h=72">
  </a>

  <div class="language-yaml highlighter-rouge"><pre class="highlight"><code><span class="s">mount</span><span class="pi">:</span>
  <span class="pi">-</span> <span class="s">from</span><span class="pi">:</span> <span class="s">tmp_dir</span>
    <span class="s">to</span><span class="pi">:</span> <span class="s">&lt;absolute_path&gt;</span>
  <span class="pi">-</span> <span class="s">from</span><span class="pi">:</span> <span class="s">build_dir</span>
    <span class="s">to</span><span class="pi">:</span> <span class="s">&lt;absolute_path&gt;</span>
  <span class="pi">-</span> <span class="s">fromPath</span><span class="pi">:</span> <span class="s">&lt;absolute_or_relative_path&gt;</span>
    <span class="s">to</span><span class="pi">:</span> <span class="s">&lt;absolute_path&gt;</span></code></pre>
  </div>
---

Довольно часто бывают случаи, когда при сборке у вас появляются файлы которые нет необходимости оставлять в образе, и их нужно исключить. Например:
- Большинство пакетных менеджеров создают в системе кэш пакетов и служебных файлов.
  - [APT](https://wiki.debian.org/Apt) хранит список пакетов в директории `/var/lib/apt/lists/`.
  - APT также хранит сами пакеты в директории `/var/cache/apt/`.
  - [YUM](http://yum.baseurl.org/) может оставлять скачанные пакеты в директории `/var/cache/yum/.../packages/`.
- Менеджеры пакетов, используемые разработчиками, например ​npm (Node.js), Glide (Go), pip (Python) хранят файлы в директории кэша.
- Компиляция приложений на C/C++ и т.п. языках оставляет объектные файлы и т.п.

Такого рода файлы:
- не нужны в образе;
- могут значительно увеличивать размер образа;
- могут быть полезны при повторнй сборке образа или других образов.

Если монтировать такого рода файлы в сборочный контейнер, то можно добиться не только уменьшения размера образа, но и ускорения процесса сборки. В Docker такой механизм реализуется с помощью функционала [томов](https://docs.docker.com/storage/volumes/).

Для указания тома используется директива `mount`. Директории узла сборки монтируются в сборочный контейнер согласно директив `from`/`fromPath` и `to` описания томов. Для указания в качестве точки монтирования на сборочном узле любого файла или директории, вы можете использовать директиву `fromPath`. Либо, используя директиву `from`, вы можете указать одну из следующих служебных директорий:
- `tmp_dir` временная директория, индивидуальная для каждого описанного образа, создаваемая заново при каждой сборке;
- `build_dir` общая директория, доступная всем образам проекта и сохраняемая между сборками (находится по пути — `~/.werf/shared_context/mounts/projects/<project name>/<mount id>/`). Вы можете использовать эту директорию для хранения, например, кэша и т.п.

> Werf монтирует служебные директории с возможностью чтения и записи при каждой сборке, но в образе содержимого этих директорий не будет. Если вам необходимо сохранить какие-либо данные из этих директорий непосредственно в образе, то вы должны их скопировать при сборке.

На стадии `from`, Werf добавляет специальные метки к образу стадии, согласно описанных точек монтирования. Затем, на каждой стадии, Werf использует эти метки при  монтировании директорий в сборочный контейнер. Такая реализация позволяет наследовать точки монтирования от [базового образа]({{ site.baseurl }}/ru/documentation/configuration/stapel_image/base_image.html).

Также, нужно иметь ввиду, что на стадии `from` Werf очищает точки монтирования в [базовом образе]({{ site.baseurl }}/ru/documentation/configuration/stapel_image/base_image.html) (т.е. эти папки будут пусты).
