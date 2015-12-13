// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 * NullTest
 *
 * NullTest represents the operation of testing a value for NULLness.
 * The appropriate test is performed and returned as a boolean Datum.
 *
 * NOTE: the semantics of this for rowtype inputs are noticeably different
 * from the scalar case.  We provide an "argisrow" flag to reflect that.
 * ----------------
 */
type NullTest struct {
	Xpr          Node         `json:"xpr"`
	Arg          Node         `json:"arg"`          /* input expression */
	Nulltesttype NullTestType `json:"nulltesttype"` /* IS NULL, IS NOT NULL */
	Argisrow     bool         `json:"argisrow"`     /* T if input is of a composite type */
}

func (node NullTest) MarshalJSON() ([]byte, error) {
	type NullTestMarshalAlias NullTest
	return json.Marshal(map[string]interface{}{
		"NullTest": (*NullTestMarshalAlias)(&node),
	})
}

func (node *NullTest) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		node.Xpr, err = UnmarshalNodeJSON(fields["xpr"])
		if err != nil {
			return
		}
	}

	if fields["arg"] != nil {
		node.Arg, err = UnmarshalNodeJSON(fields["arg"])
		if err != nil {
			return
		}
	}

	if fields["nulltesttype"] != nil {
		err = json.Unmarshal(fields["nulltesttype"], &node.Nulltesttype)
		if err != nil {
			return
		}
	}

	if fields["argisrow"] != nil {
		err = json.Unmarshal(fields["argisrow"], &node.Argisrow)
		if err != nil {
			return
		}
	}

	return
}
