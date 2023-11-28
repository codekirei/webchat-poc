package gensql

const headerTemplate = `// Do not edit this file directly!
// This code was automatically generated.
// Call "go generate" to regenerate it from latest source.

package %s
`

const constTemplate = "\nconst %s = `%s`\n"
