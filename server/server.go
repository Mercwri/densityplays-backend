package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/Mercwri/densityplays/internal"
	_ "github.com/Mercwri/densityplays/migrations"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/spf13/cobra"
)

func Init(app *pocketbase.PocketBase) {
	rootCmd := app.RootCmd
	rootCmd.AddCommand(BungieInitCmd(app))
	rootCmd.AddCommand(BungieLinkCmd(app))
	rootCmd.AddCommand(BungieEventsCmd(app))
	rootCmd.AddCommand(TruncateEvents(app))
}

func Serve() {
	app := pocketbase.New()
	Init(app)
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/topkek/:username",
			Handler: func(c echo.Context) error {
				return GetUserActivityMap(app, c, c.PathParam("username"))
			},
		})
		return nil
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func BungieInitCmd(app *pocketbase.PocketBase) *cobra.Command {
	command := &cobra.Command{
		Use:   "bungie-init",
		Short: "interact with bungie api",
		Run: func(command *cobra.Command, args []string) {
			user_coll, _ := app.Dao().FindCollectionByNameOrId("bnet_users")
			clan_coll, _ := app.Dao().FindCollectionByNameOrId("bnet_clan")
			char_coll, _ := app.Dao().FindCollectionByNameOrId("destiny_characters")
			gid := internal.GetGroupByName(args[1], 1)
			gmem := internal.GetGroupMembers(gid)
			for _, g := range gmem.Response.Results {
				clan, _ := app.Dao().FindFirstRecordByData(clan_coll.Id, "bnet_id", gid)
				user := models.NewRecord(user_coll)
				user.Set("username", g.DestinyUserInfo.LastSeenDisplayName)
				user.Set("bnet_id", g.DestinyUserInfo.MembershipId)
				user.Set("membership_type", g.DestinyUserInfo.MembershipType)
				user.Set("clan", clan.Id)
				app.Dao().SaveRecord(user)
				prof := internal.GetProfile(g.DestinyUserInfo.MembershipType, g.DestinyUserInfo.MembershipId)
				for _, c := range prof.Response.Profile.Data.CharacterIds {
					char := models.NewRecord(char_coll)
					char.Set("bnet_id", c)
					char.Set("bnet_user", user.Id)
					app.Dao().SaveRecord(char)
				}
			}
		},
	}
	return command
}

func BungieLinkCmd(app *pocketbase.PocketBase) *cobra.Command {
	command := &cobra.Command{
		Use:   "bungie-link",
		Short: "interact with bungie api",
		Run: func(command *cobra.Command, args []string) {
			user_coll, _ := app.Dao().FindCollectionByNameOrId("bnet_users")
			char_coll, _ := app.Dao().FindCollectionByNameOrId("destiny_characters")
			query := app.Dao().RecordQuery(user_coll)
			users := []*models.Record{}
			query.All(&users)
			for _, u := range users {
				query := app.Dao().RecordQuery(char_coll).Where(dbx.HashExp{"bnet_user": u.Id})
				chars := []*models.Record{}
				query.All(&chars)
				cids := []string{}
				for _, c := range chars {
					cids = append(cids, c.Id)
				}
				u.Set("characters", cids)
				app.Dao().SaveRecord(u)
			}

		},
	}
	return command
}

func BungieEventsCmd(app *pocketbase.PocketBase) *cobra.Command {
	return &cobra.Command{
		Use:   "bungie-events",
		Short: "ingest events for character",
		Run: func(command *cobra.Command, args []string) {
			user_coll, _ := app.Dao().FindCollectionByNameOrId("bnet_users")
			char_coll, _ := app.Dao().FindCollectionByNameOrId("destiny_characters")
			events_coll, _ := app.Dao().FindCollectionByNameOrId("events")
			query := app.Dao().RecordQuery(char_coll)
			chars := []*models.Record{}
			query.All(&chars)
			for _, c := range chars {
				owner, err := app.Dao().FindRecordById(user_coll.Id, c.GetString("bnet_user"))
				if err != nil {
					log.Panic(err)
				}
				pahs := internal.GetPlayerActivityHistory(int32(owner.GetInt("membership_type")), owner.GetString("bnet_id"), c.GetString("bnet_id"))
				for _, pah := range pahs {
					for _, e := range pah.Response.Activities {
						event, err := app.Dao().FindFirstRecordByData(events_coll.Id, "bnet_id", e.ActivityDetails.InstanceId)
						fmt.Println(event)
						if err != nil {
							fmt.Println(err)
							if err.Error() == "sql: no rows in result set" {
								fmt.Println(pah)
								fmt.Println("insert")
								ev := models.NewRecord(events_coll)
								ev.Set("duration", e.Values.ActivityDurationSeconds.Basic.Value)
								ev.Set("characters", owner.Id)
								ev.Set("bnet_id", e.ActivityDetails.InstanceId)
								ev.Set("type", e.ActivityDetails.ReferenceId)
								ev.Set("start_time", e.Period)
								st, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", e.Period)
								if err != nil {
									log.Panic(err)
								}
								fmt.Println(st.String())
								et := st.Add((time.Second * time.Duration(e.Values.ActivityDurationSeconds.Basic.Value)))
								fmt.Println(et.String())
								ev.Set("end_time", et.String())
								app.Dao().SaveRecord(ev)
							} else {
								log.Panic(err)
							}
						} else {
							fmt.Println("upsert")
							chars := event.GetStringSlice("characters")
							fmt.Println(chars)
							if !contains(chars, owner.Id) {
								chars = append(chars, owner.Id)
								fmt.Println(chars)
								event.Set("characters", chars)
								app.Dao().SaveRecord(event)
							}
							if event.GetInt("duration") == 0 {
								event.Set("duration", e.Values.ActivityDurationSeconds.Basic.Value)
								app.Dao().SaveRecord(event)
							}
							if event.GetInt("type") == 0 {
								event.Set("type", e.ActivityDetails.ReferenceId)
								app.Dao().SaveRecord(event)
							}
							if event.GetString("end_time") == "" {
								st, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", e.Period)
								fmt.Println(st.String())
								et := st.Add((time.Second * time.Duration(e.Values.ActivityDurationSeconds.Basic.Value)))
								fmt.Println(et.String())
								event.Set("end_time", et.String())
								app.Dao().SaveRecord(event)
							}
						}
					}
				}
			}
			users := []*models.Record{}
			query.All(&users)
			rendered_data_coll, _ := app.Dao().FindCollectionByNameOrId("rendered_data")
			for _, u := range users {
				url := fmt.Sprintf("https://db.densityplays/topkek/%s", u.Username())
				mappy, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
				}
				body, _ := io.ReadAll(mappy.Body)
				datum, err := app.Dao().FindFirstRecordByData(rendered_data_coll.Id, "bnet_id.username", u.Username())
				if err.Error() == "sql: no rows in result set" {
					dm := models.NewRecord(rendered_data_coll)
					dm.Set("bnet_user", u.Id)
					dm.Set("rendered_data", body)
					app.Dao().Save(dm)

				} else {
					datum.Set("rendered_data", body)
					app.Dao().Save(datum)
				}
			}
		},
	}
}

func TruncateEvents(app *pocketbase.PocketBase) *cobra.Command {
	command := &cobra.Command{
		Use:   "bungie-events-truncate",
		Short: "interact with bungie api",
		Run: func(command *cobra.Command, args []string) {
			events_coll, _ := app.Dao().FindCollectionByNameOrId("events")
			all_events := []*models.Record{}
			query := app.Dao().RecordQuery(events_coll)
			query.All(&all_events)
			for _, e := range all_events {
				app.Dao().DeleteRecord(e)
			}
		},
	}
	return command
}

func GetUserActivityMap(app *pocketbase.PocketBase, c echo.Context, username string) error {
	init_user, err := app.Dao().FindFirstRecordByData("bnet_users", "username", username)
	if err != nil {
		return apis.NewNotFoundError("this user does not exist", err)
	}
	events_coll, _ := app.Dao().FindCollectionByNameOrId("events")
	record := []*models.Record{}
	query := app.Dao().RecordQuery(events_coll).Where(dbx.Like("characters", init_user.Id))
	query.All(&record)
	partners := &SortedUsers{}
	raid_defs := &RaidDefs{}
	for _, r := range record {
		for _, c := range r.GetStringSlice("characters") {
			cname, _ := app.Dao().FindRecordById("bnet_users", c)
			username := cname.GetString("username")
			if username != init_user.GetString("username") {
				r_type := r.GetString("type")
				user := partners.GetOrAddUser(username)
				user.AddRaid(r_type, raid_defs)
			}
		}
	}

	sort.Sort(partners)
	for _, p := range partners.Users {
		sort.Sort(p)
	}
	return c.JSON(http.StatusOK, partners)
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

type User struct {
	Name       string
	TotalRaids int32
	Raids      []Raid
}

func (u *User) AddRaid(refId string, raid_defs *RaidDefs) {
	u.TotalRaids++
	raid := u.GetOrAddRaid(refId, raid_defs)
	raid.RaidCount++
}

func (u *User) GetOrAddRaid(refId string, raid_defs *RaidDefs) *Raid {
	var raid Raid
	raid_def := raid_defs.GetRaidDef(refId)
	raid_name := raid_def.Name
	for i, r := range u.Raids {
		if r.Name == raid_name {
			return &u.Raids[i]
		}
	}
	if raid.Name == "" {
		raid.Name = raid_name
		u.Raids = append(u.Raids, raid)
	}
	return u.GetOrAddRaid(refId, raid_defs)
}

type Raid struct {
	Name      string
	RaidCount int32
}

type SortedUsers struct {
	Users []User
}

func (p User) Len() int           { return len(p.Raids) }
func (p User) Swap(i, j int)      { p.Raids[i], p.Raids[j] = p.Raids[j], p.Raids[i] }
func (p User) Less(i, j int) bool { return p.Raids[i].RaidCount > p.Raids[j].RaidCount }

func (p SortedUsers) Len() int           { return len(p.Users) }
func (p SortedUsers) Swap(i, j int)      { p.Users[i], p.Users[j] = p.Users[j], p.Users[i] }
func (p SortedUsers) Less(i, j int) bool { return p.Users[i].TotalRaids > p.Users[j].TotalRaids }
func (p *SortedUsers) GetOrAddUser(name string) *User {
	user := User{}
	for i, u := range p.Users {
		if u.Name == name {
			return &p.Users[i]
		}
	}
	if user.Name == "" {
		user.Name = name
		p.Users = append(p.Users, user)
	}
	return p.GetOrAddUser(name)
}

type RaidDefs struct {
	RaidDefs []RaidDef
}

type RaidDef struct {
	Name string
	Id   string
}

func (rds *RaidDefs) GetRaidDef(refId string) *RaidDef {
	rd := RaidDef{}
	for i, r := range rds.RaidDefs {
		if r.Id == refId {
			return &rds.RaidDefs[i]
		}
	}
	if rd.Name == "" {
		api_def := internal.GetDestinyEntityDefinition(refId)
		rd.Name = api_def.Response.DisplayProperties.Name
		rd.Id = refId
		rds.RaidDefs = append(rds.RaidDefs, rd)
	}
	return rds.GetRaidDef(refId)
}
