package repository

import "strings"

// JSON_SEARCH matches individual labels. Escape SQL pattern characters literally.
// Match more requested labels first, then prefer whole-label matches.
func tagSearch(tags []string) (string, []any, string, []any) {
	if len(tags) == 0 {
		return "", nil, "", nil
	}
	fuzzy := []string{}
	exact := []string{}
	args := []any{}
	exactArgs := []any{}
	for _, tag := range tags {
		fuzzy = append(fuzzy, "(JSON_SEARCH(LOWER(CAST(p.tags AS CHAR)), 'one', LOWER(?), '!') IS NOT NULL)")
		exact = append(exact, "JSON_CONTAINS(LOWER(CAST(p.tags AS CHAR)), JSON_QUOTE(LOWER(?)))")
		pattern := strings.NewReplacer("!", "!!", "%", "!%", "_", "!_").Replace(tag)
		args = append(args, "%"+pattern+"%")
		exactArgs = append(exactArgs, tag)
	}
	rankArgs := append(append([]any{}, args...), exactArgs...)
	return "(" + strings.Join(fuzzy, " OR ") + ")", args, "(" + strings.Join(fuzzy, " + ") + ") DESC,(" + strings.Join(exact, " + ") + ") DESC,", rankArgs
}
