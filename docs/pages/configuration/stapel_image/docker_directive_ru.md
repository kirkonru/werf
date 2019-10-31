---
title: Добавление Docker-инструкций
sidebar: documentation
permalink: ru/documentation/configuration/stapel_image/docker_directive.html
ref: documentation_configuration_stapel_image_docker_directive
lang: ru
author: Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vTZB0BLxL7mRUFxkrOMaj310CQgb5D5H_V0gXe7QYsTu3kKkdwchg--A1EoEP2CtKbO8pp2qARfeoOK/pub?w=2031&amp;h=144" data-featherlight="image">
    <img src="https://docs.google.com/drawings/d/e/2PACX-1vTZB0BLxL7mRUFxkrOMaj310CQgb5D5H_V0gXe7QYsTu3kKkdwchg--A1EoEP2CtKbO8pp2qARfeoOK/pub?w=1016&amp;h=72">
  </a>

    <div class="language-yaml highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="na">docker</span><span class="pi">:</span>
    <span class="na">VOLUME</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;volume&gt;</span>
    <span class="na">EXPOSE</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;expose&gt;</span>
    <span class="na">ENV</span><span class="pi">:</span>
      <span class="s">&lt;env_name&gt;</span><span class="pi">:</span> <span class="s">&lt;env_value&gt;</span>
    <span class="na">LABEL</span><span class="pi">:</span>
      <span class="s">&lt;label_name&gt;</span><span class="pi">:</span> <span class="s">&lt;label_value&gt;</span>
    <span class="na">ENTRYPOINT</span><span class="pi">:</span> <span class="s">&lt;entrypoint&gt;</span>
    <span class="na">CMD</span><span class="pi">:</span> <span class="s">&lt;cmd&gt;</span>
    <span class="na">WORKDIR</span><span class="pi">:</span> <span class="s">&lt;workdir&gt;</span>
    <span class="na">USER</span><span class="pi">:</span> <span class="s">&lt;user&gt;</span>
    <span class="na">STOPSIGNAL</span><span class="pi">:</span> <span class="s">&lt;stopsignal&gt;</span>
    <span class="na">HEALTHCHECK</span><span class="pi">:</span> <span class="s">&lt;healthcheck&gt;</span></code></pre></div></div>
---

Docker can build images by [Dockerfile](https://docs.docker.com/engine/reference/builder/) instructions. These instructions can be divided into two groups: build-time instructions and other instructions that effect on an image manifest.  

Build-time instructions do not make sense in a werf build process. Therefore, werf supports only following instructions:

* `USER` to set the user and the group to use when running the image (read more [here](https://docs.docker.com/engine/reference/builder/#user)).
* `WORKDIR` to set the working directory (read more [here](https://docs.docker.com/engine/reference/builder/#workdir)).
* `VOLUME` to add mount point (read more [here](https://docs.docker.com/engine/reference/builder/#volume)).
* `ENV` to set the environment variable (read more [here](https://docs.docker.com/engine/reference/builder/#env)).
* `LABEL` to add metadata to an image (read more [here](https://docs.docker.com/engine/reference/builder/#label)).
* `EXPOSE` to inform Docker that the container listens on the specified network ports at runtime (read more [here](https://docs.docker.com/engine/reference/builder/#expose))
* `ENTRYPOINT` to configure a container that will run as an executable (read more [here](https://docs.docker.com/engine/reference/builder/#entrypoint)).
* `CMD` to provide default arguments for the `ENTRYPOINT` to configure a container that will run as an executable (read more [here](https://docs.docker.com/engine/reference/builder/#cmd)).
* `STOPSIGNAL` to set the system call signal that will be sent to the container to exit (read more [here](https://docs.docker.com/engine/reference/builder/#stopsignal))
* `HEALTHCHECK` to tell Docker how to test a container to check that it is still working (read more [here](https://docs.docker.com/engine/reference/builder/#healthcheck))

These instructions can be specified in the `docker` config directive.

Here is an example of using docker instructions:

```yaml
docker:
  WORKDIR: /app
  CMD: ['python', './index.py']
  EXPOSE: '5000'
  ENV:
    TERM: xterm
    LC_ALL: en_US.UTF-8
```

Defined docker instructions are applied on the last stage called `docker_instructions`.
Thus, instructions do not affect other stages, ones just will be applied to a built image.

If need to use special environment variables in build-time of your application image, such as `TERM` environment, you should use a [base image]({{ site.baseurl }}/documentation/configuration/stapel_image/base_image.html) with these variables.

> Tip: you can also implement exporting environment variables right in [_user stage_]({{ site.baseurl }}/documentation/configuration/stapel_image/assembly_instructions.html#what-is-user-stages) instructions
