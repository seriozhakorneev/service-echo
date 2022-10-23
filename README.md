# service-echo
JSON Echo Service with clean architecture based on
[Go-clean-template](https://github.com/evrone/go-clean-template).

Service returns the JSON-document that was passed to it.
When rewrite rules are active the document will be overwritten considering them.

- Overwriting recursion for JSON-documents & unit-tests
- Echo REST-HTTP & integration-tests
- Swagger Docs

### Rewrite Rules
Rewrite rules are stored in `config/config.yaml`

```yaml
rewriter:
  active: true
  rules:
    name: "pokemon"
    value: "pikachu"
    new: "bulbasaur"
```
To activate rewriter - set `active` field to `true`, deactivate -`false`.

 - `name` and `value` represents field name and field value of document which should be overwritten 
 - `new` represents new field value for that pair

Example above will be rewritten like this:

`"pokemon":"pikachu"` `->` `"pokemon":"bulbasaur"`

Pair of field name and field value will be found anywhere deep in document:
```json
{
  "pokemon": "pikachu",

  "some_object": {
    "another_object": {
      "pokemon":"pikachu"
    }
  },
  
  "some_array": [
    {
      "another_array": [
        {
          "pokemon": "pikachu"
        }
      ]
    }
  ]
}
```
result:
```json
{
  "pokemon": "bulbasaur",

  "some_object": {
    "another_object": {
      "pokemon": "bulbasaur"
    }
  },
  
  "some_array": [
    {
      "another_array": [
        {
          "pokemon": "bulbasaur"
        }
      ]
    }
  ]
}
```

### Local development:
```sh
# Run app
$ make run

# Or build docker container
$ make compose-up
```
