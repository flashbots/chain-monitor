package config

import (
	"reflect"

	"github.com/flashbots/chain-monitor/utils"
)

type Config struct {
	L1     *L1     `yaml:"l1"`
	L2     *L2     `yaml:"l2"`
	Log    *Log    `yaml:"log"`
	Server *Server `yaml:"server"`
}

func New() *Config {
	return &Config{
		L1:     &L1{},
		Log:    &Log{},
		Server: &Server{},

		L2: &L2{
			Monitor: &Monitor{},
		},
	}
}

func (c *Config) Validate() error {
	return validate(c)
}

func validate(item interface{}) error {
	v := reflect.ValueOf(item)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	errs := []error{}
	for idx := 0; idx < v.NumField(); idx++ {
		field := v.Field(idx)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		if v, ok := field.Interface().(validatee); ok {
			if err := v.Validate(); err != nil {
				errs = append(errs, err)
			}
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.Struct:
			if err := validate(field.Interface()); err != nil {
				errs = append(errs, err)
			}
		case reflect.Slice, reflect.Array:
			for jdx := 0; jdx < field.Len(); jdx++ {
				if err := validate(field.Index(jdx).Interface()); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	return utils.FlattenErrors(errs)
}
