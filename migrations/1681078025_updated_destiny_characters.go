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

		collection, err := dao.FindCollectionByNameOrId("n16hg12txqa3jw8")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("cwlv8t6y")

		// add
		new_bnet_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tcskhd9n",
			"name": "bnet_id",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_bnet_id)
		collection.Schema.AddField(new_bnet_id)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("n16hg12txqa3jw8")
		if err != nil {
			return err
		}

		// add
		del_bnet_id := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cwlv8t6y",
			"name": "bnet_id",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), del_bnet_id)
		collection.Schema.AddField(del_bnet_id)

		// remove
		collection.Schema.RemoveField("tcskhd9n")

		return dao.SaveCollection(collection)
	})
}
