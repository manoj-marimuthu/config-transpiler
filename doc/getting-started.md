## Getting Started

To get an executable run:

```bash
go build -o config-transpiler .
```

You will now see an executable in your current directory. To run the executable:

```bash
./config-transpiler <input_file.yaml> <output_file.yaml>
```

Use filenames of your choice where the angular brackets are specified. If you do 
not want an executable, use:

```bash
go run ./... <input_file.yaml> <output_file.yaml>
```
The project also involves a test suite which can run using:

```bash
go test ./converter
```
 All test files are inside the converter as of now. Beware that the config-transpiler
does not support all the keys and sections as specified in the documentation of cloud-config yet.
