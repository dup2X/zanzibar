namespace java com.uber.zanzibar.endpoints.nestedStructs

typedef string ItemType
const ItemType T_OPTIONAL = "optional"
const ItemType T_REQUIRED = "required"

enum ItemState {
	REQUIRED,
	OPTIONAL
}

typedef string UUID

struct RequiredExtra {
	1: required list<string> one
	2: required list<OptionalTerminal> two
	3: required set<string> three
	4: required set<OptionalTerminal> four
	5: required map<string, string> five
	6: required map<string, OptionalTerminal> six
	7: required UUID seven
	8: required ItemType eight
	9: required ItemState nine
	#10: required OptionalExtra ten
}

struct OptionalExtra {
	1: optional list<string> one
	2: optional list<OptionalTerminal> two
	3: optional set<string> three
	4: optional set<OptionalTerminal> four
	5: optional map<string, string> five
	6: optional map<string, OptionalTerminal> six
	7: optional UUID seven
	8: optional ItemType eight
	9: optional ItemState nine
	#10: optional RequiredExtra ten
}

struct OptionalTerminal {
	1: optional bool one
	2: optional byte two
	3: optional i16 three
	4: optional i32 four
	5: optional i64 five
	6: optional double six
	7: optional binary seven
	8: optional string eight
}

struct RequiredStruct {
	1: required bool one
	2: required byte two
	3: required i16 three
	4: required i32 four
	5: required i64 five
	6: required double six
	7: required binary seven
	8: required string eight
	9: required RequiredExtra nine
}

struct OptionalStruct {
	1: optional bool one
	2: optional byte two
	3: optional i16 three
	4: optional i32 four
	5: optional i64 five
	6: optional double six
	7: optional binary seven
	8: optional string eight
	9: optional OptionalExtra nine
}

service NestedStructs {
	RequiredStruct first(
		1: required bool one
		2: required byte two
		3: required i16 three
		4: required i32 four
		5: required i64 five
		6: required double six
		7: required binary seven
		8: required string eight
		9: required RequiredStruct nine
	) (
		zanzibar.http.method = "POST"
		zanzibar.http.path = "/first"
		zanzibar.http.status = "200"
	)
	OptionalStruct second(
		1: optional bool one
		2: optional byte two
		3: optional i16 three
		4: optional i32 four
		5: optional i64 five
		6: optional double six
		7: optional binary seven
		8: optional string eight
		9: optional OptionalStruct nine
	) (
		zanzibar.http.method = "POST"
		zanzibar.http.path = "/second"
		zanzibar.http.status = "200"
	)
}

