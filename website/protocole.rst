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
         CREATE		= 1;
         MODIFY		= 2;
         DELETE		= 3;
         GET		= 4;
         MOVE		= 5;
         CONNECT	= 6;
         DISCONNECT	= 7;
         ERROR		= 8;
         SUCCES		= 9;
         NONE		= 10;
         PASSWORD	= 11;
         GETBOARD	= 12;
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

Dans le cas ou le password correspond a celui de l'utilisateur en base de donnee, l'authentification est acceptee. Le serveur renvoie alors un message avec les informations necessaires pour les autres communications.

    .. code-block:: protobuf
       :emphasize-lines: 3,5

	Msg {
	    target = IDENT;
	    command = SUCCESS;
	    AuthorId:  id de l'utilisateur
	    SessionId: chaine de session unique, sert pour verifier que l'utilisateur est bien authentifie par la suite.
	    Ident: &message.Msg_Ident{
	    Login: pseudo;
	}

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

Cartes
------



3. Erreur

En cas d'erreur, le message est toujours identique, sauf le code d'erreur qui peut etre entre 5 et 8.

    .. code-block:: protobuf
       :emphasize-lines: 3,5

	5:erreur creation carte
	6:erreur modification carte
	7:erreur delete carte
	8: erreur get carte

Le message se presente sous la forme suivante :

    .. code-block:: protobuf
       :emphasize-lines: 3,5

	Msg {
	    target = CARDS;
	    command = ERROR;
	    author_id = authorId du message demandant une action;
	    session_id = sessionId du message demandant une action;
	    Error: &message.Msg_Error{
	    error_id = error_code; // Cette erreur provient de la l'enum decrit dans cette page
	}

Colonnes
--------

Erreurs
-------


message.proto
=============

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       package message;
       
       enum TARGET {
         USERS		= 1;
         COLUMNS	= 2;
         PROJECTS	= 3;
         CARDS		= 4;
         ADMIN		= 5;
         IDENT		= 6;
         NOTIF		= 7;
         METADATA      = 8;
       };
       
       enum CMD {
         CREATE	= 1;
         MODIFY	= 2;
         DELETE	= 3;
         GET		= 4;
         MOVE		= 5;
         CONNECT       = 6;
         DISCONNECT    = 7;
         ERROR		= 8;
         SUCCES        = 9;
         NONE          = 10;
       };
       
       message Msg {
         required TARGET	target = 1;
         required CMD		command = 2;
         required uint32	author_id = 3; // contains author id (who is speaking) in reception and addressee msg in sending
         required string	session_id = 4;
       
         message Users {
           required uint32	id = 1;
           required string	name = 2;
           required string	password = 3;
           required bool	admin = 4; // is SUPERadmin or not
           optional string	mail = 5;
         }
       
         message Columns {
           required uint32	project_id = 1;
           required uint32	id = 2;
           required string	name = 3;
           optional string	desc = 4;
           repeated string	tags = 5;
           repeated uint32	scripts_ids = 6; // IDs of the scripts attached to the column
           repeated uint32	write = 7; // list of the user IDs with write permission on the column (if empty: free for all)
         }
       
         message Projects {
           required uint32	id = 1;
           required string	name = 2;
           repeated uint32	admins_id = 3; // list of the administrator users of the projects
           repeated uint32	read = 4;  // list of the user IDs with read permission on the project (if empty: free for all)
         }
       
         message Cards {
           // Comments struct
           required uint32	id = 1;
           required uint32	project_id = 2;
           required uint32	column_id = 3;
           required string	name = 4;
           // repeated Comment	comments = 5; // repeated = dynamically sized array of Comments
           optional string	desc = 6;
           repeated string	tags = 7;
           optional uint32	user_id = 8; // ID of the card author
           repeated uint32	scripts_ids = 9; // IDs of the scripts attached to the card
           repeated uint32	write = 10; // // list of the user IDs with write permission on the card (if empty: free for all)
         }
       
         message Comment {
           required uint32	id = 1;
           required string	content = 2;
           required string	author_id = 3;
           required uint32     timestamp = 4;
           required uint32     card_id = 5;
         }
       
         message Metadata {
           required uint32     object_type = 1;
           required uint32     object_id = 2;
           optional string     data_key = 3;
           optional uint32     data_value = 4;
         }
       
         message Ident {
           required string	login = 1;
           optional string	password = 2;
         }
       
         message Error {
           required uint32	error_id = 1;
         }
       
         message Notif {
           optional string	msg = 1;
         }
       
         optional Users	users = 5;
         optional Columns	columns = 6;
         optional Projects	projects = 7;
         optional Cards	cards = 8;
         optional Ident	ident = 9;
         optional Error	error = 10;
         optional Notif	notif = 11;
       }



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
