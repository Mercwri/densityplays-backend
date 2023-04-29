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
			"id": "n16hg12txqa3jw8",
			"created": "2023-04-09 07:04:02.571Z",
			"updated": "2023-04-09 07:04:02.571Z",
			"name": "destiny_characters",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "cf9gdive",
					"name": "class",
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
					"id": "cwlv8t6y",
					"name": "bnet_id",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				},
				{
					"system": false,
					"id": "k6gegay0",
					"name": "field",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "df654xxn69ysgkf",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_QYrrBXj` + "`" + ` ON ` + "`" + `destiny_characters` + "`" + ` (` + "`" + `bnet_id` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("n16hg12txqa3jw8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
