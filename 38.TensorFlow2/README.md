# TensorFlow

```CMD

docker pull tzutalin/py2qt4

docker run -it \
	--user $(id -u) \
	-e DISPLAY=unix$DISPLAY \
	--workdir=$(pwd) \
	--volume="/home/$USER:/home/$USER" \
	--volume="/etc/group:/etc/group:ro" \
	--volume="/etc/passwd:/etc/passwd:ro" \
	--volume="/etc/shadow:/etc/shadow:ro" \
	--volume="/etc/sudoers.d:/etc/sudoers.d:ro" \
	-v /tmp/.X11-unix:/tmp/.X11-unix \
	tzutalin/py2qt4

make all

Erro
sudo apt-get update
sudo apt-get install pyqt5-dev-tools

You don't have enough free space in /var/cache/apt/archives/
sudo apt-get autoclean
sudo apt-get autoremove 

docker image ls
docker rmi <imagem>


make all
erro...

make qt4py2;./labelImg.py

python labelImg.py



make qt4py2;./labelImg.py






Funcionou no Ubuntu Desktop 20.04 
sudo apt-get update

sudo apt install python3-pip
pip3 install labelImg

sudo apt-get install pyqt5-dev-tools
sudo pip3 install -r requirements/requirements-linux-python3.txt

make qt5py3
python3 labelImg.py

```
