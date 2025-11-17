package grammar

import (
	"hash/fnv"
	"io"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/set"
)

// E is the empty string ε.
var E = String[Symbol]{}

var (
	EqString   = eqString
	CmpString  = cmpString
	HashString = hashFuncForString()
)

// String represent a string of grammar symbols.
type String[T Symbol] []T

// String returns a string representation of a string of symbols.
func (s String[T]) String() string {
	if len(s) == 0 {
		return "ε"
	}

	names := make([]string, len(s))
	for i, symbol := range s {
		names[i] = symbol.String()
	}

	return strings.Join(names, " ")
}

// Equal determines whether or not two strings are the same.
func (s String[T]) Equal(rhs String[T]) bool {
	if len(s) != len(rhs) {
		return false
	}

	for i := range s {
		if !s[i].Equal(rhs[i]) {
			return false
		}
	}

	return true
}

// ContainsSymbol checks whether a string contains the given symbol.
func (s String[T]) ContainsSymbol(symbol T) bool {
	for _, sym := range s {
		if sym.Equal(symbol) {
			return true
		}
	}

	return false
}

// HasPrefix checks whether a string starts with the given prefix.
func (s String[T]) HasPrefix(prefix String[T]) bool {
	ls, lp := len(s), len(prefix)
	return ls >= lp && s[:lp].Equal(prefix)
}

// HasSuffix checks whether a string ends with the given suffix.
func (s String[T]) HasSuffix(prefix String[T]) bool {
	ls, lp := len(s), len(prefix)
	return ls >= lp && s[ls-lp:].Equal(prefix)
}

// AnyMatch returns true if at least one symbol satisfies the provided predicate.
func (s String[T]) AnyMatch(pred generic.Predicate1[T]) bool {
	for _, sym := range s {
		if pred(sym) {
			return true
		}
	}

	return false
}

// Append appends new symbols to the current string and returns a new string.
func (s String[T]) Append(syms ...T) String[T] {
	newS := make(String[T], len(s)+len(syms))

	copy(newS, s)
	copy(newS[len(s):], syms)

	return newS
}

// Prepend prepends new symbols to the current string and returns a new string.
func (s String[T]) Prepend(syms ...T) String[T] {
	newS := make(String[T], len(syms)+len(s))

	copy(newS, syms)
	copy(newS[len(syms):], s)

	return newS
}

// Concat concatenates the current string with one or more strings and returns a new string.
func (s String[T]) Concat(ss ...String[T]) String[T] {
	l := len(s)
	for _, t := range ss {
		l += len(t)
	}

	newS := make(String[T], l)

	copy(newS, s)
	i := len(s)
	for _, t := range ss {
		copy(newS[i:], t)
		i += len(t)
	}

	return newS
}

// Terminals returns all terminal symbols of a string of symbols.
func (s String[Symbol]) Terminals() String[Terminal] {
	terms := String[Terminal]{}
	for _, sym := range s {
		if v, ok := any(sym).(Terminal); ok {
			terms = append(terms, v)
		}
	}
	return terms
}

// NonTerminals returns all non-terminal symbols of a string of symbols.
func (s String[Symbol]) NonTerminals() String[NonTerminal] {
	nonTerms := String[NonTerminal]{}
	for _, sym := range s {
		if v, ok := any(sym).(NonTerminal); ok {
			nonTerms = append(nonTerms, v)
		}
	}
	return nonTerms
}

// LongestCommonPrefixOf computes the longest common prefix of a list of strings.
// If the input is empty or there is no common prefix, it returns the empty string ε.
func LongestCommonPrefixOf(ss ...String[Symbol]) String[Symbol] {
	if len(ss) == 0 {
		return E
	}

	lcp := ss[0]

	for i := 1; i < len(ss); i++ {
		for !ss[i].HasPrefix(lcp) {
			lcp = lcp[:len(lcp)-1]
			if len(lcp) == 0 {
				return E
			}
		}
	}

	return lcp
}

// WriteString writes a string of symbols to the provided io.Writer.
// It returns the number of bytes written and any error encountered.
func WriteString(w io.Writer, s String[Symbol]) (n int, err error) {
	total := 0
	for _, x := range s {
		n, err := WriteSymbol(w, x)
		total += n

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

func eqString(lhs, rhs String[Symbol]) bool {
	return lhs.Equal(rhs)
}

func eqStringSet(lhs, rhs set.Set[String[Symbol]]) bool {
	return lhs.Equal(rhs)
}

// cmpString is a CompareFunc for String[Symbol] type.
//
// The comparing criteria are as follows:
//
//  1. strings contain more non-terminal symbols are prioritized first.
//  3. If two strings have the same number of non-terminals,
//     those with more terminal symbols come first.
//  4. If two strings have the same number of non-terminals and terminals,
//     they are compared alphabetically based on their symbols.
//
// This function can be used for sorting strings
// to ensure a consistent and deterministic order for any given set of strings.
func cmpString(lhs, rhs String[Symbol]) int {
	// First, compare based on the number of non-terminal symbols in the strings.
	lhsNonTermLen, rhsNonTermLen := len(lhs.NonTerminals()), len(rhs.NonTerminals())
	if lhsNonTermLen > rhsNonTermLen {
		return -1
	} else if rhsNonTermLen > lhsNonTermLen {
		return 1
	}

	// Second, if the number of non-terminals is the same, compare based on the number of terminals.
	lhsTermLen, rhsTermLen := len(lhs.Terminals()), len(rhs.Terminals())
	if lhsTermLen > rhsTermLen {
		return -1
	} else if rhsTermLen > lhsTermLen {
		return 1
	}

	// Next, if the number of terminals is also the same, compare alphabetically based on the string representations.
	lhsString, rhsString := lhs.String(), rhs.String()
	if lhsString < rhsString {
		return -1
	} else if rhsString < lhsString {
		return 1
	}

	return 0
}

func hashFuncForString() hash.HashFunc[String[Symbol]] {
	h := fnv.New64()

	return func(s String[Symbol]) uint64 {
		h.Reset()
		_, _ = WriteString(h, s) // Hash.Write never returns an error
		return h.Sum64()
	}
}
