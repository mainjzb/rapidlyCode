package main

import (
	"context"
	"log"
	"strings"
	"sync"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"maplestorySkill/ent"
	"maplestorySkill/ent/migrate"
)

type Skill struct {
	Name           string
	Type           string // 主动技能/被动技能
	JobAdvancement string
	Icons          []string
	IconName       string
	RequireLevel   int
	Detail         string // 技能详情连接
	MaxLevel       int
	Description    string

	MechanicsLevel  string
	MechanicsDetail string
}

var client *ent.Client
var createClient sync.Once

func SqlClient() *ent.Client {
	createClient.Do(func() {
		var err error
		client, err = ent.Open(dialect.SQLite, "./skill.db?_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		// defer client.Close()
		ctx := context.Background()

		// Run the automatic migration tool to create all schema resources.
		if err := client.Schema.Create(ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	})
	return client
}

func Close() {
	if client != nil {
		client.Close()
	}
}

func AddSkill(skill Skill) {
	if client == nil {
		SqlClient()
	}
	ctx := context.Background()

	client.HeroSkill.Query().Where()

	client.HeroSkill.Create().
		SetName(skill.Name).
		SetType(skill.Type).
		SetJobAdvancement(skill.JobAdvancement).
		SetIcon(strings.Join(skill.Icons, "|")).
		SetIconName(skill.IconName).
		SetRequireLevel(skill.RequireLevel).
		SetDetail(skill.Detail).
		SetMaxLevel(skill.MaxLevel).
		SetDescription(skill.Description).
		SetMechanicsLevel(skill.MechanicsLevel).
		SetMechanicsDetail(skill.MechanicsDetail).
		Save(ctx)
}
