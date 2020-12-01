# Java 8 CSML Studio Function boilerplate

## gradle build
add a zip task in the gradle.build

```
task buildZip(type: Zip) {
    from compileJava
    from processResources
    into('lib') {
        from configurations.runtimeClasspath
    }
}
```

```
$ gradle build -i
```
the .zip file will be found in ./build/distributions

#### IMPORTANT NOTE
In CSML Studio, under 'Add a custom function', your function's handler is the method that is called when your function is invoked. In this case, the handler is 'example.Handler::handleRequest'.