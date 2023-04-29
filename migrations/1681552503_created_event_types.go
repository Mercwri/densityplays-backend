package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "ghfidcu6xqj9279",
			"created": "2023-04-15 09:55:03.673Z",
			"updated": "2023-04-15 09:55:03.673Z",
			"name": "event_types",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "0jvsszg9",
					"name": "bnet_id",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "pdgxtbxf",
					"name": "event_name",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_DpxzLBx` + "`" + ` ON ` + "`" + `event_types` + "`" + ` (` + "`" + `bnet_id` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ghfidcu6xqj9279")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
