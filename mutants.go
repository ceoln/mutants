package mutants

import "github.com/ceoln/expressions"
import "math/rand"

type Mutant struct {
	expressions.Expression
}

var Zero = Mutant{expressions.Zero}

type roughCopyVisitor struct {
	accuracy expressions.Float
	m        map[string]expressions.Float
}

func (r roughCopyVisitor) VisitConstant(value expressions.Float) (expressions.ExpressionLike, bool) {
	accuracy := r.accuracy
	if expressions.Float(rand.Float32()) < accuracy {
		return Mutant{expressions.NewConstant(value)}, true
	} else {
		return Mutant{expressions.NewRandomConstant()}, true
	}
}

func (r roughCopyVisitor) VisitVariableRef(s string) (expressions.ExpressionLike, bool) {
	accuracy := r.accuracy
	if expressions.Float(rand.Float32()) < accuracy {
		return Mutant{expressions.NewVariableRef(s)}, true
	} else {
		return Mutant{expressions.NewRandomVariableRef(r.m)}, true
	}
}

func (r roughCopyVisitor) VisitBinaryOperation(op byte, lhsv, rhsv expressions.ExpressionLike) (expressions.ExpressionLike, bool) {
	accuracy := r.accuracy
	accurate := expressions.Float(rand.Float32()) < accuracy
	if !accurate {
		if expressions.Float(rand.Float32()) < 0.5 {
			return Mutant{expressions.NewRandomBinOp(r.m)}, true
		} else {
			op = "+-*/"[rand.Intn(4)]
		}
	}
	lhs, okay := lhsv.(Mutant)
	if !okay {
		return Zero, false
	}
	rhs, okay := rhsv.(Mutant)
	if !okay {
		return Zero, false
	}
	left, okay := lhs.RoughCopy(r.accuracy, r.m)
	if !okay {
		return Zero, false
	}
	right, okay := rhs.RoughCopy(r.accuracy, r.m)
	if !okay {
		return Zero, false
	}
	return Mutant{expressions.NewBinaryOperation(op, left, right)}, true
}

func (mut Mutant) RoughCopy(accuracy expressions.Float, m map[string]expressions.Float) (Mutant, bool) {
	v := roughCopyVisitor{accuracy, m}
	answerv, okay := mut.Visit(v)
	if !okay {
		return Zero, false
	}
	answer, okay := answerv.(Mutant)
	if !okay {
		return Zero, false
	}
	return answer, true
}