.. highlightlang:: protobuf

.. _protocole:

Protocole
=========

Ceci est le protocole utilisé pour faire communiquer notre serveur et notre client.

Ce protocole utilise les Protobuf_ pour communiquer.

.. Ouverture d'une socket par client (client logiciel).
.. Chaque client peut parler au nom de plusieurs utilisateurs (ou author).

.. Tout message est precede d'un unsigned int pour preciser la taille du message qui va suivre.

.. - Identitifaction
.. Client -> Premier message d'identification a la connexion
.. Le author_id n'est pas connu au moment de l'identification, il importe donc peu
.. Idem pour le session_id
.. le login et le password sont en clair
.. Msg {
..     target = IDENT;
..     command = CONNECT;
..     author_id = -1
..     session_id = "";
..     Message Ident {
..     login = pseudo;
..     password = password;
..     }
.. }

.. Serveur -> Deux reponses possibles selon la validite de l'ident
.. En cas d'erreur:
.. Msg {
..     target = ERROR;
..     command = CONNECT;
..     author_id = -1
..     session_id = "";
..     Message Error {
..     error_id = error_connexion_id;
..     }
.. }

.. En cas de reussite:
.. Msg {
..     target = IDENT;
..     command = CONNECT;
..     author_id = id calcule par le serveur
..     session_id = session_id calcule par le serveur
.. }

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
