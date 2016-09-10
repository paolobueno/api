## This is the example web server from lesson two in Golang intro.

[Golang intro](https://github.com/fheng/golang-intro/blob/master/chapter-two.md)

## Homework

- Add a new route to the api
```
/api/time
```
That returns a json structure that looks like :

```json

{"time":1473438486}

``` 
- Add a new test that calls the new api 


## Hints

- Start with the router and look at the echo handler
- you will need a new type that matches has a time property
