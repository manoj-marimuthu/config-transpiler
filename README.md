## Config Transpiler

A go transpiler that converts cloud-config YAML format into Butane YAML format.

### Why I built it ?

While exploring the projects for LFX 2026, I noticed CNCF's project "Flatcar container Linux:
Cloud-Init to Butane YAML config transpiler". It was genuinely interesting, Hence To learn the 
architecture, programming practices and design choices, I decided to build a prototype.

### Documentation and Examples

To get started, checkout the [getting started](https://github.com/manoj-marimuthu/config-transpiler/blob/main/doc/getting-started.md)
documentation under the ```doc``` directory. The directory also holds examples for reference in  
[examples](https://github.com/manoj-marimuthu/config-transpiler/blob/main/doc/examples.md).

### Building guide

To use the tool, clone this project using git and build the binary.
```bash
git clone https://www.github.com/manoj-marimuthu/config-transpiler.git
cd config-transpiler
go build -o config-transpiler
```

To convert a cloud-config YAML file into Butane YAML file, run:
```bash
./config-transpiler input.yaml output.yaml
```

Reads ```input.yaml``` file and converts it to Butane YAML format and writes 
it to ```output.yaml```. To learn more, checkout the documentation.

### Inspiration

The project uses programming styles and choices from:

[container-linux-config-transpiler](https://www.github.com/flatcar/container-linux-config-transpiler)

I used the project structure, The idea of registering and using converter functions from here and will
continue to do so. 
