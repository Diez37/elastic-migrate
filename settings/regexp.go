package settings

// see https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html#field.summary
const (
    RegularFlagCanonEq               JavaRegularFlag = "CANON_EQ"
    RegularFlagCaseInsensitive       JavaRegularFlag = "CASE_INSENSITIVE"
    RegularFlagComments              JavaRegularFlag = "COMMENTS"
    RegularFlagDotAll                JavaRegularFlag = "DOTALL"
    RegularFlagLiteral               JavaRegularFlag = "LITERAL"
    RegularFlagMultiline             JavaRegularFlag = "MULTILINE"
    RegularFlagUnicodeCase           JavaRegularFlag = "UNICODE_CASE"
    RegularFlagUnicodeCharacterClass JavaRegularFlag = "UNICODE_CHARACTER_CLASS"
    RegularFlagUnixLines             JavaRegularFlag = "UNIX_LINES"

    JavaRegularFlagSeparator string = "|"
)

type JavaRegularFlag string
