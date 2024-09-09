// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EntityChanges {
    #[prost(message, repeated, tag="5")]
    pub entity_changes: ::prost::alloc::vec::Vec<EntityChange>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EntityChange {
    #[prost(string, tag="1")]
    pub entity: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub id: ::prost::alloc::string::String,
    /// Deprecated, this is not used within `graph-node`.
    #[prost(uint64, tag="3")]
    #[deprecated(since="1.3.0",note="not used by 'graph-node', will be removed")]
    pub ordinal: u64,
    #[prost(enumeration="entity_change::Operation", tag="4")]
    pub operation: i32,
    #[prost(message, repeated, tag="5")]
    pub fields: ::prost::alloc::vec::Vec<Field>,
}
/// Nested message and enum types in `EntityChange`.
pub mod entity_change {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Operation {
        /// Protobuf default should not be used, this is used so that the consume can ensure that the value was actually specified
        Unspecified = 0,
        Create = 1,
        Update = 2,
        Delete = 3,
        Final = 4,
    }
    impl Operation {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Operation::Unspecified => "OPERATION_UNSPECIFIED",
                Operation::Create => "OPERATION_CREATE",
                Operation::Update => "OPERATION_UPDATE",
                Operation::Delete => "OPERATION_DELETE",
                Operation::Final => "OPERATION_FINAL",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "OPERATION_UNSPECIFIED" => Some(Self::Unspecified),
                "OPERATION_CREATE" => Some(Self::Create),
                "OPERATION_UPDATE" => Some(Self::Update),
                "OPERATION_DELETE" => Some(Self::Delete),
                "OPERATION_FINAL" => Some(Self::Final),
                _ => None,
            }
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Value {
    #[prost(oneof="value::Typed", tags="1, 2, 3, 4, 5, 6, 7, 10")]
    pub typed: ::core::option::Option<value::Typed>,
}
/// Nested message and enum types in `Value`.
pub mod value {
    #[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Typed {
        #[prost(int32, tag="1")]
        Int32(i32),
        #[prost(string, tag="2")]
        Bigdecimal(::prost::alloc::string::String),
        #[prost(string, tag="3")]
        Bigint(::prost::alloc::string::String),
        #[prost(string, tag="4")]
        String(::prost::alloc::string::String),
        #[prost(string, tag="5")]
        Bytes(::prost::alloc::string::String),
        #[prost(bool, tag="6")]
        Bool(bool),
        #[prost(int64, tag="7")]
        Timestamp(i64),
        // reserved 8 to 9;  // For future types

        #[prost(message, tag="10")]
        Array(super::Array),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Array {
    #[prost(message, repeated, tag="1")]
    pub value: ::prost::alloc::vec::Vec<Value>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Field {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, optional, tag="3")]
    pub new_value: ::core::option::Option<Value>,
    /// Deprecated, this is not used within `graph-node`.
    #[prost(message, optional, tag="5")]
    #[deprecated(since="1.3.0",note="not used by 'graph-node', will be removed")]
    pub old_value: ::core::option::Option<Value>,
}
// @@protoc_insertion_point(module)
