package database

type DatabaseTypes struct {
	table []DatabaseType
}

var _ IDatabaseType = (*DatabaseTypes)(nil)

func NewDatabaseTypes(tp ...func(*DatabaseTypes)) *DatabaseTypes {
	t := &DatabaseTypes{}
	for _, f := range tp {
		f(t)
	}
	return t
}

func (d *DatabaseTypes) TableName() string {
	return d.table[0].name
}

func (d *DatabaseTypes) addTable(t DatabaseType) {
	d.table = append(d.table, t)
}

func (d *DatabaseTypes) GetTables() []DatabaseType {
	return d.table
}

type DatabaseType struct {
	name        string
	tp          string
	constraints []string
}

func (d *DatabaseType) TableName() string {
	return d.name
}

func (d *DatabaseTypes) Id(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		t.addTable(DatabaseType{name: name, tp: "int", constraints: []string{"PRIMARY KEY", "AUTOINCREMENT"}})
	}
}

func (d *DatabaseTypes) String(name string, length int) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if length == 0 {
			length = 255
		}
		t.addTable(DatabaseType{name: name, tp: "varchar(" + string(length) + ")"})
	}
}

func (d *DatabaseTypes) Bool(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		t.addTable(DatabaseType{name: name, tp: "bool"})
	}
}

func (d *DatabaseTypes) Int(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		t.addTable(DatabaseType{name: name, tp: "int"})
	}
}

func (d *DatabaseTypes) Float(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		t.addTable(DatabaseType{name: name, tp: "float"})
	}
}

func (d *DatabaseTypes) Null(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if len(d.table) > 0 {
			d.table[len(d.table)-1].constraints = append(d.table[len(d.table)-1].constraints, "NULL")
		}
	}
}

func (d *DatabaseTypes) NotNull(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if len(d.table) > 0 {
			d.table[len(d.table)-1].constraints = append(d.table[len(d.table)-1].constraints, "NOT NULL")
		}
	}
}

func (d *DatabaseTypes) PrimaryKey(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if len(d.table) > 0 {
			d.table[len(d.table)-1].constraints = append(d.table[len(d.table)-1].constraints, "PRIMARY KEY")
		}
	}
}

func (d *DatabaseTypes) AutoIncrement(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if len(d.table) > 0 {
			d.table[len(d.table)-1].constraints = append(d.table[len(d.table)-1].constraints, "AUTOINCREMENT")
		}
	}
}

func (d *DatabaseTypes) Unique(name string) func(*DatabaseTypes) {
	return func(t *DatabaseTypes) {
		if len(d.table) > 0 {
			d.table[len(d.table)-1].constraints = append(d.table[len(d.table)-1].constraints, "UNIQUE")
		}
	}
}

func (d *DatabaseTypes) Timestamp() func(*DatabaseTypes) {
	// criar por padrao created_at e updated_at
	return func(t *DatabaseTypes) {
		t.addTable(DatabaseType{name: "created_at", tp: "timestamp"})
		t.addTable(DatabaseType{name: "updated_at", tp: "timestamp"})
	}
}
