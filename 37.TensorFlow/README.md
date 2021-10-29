# TensorFlow

```CMD
https://www.server-world.info/en/note?os=Ubuntu_20.04&p=tensorflow&f=3
docker pull tensorflow/tensorflow:2.0.0-py3

docker run --rm tensorflow/tensorflow:2.0.0-py3 \
python -c "import tensorflow as tf; print(tf.reduce_sum(tf.random.normal([1000, 1000])))"

Ou criar o arquivo e subir via FTP
sudo nano hello_tensorflow.py 

import tensorflow as tf
hello = tf.constant('Hello, TensorFlow World!')
tf.print(hello)



docker run --rm -v $PWD:/tmp -w /tmp tensorflow/tensorflow:2.0.0-py3 python ./hello_tensorflow.py

docker pull tensorflow/tensorflow:2.0.0-py3-jupyter

docker images

docker run -dt -p 8888:8888 tensorflow/tensorflow:2.0.0-py3-jupyter

docker ps


docker exec a4b3b2b68ef2 bash -c "jupyter notebook list"

Acesse 192.168.1.106:8888
Digite o token que apareceu na tela de SSH



```
