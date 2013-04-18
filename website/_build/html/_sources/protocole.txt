.. highlight:: rest

.. _protocole:

=========
Protocole
=========

Ceci est le protocole utilisé pour faire communiquer notre serveur et notre client.

Ce protocole utilise les Protobuf_ pour communiquer. Il le fait sur des sockets en tcp.

Chaque client ouvre une socket pour communiquer mais peut parler au nom de plusieurs utilisateurs. Ce qui permet par exemple a un client web de se connecter une seule fois au serveur, mais de gerer plusieurs utilisateurs qui se connectent en parallele.

La base de donnee obi-wan-kanbanobi est repartie en projets, contenant des colonnes contenant des cartes.

Information de base
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

1. Creation

Pour la creation d'une carte, le message doit etre le suivant :

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = CARDS;
           command = CREATE;
           author_id = id
           session_id = chaine de session unique
           message Cards {
             id = id sans importance a la creation;
             project_id = id du projet;
             column_id = id de la colonne;
             name = nom de la carte;
             desc = contenu de la carte;
             tags = les tags de la carte;
             user_id = id du createur;
             scripts_ids = id non utilise pour le moment; // IDs of the scripts attached to the card
             write = liste d'id d'utilisateurs qui ont les droits sur la carte;
           }
       }

2. Update

Les informations demandees pour un update sont les memes que celles pour la creation. Seule la partie de presentation du message change :

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = CARDS;
           command = UPDATE;
           author_id = id
           session_id = chaine de session unique
           message Cards {
             id = id de la carte;
             project_id = id du projet;
           }
       }


3. Delete

Pour le delete, les seules informations necessaires sont l'id et le project id :

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = CARDS;
           command = DELETE;
           author_id = id
           session_id = chaine de session unique
           message Cards {
             id = id de la carte;
             project_id = id du projet;
           }
       }

4. Get

Pour le get, les seules informations necessaires sont l'id et le project id :

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = CARDS;
           command = GET;
           author_id = id
           session_id = chaine de session unique
           message Cards {
             id = id de la carte;
             project_id = id du projet;
           }
       }

5. Success

    .. code-block:: protobuf
       :emphasize-lines: 3,5

       Msg {
           target = CARDS;
           command = SUCCESS;
           author_id = id
           session_id = chaine de session unique
       }

6. Erreur

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

La reception differentes colonnes composannt un projet s'effecture avec l'envoi d'un paquet MSG_COLUMN contenant l'identifiant du projet cible.
Le serveur peut repondre 


Erreurs
-------

    .. code-block:: protobuf
       :emphasize-lines: 3,5

        1:Utilisateur non connecte
        2:Pas les droits
        3:Commande invalide
        5:erreur creation carte
        6:erreur modification carte
        7:erreur delete carte
        8: erreur get carte
        11 - 20 -> user :
        - 11: erreur creation utilisateur
        - 12: erreur modification utilisateur
        - 13: erreur modification du mot de passe utilisateur
        - 14: erreur suppression utilisateur
        - 15: erreur get utilisateur
        - 16: erreur get user board
        - 17: erreur set admin
        - 18: erreur unset admin
        21-30 -> colonnes
        31 - 40 -> projet
        - 31 unspecified project error
        - 32 db_error.
        - 33 migrating (changement d'admin)
        41 -50 -> metadata
        - 41 unspecified metadata error
        - 42 erreur 42, serviette non trouvée
        - 43 metadataID inconnu
        - 44
        51-60 -> ident {
        51 -> Login error
        52 -> disconnect error
        }

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
          PASSWORD      = 11;
          GETBOARD	= 12;
        };
        
        message Msg {
          required TARGET	target = 1;
          required CMD		command = 2;
          required uint32	author_id = 3; // contains author id (who is speaking) in reception and addressee msg in sending
          required string	session_id = 4;
        
          message Password {
            required uint32	id = 1;
            required string	oldpassword = 2;
            required string	newpassword = 3;
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
          
          message Columns {
            required uint32	project_id = 1;
            required uint32	id = 2;
            required string	name = 3;
            optional string	desc = 4;
            repeated string	tags = 5;
            repeated uint32	scripts_ids = 6; // IDs of the scripts attached to the column
            repeated uint32	write = 7; // list of the user IDs with write permission on the column (if empty: free for all)
            repeated Cards 	ColumnCards = 8;
          }
        
          message Projects {
            required uint32	id = 1;
            required string	name = 2;
            required string	content = 3;
            repeated uint32	admins_id = 4; // list of the administrator users of the projects
            repeated uint32	read = 5;  // list of the user IDs with read permission on the project (if empty: free for all)
            repeated Columns	projectColumns = 6; 
          }
        
          message Comment {
            required uint32	id = 1;
            required string	content = 2;
            required string	author_id = 3;
            required uint32     timestamp = 4;
            required uint32     card_id = 5;
          }
        
          message Metadata {
            required uint32     id = 1;
            required uint32     object_type = 2;
            required uint32     object_id = 3;
            optional string     data_key = 4;
            optional string     data_value = 5;
          }
        
          message Users {
            required uint32	id = 1;
            required string	name = 2;
            optional string	password = 3;
            required bool	admin = 4; // is SUPERadmin or not
            optional string	mail = 5;
            repeated Projects	userProject = 6;  // permit to send the user the list of the project he can read
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
          optional Password     password = 12;
          optional Metadata	metadata = 13;
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
