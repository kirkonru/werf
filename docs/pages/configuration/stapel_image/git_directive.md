---
title: Adding source code from git repositories
sidebar: documentation
permalink: documentation/configuration/stapel_image/git_directive.html
ref: documentation_configuration_stapel_image_git_directive
summary: |
  <a class="google-drawings" href="https://docs.google.com/drawings/d/e/2PACX-1vRUYmRNmeuP14OcChoeGzX_4soCdXx7ZPgNqm5ePcz9L_ItMUqyolRoJyPL7baMNoY7P6M0B08eMtsb/pub?w=2031&amp;h=144" data-featherlight="image">
      <img src="https://docs.google.com/drawings/d/e/2PACX-1vRUYmRNmeuP14OcChoeGzX_4soCdXx7ZPgNqm5ePcz9L_ItMUqyolRoJyPL7baMNoY7P6M0B08eMtsb/pub?w=1016&amp;h=72">
  </a>

  <div class="tabs">
    <a href="javascript:void(0)" class="tabs__btn active" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'local')">Local</a>
    <a href="javascript:void(0)" class="tabs__btn" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'remote')">Remote</a>
  </div>

  <div id="local" class="tabs__content active">
  <div class="language-yaml highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="na">git</span><span class="pi">:</span>
  <span class="pi">-</span> <span class="na">add</span><span class="pi">:</span> <span class="s">&lt;absolute path in git repository&gt;</span>
    <span class="na">to</span><span class="pi">:</span> <span class="s">&lt;absolute path inside image&gt;</span>
    <span class="na">owner</span><span class="pi">:</span> <span class="s">&lt;owner&gt;</span>
    <span class="na">group</span><span class="pi">:</span> <span class="s">&lt;group&gt;</span>
    <span class="na">includePaths</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
    <span class="na">excludePaths</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
    <span class="na">stageDependencies</span><span class="pi">:</span>
      <span class="na">install</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
      <span class="na">beforeSetup</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
      <span class="na">setup</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span></code></pre>
  </div></div>     
  </div>

  <div id="remote" class="tabs__content">
  <div class="language-yaml highlighter-rouge"><div class="highlight"><pre class="highlight"><code><span class="na">git</span><span class="pi">:</span>
  <span class="pi">-</span> <span class="na">url</span><span class="pi">:</span> <span class="s">&lt;git repo url&gt;</span>
    <span class="na">branch</span><span class="pi">:</span> <span class="s">&lt;branch name&gt;</span>
    <span class="na">commit</span><span class="pi">:</span> <span class="s">&lt;commit&gt;</span>
    <span class="na">tag</span><span class="pi">:</span> <span class="s">&lt;tag&gt;</span>
    <span class="na">add</span><span class="pi">:</span> <span class="s">&lt;absolute path in git repository&gt;</span>
    <span class="na">to</span><span class="pi">:</span> <span class="s">&lt;absolute path inside image&gt;</span>
    <span class="na">owner</span><span class="pi">:</span> <span class="s">&lt;owner&gt;</span>
    <span class="na">group</span><span class="pi">:</span> <span class="s">&lt;group&gt;</span>
    <span class="na">includePaths</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
    <span class="na">excludePaths</span><span class="pi">:</span>
    <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
    <span class="na">stageDependencies</span><span class="pi">:</span>
      <span class="na">install</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
      <span class="na">beforeSetup</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
      <span class="na">setup</span><span class="pi">:</span>
      <span class="pi">-</span> <span class="s">&lt;path or glob relative to path in add&gt;</span>
  </code></pre>
  </div></div>
  </div>
---

## What is git mapping?

***Git mapping*** describes a file or directory from the git repository that should be added to the image by a specific path. The repository may be a local one, hosted in the directory that contains the config, or a remote one, and in this case, the configuration of the _git mapping_ contains the repository address and the version (branch, tag or commit hash).

Werf adds the files from the repository to the image by using the full transfer of files with git archive or by applying patches between commits.
The full transfer is used for the initial adding of files. The subsequent builds use applying patches to reflect changes in a git repository. The algorithm behind the full transfer and applying patches is reviewed the [More details: git_archive...](#more-details-gitarchive-gitcache-gitlatestpatch) section.

The configuration of the _git mapping_ supports filtering files, and you can use the set of _git mappings_ to create virtually any resulting file structure in the image. Also, you can specify the owner and the group of files in the _git mapping_ configuration — no subsequent `chown` required.

Werf has support for submodules. Werf detects if files specified with _git mapping_ configuration are contained in submodules and does the very best it could to handle the changes of files in submodules correctly.

> All project's submodules locked to a specific commit, so all collaborators receive the same content. Therefore, Werf **does not initialize, update submodules** and just uses these commits

An example of a _git mapping_ configuration for adding source files from a local repository from the `/src` into the `/app` directory, and remote phantomjs source files to `/src/phantomjs`:

```yaml
git:
- add: /src
  to: /app
- url: https://github.com/ariya/phantomjs
  add: /
  to: /src/phantomjs
```

## Motivation for git mappings

The main idea is to bring git history into the build process.

### Patching instead of copying

Most commits in the real application repository relate to updating the code of the application itself. In this case, if the compilation is not required, assembling a new image shall be nothing more than applying patches to the files in the previous image.

### Remote repositories

Building an application image may depend on source files in other repositories. Werf provides the ability to add files from remote repositories too. Werf can detect changes in local repositories and remote repositories.

## Syntax of a git mapping

The _git mapping_ configuration for a local repository has the following parameters:

- `add` — the path to a directory or file whose contents must be copied to the image. The path is specified relative to the repository root, and the path is absolute (i.e., it must start with `/`). This parameter is optional, the content of the entire repository is transferred by default, i.e., an empty `add` is equal to `add: /`;
- `to` — the path in the image, where the content specified with `add` will be copied;
- `owner` — the name or uid of the owner of the copied files;
- `group` — the name or gid of the group of the owner;
- `excludePaths` — a set of masks to ignore the files or directories during recursive copying. Paths in masks are specified relative to add;
- `includePaths` — a set of masks to include the files or directories during recursive copying. Paths in masks are specified relative to add;
- `stageDependencies` — a set of masks to detect changes that lead to the user stages rebuilds. This is reviewed in detail in the [Running assembly instructions]({{ site.baseurl }}/documentation/configuration/stapel_image/assembly_instructions.html) reference.

The _git mapping_ configuration for a remote repository has some additional parameters:
- `url` — remote repository address;
- `branch`, `tag`, `commit` — a name of branch, tag or commit hash that will be used. If these parameters are not specified, the master branch is used.

## Uses of git mappings

### Copying of directories

The `add` parameter specifies a path in a repository, from which all files must be recursively retrieved and added to the image with the `to` path; if the parameter is not specified, then the default path — `/` is used, i.e., the entire repository is transferred.
For example:

```yaml
git:
- add: /
  to: /app
```

This is the simple _git mapping_ configuration that adds the entire content from the repository to the `/app` directory in the image.

If the repository contains the following structure:

![git repository files tree]({{ site.baseurl }}/images/build/git_mapping_01.png)

Then the image contains this structure:

![image files tree]({{ site.baseurl }}/images/build/git_mapping_02.png)

Multiple _git mappings_ may be specified:

```yaml
git:
- add: /src
  to: /app/src
- add: /assets
  to: /static
```

If the repository contains the following structure:

![git repository files tree]({{ site.baseurl }}/images/build/git_mapping_03.png)

Then the image contains this structure:

![image files tree]({{ site.baseurl }}/images/build/git_mapping_04.png)

It should be noted, that _git mapping_ configuration doesn't specify a directory to be transferred like `cp -r /src /app`. `add` parameter specifies a directory content that will be recursively transferred from the repository. That is if the `/assets` directory needs to be transferred to the `/app/assets` directory, then the name **assets** should be written twice, or `includePaths` [filter](#using-filters) can be used.

```yaml
git:
- add: /assets
  to: /app/assets
```

or

```yaml
git:
- add: /
  to: /app
  includePaths: assets
```

> Werf has no convention for trailing `/` that is available in rsync, i.e. `add: /src` and `add: /src/` are the same.

### Copying of file

Copying the content, not the specified directory, from `add` path also applies to files. To transfer the file to the image, you must specify its name twice — once in `add`, and again in `to`. This provides an ability to rename the file:

```yaml
git:
- add: /config/prod.yaml
  to: /app/conf/production.yaml
```

### Changing an owner

The _git mapping_ configuration provides parameters `owner` and `group`. These are the names or numerical ids of the owner and group used for all files and directories transferred to the image.

```yaml
git:
- add: /src/index.php
  to: /app/index.php
  owner: www-data
```

![index.php owned by www-data user and group]({{ site.baseurl }}/images/build/git_mapping_05.png)

If only the `owner` parameter is specified, the group for files is the same as the primary group of the specified user.

If `owner` or `group` value is a string, then the specified user or group must be added to the system by the time of the full transfer of files, otherwise, build ends with an error.

```yaml
git:
- add: /src/index.php
  to: /app/index.php
  owner: wwwdata
```



### Using filters

`includePaths` and `excludePaths` parameters are used when processing the file list. These are the sets of masks that can be used to include and exclude files and directories from/to the list of files that will be transferred to the image. Simply stated, the `excludePaths` filter works as follows: masks are applied to each file found in `add` path. If at least one mask matches, then the file is ignored; if no matches are found, then the file gets added to the image. `includePaths` works the opposite way: if at least one mask is a match, the file gets added to the image.

_Git mapping_ configuration can contain both filters. In this case, a file is added to the image if the path matches with one of `includePaths` masks and not match with all `excludePaths` masks.

For example:

```yaml
git:
- add: /src
  to: /app
  includePaths:
  - '**/*.php'
  - '**/*.js'
  excludePaths:
  - '**/*-dev.*'
  - '**/*-test.*'
```

This is the git mapping configuration that adds `.php` and `.js` files from `/src` except files with suffixes that starts with `-dev.` or `-test.`.

To determine whether the file matches the mask the following algorithm is applied:
 - take for the check the next absolute file path inside the repository;
 - compare this path with configured include or exclude path mask or plain path:
   - the path in `add` is concatenated with the mask or raw path from include or exclude config directive;
   - two paths are compared with the use of glob patterns: if file matches the mask, then it will be included (for `includePaths`) or excluded (for `excludePaths`), the algorithm is ended.
 - compare this path with configured include or exclude path mask or plain path with additional pattern:
   - the path in `add` is concatenated with the mask or raw path from include or exclude config directive and concatenated with additional suffix pattern `**/*`;
   - two paths are compared with the use of glob patterns: if file matches the mask, then it will be included (for `includePaths`) or excluded (for `excludePaths`), the algorithm is ended.

> The second step with adding `**/*` template is for convenience: the most frequent use case of a _git mapping_ with filters is to configure recursive copying for the directory. Adding `**/*` makes enough to specify the directory name only, and its entire content matches the filter.

Masks may contain the following patterns:

- `*` — matches any file. This pattern includes `.` and exclude `/`
- `**` — matches directories recursively or files expansively
- `?` — matches any one character. Equivalent to /.{1}/ in regexp
- `[set]` — matches any one character in the set. Behaves exactly like character sets in regexp, including set negation ([^a-z])
- `\` — escapes the next metacharacter

Mask that starts with `*` or `**` patterns should be escaped with quotes in `werf.yaml` file:
 - `"*.rb"` — with double quotes
- `'**/*'` — with single quotes

Examples of filters:

```yaml
add: /src
to: /app
includePaths:
# match all php files residing directly in /src
- '*.php'

# matches recursively all php files from /src
# (also matches *.php because '.' is included in **)
- '**/*.php'

# matches all files from /src/module1 recursively
# an example of implicit adding of **/*
- module1
```

`includePaths` filter can be used to copy one file without renaming:
```yaml
git:
- add: /src
  to: /app
  includePaths: index.php
```

### Target paths overlapping

If multiple git mappings are added, you should remember those intersecting paths defined in `to` may result in the inability to add files to the image. For example:

```yaml
git:
- add: /src
  to: /app
- add: /assets
  to: /app/assets
```

When processing a config, werf calculates the possible intersections among all git mappings concerning `includePaths` and `excludePaths` filters. If an intersection is detected, then werf can resolve simple conflicts with implicitly adding `excludePaths` into the git mapping. In other cases, the build ends with an error. However, implicit `excludePaths` filter can have undesirable effects, so try to avoid conflicts of intersecting paths between configured git mappings.

Implicit `excludePaths` example:

```yaml
git:
- add: /src
  to: /app
  excludePaths:  # werf add this filter to resolve a conflict
  - assets       # between paths /src/assets and /assets
- add: /assets
  to: /app/assets
```

## Working with remote repositories

Werf may use remote repositories as file sources. For this purpose, the _git mapping_ configuration contains an `url` parameter where you should specify the repository address. Werf supports `https` and `git+ssh` protocols.

### https

The syntax for https protocol is:

{% raw %}
```yaml
git:
- url: https://[USERNAME[:PASSWORD]@]repo_host/repo_path[.git/]
```
{% endraw %}

`https` access may require login and password.

For example, login and password from GitLab CI variables:

{% raw %}
```yaml
git:
- url: https://{{ env "CI_REGISTRY_USER" }}:{{ env "CI_JOB_TOKEN" }}@registry.gitlab.company.name/common/helper-utils.git
```
{% endraw %}

In this example, the [env](http://masterminds.github.io/sprig/os.html) method from the sprig library is used to access the environment variables.

### git, ssh

Werf supports access to the repository via the git protocol. Access via this protocol is typically protected using ssh tools: this feature is used by GitHub, Bitbucket, GitLab, Gogs, Gitolite, etc. Most often the repository address looks as follows:

```yaml
git:
- url: git@gitlab.company.name:project_group/project.git
```

To successfully work with remote repositories via ssh, you should understand how werf searches for access keys.


#### Working with ssh keys

Keys for ssh connects are provided by ssh-agent. The ssh-agent is a daemon that operates via file socket, the path to which is stored in the environment variable `SSH_AUTH_SOCK`. Werf mounts this file socket to all _assembly containers_ and sets the environment variable `SSH_AUTH_SOCK`, i.e., connection to remote git repositories is established with the use of keys that are registered in the running ssh-agent.

The ssh-agent is determined as follows:

- If werf is started with `--ssh-key` flags (there may be multiple flags):
  - A temporary ssh-agent runs with defined keys, and it is used for all git operations with remote repositories.
  - The already running ssh-agent is ignored in this case.
- No `--ssh-key` flags specified and ssh-agent is running:
  - `SSH_AUTH_SOCK` environment variable is used, and the keys added to this agent is used for git operations.
- No `--ssh-key` flags specified and ssh-agent is not running:
  - If `~/.ssh/id_rsa` file exists, then werf will run the temporary ssh-agent with the  key from `~/.ssh/id_rsa` file.
- If none of the previous options is applicable, then the ssh-agent is not started, and no keys for git operation are available. Build images with remote _git mappings_ ends with an error.

## More details: gitArchive, gitCache, gitLatestPatch

Let us review adding files to the resulting image in more detail. As stated earlier, the docker image contains multiple layers. To understand what layers werf create, let's consider the building actions based on three sample commits: `1`, `2` and `3`:

- Build of commit No. 1. All files are added to a single layer based on the configuration of the _git mappings_. This is done with the help of the git archive. This is the layer of the _gitArchive_ stage.
- Build of commit No. 2. Another layer is added where the files are changed by applying a patch. This is the layer of the _gitLatestPatch_ stage.
- Build of commit No. 3. Files have already added, so werf apply patches in the _gitLatestPatch_ stage layer.

Build sequence for these commits may be represented as follows:

| | gitArchive | --- | gitLatestPatch |
|---|:---:|:---:|:---:|
| Commit No. 1 is made, build at 10:00 |  files as in commit No. 1 | --- | - |
| Commit No. 2 is made, build at 10:05 |  files as in commit No. 1 | --- | files as in commit No. 2 |
| Commit No. 3 is made, build at 10:15 |  files as in commit No. 1 | --- | files as in commit No. 3 |

A space between the layers in this table is not accidental. After a while, the number of commits grows, and the patch between commit No. 1 and the current commit may become quite large, which will further increase the size of the last layer and the total _stages_ size. To prevent the growth of the last layer werf provides another intermediary stage — _gitCache_.
How does werf work with these three stages? Now we are going to need more commits to illustrate this, let it be `1`, `2`, `3`, `4`, `5`, `6` and `7`.

- Build of commit No. 1. As before, files are added to a single layer based on the configuration of the _git mappings_. This is done with the help of the git archive. This is the layer of the _gitArchive_ stage.
- Build of commit No. 2. The size of the patch between `1` and `2` does not exceed 1 MiB, so only the layer of the _gitLatestPatch_ stage is modified by applying the patch between `1` and `2`.
- Build of commit No. 3. The size of the patch between `1` and `3` does not exceed 1 MiB, so only the layer of the _gitLatestPatch_ stage is modified by applying the patch between `1` and `3`.
- Build of commit No. 4. The size of the patch between `1` and `4` exceeds 1 MiB. Now _gitCache_ stage layer is added by applying the patch between `1` and `4`.
- Build of commit No. 5. The size of the patch between `4` and `5` does not exceed 1 MiB, so only the layer of the _gitLatestPatch_ stage is modified by applying the patch between `4` and `5`.

This means that as commits are added starting from the moment the first build is done, big patches are gradually accumulated into the layer for the _gitCache_ stage, and only patches with moderate size are applied in the layer for the last _gitLatestPatch_ stage. This algorithm reduces the size of _stages_.

| | gitArchive | gitCache | gitLatestPatch |
|---|:---:|:---:|:---:|
| Commit No. 1 is made, build at 12:00 |  1 |  - | - |
| Commit No. 2 is made, build at 12:19 |  1 |  - | 2 |
| Commit No. 3 is made, build at 12:25 |  1 |  - | 3 |
| Commit No. 4 is made, build at 12:45 |  1 | *4 | - |
| Commit No. 5 is made, build at 12:57 |  1 |  4 | 5 |

\* — the size of the patch for commit `4` exceeded 1 MiB, so this patch is applied in the layer for the _gitCache_ stage.

### Rebuild of gitArchive stage

For various reasons, you may want to reset the _gitArchive_ stage, for example, to decrease the size of _stages_ and the image.

To illustrate the unnecessary growth of image size assume the rare case of 2GiB file in git repository. First build tranfers this file in the layer of the _gitArchive_ stage. Then some optimization occured and file is recompiled and it's size is decreased to 1.6GiB. The build of this new commit applies patch in the layer of the _gitCache_ stage. The image size become 3.6GiB of which 2GiB is a cached old version of the big file. Rebuilding from _gitArchive_ stage can reduce image size to 1.6GiB. This situation is quite rare but gives a good explanation of correlation between the layers of the _git stages_.

You can reset the _gitArchive_ stage specifying the **[werf reset]** or **[reset werf]** string in the commit message. Let us assume that, in the previous example commit `6` contains **[werf reset]** in its message, and then the builds would look as follows:

| | gitArchive | gitCache | gitLatestPatch |
|---|:---:|:---:|:---:|
| Commit No. 1 is made, build at 12:00 |  1 |  - | - |
| Commit No. 2 is made, build at 12:19 |  1 |  - | 2 |
| Commit No. 3 is made, build at 12:25 |  1 |  - | 3 |
| Commit No. 4 is made, build at 12:45 |  1 |  4 | - |
| Commit No. 5 is made, build at 12:57 |  1 |  4 | 5 |
| Commit No. 6 is made, build at 13:22 |  *6 |  - | - |

\* — commit `6` contains the **[werf reset]** string in its message, so the _gitArchive_ stage is rebuilt.

### _git stages_ and rebasing

Each _git stage_ stores service labels with commits SHA from which this _stage_ was built.
These commits are used for creating patches on the next _git stage_ (in a nutshell, `git diff COMMIT_FROM_PREVIOUS_GIT_STAGE LATEST_COMMIT` for each described _git mapping_).
So, if any saved commit is not in a git repository (e.g., after rebasing) then werf rebuilds that stage with latest commits at the next build.
