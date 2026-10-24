package boa

import flag "github.com/spf13/pflag"

// RequiredFlagAnnotation marks a flag as required. Set it with MarkFlagRequired.
const RequiredFlagAnnotation = "boa_annotation_required_flag"

// MarkFlagRequired causes the command to report an error if invoked without the named flag.
func (c *Command) MarkFlagRequired(name string) error {
	return MarkFlagRequired(c.Flags(), name)
}

// MarkPersistentFlagRequired causes the command to report an error if invoked without the named persistent flag.
func (c *Command) MarkPersistentFlagRequired(name string) error {
	return MarkFlagRequired(c.PersistentFlags(), name)
}

// MarkFlagRequired causes a command to report an error if invoked without the named flag in flags.
func MarkFlagRequired(flags *flag.FlagSet, name string) error {
	return flags.SetAnnotation(name, RequiredFlagAnnotation, []string{"true"})
}
