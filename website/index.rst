.. Obi-wan Kanbanobi documentation master file, created by
   sphinx-quickstart on Wed Jan  2 05:59:36 2013.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Welcome to Obi-wan Kanbanobi's documentation!
=============================================

Obi-wan Kanbanobi est un projet de serveur et client libre pour faire du Kanban_ réalisé dans le cadre du projet libre de troisième année à EPITECH_ par :

* Louis Roché
* Alexandre Baron
* Arnaud Baussart
* Maxime Constantinian
* Timothée Maurin
* Pierre-Jean Prost

Le serveur est réalisé en Go_ et le client en python. Notre serveur et notre protocole sont concus de manière et ce que la création d'un client soit la plus simple possible. Il s'agit en fait d'un simple visualisateur et toute « l'intelligence » est dans le serveur.

Le stockage des données côté serveur se fait sur une base de données Postgresql_.

La communication entre le client et le serveur se fait avec des Protobuf_ en tcp. le :ref:`protocole` est documenté sur ce site.

Les sources sont disponibles sur Bitbucket_ et un miroir sur Github_. Le depot contient les sources du serveur, du client et du site.

Des tests sont effectues a chaque commit et sont disponibles sur Travis_.

Contents:

.. toctree::
   :maxdepth: 1

   protocole.rst

.. _Kanban: http://fr.wikipedia.org/wiki/Kanban
.. _Go: http://golang.org
.. _Postgresql: http://www.postgresql.org/
.. _Protobuf: http://code.google.com/p/protobuf/
.. _EPITECH: http://epitech.eu
.. _Bitbucket: https://bitbucket.org/ongisnotaguild/obi-wan-kanbanobi
.. _Github: https://github.com/Khady/obiwan-kanbanobi
.. _Travis: https://travis-ci.org/Khady/obiwan-kanbanobi
