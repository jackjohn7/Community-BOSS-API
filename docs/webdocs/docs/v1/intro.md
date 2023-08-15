---
sidebar_position: 1
---

# Getting Started

Our API follows standard, RESTful practices. As such, to consume our data, you must utilize HTTP requests.

## Note on Versioning

In order to maintain compatibility and not break existing projects, our API is versioned. At the moment, our API is in beta. It is not recommended to consume the beta version of the API over V1.

## Available Data

There are four sections of the API that you can consume data from:

- Quarters
- Subjects
- Courses
- Sections

### Swagger

While these docs are incomplete, feel free to check out the swagger auto-generated documentation [here](https://community-boss-api-development.up.railway.app/swagger/index.html).

## Examples

Depending on what language you're using, you'll need to figure out what you need to use to make HTTP Requests. In these examples, I'll be consuming some data from our quarters API. It will return the currently stored quarters.

### Python

*The unofficial official language of Louisiana Tech University*

In python, you will need the [requests](https://pypi.org/project/requests/) module.

You can install this module with the following command:

```bash
pip install requests
```

```python
import requests
r = requests.get("https://community-boss-api-development.up.railway.app/beta/quarters")
```

There is also a full python file example in the [GitHub repo](https://github.com/jackjohn7/Community-BOSS-API/) in the **docs/examples/** directory.

### JavaScript

How you make this request will depend on your project and where it's meant to run. For this example, we'll be using the native [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) API, so it should work anywhere that supports the fetch API.

```js
fetch("https://community-boss-api-development.up.railway.app/beta/quarters").
  then(res => res.json()).
  then(parsedResponse => {
    console.log(parsedResponse);
  });
```
