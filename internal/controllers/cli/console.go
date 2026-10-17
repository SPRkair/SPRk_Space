package cli

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type console struct {
	sc *bufio.Scanner
}

func newConsole(sc *bufio.Scanner) *console {
	return &console{sc: sc}
}

func (c console) ReadLine(prompt string) (string, bool) {
	fmt.Printf(prompt)
	if !c.sc.Scan() {
		return "", false
	}
	return strings.TrimSpace(c.sc.Text()), true
}

func (c console) ReadInt64(prompt string) (int64, bool) {
	s, ok := c.ReadLine(prompt)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("\nInvalid number\n")
		return 0, false
	}
	return n, true
}

func (c console) Pause() {
	fmt.Println("Press Enter")
	c.sc.Scan()
}

type editField struct {
	label    string
	value    reflect.Value // адресуемое значение внутри draft
	readonly bool
}

var (
	durationType = reflect.TypeOf(time.Duration(0))
	timeType     = reflect.TypeOf(time.Time{})
)

func editStruct[T any](c console, title string, src *T) (*T, bool) {
	tmp := *src
	fields := collectFields(reflect.ValueOf(&tmp).Elem(), "")

	for {
		fmt.Printf("\n== %s ==\n", title)
		for i, f := range fields {
			note := ""
			if f.readonly {
				note = "  (read-only)"
			}
			fmt.Printf("%d. %s: %v%s\n", i+1, f.label, f.value.Interface(), note)
		}
		fmt.Println("s. Save")
		fmt.Println("c. Cancel")

		line, ok := c.ReadLine("> ")
		if !ok {
			return nil, false
		}

		switch strings.ToLower(line) {
		case "c":
			return nil, false
		case "s":
			if v, ok := any(&tmp).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					fmt.Println("Cli: editStruct: Validation error:", err)
					continue
				}
			}
			return &tmp, true
		}

		n, err := strconv.Atoi(line)
		if err != nil || n < 1 || n > len(fields) {
			fmt.Println("Unknown choice")
			continue
		}
		f := fields[n-1]
		if f.readonly {
			fmt.Println("Field is read-only")
			continue
		}

		s, ok := c.ReadLine(fmt.Sprintf("New value for %s: ", f.label))
		if !ok {
			return nil, false
		}
		err = setFromString(f.value, s)
		if err != nil {
			fmt.Println("Invalid value:", err)
		}
	}
}

func collectFields(v reflect.Value, prefix string) []editField {
	var res []editField
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		// тег: cli:"Метка,ro" или cli:"-"
		label, readonly := sf.Name, false
		if tag, ok := sf.Tag.Lookup("cli"); ok {
			if tag == "-" {
				continue
			}
			parts := strings.Split(tag, ",")
			if parts[0] != "" {
				label = parts[0]
			}
			for _, opt := range parts[1:] {
				if opt == "ro" {
					readonly = true
				}
			}
		}
		label = prefix + label

		fv := v.Field(i)
		// вложенная структура (кроме time.Time) раскрывается рекурсивно
		if fv.Kind() == reflect.Struct && fv.Type() != timeType {
			res = append(res, collectFields(fv, label+".")...)
			continue
		}
		res = append(res, editField{label: label, value: fv, readonly: readonly})
	}
	return res
}

func setFromString(v reflect.Value, s string) error {
	switch {
	case v.Type() == durationType:
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		v.SetInt(int64(d))
		return nil
	case v.Type() == timeType:
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return fmt.Errorf("expected RFC3339, e.g. 2026-01-31T15:04:05Z")
		}
		v.Set(reflect.ValueOf(t))
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 10, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetUint(n)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetFloat(f)
	default:
		return fmt.Errorf("type %s is not editable", v.Type())
	}
	return nil
}
