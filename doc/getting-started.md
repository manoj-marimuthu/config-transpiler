## Getting Started

If you have followed the instructions in the project's ```README.md``` file,
You will have successfully built the binary. Before using the tool,
know that it supports only a subset.

In Cloud-config, It supports a range of keys (see: [features](https://www.github.com/manoj-marimuthu/config-transpiler/blob/main/doc/supported-features.md)). For example-

```yaml
#cloud-config

users:
    - name: foobar
      gecos: Foo B. Bar
    - name: barfoo
      gecos: Bar B. Foo
```

So, The generated butane configuration file will have the ```passwd``` section that has the
```users``` section. Each entry in the ```users``` section may contain the ```name``` and ```gecos```
fields. So the converted format depends on the input's features. Hence we get:

```yaml
passwd:
    users: 
        - name: foobar
          gecos: Foo B. Bar
        - name: barfoo
          gecos: Bar B. Foo
```

The point is, You can perform conversion but the scope of conversion depends on the features supported.
### How to use the CLI ?

To convert a cloud-config YAML file into a Butane YAML file, run:

```bash
./config-transpiler <input_file_name.yaml> <output_file_name.yaml>
```

If the input file cannot be read/found, the program exits with an error.
