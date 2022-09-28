package schema

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/eolinker/eosc/log"
	"github.com/stretchr/testify/assert"
)

func Example() {

	type myDate struct {
		Day   string `json:"day,omitempty"`
		Month string `json:"month,omitempty"`
		Year  string `json:"year,omitempty"`
	}

	type MyObject struct {
		ID     string  `json:"id,omitempty" doc:"Object ID" readOnly:"true"`
		Rate   float64 `doc:"Rate of change" minimum:"0"`
		Coords []int   `doc:"X,Y coordinates" minItems:"2" maxItems:"2"`
		myDate `json:"date,omitempty" dependencies:"day:month;year month:year"`
		Bucket map[string]interface{} `json:"bucket,omitempty" dependencies:"apple:banana;peach banana:melon"`
		Target RequireId              `json:"target" skill:"github.com/eolinker/apinto/service.service.IService"`
	}

	generated, err := Generate(reflect.TypeOf(MyObject{}), nil)
	if err != nil {
		panic(err)
	}
	bytes, _ := json.MarshalIndent(generated, "", "\t")
	log.Debug(string(bytes))
	// output: {"type":"object","properties":{"bucket":{"type":"object","additionalProperties":{},"dependencies":{"apple":["banana","peach"],"banana":["melon"]}},"coords":{"type":"array","description":"X,Y coordinates","items":{"type":"integer","format":"int32"},"minItems":2,"maxItems":2},"date":{"type":"object","properties":{"day":{"type":"string"},"month":{"type":"string"},"year":{"type":"string"}},"additionalProperties":false,"dependencies":{"day":["month","year"],"month":["year"]}},"id":{"type":"string","description":"Object ID","readOnly":true},"rate":{"type":"number","description":"Rate of change","format":"double","minimum":0}},"additionalProperties":false,"required":["rate","coords"],"dependencies":{"id":["rate"],"rate":["coords","date"]}}
}

var types = []struct {
	in     interface{}
	out    string
	format string
}{
	{false, "boolean", ""},
	{0, "integer", "int32"},
	{int64(0), "integer", "int64"},
	{uint64(0), "integer", "int64"},
	{float32(0), "number", "float"},
	{0.0, "number", "double"},
	{F(0.0), "number", "double"},
	{"hello", "string", ""},
	{struct{}{}, "object", ""},
	{[]string{"foo"}, "array", ""},
	{net.IP{}, "string", "ipv4"},
	{time.Time{}, "string", "date-time"},
	{url.URL{}, "string", "uri"},
	{[]byte{}, "string", ""},
	// TODO: map
}

func TestSchemaTypes(outer *testing.T) {
	outer.Parallel()
	for _, tt := range types {
		local := tt
		outer.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			t.Parallel()
			s, err := Generate(reflect.ValueOf(local.in).Type(), nil)
			assert.NoError(t, err)
			assert.Equal(t, local.out, s.Type)
			assert.Equal(t, local.format, s.Format)
		})
	}
}

func TestSchemaRequiredFields(t *testing.T) {
	type Example struct {
		Optional string `json:"optional,omitempty"`
		Required string `json:"required"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Len(t, s.Properties, 2)
	assert.NotContains(t, s.Required, "optional")
	assert.Contains(t, s.Required, "required")
}

func TestSchemaRenameField(t *testing.T) {
	type Example struct {
		Foo string `json:"bar"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Empty(t, s.findProperties("foo"))
	assert.NotEmpty(t, s.findProperties("bar"))
}

func TestSchemaDescription(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" description:"I am a test"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "I am a test", s.Properties["foo"].Description)
}

func TestSchemaFormat(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" format:"date-time"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "date-time", s.Properties["foo"].Format)
}

func TestSchemaEnum(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" enum:"one,two,three"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"one", "two", "three"}, s.findProperties("foo").Enum)
}

func TestSchemaArrayEnum(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" enum:"one,two,three"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)

	// Array itself should not have an enum member set.
	assert.Equal(t, []interface{}{}, s.findProperties("foo").Enum)

	// Items in the array should be one of the allowed enum values.
	assert.Equal(t, []interface{}{"one", "two", "three"}, s.findProperties("foo").Items.Enum)
}

func TestSchemaDefault(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" default:"def"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "def", s.findProperties("foo").Default)
}

func TestSchemaExample(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" example:"ex"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "ex", s.findProperties("foo").Example)
}

func TestSchemaNullable(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" nullable:"true"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, true, s.findProperties("foo").Nullable)
}

func TestSchemaNullableError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" nullable:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaReadOnly(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" readOnly:"true"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, true, s.findProperties("foo").ReadOnly)
}

func TestSchemaReadOnlyError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" readOnly:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaWriteOnly(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" writeOnly:"true"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, true, s.findProperties("foo").WriteOnly)
}

func TestSchemaWriteOnlyError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" writeOnly:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaDeprecated(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" deprecated:"true"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, true, s.findProperties("foo").Deprecated)
}

func TestSchemaDeprecatedError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" deprecated:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMinimum(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" minimum:"1"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1.0, *s.findProperties("foo").Minimum)
}

func TestSchemaMinimumError(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" minimum:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaExclusiveMinimum(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" exclusiveMinimum:"1"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1.0, *s.findProperties("foo").Minimum)
	assert.Equal(t, true, *s.findProperties("foo").ExclusiveMinimum)
}

func TestSchemaExclusiveMinimumError(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" exclusiveMinimum:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMaximum(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" maximum:"0"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, *s.findProperties("foo").Maximum)
}

func TestSchemaMaximumError(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" maximum:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaExclusiveMaximum(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" exclusiveMaximum:"0"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, *s.findProperties("foo").Maximum)
	assert.Equal(t, true, *s.findProperties("foo").ExclusiveMaximum)
}

func TestSchemaExclusiveMaximumError(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" exclusiveMaximum:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMultipleOf(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" multipleOf:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, s.findProperties("foo").MultipleOf)
}

func TestSchemaMultipleOfError(t *testing.T) {
	type Example struct {
		Foo float64 `json:"foo" multipleOf:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMinLength(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" minLength:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MinLength)
}

func TestSchemaMinLengthError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" minLength:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMaxLength(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" maxLength:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MaxLength)
}

func TestSchemaMaxLengthError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" maxLength:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaPattern(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" pattern:"a-z+"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, "a-z+", s.findProperties("foo").Pattern)
}

func TestSchemaPatternError(t *testing.T) {
	type Example struct {
		Foo string `json:"foo" pattern:"(.*"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMinItems(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" minItems:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MinItems)
}

func TestSchemaMinItemsError(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" minItems:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMaxItems(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" maxItems:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MaxItems)
}

func TestSchemaMaxItemsError(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" maxItems:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaUniqueItems(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" uniqueItems:"true"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, true, s.findProperties("foo").UniqueItems)
}

func TestSchemaUniqueItemsError(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" uniqueItems:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMinProperties(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" minProperties:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MinProperties)
}

func TestSchemaMinPropertiesError(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" minProperties:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaMaxProperties(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" maxProperties:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), *s.findProperties("foo").MaxProperties)
}

func TestSchemaMaxPropertiesError(t *testing.T) {
	type Example struct {
		Foo []string `json:"foo" maxProperties:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

//func TestSchemaMap(t *testing.T) {
//	s, err := Generate(reflect.TypeOf(map[string]string{}), nil)
//	assert.NoError(t, err)
//	assert.Equal(t, &Schema{
//		Type: "object",
//		AdditionalProperties: &Schema{
//			Type: "string",
//		},
//	}, s)
//}

func TestSchemaSlice(t *testing.T) {
	s, err := Generate(reflect.TypeOf([]string{}), nil)
	assert.NoError(t, err)
	assert.Equal(t, &Schema{
		Type: "array",
		Items: &Schema{
			Type: "string",
		},
	}, s)
}

func TestSchemaUnsigned(t *testing.T) {
	s, err := Generate(reflect.TypeOf(uint(10)), nil)
	assert.NoError(t, err)
	min := 0.0
	assert.Equal(t, &Schema{
		Type:    "integer",
		Format:  "int32",
		Minimum: &min,
	}, s)
}

func TestSchemaNonStringExample(t *testing.T) {
	type Example struct {
		Foo uint32 `json:"foo" example:"10"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, uint32(10), s.findProperties("foo").Example)
}

func TestSchemaNonStringExampleErrorUnmarshal(t *testing.T) {
	type Example struct {
		Foo uint32 `json:"foo" example:"bad"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaNonStringExampleErrorCast(t *testing.T) {
	type Example struct {
		Foo bool `json:"foo" example:"1"`
	}

	_, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.Error(t, err)
}

func TestSchemaFieldFilteredOut(t *testing.T) {
	type Example struct {
		Foo bool `json:"-"`
	}

	s, err := Generate(reflect.ValueOf(Example{}).Type(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(s.Properties))
}

func TestPointerHelpers(t *testing.T) {
	// Just confirm this compiles.
	_ = Schema{
		Minimum:   F(98.6),
		MinLength: I(5),
	}
}

func TestHasValidation(t *testing.T) {
	s := Schema{
		Type: "string",
	}
	assert.Equal(t, false, s.HasValidation())

	s.Pattern = "^[a-z]+$"
	assert.Equal(t, true, s.HasValidation())
}

func TestEmbedded(t *testing.T) {
	type Foo struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	type Bar struct {
		*Foo
		B int    `json:"b"`
		C string `json:"c"`
	}

	s, err := Generate(reflect.TypeOf(Bar{}), nil)
	assert.NoError(t, err)

	assert.Len(t, s.Properties, 3)
	assert.Equal(t, "integer", s.findProperties("b").Type)
}

func TestStringArrayExample(t *testing.T) {
	type Foo struct {
		A []string `json:"a" example:"[\"a\",\"b\",\"c\"]"`
		B []string `json:"b" example:"a, b, c"`
	}

	s, err := Generate(reflect.TypeOf(Foo{}), nil)
	assert.NoError(t, err)

	assert.Equal(t, s.findProperties("a").Example, []string{"a", "b", "c"})
	assert.Equal(t, s.findProperties("b").Example, []string{"a", "b", "c"})
}

func TestExampleBadJSON(t *testing.T) {
	type Foo struct {
		A []string `json:"a" example:"[\"a\", 1]"`
	}

	_, err := Generate(reflect.TypeOf(Foo{}), nil)
	assert.Error(t, err)
}

func TestGenerate(t *testing.T) {
	type myDate struct {
		Day   string `json:"day,omitempty"`
		Month string `json:"month,omitempty"`
		Year  string `json:"year,omitempty"`
	}
	type MyObject struct {
		ID   string  `json:"id"`
		Name string  `json:"name,omitempty"`
		Date *myDate `json:"date,omitempty" dependencies:"day:month;year"`
	}

	type args struct {
		t      reflect.Type
		schema *Schema
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "对MyObject和date加上依赖性限制",
			args:    args{reflect.TypeOf(MyObject{}), &Schema{Dependencies: map[string][]string{"name": {"date"}}}},
			want:    "{\"type\":\"object\",\"properties\":{\"date\":{\"type\":\"object\",\"properties\":{\"day\":{\"type\":\"string\"},\"month\":{\"type\":\"string\"},\"year\":{\"type\":\"string\"}},\"additionalProperties\":false,\"dependencies\":{\"day\":[\"month\",\"year\"]}},\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"additionalProperties\":false,\"required\":[\"id\"],\"dependencies\":{\"name\":[\"date\"]}}",
			wantErr: false,
		},
		{
			name:    "对date加上依赖性限制",
			args:    args{reflect.TypeOf(MyObject{}), nil},
			want:    "{\"type\":\"object\",\"properties\":{\"date\":{\"type\":\"object\",\"properties\":{\"day\":{\"type\":\"string\"},\"month\":{\"type\":\"string\"},\"year\":{\"type\":\"string\"}},\"additionalProperties\":false,\"dependencies\":{\"day\":[\"month\",\"year\"]}},\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"additionalProperties\":false,\"required\":[\"id\"]}",
			wantErr: false,
		},
		{
			name:    "在结构体的dependencies加上不存在的key",
			args:    args{reflect.TypeOf(MyObject{}), &Schema{Dependencies: map[string][]string{"name": {"date", "coords"}}}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.t, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr == true {
				return
			}

			schema, _ := json.Marshal(got)

			if string(schema) != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
