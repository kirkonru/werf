---
title: Releases
permalink: releases.html
sidebar: documentation
layout: default
---

<link rel="stylesheet" href="{{ site.baseurl }}/css/releases.css">

{% assign releases = site.data.releases.releases %}

<div class="page__container page_releases">

<div class="releases__block-title">
    Release channels
</div>

<!-- Releases description -->
<div class="releases__info">
    Each werf release progresses through all release channels, starting with Alpha → Beta → Early-Access → Stable → Rock-Solid. You can think of each release on a lower channel as a release-candidate for the higher one. Once a release is considered bug-free, it is promoted to the next channel.
</div>

<!-- latest version per channel -->
{% assign channels_sorted = site.data.channels_info.channels | sort: "stability" %}
{% assign channels_sorted_reverse = site.data.channels_info.channels | sort: "stability" | reverse  %}
{% assign latest_group_versions = site.data.releases_latest.latest | where: "group", "1.0" | first | map: "channels" | first %}

<div class="releases__menu">
{% for channel in channels_sorted_reverse %}
{% assign channel_latest_version = latest_group_versions | where: "name",  channel.name | map: "version" | first %}
{% assign channel_latest_version_info = site.data.releases.releases | where: "tag_name", channel_latest_version | first %}
    <div class="releases__menu-item">
        <div class="releases__menu-item-header">            
            <div class="releases__menu-item-title">
                {{ channel.title }}
            </div>
            {% if version != empty  %}
                <a href="{{ channel_latest_version_info.html_url }}" class="releases__btn">
                {{ channel_latest_version }}
                </a>
            {% endif %}
        </div>        
        <div class="releases__menu-item-description">
            {{ channel.description[page.lang] }}
        </div>
    </div>
  {% endfor %}
</div>

<div class="releases__block-title">
    Releases
</div>

<div class="releases">
    <div class="tabs">
    {% for channel in channels_sorted_reverse %}
    <a href="javascript:void(0)" class="tabs__btn{% if channel == channels_sorted_reverse[1] %} active{% endif %}" onclick="openTab(event, 'tabs__btn', 'tabs__content', '{{channel.name}}')">{{channel.title}}</a>
    {% endfor %}
    <a href="javascript:void(0)" class="tabs__btn" onclick="openTab(event, 'tabs__btn', 'tabs__content', 'all')">All channels</a>
    </div>

    {% for channel in channels_sorted_reverse %}
    <div id="{{ channel.name }}" class="tabs__content{% if channel == channels_sorted_reverse[1] %} active{% endif %}">
    <div class="releases__info">
        <p>{{ channel.tooltip[page.lang] }}</p>
        <p class="releases__info-text">{{ channel.description[page.lang] }}</p>
    </div>

    {% assign channel_history = site.data.releases_history.history | reverse | where: "group", "1.0"  | where: "name", channel.name %}

    {% if channel_history.size > 0 %}
        {% for channel_action in channel_history %}
           {% assign release = site.data.releases.releases | where: "tag_name", channel_action.version | first %}
            <div class="releases__title">
                <a href="{{ release.html_url }}">
                    {{ release.tag_name }}
                </a>
            </div>
            <div class="releases__body">
                {{ release.body | markdownify }}
            </div>
        {% endfor %}
    {% else %}
        <div class="releases__info releases__info_notification">
            <p>There are no versions on the channel yet, but they will appear soon.</p>
        </div>
    {% endif %}

    </div>
    {% endfor %}

    <div id="all" class="tabs__content">
        <div class="releases__info">
            <p>This is a list of all of the releases (Alpha, Beta, Early-Access, Stable and Rock-Solid) combined in chronological order.</p>
        </div>
    {% for release in site.data.releases.releases %}
            <div class="releases__title">
                <a href="{{ release.html_url }}">
                    {{ release.tag_name }}
                </a>
            </div>
            <div class="releases__body">
                {{ release.body | markdownify }}
            </div>
    {% endfor %}
    </div>
</div>
