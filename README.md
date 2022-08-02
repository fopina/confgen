# confgen
Generate configuration files from Go templates

As the (internet) world is all going serverless and [12-factor](https://12factor.net/) for scalability and resilience and such, `envsubst` became quite popular to "12-factor" some apps that only support configuration through files.

Issue with `envsubst` is that it only does... env substitutions, nothing else. No logic control to allow for more complex templating (to eventually simplify/minify the required env vars).

An easy (and probably popular) option is to use [Jinja2](https://palletsprojects.com/p/jinja/) templates.  
(python) Developers already know it from Django/Flask, devops know it from ansible.

But for a python-less image (such as nginx), adding ~50M of python runtime just to have templating sounds a bit too much.

`confgen` tries to help with that.
Instead of Jinja2 templating, it's [Go templating](https://golang.org/pkg/text/template/).
Instead of 50M footprint, it's 2M.

`confgen` simply injects the `env` function into the templates mapped to Go `os.Getenv`.

If you need more functions in your templates and you don't mind having twice the footprint (a bit over 4M instad of 2M), you have the `confgen_sprig` builds which uses [sprig](http://masterminds.github.io/sprig/) funcmap instead, exposing not only the same `env` function but a lot more.  
Check its [documentation](http://masterminds.github.io/sprig/) for full list.

## Usage

* Add confgen in your Dockerfile
  ```
  ...
  # multi-platform image so it is simpler than having to fetch the right binary from releases
  # immutable tags also available, eg fopina/confgen:0.1.6
  COPY --from=fopina/confgen:latest /confgen /usr/local/bin/confgen
  # or with sprig, if you need more template functions than "env"
  # immutable tags also available, eg fopina/confgen:0.1.6-sprig
  # COPY --from=fopina/confgen:sprig /confgen /usr/local/bin/confgen
  RUN chmod a+x /usr/local/bin/confgen
  ...
  ```

* Call it from your entrypoint
  ```
  ...
  confgen -o /path/to/used/conf /path/to/template
  ...
  ```

* With your [template](https://golang.org/pkg/text/template/) using the `env` function
  ```
  config_key={{ "MY_CONFIG" | env }};
  ```

* Set those env vars when creating the containers
  ```
  docker run -e MY_CONFIG=wtv ...
  ```

## Docs

Check [examples](examples) for actual usage with nginx paired with `envsubst`.  
And check [Go template](https://golang.org/pkg/text/template/) documentation for full list of logical controls.

Check [sprig documentation](http://masterminds.github.io/sprig/) for its templating functions.

Find below the ones implemented in this code (without sprig, the lite binary):

### env

The `env` function reads an environment variable:

```
env "HOME"
```
