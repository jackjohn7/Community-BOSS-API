# Community-BOSS-API

Open Source software created by John (Jack) Branch and Cameron Thomas at Louisiana Tech University. This is a companion project to the [web app](https://github.com/CamJak/Boss-Integrated-Scheduler). The original web app uses an internal API, but this publicly available API serves nearly all of the same data.

## Purpose

The project serves two main purposes:

1. Enable Louisiana Tech Students to consume BOSS data in their projects. This can enable new visualization and data analysis tools for tech students for example.

2. Provide an opportunity for Tech students to learn the fundamentals of API consumption through a simple, RESTful API serving data relevant to them. We plan on creating some basic learning materials to go along side this API (interactive demos on the documentation site).

## Technologies

The project makes use of the following technologies:

- Golang or Rust (still deciding)
  - If Golang:
    - Fiber or Gin
    - Rate limiting middleware (using Redis likely)
  - If Rust:
    - Actix Web Framework
    - [Rate limiting library](https://github.com/jacob-pro/actix-extensible-rate-limit) for actix
- Astro Web Framework for Documentation
- Playwright for E2E testing of the documentation site

## About the Team

We're both incredibly excited to create software to improve the experience at Louisiana Tech University.

### Jack Branch

I'm a computer science student at Louisiana Tech University. I work as a full-stack web developer at the Cyber Innovation Center in Bossier City, LA. Super enthusiastic about new development technologies

### Cameron Thomas

Cyber Engineering student at Louisiana Tech University. _more information_

