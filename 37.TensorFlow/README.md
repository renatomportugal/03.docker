# TensorFlow

```CMD
Para rodar o tensorflow é necessário que o processador tenha instruções AVX.
Advanced Vector Extensions (AVX) are extensions to the x86 instruction set architecture for microprocessors from Intel and AMD proposed by Intel in March 2008 and first supported by Intel with the Sandy Bridge processor shipping in Q1 2011 and later on by AMD with the Bulldozer processor shipping in Q3 2011. AVX provides new features, new instructions and a new coding scheme.

https://www.server-world.info/en/note?os=Ubuntu_20.04&p=tensorflow&f=3

01. Rodar
docker run -it -p 8888:8888 tensorflow/tensorflow:2.0.0-py3-jupyter

02. Caso precise pegar o token
docker exec 5419643d0af7 bash -c "jupyter notebook list"

03. Acessar
http://192.168.1.106:8888
http://localhost:8888

04. Esta é a tela do Jupyter
New, Notebook Python 3

import tensorflow as tf

Rodar....


Se não tiver AVX...
The kernel appears to have died. It will restart automatically.
```
