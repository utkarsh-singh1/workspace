FROM ubuntu
Run apt-get update && apt-get -y upgrade
Run apt-get -y install python3
Run mkdir /home/python
Copy . /home/python
ENTRYPOINT python3 /home/python/simple.py
