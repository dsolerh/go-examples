package types

type NumericTypes interface{ int | float64 }

type LiteralTypes interface{ NumericTypes | string | bool }
