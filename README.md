## Taurl

Fast url shortner written entirely in Go using the gin framework.

### Provided routes:

#### Retrives an existing shortend url.

<details>
 <summary><code>GET</code> <code><b>/api/get</b></code> <code>(first reads the cache, then the db.)</code></summary>

##### Parameters

> | name | type     | data type | description  |
> | ---- | -------- | --------- | ------------ |
> | Url  | required | String    | shortend url |

##### Responses

> | http code | content-type       | response                           |
> | --------- | ------------------ | ---------------------------------- |
> | `200`     | `application/json` | `Url body`                         |
> | `400`     | `application/json` | `{"message":"Reason"}`             |
> | `404`     | `application/json` | `{"message":"Url was not found."}` |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:3000/api/get?Url=http://localhost:3000/hash
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/api/get_from_id</b></code> <code>(first reads the cache, then the db, but uses unique id of a shortend url.)</code></summary>

##### Parameters

> | name      | type     | data type | description       |
> | --------- | -------- | --------- | ----------------- |
> | unique_id | required | String    | shortend url's id |

##### Responses

> | http code | content-type       | response                           |
> | --------- | ------------------ | ---------------------------------- |
> | `200`     | `application/json` | `Url body`                         |
> | `400`     | `application/json` | `{"message":"Reason"}`             |
> | `404`     | `application/json` | `{"message":"Url was not found."}` |

##### Example cURL

> ```javascript
>  curl -X GET -H "Content-Type: application/json" http://localhost:3000/api/get_from_id?Url=http://localhost:3000/hash
> ```

</details>

#### Sets a new shortend url and returns it.

<details>
 <summary><code>POST</code> <code><b>/api/set</b></code> <code>(Creates and stores a new shortend url.)</code></summary>

##### Parameters

> | name | type     | data type | description  |
> | ---- | -------- | --------- | ------------ |
> | Url  | required | String    | Original url |

##### Responses

> | http code | content-type       | response                          |
> | --------- | ------------------ | --------------------------------- |
> | `201`     | `application/json` | `Url body`                        |
> | `400`     | `application/json` | `{"message":"Reason"}`            |
> | `500`     | `application/json` | `{"message":"Failed to set db."}` |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data '{"Url":"https://example.com"}'  http://localhost:3000/api/set
> ```

</details>

### Building

The program expects these environment variables :

```sh
# defaults to 3000
PORT=3000

# defaults to http://localhost:3000
# this is what you would set as your base domain.
BASE_URL="http://localhost:3000"

# optionally you can provide GIN_MODE to get or remove log messages
# from the gin framework.
# defaults to debug
GIN_MODE="debug" # release, debug
```

clone and run using go cli.

```sh
git clone --depth 1 https://github.com/prashantrahul141/taurl && cd taurl && go run .
```

![One](https://raw.githubusercontent.com/prashantrahul141/taurl/main/assets/meta/01.png)
![Two](https://raw.githubusercontent.com/prashantrahul141/taurl/main/assets/meta/02.png)
