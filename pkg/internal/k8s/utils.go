package k8s

import (
	"fmt"
	"reflect"

	un "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//UnwrapNodePort Retrieves the nodeport assined from the k8s api result
func unwrapNodePort(result *un.Unstructured) (map[string]int64, error) {
	res := make(map[string]int64)

	keys := []string{"nodePort", "targetPort"}
	ports, ok, err := un.NestedSlice(result.UnstructuredContent(), "spec", "ports")

	if ok && err == nil && len(ports) > 0 {
		for _, ele := range ports {
			port, _ := ele.(map[string]interface{})
			for _, key := range keys {
				num, _ := port[key].(int64)
				res[key] = num
			}
		}
	}

	return res, nil
}

func (r *Response) GetString(k ...string) (string, error) {
	s, err := r.GetValue(reflect.String, k...)
	if err != nil {
		return "", err
	}
	return s.(string), nil
}

func (r *Response) GetInt64(k ...string) (int64, error) {
	s, err := r.GetValue(reflect.Int64, k...)
	if err != nil {
		return 0, err
	}
	return s.(int64), nil
}

func (r *Response) GetBool(k ...string) (bool, error) {
	s, err := r.GetValue(reflect.Bool, k...)
	if err != nil {
		return false, err
	}
	return s.(bool), nil
}

func (r *Response) GetMap(k ...string) (map[string]interface{}, error) {
	s, err := r.GetValue(reflect.Map, k...)
	if err != nil {
		return nil, err
	}
	return s.(map[string]interface{}), nil
}

func (r *Response) GetSlice(k ...string) ([]interface{}, error) {
	s, err := r.GetValue(reflect.Slice, k...)
	if err != nil {
		return nil, err
	}
	return s.([]interface{}), nil
}

func (r *Response) GetValue(t reflect.Kind, k ...string) (interface{}, error) {
	var ok bool
	var err error
	var v interface{}

	switch t {
	case reflect.Slice:
		v, ok, err = un.NestedSlice(r.Value, k...)
	case reflect.Map:
		v, ok, err = un.NestedMap(r.Value, k...)
	case reflect.String:
		v, ok, err = un.NestedString(r.Value, k...)
	case reflect.Bool:
		v, ok, err = un.NestedBool(r.Value, k...)
	case reflect.Int64:
		v, ok, err = un.NestedInt64(r.Value, k...)
	default:
		err = fmt.Errorf("invalid kind")
	}

	if err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("key %s not found", k)
	}

	return v, nil
}
