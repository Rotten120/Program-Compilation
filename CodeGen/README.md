# CODEGEN

Generates nesting if statements for multiplying positive integers by two

To run, build with the following command
```
go build -o codegen.exe main.go
```

To execute, input the highest number it can multiply
```
./codegen.exe 10
```

You can also modify the filepath
```
./codegen.exe --dest="newFile.go" 5
```

If no number is added, default value is applied (value of 5)
```
./codegen.exe
```
