package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("csbbkdqwtdwxdmq")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT COUNT(type) as type,(ROW_NUMBER() OVER()) as id from events;"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("wf9iny4q")

		// add
		new_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ayvmsmeg",
			"name": "type",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_type)
		collection.Schema.AddField(new_type)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("csbbkdqwtdwxdmq")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT type,(ROW_NUMBER() OVER()) as id from events;"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wf9iny4q",
			"name": "type",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_type)
		collection.Schema.AddField(del_type)

		// remove
		collection.Schema.RemoveField("ayvmsmeg")

		return dao.SaveCollection(collection)
	})
}
