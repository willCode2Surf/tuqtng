//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"testing"

	"github.com/couchbaselabs/dparval"
)

func TestCompare(t *testing.T) {

	nonExistantProperty := NewProperty("dne")
	null := NewLiteralNull()
	numberSixty := NewLiteralNumber(60.0)
	numberNine := NewLiteralNumber(9.0)
	stringBob := NewLiteralString("bob")
	stringCat := NewLiteralString("cat")
	patternMatchSingle := NewLiteralString("b_b")
	patternMatchMulti := NewLiteralString("b%")
	patternNoMatch := NewLiteralString("____")

	tests := ExpressionTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), false, nil},
		{NewGreaterThanOperator(numberSixty, numberNine), true, nil},
		{NewGreaterThanOperator(numberNine, numberSixty), false, nil},
		{NewGreaterThanOperator(null, numberSixty), nil, nil},
		{NewGreaterThanOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewGreaterThanOperator(stringBob, numberSixty), false, nil},

		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), true, nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberNine), true, nil},
		{NewGreaterThanOrEqualOperator(numberNine, numberSixty), false, nil},
		{NewGreaterThanOrEqualOperator(null, numberSixty), nil, nil},
		{NewGreaterThanOrEqualOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewGreaterThanOrEqualOperator(stringBob, numberSixty), false, nil},

		{NewLessThanOperator(numberSixty, numberSixty), false, nil},
		{NewLessThanOperator(numberSixty, numberNine), false, nil},
		{NewLessThanOperator(numberNine, numberSixty), true, nil},
		{NewLessThanOperator(null, numberSixty), nil, nil},
		{NewLessThanOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOperator(stringBob, numberSixty), false, nil},

		{NewLessThanOrEqualOperator(numberSixty, numberSixty), true, nil},
		{NewLessThanOrEqualOperator(numberSixty, numberNine), false, nil},
		{NewLessThanOrEqualOperator(numberNine, numberSixty), true, nil},
		{NewLessThanOrEqualOperator(null, numberSixty), nil, nil},
		{NewLessThanOrEqualOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOrEqualOperator(stringBob, numberSixty), false, nil},

		{NewEqualToOperator(numberSixty, numberSixty), true, nil},
		{NewEqualToOperator(numberSixty, numberNine), false, nil},
		{NewEqualToOperator(numberNine, numberSixty), false, nil},
		{NewEqualToOperator(null, numberSixty), nil, nil},
		{NewEqualToOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewEqualToOperator(stringBob, numberSixty), false, nil},

		{NewNotEqualToOperator(numberSixty, numberSixty), false, nil},
		{NewNotEqualToOperator(numberSixty, numberNine), true, nil},
		{NewNotEqualToOperator(numberNine, numberSixty), true, nil},
		{NewNotEqualToOperator(null, numberSixty), nil, nil},
		{NewNotEqualToOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewNotEqualToOperator(stringBob, numberSixty), false, nil},

		{NewLikeOperator(stringBob, patternMatchSingle), true, nil},
		{NewLikeOperator(stringCat, patternMatchSingle), false, nil},
		{NewLikeOperator(stringBob, patternMatchMulti), true, nil},
		{NewLikeOperator(stringCat, patternMatchMulti), false, nil},
		{NewLikeOperator(stringBob, patternNoMatch), false, nil},
		{NewLikeOperator(stringBob, numberNine), nil, nil},
		{NewLikeOperator(numberNine, patternMatchSingle), nil, nil},
		{NewLikeOperator(stringBob, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewLikeOperator(nonExistantProperty, stringBob), nil, &dparval.Undefined{"dne"}},

		{NewNotLikeOperator(stringBob, patternMatchSingle), false, nil},
		{NewNotLikeOperator(stringCat, patternMatchSingle), true, nil},
		{NewNotLikeOperator(stringBob, patternMatchMulti), false, nil},
		{NewNotLikeOperator(stringCat, patternMatchMulti), true, nil},
		{NewNotLikeOperator(stringBob, patternNoMatch), true, nil},
		{NewNotLikeOperator(stringBob, numberNine), nil, nil},
		{NewNotLikeOperator(numberNine, patternMatchSingle), nil, nil},
		{NewNotLikeOperator(stringBob, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewNotLikeOperator(nonExistantProperty, stringBob), nil, &dparval.Undefined{"dne"}},

		// these tests all conform to the table in the specification
		{NewIsNullOperator(stringBob), false, nil},
		{NewIsNullOperator(null), true, nil},
		{NewIsNullOperator(nonExistantProperty), false, nil},

		{NewIsNotNullOperator(stringBob), true, nil},
		{NewIsNotNullOperator(null), false, nil},
		{NewIsNotNullOperator(nonExistantProperty), false, nil},

		{NewIsMissingOperator(stringBob), false, nil},
		{NewIsMissingOperator(null), false, nil},
		{NewIsMissingOperator(nonExistantProperty), true, nil},

		{NewIsNotMissingOperator(stringBob), true, nil},
		{NewIsNotMissingOperator(null), true, nil},
		{NewIsNotMissingOperator(nonExistantProperty), false, nil},

		{NewIsValuedOperator(stringBob), true, nil},
		{NewIsValuedOperator(null), false, nil},
		{NewIsValuedOperator(nonExistantProperty), false, nil},

		{NewIsNotValuedOperator(stringBob), false, nil},
		{NewIsNotValuedOperator(null), true, nil},
		{NewIsNotValuedOperator(nonExistantProperty), false, nil},

		// new tests to check ordering of missing/null values
		{NewLessThanOperator(numberSixty, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOperator(numberSixty, null), nil, nil},

		//comparing booleans
		{NewEqualToOperator(NewLiteralBool(true), NewLiteralBool(false)), false, nil},
		{NewEqualToOperator(NewLiteralBool(false), NewLiteralBool(true)), false, nil},
	}

	tests.Run(t)

}

func TestCompareStringRepresentation(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionStringTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), "60 > 60"},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), "60 >= 60"},
		{NewLessThanOperator(numberSixty, numberSixty), "60 < 60"},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), "60 <= 60"},
		{NewEqualToOperator(numberSixty, numberSixty), "60 = 60"},
		{NewNotEqualToOperator(numberSixty, numberSixty), "60 != 60"},
		{NewLikeOperator(stringBob, stringBob), `"bob" LIKE "bob"`},
		{NewNotLikeOperator(stringBob, stringBob), `"bob" NOT LIKE "bob"`},

		{NewIsNullOperator(stringBob), `"bob" IS NULL`},
		{NewIsNotNullOperator(stringBob), `"bob" IS NOT NULL`},
		{NewIsMissingOperator(stringBob), `"bob" IS MISSING`},
		{NewIsNotMissingOperator(stringBob), `"bob" IS NOT MISSING`},
		{NewIsValuedOperator(stringBob), `"bob" IS VALUED`},
		{NewIsNotValuedOperator(stringBob), `"bob" IS NOT VALUED`},
	}

	tests.Run(t)
}

func TestCompareValidate(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionValidateTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewLessThanOperator(numberSixty, numberSixty), nil},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewEqualToOperator(numberSixty, numberSixty), nil},
		{NewNotEqualToOperator(numberSixty, numberSixty), nil},
		{NewLikeOperator(stringBob, stringBob), nil},
		{NewNotLikeOperator(stringBob, stringBob), nil},

		{NewIsNullOperator(stringBob), nil},
		{NewIsNotNullOperator(stringBob), nil},
		{NewIsMissingOperator(stringBob), nil},
		{NewIsNotMissingOperator(stringBob), nil},
		{NewIsValuedOperator(stringBob), nil},
		{NewIsNotValuedOperator(stringBob), nil},

		// first arg not valid
		{NewGreaterThanOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewGreaterThanOrEqualOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLessThanOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLessThanOrEqualOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewEqualToOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewNotEqualToOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLikeOperator(notValidExpression, stringBob), notValidExpressionError},
		{NewNotLikeOperator(notValidExpression, stringBob), notValidExpressionError},

		{NewIsNullOperator(notValidExpression), notValidExpressionError},
		{NewIsNotNullOperator(notValidExpression), notValidExpressionError},
		{NewIsMissingOperator(notValidExpression), notValidExpressionError},
		{NewIsNotMissingOperator(notValidExpression), notValidExpressionError},
		{NewIsValuedOperator(notValidExpression), notValidExpressionError},
		{NewIsNotValuedOperator(notValidExpression), notValidExpressionError},

		// second arg not valid
		{NewGreaterThanOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewGreaterThanOrEqualOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLessThanOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLessThanOrEqualOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewEqualToOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewNotEqualToOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLikeOperator(stringBob, notValidExpression), notValidExpressionError},
		{NewNotLikeOperator(stringBob, notValidExpression), notValidExpressionError},
	}

	tests.Run(t)
}

func TestCompareVerifyFormalNotation(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionVerifyFormalNotationTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), NewGreaterThanOperator(numberSixty, numberSixty), nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), NewGreaterThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewLessThanOperator(numberSixty, numberSixty), NewLessThanOperator(numberSixty, numberSixty), nil},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), NewLessThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewEqualToOperator(numberSixty, numberSixty), NewEqualToOperator(numberSixty, numberSixty), nil},
		{NewNotEqualToOperator(numberSixty, numberSixty), NewNotEqualToOperator(numberSixty, numberSixty), nil},
		{NewLikeOperator(stringBob, stringBob), NewLikeOperator(stringBob, stringBob), nil},
		{NewNotLikeOperator(stringBob, stringBob), NewNotLikeOperator(stringBob, stringBob), nil},

		{NewIsNullOperator(stringBob), NewIsNullOperator(stringBob), nil},
		{NewIsNotNullOperator(stringBob), NewIsNotNullOperator(stringBob), nil},
		{NewIsMissingOperator(stringBob), NewIsMissingOperator(stringBob), nil},
		{NewIsNotMissingOperator(stringBob), NewIsNotMissingOperator(stringBob), nil},
		{NewIsValuedOperator(stringBob), NewIsValuedOperator(stringBob), nil},
		{NewIsNotValuedOperator(stringBob), NewIsNotValuedOperator(stringBob), nil},

		// first arg not formal
		{NewGreaterThanOperator(notFormalExpression, numberSixty), NewGreaterThanOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewGreaterThanOrEqualOperator(notFormalExpression, numberSixty), NewGreaterThanOrEqualOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewLessThanOperator(notFormalExpression, numberSixty), NewLessThanOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewLessThanOrEqualOperator(notFormalExpression, numberSixty), NewLessThanOrEqualOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewEqualToOperator(notFormalExpression, numberSixty), NewEqualToOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewNotEqualToOperator(notFormalExpression, numberSixty), NewNotEqualToOperator(notFormalExpression, numberSixty), notFormalExpressionError},
		{NewLikeOperator(notFormalExpression, stringBob), NewLikeOperator(notFormalExpression, stringBob), notFormalExpressionError},
		{NewNotLikeOperator(notFormalExpression, stringBob), NewNotLikeOperator(notFormalExpression, stringBob), notFormalExpressionError},

		{NewIsNullOperator(notFormalExpression), NewIsNullOperator(notFormalExpression), notFormalExpressionError},
		{NewIsNotNullOperator(notFormalExpression), NewIsNotNullOperator(notFormalExpression), notFormalExpressionError},
		{NewIsMissingOperator(notFormalExpression), NewIsMissingOperator(notFormalExpression), notFormalExpressionError},
		{NewIsNotMissingOperator(notFormalExpression), NewIsNotMissingOperator(notFormalExpression), notFormalExpressionError},
		{NewIsValuedOperator(notFormalExpression), NewIsValuedOperator(notFormalExpression), notFormalExpressionError},
		{NewIsNotValuedOperator(notFormalExpression), NewIsNotValuedOperator(notFormalExpression), notFormalExpressionError},

		// second arg not formal
		{NewGreaterThanOperator(numberSixty, notFormalExpression), NewGreaterThanOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewGreaterThanOrEqualOperator(numberSixty, notFormalExpression), NewGreaterThanOrEqualOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewLessThanOperator(numberSixty, notFormalExpression), NewLessThanOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewLessThanOrEqualOperator(numberSixty, notFormalExpression), NewLessThanOrEqualOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewEqualToOperator(numberSixty, notFormalExpression), NewEqualToOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewNotEqualToOperator(numberSixty, notFormalExpression), NewNotEqualToOperator(numberSixty, notFormalExpression), notFormalExpressionError},
		{NewLikeOperator(stringBob, notFormalExpression), NewLikeOperator(stringBob, notFormalExpression), notFormalExpressionError},
		{NewNotLikeOperator(stringBob, notFormalExpression), NewNotLikeOperator(stringBob, notFormalExpression), notFormalExpressionError},
	}

	tests.Run(t, []string{}, []string{"bucket", "child"}, "")

	// again with single bucket
	tests = ExpressionVerifyFormalNotationTestSet{

		// first arg not formal
		{NewGreaterThanOperator(notFormalExpression, numberSixty), NewGreaterThanOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewGreaterThanOrEqualOperator(notFormalExpression, numberSixty), NewGreaterThanOrEqualOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewLessThanOperator(notFormalExpression, numberSixty), NewLessThanOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewLessThanOrEqualOperator(notFormalExpression, numberSixty), NewLessThanOrEqualOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewEqualToOperator(notFormalExpression, numberSixty), NewEqualToOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewNotEqualToOperator(notFormalExpression, numberSixty), NewNotEqualToOperator(notFormalExpressionAfter, numberSixty), nil},
		{NewLikeOperator(notFormalExpression, stringBob), NewLikeOperator(notFormalExpressionAfter, stringBob), nil},
		{NewNotLikeOperator(notFormalExpression, stringBob), NewNotLikeOperator(notFormalExpressionAfter, stringBob), nil},

		{NewIsNullOperator(notFormalExpression), NewIsNullOperator(notFormalExpressionAfter), nil},
		{NewIsNotNullOperator(notFormalExpression), NewIsNotNullOperator(notFormalExpressionAfter), nil},
		{NewIsMissingOperator(notFormalExpression), NewIsMissingOperator(notFormalExpressionAfter), nil},
		{NewIsNotMissingOperator(notFormalExpression), NewIsNotMissingOperator(notFormalExpressionAfter), nil},
		{NewIsValuedOperator(notFormalExpression), NewIsValuedOperator(notFormalExpressionAfter), nil},
		{NewIsNotValuedOperator(notFormalExpression), NewIsNotValuedOperator(notFormalExpressionAfter), nil},

		// second arg not formal
		{NewGreaterThanOperator(numberSixty, notFormalExpression), NewGreaterThanOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewGreaterThanOrEqualOperator(numberSixty, notFormalExpression), NewGreaterThanOrEqualOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewLessThanOperator(numberSixty, notFormalExpression), NewLessThanOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewLessThanOrEqualOperator(numberSixty, notFormalExpression), NewLessThanOrEqualOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewEqualToOperator(numberSixty, notFormalExpression), NewEqualToOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewNotEqualToOperator(numberSixty, notFormalExpression), NewNotEqualToOperator(numberSixty, notFormalExpressionAfter), nil},
		{NewLikeOperator(stringBob, notFormalExpression), NewLikeOperator(stringBob, notFormalExpressionAfter), nil},
		{NewNotLikeOperator(stringBob, notFormalExpression), NewNotLikeOperator(stringBob, notFormalExpressionAfter), nil},
	}

	tests.Run(t, []string{}, []string{"bucket"}, "bucket")
}

func TestTypeMismatchError(t *testing.T) {
	err := TypeMismatch{4, 3}

	if err.Error() != "Types do not match, 4 3" {
		t.Errorf("Expected Types do not match 4,3, got %v", err.Error())
	}
}
