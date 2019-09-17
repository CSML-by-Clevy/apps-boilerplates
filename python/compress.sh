#!/bin/bash
cd src/package
zip -r9 ../../mypythonscript.zip .
cd ..
zip -g ../mypythonscript.zip index.py
cd ..
echo "Compressed file ready to upload : ${PWD}/mypythonscript.zip"