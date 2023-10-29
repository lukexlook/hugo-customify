> [!WARNING]
> This repository was created to allow @lukexlook to introduce custom functions (as simple as `Plainify`)
> to Hugo shortcodes. The initial focus is on a straightforward approach for beginners
> rather than idiomatic practices.

## Sourcegraph

Based on [the search results regarding Hugo@v0.119.0](https://sourcegraph.com/search?q=context%3Aglobal+repo%3A%5Egithub%5C.com%2Fgohugoio%2Fhugo%24%40v0.119.0+Plainify+&patternType=standard&sm=1&groupBy=path):

- **Plainify** is first defined as a method on the **Namespace** struct in `tpl/transform/transform.go`.
- It is then registered to the **TemplateFuncsNamespace** in `tpl/transform/init.go`.

Based on the above information, it is believed that registering custom functions to
**TemplateFuncsNamespace** in `tpl/transform/init.go` is enough for simple cases.

## Install the Modified Hugo

```bash
cd hugo
CGO_ENABLED=1 go install -tags extended .
```
