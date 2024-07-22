package migrations

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type Command struct {
	TableName string
	Schema    string
}

type ConfigCommand struct {
	Migrations struct {
		Dir string `yaml:"dir"`
	} `yaml:"migrations"`
}

var config ConfigCommand

const templateMigrations = `package migrations

import (
	"github.com/diogoX451/gopherORM/internal/database"
	"github.com/diogoX451/gopherORM/internal/database/migrations"
)

var _ migrations.IMigration = (&{{.TableNameUpper}})(nil)

func (m {{.TableNameUpper}}) GetTableName() string {
	return "{{.TableName}}"
}

type {{.TableNameUpper}} struct {
	columns database.IDatabaseType
}

func New{{.TableNameUpper}}() *{{.TableNameUpper}} {
	return &{{.TableName}}{}
}

func (m {{.TableNameUpper}}) Up() database.DatabaseTypes {
	return database.NewDatabaseTypes()
}

func (m {{.TableNameUpper}}) Down() database.DatabaseTypes {
	return database.NewDatabaseTypes()
}
`

func NewCommand(name string, schema string) *Command {
	return &Command{
		TableName: name,
		Schema:    schema,
	}
}

func (c *Command) Run() {
	data, err := os.ReadFile("gopher.yaml")
	if err != nil {
		panic("Error loading gopher.yaml file: " + err.Error())
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Error unmarshalling YAML: " + err.Error())
	}

	fmt.Println(config.Migrations.Dir)

	c.createArchive()
}

func (c Command) createArchive() {

	title := cases.Title(language.Make(c.TableName)).String(c.TableName)

	data := struct {
		TableName      string
		TableNameUpper string
	}{
		TableName:      c.TableName,
		TableNameUpper: title,
	}

	tmp, err := template.New("migrations").Parse(templateMigrations)
	if err != nil {
		panic("Error parsing template: " + err.Error())
	}

	err = os.MkdirAll(config.Migrations.Dir, os.ModePerm)
	if err != nil {
		panic("Error creating migrations directory: " + err.Error())
	}

	filePath := fmt.Sprintf("%s/%s_%s.go", config.Migrations.Dir, time.Now().Format("2006_01_02_15_04_05"), c.TableName)
	file, err := os.Create(filePath)
	if err != nil {
		panic("Error creating file: " + err.Error())
	}
	defer file.Close()

	err = tmp.Execute(file, data)
	if err != nil {
		panic("Error executing template: " + err.Error())
	}

	fmt.Println("Migration file created successfully:", filePath)
}
