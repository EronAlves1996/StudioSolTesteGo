package validators

import "testing"

func TestValidators(t *testing.T) {
	t.Run("test minSize with 5 characters", func(t *testing.T) {
		got := Validators["minSize"]("onefunct", 5)
		expect := true

		assert(got, expect, t)

		got2 := Validators["minSize"]("one3", 5)
		expect2 := false

		assert(got2, expect2, t)
	})

	t.Run("test minSize with 10 characters", func(t *testing.T) {
		got := Validators["minSize"]("with10Chars", 10)
		expect := true
		assert(got, expect, t)
	})

	t.Run("test minUppercase with 2 characters", func(t *testing.T) {
		got := Validators["minUppercase"]("Ithavenoneupperacse", 2)
		expect := false

		assert(got, expect, t)
	})

	t.Run("test minUppercase with 5 characters", func(t *testing.T) {
		got := Validators["minUppercase"]("ITHAVE5uppeCASAcase", 5)
		expect := true
		assert(got, expect, t)
	})

	t.Run("test minLowercase with 6 characters", func(t *testing.T) {
		got := Validators["minLowerCase"]("IhavesTHeexpectect", 6)
		expect := true
		assert(got, expect, t)
	})

	t.Run("test minLowercase with 12 characters", func(t *testing.T) {
		got := Validators["minLowerCase"]("IdontHaveMINIMUMCHARS", 12)
		expect := false
		assert(got, expect, t)
	})
}

func assert(got bool, expect bool, t *testing.T) {
	t.Helper()
	if got != expect {
		t.Errorf("Got %t, expect %t", got, expect)
	}
}
