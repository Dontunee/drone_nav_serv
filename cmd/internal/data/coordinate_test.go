package data

import "testing"

func TestCoordinate_Getloc(t *testing.T) {
	type arg struct {
		coordinate Coordinate
		vec        string
	}

	type result struct {
		loc   float64
		error string
	}

	tests := []struct {
		name string
		args arg
		res  result
	}{
		{
			"valid input returns valid response and nil error ",
			arg{
				coordinate: Coordinate{
					X: 123.12,
					Y: 456.56,
					Z: 789.89,
				},
				vec: "20.0",
			},
			result{
				1389.57,
				"",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := test.args.coordinate.Getloc(test.args.vec)
			if resp != test.res.loc {
				t.Errorf("Test failed with unexpected loc response %f", resp)
			}
			if err != test.res.error {
				t.Errorf("Test failed with unexpected error response %s", err)
			}
		})
	}

}
func TestCoordinate_Init(t *testing.T) {
	type arg struct {
		coordinate Coordinate
		x,
		y,
		z string
	}
	tests := []struct {
		name string
		args arg
		err  string
	}{
		{
			"Valid float input returns no error",
			arg{Coordinate{}, "20.57", "10.93", "30.84"},
			"",
		},
		{
			"Invalid input should return error",
			arg{Coordinate{}, "20.57fsd", "10.93dsd", "30.8dsds4"},
			invalidInput,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.args.coordinate.Init(test.args.x, test.args.y, test.args.z)
			if res != test.err {
				t.Errorf("Test failed")
			}
		})

	}

}
func TestIsFixed(t *testing.T) {
	type input struct {
		num       float64
		precision int
	}
	tests := []struct {
		name   string
		args   input
		result float64
	}{
		{
			"valid input to 2 d.p",
			input{
				1369.9405,
				2,
			},
			1369.94,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := toFixed(test.args.num, test.args.precision); res != test.result {
				t.Errorf("Test failed with this result %f", test.result)
			}
		})
	}
}
