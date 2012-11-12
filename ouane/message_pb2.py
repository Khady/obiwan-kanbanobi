# Generated by the protocol buffer compiler.  DO NOT EDIT!

from google.protobuf import descriptor
from google.protobuf import message
from google.protobuf import reflection
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)



DESCRIPTOR = descriptor.FileDescriptor(
  name='message.proto',
  package='message',
  serialized_pb='\n\rmessage.proto\x12\x07message\"\xda\x07\n\x03Msg\x12\x1f\n\x06target\x18\x01 \x02(\x0e\x32\x0f.message.TARGET\x12\x1d\n\x07\x63ommand\x18\x02 \x02(\x0e\x32\x0c.message.CMD\x12\x11\n\tauthor_id\x18\x03 \x02(\r\x12\x12\n\nsession_id\x18\x04 \x02(\t\x12!\n\x05users\x18\x05 \x01(\x0b\x32\x12.message.Msg.Users\x12%\n\x07\x63olumns\x18\x06 \x01(\x0b\x32\x14.message.Msg.Columns\x12\'\n\x08projects\x18\x07 \x01(\x0b\x32\x15.message.Msg.Projects\x12!\n\x05\x63\x61rds\x18\x08 \x01(\x0b\x32\x12.message.Msg.Cards\x12!\n\x05ident\x18\t \x01(\x0b\x32\x12.message.Msg.Ident\x12!\n\x05\x65rror\x18\n \x01(\x0b\x32\x12.message.Msg.Error\x12!\n\x05notif\x18\x0b \x01(\x0b\x32\x12.message.Msg.Notif\x1aP\n\x05Users\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0c\n\x04name\x18\x02 \x02(\t\x12\x10\n\x08password\x18\x03 \x02(\t\x12\r\n\x05\x61\x64min\x18\x04 \x02(\x08\x12\x0c\n\x04mail\x18\x05 \x01(\t\x1aw\n\x07\x43olumns\x12\x12\n\nproject_id\x18\x01 \x02(\r\x12\n\n\x02id\x18\x02 \x02(\r\x12\x0c\n\x04name\x18\x03 \x02(\t\x12\x0c\n\x04\x64\x65sc\x18\x04 \x01(\t\x12\x0c\n\x04tags\x18\x05 \x01(\t\x12\x13\n\x0bscripts_ids\x18\x06 \x01(\r\x12\r\n\x05write\x18\x07 \x03(\r\x1a\x45\n\x08Projects\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0c\n\x04name\x18\x02 \x02(\t\x12\x11\n\tadmins_id\x18\x03 \x03(\r\x12\x0c\n\x04read\x18\x04 \x03(\r\x1a\x83\x02\n\x05\x43\x61rds\x12\n\n\x02id\x18\x01 \x02(\r\x12\x12\n\nproject_id\x18\x02 \x02(\r\x12\x11\n\tcolumn_id\x18\x03 \x02(\r\x12\x0c\n\x04name\x18\x04 \x02(\t\x12,\n\x08\x63omments\x18\x05 \x03(\x0b\x32\x1a.message.Msg.Cards.Comment\x12\x0c\n\x04\x64\x65sc\x18\x06 \x01(\t\x12\x0c\n\x04tags\x18\x07 \x01(\t\x12\x10\n\x08users_id\x18\x08 \x01(\t\x12\x13\n\x0bscripts_ids\x18\t \x01(\r\x12\r\n\x05write\x18\n \x03(\r\x1a\x39\n\x07\x43omment\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0f\n\x07\x63ontent\x18\x02 \x02(\t\x12\x11\n\tauthor_id\x18\x03 \x02(\t\x1a$\n\x05Ident\x12\r\n\x05login\x18\x01 \x02(\t\x12\x0c\n\x04hash\x18\x02 \x02(\t\x1a\x19\n\x05\x45rror\x12\x10\n\x08\x65rror_id\x18\x01 \x02(\r\x1a\x14\n\x05Notif\x12\x0b\n\x03msg\x18\x01 \x01(\t*U\n\x03\x43MD\x12\n\n\x06\x43REATE\x10\x01\x12\n\n\x06MODIFY\x10\x02\x12\n\n\x06\x44\x45LETE\x10\x03\x12\x07\n\x03GET\x10\x04\x12\x08\n\x04MOVE\x10\x05\x12\x0b\n\x07\x43ONNECT\x10\x06\x12\n\n\x06LOGOUT\x10\x07*e\n\x06TARGET\x12\t\n\x05USERS\x10\x01\x12\x0b\n\x07\x43OLUMNS\x10\x02\x12\x0c\n\x08PROJECTS\x10\x03\x12\t\n\x05\x43\x41RDS\x10\x04\x12\t\n\x05\x41\x44MIN\x10\x05\x12\t\n\x05IDENT\x10\x06\x12\t\n\x05\x45RROR\x10\x07\x12\t\n\x05NOTIF\x10\x08')

_CMD = descriptor.EnumDescriptor(
  name='CMD',
  full_name='message.CMD',
  filename=None,
  file=DESCRIPTOR,
  values=[
    descriptor.EnumValueDescriptor(
      name='CREATE', index=0, number=1,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='MODIFY', index=1, number=2,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='DELETE', index=2, number=3,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='GET', index=3, number=4,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='MOVE', index=4, number=5,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='CONNECT', index=5, number=6,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='LOGOUT', index=6, number=7,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1015,
  serialized_end=1100,
)


_TARGET = descriptor.EnumDescriptor(
  name='TARGET',
  full_name='message.TARGET',
  filename=None,
  file=DESCRIPTOR,
  values=[
    descriptor.EnumValueDescriptor(
      name='USERS', index=0, number=1,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='COLUMNS', index=1, number=2,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='PROJECTS', index=2, number=3,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='CARDS', index=3, number=4,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='ADMIN', index=4, number=5,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='IDENT', index=5, number=6,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='ERROR', index=6, number=7,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='NOTIF', index=7, number=8,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1102,
  serialized_end=1203,
)


CREATE = 1
MODIFY = 2
DELETE = 3
GET = 4
MOVE = 5
CONNECT = 6
LOGOUT = 7
USERS = 1
COLUMNS = 2
PROJECTS = 3
CARDS = 4
ADMIN = 5
IDENT = 6
ERROR = 7
NOTIF = 8



_MSG_USERS = descriptor.Descriptor(
  name='Users',
  full_name='message.Msg.Users',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Users.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='name', full_name='message.Msg.Users.name', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='password', full_name='message.Msg.Users.password', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='admin', full_name='message.Msg.Users.admin', index=3,
      number=4, type=8, cpp_type=7, label=2,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='mail', full_name='message.Msg.Users.mail', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=392,
  serialized_end=472,
)

_MSG_COLUMNS = descriptor.Descriptor(
  name='Columns',
  full_name='message.Msg.Columns',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='project_id', full_name='message.Msg.Columns.project_id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Columns.id', index=1,
      number=2, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='name', full_name='message.Msg.Columns.name', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='desc', full_name='message.Msg.Columns.desc', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='tags', full_name='message.Msg.Columns.tags', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='scripts_ids', full_name='message.Msg.Columns.scripts_ids', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='write', full_name='message.Msg.Columns.write', index=6,
      number=7, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=474,
  serialized_end=593,
)

_MSG_PROJECTS = descriptor.Descriptor(
  name='Projects',
  full_name='message.Msg.Projects',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Projects.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='name', full_name='message.Msg.Projects.name', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='admins_id', full_name='message.Msg.Projects.admins_id', index=2,
      number=3, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='read', full_name='message.Msg.Projects.read', index=3,
      number=4, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=595,
  serialized_end=664,
)

_MSG_CARDS_COMMENT = descriptor.Descriptor(
  name='Comment',
  full_name='message.Msg.Cards.Comment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Cards.Comment.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='content', full_name='message.Msg.Cards.Comment.content', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='author_id', full_name='message.Msg.Cards.Comment.author_id', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=869,
  serialized_end=926,
)

_MSG_CARDS = descriptor.Descriptor(
  name='Cards',
  full_name='message.Msg.Cards',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Cards.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='project_id', full_name='message.Msg.Cards.project_id', index=1,
      number=2, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='column_id', full_name='message.Msg.Cards.column_id', index=2,
      number=3, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='name', full_name='message.Msg.Cards.name', index=3,
      number=4, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='comments', full_name='message.Msg.Cards.comments', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='desc', full_name='message.Msg.Cards.desc', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='tags', full_name='message.Msg.Cards.tags', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='users_id', full_name='message.Msg.Cards.users_id', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='scripts_ids', full_name='message.Msg.Cards.scripts_ids', index=8,
      number=9, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='write', full_name='message.Msg.Cards.write', index=9,
      number=10, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_MSG_CARDS_COMMENT, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=667,
  serialized_end=926,
)

_MSG_IDENT = descriptor.Descriptor(
  name='Ident',
  full_name='message.Msg.Ident',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='login', full_name='message.Msg.Ident.login', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='hash', full_name='message.Msg.Ident.hash', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=928,
  serialized_end=964,
)

_MSG_ERROR = descriptor.Descriptor(
  name='Error',
  full_name='message.Msg.Error',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='error_id', full_name='message.Msg.Error.error_id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=966,
  serialized_end=991,
)

_MSG_NOTIF = descriptor.Descriptor(
  name='Notif',
  full_name='message.Msg.Notif',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='msg', full_name='message.Msg.Notif.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=993,
  serialized_end=1013,
)

_MSG = descriptor.Descriptor(
  name='Msg',
  full_name='message.Msg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='target', full_name='message.Msg.target', index=0,
      number=1, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=1,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='command', full_name='message.Msg.command', index=1,
      number=2, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=1,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='author_id', full_name='message.Msg.author_id', index=2,
      number=3, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='session_id', full_name='message.Msg.session_id', index=3,
      number=4, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='users', full_name='message.Msg.users', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='columns', full_name='message.Msg.columns', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='projects', full_name='message.Msg.projects', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='cards', full_name='message.Msg.cards', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='ident', full_name='message.Msg.ident', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='error', full_name='message.Msg.error', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='notif', full_name='message.Msg.notif', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_MSG_USERS, _MSG_COLUMNS, _MSG_PROJECTS, _MSG_CARDS, _MSG_IDENT, _MSG_ERROR, _MSG_NOTIF, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=27,
  serialized_end=1013,
)

_MSG_USERS.containing_type = _MSG;
_MSG_COLUMNS.containing_type = _MSG;
_MSG_PROJECTS.containing_type = _MSG;
_MSG_CARDS_COMMENT.containing_type = _MSG_CARDS;
_MSG_CARDS.fields_by_name['comments'].message_type = _MSG_CARDS_COMMENT
_MSG_CARDS.containing_type = _MSG;
_MSG_IDENT.containing_type = _MSG;
_MSG_ERROR.containing_type = _MSG;
_MSG_NOTIF.containing_type = _MSG;
_MSG.fields_by_name['target'].enum_type = _TARGET
_MSG.fields_by_name['command'].enum_type = _CMD
_MSG.fields_by_name['users'].message_type = _MSG_USERS
_MSG.fields_by_name['columns'].message_type = _MSG_COLUMNS
_MSG.fields_by_name['projects'].message_type = _MSG_PROJECTS
_MSG.fields_by_name['cards'].message_type = _MSG_CARDS
_MSG.fields_by_name['ident'].message_type = _MSG_IDENT
_MSG.fields_by_name['error'].message_type = _MSG_ERROR
_MSG.fields_by_name['notif'].message_type = _MSG_NOTIF
DESCRIPTOR.message_types_by_name['Msg'] = _MSG

class Msg(message.Message):
  __metaclass__ = reflection.GeneratedProtocolMessageType
  
  class Users(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_USERS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Users)
  
  class Columns(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_COLUMNS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Columns)
  
  class Projects(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_PROJECTS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Projects)
  
  class Cards(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    
    class Comment(message.Message):
      __metaclass__ = reflection.GeneratedProtocolMessageType
      DESCRIPTOR = _MSG_CARDS_COMMENT
      
      # @@protoc_insertion_point(class_scope:message.Msg.Cards.Comment)
    DESCRIPTOR = _MSG_CARDS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Cards)
  
  class Ident(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_IDENT
    
    # @@protoc_insertion_point(class_scope:message.Msg.Ident)
  
  class Error(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_ERROR
    
    # @@protoc_insertion_point(class_scope:message.Msg.Error)
  
  class Notif(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_NOTIF
    
    # @@protoc_insertion_point(class_scope:message.Msg.Notif)
  DESCRIPTOR = _MSG
  
  # @@protoc_insertion_point(class_scope:message.Msg)

# @@protoc_insertion_point(module_scope)
