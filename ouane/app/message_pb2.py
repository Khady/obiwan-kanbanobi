# Generated by the protocol buffer compiler.  DO NOT EDIT!

from google.protobuf import descriptor
from google.protobuf import message
from google.protobuf import reflection
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)



DESCRIPTOR = descriptor.FileDescriptor(
  name='message.proto',
  package='message',
  serialized_pb='\n\rmessage.proto\x12\x07message\"\xae\n\n\x03Msg\x12\x1f\n\x06target\x18\x01 \x02(\x0e\x32\x0f.message.TARGET\x12\x1d\n\x07\x63ommand\x18\x02 \x02(\x0e\x32\x0c.message.CMD\x12\x11\n\tauthor_id\x18\x03 \x02(\r\x12\x12\n\nsession_id\x18\x04 \x02(\t\x12!\n\x05users\x18\x05 \x01(\x0b\x32\x12.message.Msg.Users\x12%\n\x07\x63olumns\x18\x06 \x01(\x0b\x32\x14.message.Msg.Columns\x12\'\n\x08projects\x18\x07 \x01(\x0b\x32\x15.message.Msg.Projects\x12!\n\x05\x63\x61rds\x18\x08 \x01(\x0b\x32\x12.message.Msg.Cards\x12!\n\x05ident\x18\t \x01(\x0b\x32\x12.message.Msg.Ident\x12!\n\x05\x65rror\x18\n \x01(\x0b\x32\x12.message.Msg.Error\x12!\n\x05notif\x18\x0b \x01(\x0b\x32\x12.message.Msg.Notif\x12\'\n\x08password\x18\x0c \x01(\x0b\x32\x15.message.Msg.Password\x1a@\n\x08Password\x12\n\n\x02id\x18\x01 \x02(\r\x12\x13\n\x0boldpassword\x18\x02 \x02(\t\x12\x13\n\x0bnewpassword\x18\x03 \x02(\t\x1a\x99\x01\n\x05\x43\x61rds\x12\n\n\x02id\x18\x01 \x02(\r\x12\x12\n\nproject_id\x18\x02 \x02(\r\x12\x11\n\tcolumn_id\x18\x03 \x02(\r\x12\x0c\n\x04name\x18\x04 \x02(\t\x12\x0c\n\x04\x64\x65sc\x18\x06 \x01(\t\x12\x0c\n\x04tags\x18\x07 \x03(\t\x12\x0f\n\x07user_id\x18\x08 \x01(\r\x12\x13\n\x0bscripts_ids\x18\t \x03(\r\x12\r\n\x05write\x18\n \x03(\r\x1a\xa0\x01\n\x07\x43olumns\x12\x12\n\nproject_id\x18\x01 \x02(\r\x12\n\n\x02id\x18\x02 \x02(\r\x12\x0c\n\x04name\x18\x03 \x02(\t\x12\x0c\n\x04\x64\x65sc\x18\x04 \x01(\t\x12\x0c\n\x04tags\x18\x05 \x03(\t\x12\x13\n\x0bscripts_ids\x18\x06 \x03(\r\x12\r\n\x05write\x18\x07 \x03(\r\x12\'\n\x0b\x43olumnCards\x18\x08 \x03(\x0b\x32\x12.message.Msg.Cards\x1a\x84\x01\n\x08Projects\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0c\n\x04name\x18\x02 \x02(\t\x12\x0f\n\x07\x63ontent\x18\x03 \x02(\t\x12\x11\n\tadmins_id\x18\x04 \x03(\r\x12\x0c\n\x04read\x18\x05 \x03(\r\x12,\n\x0eprojectColumns\x18\x06 \x03(\x0b\x32\x14.message.Msg.Columns\x1a]\n\x07\x43omment\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0f\n\x07\x63ontent\x18\x02 \x02(\t\x12\x11\n\tauthor_id\x18\x03 \x02(\t\x12\x11\n\ttimestamp\x18\x04 \x02(\r\x12\x0f\n\x07\x63\x61rd_id\x18\x05 \x02(\r\x1aX\n\x08Metadata\x12\x13\n\x0bobject_type\x18\x01 \x02(\r\x12\x11\n\tobject_id\x18\x02 \x02(\r\x12\x10\n\x08\x64\x61ta_key\x18\x03 \x01(\t\x12\x12\n\ndata_value\x18\x04 \x01(\r\x1a|\n\x05Users\x12\n\n\x02id\x18\x01 \x02(\r\x12\x0c\n\x04name\x18\x02 \x02(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\r\n\x05\x61\x64min\x18\x04 \x02(\x08\x12\x0c\n\x04mail\x18\x05 \x01(\t\x12*\n\x0buserProject\x18\x06 \x03(\x0b\x32\x15.message.Msg.Projects\x1a(\n\x05Ident\x12\r\n\x05login\x18\x01 \x02(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x1a\x19\n\x05\x45rror\x12\x10\n\x08\x65rror_id\x18\x01 \x02(\r\x1a\x14\n\x05Notif\x12\x0b\n\x03msg\x18\x01 \x01(\t*h\n\x06TARGET\x12\t\n\x05USERS\x10\x01\x12\x0b\n\x07\x43OLUMNS\x10\x02\x12\x0c\n\x08PROJECTS\x10\x03\x12\t\n\x05\x43\x41RDS\x10\x04\x12\t\n\x05\x41\x44MIN\x10\x05\x12\t\n\x05IDENT\x10\x06\x12\t\n\x05NOTIF\x10\x07\x12\x0c\n\x08METADATA\x10\x08*\x96\x01\n\x03\x43MD\x12\n\n\x06\x43REATE\x10\x01\x12\n\n\x06MODIFY\x10\x02\x12\n\n\x06\x44\x45LETE\x10\x03\x12\x07\n\x03GET\x10\x04\x12\x08\n\x04MOVE\x10\x05\x12\x0b\n\x07\x43ONNECT\x10\x06\x12\x0e\n\nDISCONNECT\x10\x07\x12\t\n\x05\x45RROR\x10\x08\x12\n\n\x06SUCCES\x10\t\x12\x08\n\x04NONE\x10\n\x12\x0c\n\x08PASSWORD\x10\x0b\x12\x0c\n\x08GETBOARD\x10\x0c')

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
      name='NOTIF', index=6, number=7,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='METADATA', index=7, number=8,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1355,
  serialized_end=1459,
)


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
      name='DISCONNECT', index=6, number=7,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='ERROR', index=7, number=8,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='SUCCES', index=8, number=9,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='NONE', index=9, number=10,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='PASSWORD', index=10, number=11,
      options=None,
      type=None),
    descriptor.EnumValueDescriptor(
      name='GETBOARD', index=11, number=12,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1462,
  serialized_end=1612,
)


USERS = 1
COLUMNS = 2
PROJECTS = 3
CARDS = 4
ADMIN = 5
IDENT = 6
NOTIF = 7
METADATA = 8
CREATE = 1
MODIFY = 2
DELETE = 3
GET = 4
MOVE = 5
CONNECT = 6
DISCONNECT = 7
ERROR = 8
SUCCES = 9
NONE = 10
PASSWORD = 11
GETBOARD = 12



_MSG_PASSWORD = descriptor.Descriptor(
  name='Password',
  full_name='message.Msg.Password',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Password.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='oldpassword', full_name='message.Msg.Password.oldpassword', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='newpassword', full_name='message.Msg.Password.newpassword', index=2,
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
  serialized_start=433,
  serialized_end=497,
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
      name='desc', full_name='message.Msg.Cards.desc', index=4,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='tags', full_name='message.Msg.Cards.tags', index=5,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='user_id', full_name='message.Msg.Cards.user_id', index=6,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='scripts_ids', full_name='message.Msg.Cards.scripts_ids', index=7,
      number=9, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='write', full_name='message.Msg.Cards.write', index=8,
      number=10, type=13, cpp_type=3, label=3,
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
  serialized_start=500,
  serialized_end=653,
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
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='scripts_ids', full_name='message.Msg.Columns.scripts_ids', index=5,
      number=6, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
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
    descriptor.FieldDescriptor(
      name='ColumnCards', full_name='message.Msg.Columns.ColumnCards', index=7,
      number=8, type=11, cpp_type=10, label=3,
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
  serialized_start=656,
  serialized_end=816,
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
      name='content', full_name='message.Msg.Projects.content', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='admins_id', full_name='message.Msg.Projects.admins_id', index=3,
      number=4, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='read', full_name='message.Msg.Projects.read', index=4,
      number=5, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='projectColumns', full_name='message.Msg.Projects.projectColumns', index=5,
      number=6, type=11, cpp_type=10, label=3,
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
  serialized_start=819,
  serialized_end=951,
)

_MSG_COMMENT = descriptor.Descriptor(
  name='Comment',
  full_name='message.Msg.Comment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='id', full_name='message.Msg.Comment.id', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='content', full_name='message.Msg.Comment.content', index=1,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='author_id', full_name='message.Msg.Comment.author_id', index=2,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='timestamp', full_name='message.Msg.Comment.timestamp', index=3,
      number=4, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='card_id', full_name='message.Msg.Comment.card_id', index=4,
      number=5, type=13, cpp_type=3, label=2,
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
  serialized_start=953,
  serialized_end=1046,
)

_MSG_METADATA = descriptor.Descriptor(
  name='Metadata',
  full_name='message.Msg.Metadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    descriptor.FieldDescriptor(
      name='object_type', full_name='message.Msg.Metadata.object_type', index=0,
      number=1, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='object_id', full_name='message.Msg.Metadata.object_id', index=1,
      number=2, type=13, cpp_type=3, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='data_key', full_name='message.Msg.Metadata.data_key', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    descriptor.FieldDescriptor(
      name='data_value', full_name='message.Msg.Metadata.data_value', index=3,
      number=4, type=13, cpp_type=3, label=1,
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
  serialized_start=1048,
  serialized_end=1136,
)

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
      number=3, type=9, cpp_type=9, label=1,
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
    descriptor.FieldDescriptor(
      name='userProject', full_name='message.Msg.Users.userProject', index=5,
      number=6, type=11, cpp_type=10, label=3,
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
  serialized_start=1138,
  serialized_end=1262,
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
      name='password', full_name='message.Msg.Ident.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=1264,
  serialized_end=1304,
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
  serialized_start=1306,
  serialized_end=1331,
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
  serialized_start=1333,
  serialized_end=1353,
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
    descriptor.FieldDescriptor(
      name='password', full_name='message.Msg.password', index=11,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_MSG_PASSWORD, _MSG_CARDS, _MSG_COLUMNS, _MSG_PROJECTS, _MSG_COMMENT, _MSG_METADATA, _MSG_USERS, _MSG_IDENT, _MSG_ERROR, _MSG_NOTIF, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=27,
  serialized_end=1353,
)

_MSG_PASSWORD.containing_type = _MSG;
_MSG_CARDS.containing_type = _MSG;
_MSG_COLUMNS.fields_by_name['ColumnCards'].message_type = _MSG_CARDS
_MSG_COLUMNS.containing_type = _MSG;
_MSG_PROJECTS.fields_by_name['projectColumns'].message_type = _MSG_COLUMNS
_MSG_PROJECTS.containing_type = _MSG;
_MSG_COMMENT.containing_type = _MSG;
_MSG_METADATA.containing_type = _MSG;
_MSG_USERS.fields_by_name['userProject'].message_type = _MSG_PROJECTS
_MSG_USERS.containing_type = _MSG;
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
_MSG.fields_by_name['password'].message_type = _MSG_PASSWORD
DESCRIPTOR.message_types_by_name['Msg'] = _MSG

class Msg(message.Message):
  __metaclass__ = reflection.GeneratedProtocolMessageType
  
  class Password(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_PASSWORD
    
    # @@protoc_insertion_point(class_scope:message.Msg.Password)
  
  class Cards(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_CARDS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Cards)
  
  class Columns(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_COLUMNS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Columns)
  
  class Projects(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_PROJECTS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Projects)
  
  class Comment(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_COMMENT
    
    # @@protoc_insertion_point(class_scope:message.Msg.Comment)
  
  class Metadata(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_METADATA
    
    # @@protoc_insertion_point(class_scope:message.Msg.Metadata)
  
  class Users(message.Message):
    __metaclass__ = reflection.GeneratedProtocolMessageType
    DESCRIPTOR = _MSG_USERS
    
    # @@protoc_insertion_point(class_scope:message.Msg.Users)
  
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
