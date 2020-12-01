# Ruby lambda function boilerplate

## Install dependencies
Install libraries in the vendor directory with the bundle command.

```
$ bundle config set path 'vendor/bundle'
$ bundle install
```

## Compress function
We need to compress the function and all its dependencies if any.

```
$ zip -r9 function.zip index.rb vendor
```

#### IMPORTANT NOTE
In the CSML Studio in the section 'Add a custom function' make sure that the handler name is the same as your script name, in this example the handler is 'index.handler'
