package slugid

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/taskcluster/taskcluster-cli/extpoints"
)

func TestSlugidV4(t *testing.T) {
	assert := assert.New(t)

	// Let's test 100 generated slugs, to increase confidence
	for i := 0; i < 1000; i++ {
		slug, _ := v4()
		match := V4_SLUG_REGEXP.MatchString(slug)
		assert.Equal(match, true, fmt.Sprintf("Slug generated by v4() is invalid."))
	}
}

func TestSlugidNice(t *testing.T) {
	assert := assert.New(t)

	// Let's test 100 generated slugs, to increase confidence
	for i := 0; i < 1000; i++ {
		slug, _ := nice()
		match := NICE_SLUG_REGEXP.MatchString(slug)
		assert.Equal(match, true, fmt.Sprintf("Slug generated by nice() is invalid."))
	}
}

func TestSlugidValidDecode(t *testing.T) {
	assert := assert.New(t)

	context := extpoints.Context{}
	context.Arguments = make(map[string]interface{})
	context.Arguments["<slug>"] = "FOH9mI0mQ1C90yoMo3ajsg"

	out, err := decode(context)
	expected := "14e1fd98-8d26-4350-bdd3-2a0ca376a3b2"

	assert.Nil(err, fmt.Sprintf("Error decoding test slug.\n%s", err))
	assert.Equal(out, expected, fmt.Sprintf("Got wrong output when decoding slug. Expected %s, got %s", expected, out))
}

func TestSlugidInvalidDecode(t *testing.T) {
	assert := assert.New(t)

	slug := "abc"

	context := extpoints.Context{}
	context.Arguments = make(map[string]interface{})
	context.Arguments["<slug>"] = slug

	_, err := decode(context)

	assert.NotNil(err, fmt.Sprintf("Expected error when decoding invalid slug '%s'.", slug))
}

func TestSlugidValidEncode(t *testing.T) {
	assert := assert.New(t)

	context := extpoints.Context{}
	context.Arguments = make(map[string]interface{})
	context.Arguments["<uuid>"] = "14e1fd98-8d26-4350-bdd3-2a0ca376a3b2"

	out, err := encode(context)
	expected := "FOH9mI0mQ1C90yoMo3ajsg"

	assert.Nil(err, fmt.Sprintf("Error encoding test uuid.\n%s", err))
	assert.Equal(out, expected, fmt.Sprintf("Got wrong output when encoding uuid. Expected %s, got %s", expected, out))
}

func TestSlugidInvalidEncode(t *testing.T) {
	assert := assert.New(t)

	uuid := "xyz"

	context := extpoints.Context{}
	context.Arguments = make(map[string]interface{})
	context.Arguments["<uuid>"] = uuid

	_, err := encode(context)

	assert.NotNil(err, fmt.Sprintf("Expected error when encoding invalid uuid '%s'.", uuid))
}
