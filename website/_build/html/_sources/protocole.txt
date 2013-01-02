.. highlight:: rest

.. _protocole:

=========
Protocole
=========

Ceci est le protocole utilisé pour faire communiquer notre serveur et notre client.

Ce protocole utilise les Protobuf_ pour communiquer. Il le fait sur des sockets en tcp.

Chaque client ouvre une socket pour communiquer mais peut parler au nom de plusieurs utilisateurs. Ce qui permet par exemple a un client web de se connecter une seule fois au serveur, mais de gerer plusieurs utilisateurs qui se connectent en parallele.

Informations de base
====================

Chaque message est précédé d'un int décrivant la taille du message à lire.

Tous les messages entre le serveur et les clients sont de type Msg et doivent contenir des informations de base :

* Une valeur de contexte provenant de l'enum TARGET qui va décrire sur le type d'objet dont va traiter le message.
* Une commande provenant de l'enum CMD qui est l'action à effectuer sur l'objet du type précédement défini.
* L'id de l'auteur du message.
* L'id de session de l'auteur pour valider l'authentification de l'auteur.


    .. code-block:: protobuf
       :emphasize-lines: 3,5

       enum TARGET {
          USERS		= 1;
          COLUMNS	= 2;
          PROJECTS	= 3;
          CARDS		= 4;
          ADMIN		= 5;
          IDENT		= 6;
          NOTIF		= 7;
          METADATA	= 8;
       }
       
       enum CMD {
          CREATE	= 1;
          MODIFY	= 2;
          DELETE	= 3;
          GET		= 4;
          MOVE		= 5;
          CONNECT	= 6;
          DISCONNECT	= 7;
          ERROR		= 8;
          SUCCES	= 9;
          NONE		= 10;
       };
       
       message Msg {
          required TARGET	target		= 1;	// type de la cible du message .. 
          required CMD		command		= 2;	// commande a effectuer sur la cible
          required uint32	author_id	= 3;	// id de l'auteur, fourni par le serveur après l'auth
          required string	session_id	= 4;	// id de session pour valider l'auth, fourni par le serveur
       }

Description des communications
==============================

Authentification
----------------

1. Message initial

Le premier message attendu de la part d'un utilisateur est le message d'authentification. Tous les messages du client avant authentification sont ignores.

Le message a envoyer doit etre de type Msg et contenir un message de type Ident. Les variables target et command doivent avoir les valeurs IDENT et CONNECT.

L'author_id et le session_id sont inconnus a ce moment la puisque c'est le serveur qui les attribue. Leurs valeurs sont ignorees.

Le message Ident doit contenir le login et le password de l'utilisateur. Le password est envoye en clair sur le reseau, mais il est stocke de maniere chiffre et n'est pas garde en memoire une fois l'authentification realisee.

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = IDENT;
           command = CONNECT;
           author_id = -1
           session_id = "";
           Message Ident {
           login = pseudo;
           password = password;
           }
       }

2. Reponse positive



3. Erreur

Cette erreur est renvoye sur une mauvaise authentification ou quand un message est envoye par une personne non authentifie.

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = ERROR;
           command = CONNECT;
           author_id = -1
           session_id = "";
           Message Error {
               error_id = error_connexion_id; // Cette erreur provient de la l'enum decrit dans cette page
           }
       }

.. - Identitifaction
.. Client -> Premier message d'identification a la connexion
.. Le author_id n'est pas connu au moment de l'identification, il importe donc peu
.. Idem pour le session_id
.. le login et le password sont en clair

.. Serveur -> Deux reponses possibles selon la validite de l'ident
.. En cas d'erreur:


.. En cas de reussite:
.. Msg {
..     target = IDENT;
..     command = CONNECT;
..     author_id = id calcule par le serveur
..     session_id = session_id calcule par le serveur
.. }


Cartes
------

Colonnes
--------

Erreurs
-------




.. Tout message est precede d'un unsigned int pour preciser la taille du message qui va suivre.

.. - Creation de compte
..   target = IDENT;
..   command = CREATE;
..   author_id = id
..   session_id = session_id;

.. ------------------------------------------------------------------------------------------
.. erreurs:
.. - erreur a la connexion
.. - n'a pas les droits
.. - target invalid
.. - cmd invalid
.. - session invalide

.. _Protobuf: http://code.google.com/p/protobuf/
