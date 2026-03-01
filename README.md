# STURE (simplte template unification runtime engine)

Import the engines you target, load your templates the way you prefer. Then render them with minimum amount of metadata and effort.

This lib values simplicity above all. It will generate garbage like it's nobodies bussiness. But it will be easy to use.

## Supported engines (in the lib)

Currently we support:
* Go templates through Go:s default `html/templates` package.
* Markdown through [yuin/goldmark](https://github.com/yuin/goldmark)
* Mustache through [cbroglie/mustache](https://github.com/cbroglie/mustache)

If you dont fancy changing the default, then simply `_` import them and they should register themselfes and be ready for consumption.

### Go templates (engines/gotmpl)

`gotempl`-package exposes 2 methods (`SetFuncs(them map[string]any)` & `SetFunc(name string, it any)`) that can be used to set the funcs available during template execution.

### Markdown (engines/markdown)

`markdown`-package exposes a method that allows you to set your own preconfigured instance of `goldmark.Markdown`. The default one comes with only `extention.GFM` enabled.

### Mustasche (engines/mustache)

`mustache`-package exposes a method where you can register a `mustsche.PartialProvider` that allows you to load template partials on the run if needed.

## Extending

If you want to use an engine not supplied by the package or dont like how the current ones are implemented. Implement the `model.Engine` interface and register your instance with the `model.Kind` you target. Ie `sture.Register(JINJA, &MyJinja{})`.

## TODO
* Unset a func by name in gotmpl package?
  > That's a VERY limited use case...?