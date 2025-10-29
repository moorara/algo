// Package dot provides a basic implementation of the DOT language for visualizing graphs.
package dot

import "strings"

// fieldReplacer escapes special characters in DOT field labels.
var fieldReplacer = strings.NewReplacer(
	`|`, `\|`,
	`{`, `\{`,
	`}`, `\}`,
	`<`, `\<`,
	`>`, `\>`,
)

// labelReplacer escapes special characters in DOT node and edge labels.
var labelReplacer = strings.NewReplacer(
	`\t`, `\\t`,
	`\n`, `\\n`,
	`\v`, `\\v`,
	`\f`, `\\f`,
	`\r`, `\\r`,
	`\x`, `0x`,
	`\u`, `0x`,
	`\U`, `0x`,
)
