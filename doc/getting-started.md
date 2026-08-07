## Getting Started

If you have followed the instructions in the project's ```README.md``` file,
You will have successfully built the binary. Before using the tool,
know the current features provided:

In Cloud-config, It supports the conversion of ```users``` section. Each user entry may contain the
```name``` and ```gecos``` fields. Example config file,

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
fields. For example, the above cloud-config file will be converted to a butane YAML file that has:

```yaml
passwd:
    users: 
        - name: foobar
          gecos: Foo B. Bar
        - name: barfoo
          gecos: Bar B. Foo
```

### How to use the CLI ?

To convert a cloud-config YAML file into a Butane YAML file, run:

```bash
./config-transpiler <input_file_name.yaml> <output_file_name.yaml>
```

If the input file cannot be read/found, the program exits with an error.
