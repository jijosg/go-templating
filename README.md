# go-templating

This project is to demonstrate the use of go templating engine using files.

The `example.gotmpl` is the template file and `values.yaml` is the config values passed to process the template

## Usage
```bash
Usage:
  template [flags]

Flags:
      --config string   config file (default is $HOME/.template.yaml)
  -h, --help            help for template
  -t, --toggle          Help message for toggle
  -f, --values string   Values file Path (default "../resources/values.yaml")
  ```

Custom values file can be passed using the -f flag.

## Example
```bash
> go run main.go
template called
docker1: docker.io
image: nginx1.0.0
```
                                    