package common

type FilterBy struct {
	Field string
	Value interface{}
	Op    string
}

func Equal(field string, value interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: value,
		Op:    "=",
	}
}

func GreaterThan(field string, value interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: value,
		Op:    ">",
	}
}

func LessThan(field string, value interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: value,
		Op:    "<",
	}
}

func NotEqual(field string, value interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: value,
		Op:    "!=",
	}
}

func Like(field string, value interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: value,
		Op:    "LIKE",
	}
}

func In(field string, values []interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: values,
		Op:    "IN",
	}
}

func NotIn(field string, values []interface{}) FilterBy {
	return FilterBy{
		Field: field,
		Value: values,
		Op:    "NOT IN",
	}
}

func And(filters ...FilterBy) FilterBy {
	return FilterBy{
		Field: "",
		Value: filters,
		Op:    "AND",
	}
}

func Or(filters ...FilterBy) FilterBy {
	return FilterBy{
		Field: "",
		Value: filters,
		Op:    "OR",
	}
}
