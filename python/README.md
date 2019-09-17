# Python lambda function boilerplate


## Install dependencies
The compressed function file will need all its dependencies so all install them in a folder called package.
```
$ cd src
$ pip install --target ./package -r requirements.txt
```
#### IMPORTANT NOTE
Make sure you do not import your dependencies from the package folder in your python script (~~`import package.mypippackage`~~), once packaged this python file will be at the same folder level as your dependencies, use `import mypippackage` instead

## Compress function
We need to compress the function and all its dependencies at the same root level. We first `cd` to the package folder, compress the dependencies, then `cd` back to `src` and add the python file(s) to the zip file.
```
$ cd src/package
$ zip -r9 ../../mypythonscript.zip .
$ cd ..
$ zip -g ../mypythonscript.zip index.py
```