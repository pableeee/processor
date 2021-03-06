package ds

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
)

var ErrorNotStruct = fmt.Errorf("not a struct")

/*
type RedisClient interface {
	SaveDocument(id string, document interface{}) error
	SaveDocumentWithVersion(id string, document interface{}, version uint64) error
	BulkBuilder() godsclient.BulkBuilder

	SearchBuilder() godsclient.SearchBuilder
	ScrollBuilder() godsclient.ScrollBuilder
	CountBuilder() godsclient.CountBuilder

	DeleteDocument(id string) error
	DeleteDocumentWithVersion(id string, version uint64) error
	DeleteDocumentsByQuery(queryBuilder querybuilders.QueryBuilder) (*godsclient.DeleteByQueryResponse, error)

	CreateSchema(prefix string, i interface{}) error
}

*/

type RedisClientImpl struct {
	client *redisearch.Client
}

func MakeRedisClient(address string, port uint16) *RedisClientImpl {
	r := RedisClientImpl{}
	r.client = redisearch.NewClient(fmt.Sprintf("%s:%d", address, port), "dsIndex")

	return &r
}

func (rd *RedisClientImpl) CreateSchema(prefix string, i interface{}) error {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ErrorNotStruct
	}

	sch := redisearch.NewSchema(redisearch.DefaultOptions)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Type.Kind() {
		case reflect.String:
			sch.AddField(redisearch.NewTextField(f.Name))
		case reflect.Int:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			sch.AddField(redisearch.NewNumericField(f.Name))
		case reflect.Struct:
			// aca quiero procesar algun struct emb, tipo 'time'
			_, ok := v.Field(i).Interface().(time.Time)
			if !ok {
				continue
			}

			sch.AddField(redisearch.NewNumericField(f.Name))
		default:
			fmt.Printf("field %s not indexable", f.Name)
		}
	}

	err := rd.client.Drop()
	if err != nil {
		log.Println("no index to drop")
	}

	if err := rd.client.CreateIndex(sch); err != nil {
		log.Println("Error creating the index")

		return err
	}

	return nil
}

func (rd *RedisClientImpl) buildDocument(id string, i interface{}) (redisearch.Document, error) {
	doc := redisearch.NewDocument(id, 1.0)

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return doc, ErrorNotStruct
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		doc.Set(f.Name, v.Field(i).Interface())
	}

	return doc, nil
}

func (rd *RedisClientImpl) SaveDocument(id string, i interface{}) error {
	doc, err := rd.buildDocument(id, i)
	if err != nil {
		return err
	}
	// Index the document. The API accepts multiple documents at a time
	if err = rd.client.Index(doc); err != nil {
		log.Println("Error creating the index")

		return err
	}

	return nil
}

func (rd *RedisClientImpl) SaveDocumentWithVersion(id string, i interface{}, version uint64) error {
	doc, err := rd.buildDocument(id, i)
	if err != nil {
		return err
	}

	doc.Set("version", version)
	// Index the document. The API accepts multiple documents at a time
	if err = rd.client.Index(doc); err != nil {
		log.Println("Error creating the index")

		return err
	}

	return nil
}

/*
func (rd *RedisClientImpl) BulkBuilder() godsclient.BulkBuilder {
	return nil
}

func (rd *RedisClientImpl) SearchBuilder() godsclient.SearchBuilder {
	return nil
}

func (rd *RedisClientImpl) ScrollBuilder() godsclient.ScrollBuilder {
	return &redisScrollBuilder{client: rd.client}
}

func (rd *RedisClientImpl) CountBuilder() godsclient.CountBuilder {
	return nil
}

func (rd *RedisClientImpl) DeleteDocument(id string) error {
	return rd.client.Delete(id, true)
}

func (rd *RedisClientImpl) DeleteDocumentWithVersion(id string, version uint64) error {
	return fmt.Errorf("not supported")
}

func (rd *RedisClientImpl) DeleteDocumentsByQuery(queryBuilder querybuilders.QueryBuilder) (*godsclient.DeleteByQueryResponse, error) {
	return nil, nil
}
*/
