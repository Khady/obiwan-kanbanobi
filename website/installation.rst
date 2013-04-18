.. highlight:: rest

.. _installation:

=======================
Installation du serveur
=======================

Nous décrirons dans cette partie la procédure à respecter pour installer le serveur.

Celle-ci se déroule en plusieurs point :

* Il faut tout d'abord ajouter un utilisateur kanban dans son fichier pg_hba.conf.

  Celui-ci se trouve généralement dans /etc/postgresql/<sub-dir eventuel>/pg_hba.conf

  ligne à rajouter dans le fichier :

  .. code-block:: none

     local kanban kanban password

* Recharger la base postgressql :

  .. code-block:: none

     /etc/init.d/postgresql restart
* exécuter le script initdb.sh avec l'utilisateur postgres :

  .. code-block:: none

     sudo -i -u postgres
     cd /path/to/project
     ./initdb.sh
     exit

======================
Installation du client
======================

* Vous devez posséder les librairies de développement python ainsi que libevent

  .. code-block:: none

     apt-get install libevent-dev
     apt-get install python-dev
* Vous devez également installer les librairies suivantes à l'aide de pip :

  .. code-block:: none

     pip install Flask
     pip install Flask-SQLAlchemy
     pip install Flask-WTF
     pip install redis
     pip install gevent
     pip install gunicorn
     pip install protobuf
* Vous devez installer redis-server

  Sur ubuntu ou debian :

  .. code-block:: none

     apt-get install redis-server

* Pour lancer le client web, vous devez vous placer à la racine de l'application ouane puis lancer la commande suivante

  .. code-block:: none

     gunicorn --debug --worker-class=gevent main:app

.. toctree::
   :maxdepth: 1
