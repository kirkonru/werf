---
title: Slug
sidebar: documentation
permalink: documentation/reference/toolbox/slug.html
author: Timofey Kirillov <timofey.kirillov@flant.com>
---

In some cases, environment variables or parameters can't be used "as is" since they contain invalid characters.

To meet the requirements for naming docker images, helm releases, and Kubernetes namespaces, werf applies a unified slug algorithm when producing these names. This algorithm excludes unacceptable symbols and ensures the uniqueness of the result for each unique input.

There are 3 types of slugs built into werf:

1. Helm Release name slug.
2. Kubernetes Namespace slug.
3. Docker tag slug.

There are commands for the every type of slug available. They apply algorithms to the provided input. You can use these commands depending on your needs.

## Basic algorithm

werf checks if the input meets slug **requirements**. If the input complies with them, werf leaves it unaltered. Otherwise, werf **transforms** characters to comply with the requirements while adding a dash symbol followed by a source-based hash suffix in the end. The [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash) algorithm is used.

While transforming the input in the slug, werf performs the following actions:
* Converts UTF-8 Latin characters to their ASCII counterparts;
* Replaces some special symbols with dash symbol (`~><+=:;.,[]{}()_&`);
* Removes all non-recognizable characters (only lowercase alphanumerical characters and dashes remain);
* Removes starting and ending dashes;
* Reduces multiple dashes to a single dash.
* Trims the length of the data so that result stays within the maximum bytes limit.

The actions are the same for all slugs since they are restrictive enough to satisfy the requirements of any slug.

### The requirements for naming Helm Releases
* only alphanumerical ASCII characters, underscores, and dashes are allowed;
* the length is limited to 53 bytes.

### Kubernetes Namespace requirements (a [DNS Label](https://www.ietf.org/rfc/rfc1035.txt) requirements)
* only alphanumerical ASCII characters, and dashes are allowed;
* the length is limited to 63 bytes.

### The requirements for Docker Tags
* the tag must be made of valid ASCII and can contain lowercase and uppercase ASCII characters, digits, underscores, periods, and dashes;
* the length is limited to 128 bytes.

## Usage

The slug may be applied to an arbitrary string with the [`werf slugify` command]({{ site.baseurl }}/documentation/cli/toolbox/slugify.html).

werf applies slug automatically in the CI/CD systems such as GitLab CI. See [plugging into the CI/CD]({{ site.baseurl }}/documentation/reference/plugging_into_cicd/overview.html) for more details. The basic principles are:
 * the slug is automatically applied to parameters that are automatically derived from the environment of the CI/CD system;
 * the slug isn't automatically applied to params that are manually specified via `--tag-*`, `--release`, or `--namespace`; in this case, parameters are only validated to comply with the requirements.

The user should run the [`werf slugify` command]({{ site.baseurl }}/documentation/cli/toolbox/slugify.html) explicitly to apply slug to parameters specified manually (via `--tag-*`, `--release`, or `--namespace`), for example:

```
werf publish --tag-git-branch $(werf slugify --format docker-tag "Features/MyBranch#123") ...
werf deploy --release $(werf slugify --format helm-release "MyProject/1") --namespace $(werf slugify --format kubernetes-namespace "MyProject/1") ...
```
