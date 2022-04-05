<h1 style="text-align: center">
  <br>
  <a href="http://github.com/go-task/task"><img src="../Logo.png" alt="task" width="200px" /></a>
  <br>
  Task
  <br>
</h1>

<p style="text-align: center">
  <a href="#introduction">Introduction</a> •
  <a href="#getting-started">Getting Started</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#development">Development</a> •
  <a href="#roadmap">Roadmap</a>
</p>

## Introduction

## Getting Started

```yaml
version: '4'

tasks:
  ham:
    labels:
      type: cookable
    antiAffinity:
      matchLabels:
        type: cookable
    vars:
      WHO: "{{ env "WHO" | default "World" }}"
    do:
    - echo "Hello, {{.V.WHO}}"

  eggs:
    labels:
      type: cookable
    antiAffinity:
      matchExpressions:
      - key: type
        operator: In
        values:
        - cookable

  milk:

  espresso:

  coffee:
    needs:
      - milk


  breakfast:
    needs:
      - ham
      - eggs

```

## Overview

Task uses Go templating to provide advanced functionality outside of what is normally possible in straight yaml. A Taskfile however is _not_ a template, and mustache (`{{ }}`) syntax is limited to fields in which a regular string, number, or boolean value is accepted. However, there are special fields which can be evaluated dynamically!

## Features

| Features                     | Redoc     | Docs         |
|------------------------------|:---------:|:------------:|
| **Specs**                    |           |
| Swagger 2.0                  | √         |
| OpenAPI 3.0                  | √         |
| OpenAPI 3.1                  | √ (basic) |
|                              |           |
| **Theming**                  |           |
| Fonts/colors                 | √         |
| Extra theme options          |           |
|                              |           |
| **Performance**              |           |
| Pagination                   |           |
| Search (enhanced)            |           |
| Search (server-side)         |           |
|                              |           |
| **Multiple APIs**            |           |
| Multiple versions            |           |
| Multiple APIs                |           |
| API catalog                  |           |
|                              |           |
| **Additional features**      |           |
| Try-it console               |           |
| Automated code samples       |           |
| Deep links                   |           |
| More SEO control             |           |
| Contextual docs              |           |
| Landing pages                |           |
| React hooks for more control |           |
| Personalization              |           |
| Analytics integrations       |           |
| Feedback                     |           |


### Lazy Evaluation

### Order of Operations
`
